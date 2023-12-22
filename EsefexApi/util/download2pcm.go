package util

import (
	"fmt"
	"net/http"
	"os/exec"
	"regexp"
	"slices"
)

var validExt = []string{"mp3", "wav", "ogg", "flac", "m4a", "aac"}

func Download2PCM(url string) ([]int16, error) {
	ext, err := ExtFromUrl(url)
	if err != nil {
		return nil, err
	}

	if !slices.Contains(validExt, ext) {
		return nil, fmt.Errorf("invalid file extension: %v", ext)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 || resp.Header.Get("Content-Type") != "audio/mpeg" {
		return nil, fmt.Errorf("status code error: %v", resp.StatusCode)
	}

	cmd := exec.Command("ffmpeg", "-hide_banner", "-loglevel", "error", "-i", "pipe:0", "-f", "s16le", "-ac", "2", "-ar", "48000", "-")
	cmd.Stdin = resp.Body

	out, err := cmd.Output()
	if err != nil {
		return nil, err
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
		return "", err
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
		return "", err
	}

	m := r.FindStringSubmatch(url)
	if len(m) != 2 {
		return "", fmt.Errorf("no extension found in url")
	}

	return m[1], nil
}
