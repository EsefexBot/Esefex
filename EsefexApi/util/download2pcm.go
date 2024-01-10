package util

import (
	"fmt"
	// "log"
	"net/http"
	"os/exec"
	"regexp"
	"slices"

	"github.com/pkg/errors"
)

var validExt = []string{"mp3", "wav", "ogg", "flac", "m4a", "aac"}

func Download2PCM(url string) ([]int16, error) {
	ext, err := ExtFromUrl(url)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting extension from url")
	}

	if !slices.Contains(validExt, ext) {
		return nil, fmt.Errorf("invalid file extension: %v", ext)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "Error downloading sound")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 || resp.Header.Get("Content-Type") != "audio/mpeg" {
		return nil, errors.Wrapf(err, "Error downloading sound: %v", resp.StatusCode)
	}

	// reject if content length is too large

	var maxFileSize int64 = 5_000_000

	if resp.ContentLength > maxFileSize {
		return nil, fmt.Errorf("file is too large (%v bytes > %v bytes)", resp.ContentLength, maxFileSize)
	}

	cmd := exec.Command("ffmpeg", "-hide_banner", "-loglevel", "error", "-i", "pipe:0", "-f", "s16le", "-ac", "2", "-ar", "48000", "-")
	cmd.Stdin = resp.Body

	out, err := cmd.Output()
	if err != nil {
		return nil, errors.Wrap(err, "Error converting sound")
	}

	pcm := make([]int16, len(out)/2)
	for i := 0; i < len(out); i += 2 {
		pcm[i/2] = int16(out[i]) | int16(out[i+1])<<8
	}

	return pcm, nil
}

func ExtFromDisposition(disposition string) (string, error) {
	r, err := regexp.Compile(`filename=".+\.(.+)"`)
	if err != nil {
		return "", errors.Wrap(err, "Error compiling regex")
	}

	m := r.FindStringSubmatch(disposition)
	if len(m) != 2 {
		return "", fmt.Errorf("no extension found in disposition")
	}

	return m[1], nil
}

func ExtFromUrl(url string) (string, error) {
	r, err := regexp.Compile(`^.+\.(.+)\?.+$`)
	if err != nil {
		return "", errors.Wrap(err, "Error compiling regex")
	}

	m := r.FindStringSubmatch(url)
	if len(m) != 2 {
		return "", fmt.Errorf("no extension found in url")
	}

	return m[1], nil
}
