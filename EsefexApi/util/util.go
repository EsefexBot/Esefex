package util

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func ExtractIconUrl(icon *discordgo.ApplicationCommandInteractionDataOption) string {
	r, err := regexp.Compile(`<:.+:\d+>`)
	if err != nil {
		panic(err)
	}

	m := r.FindString(fmt.Sprint(icon.Value))

	rn, err := regexp.Compile(`\d+`)
	if err != nil {
		panic(err)
	}

	id := rn.FindString(m)

	return fmt.Sprintf("https://cdn.discordapp.com/emojis/%v.webp", id)
}

func GetSoundURL(guildID, name string) string {
	return fmt.Sprintf("https://cdn.discordapp.com/attachments/%v/%v.mp3", guildID, name)
}

func GetEmojiURL(emoji string) string {
	runes := []rune(emoji)

	name := strings.Join(lo.Map(runes, func(r rune, index int) string { return fmt.Sprintf("%x", r) }), "-")

	url := fmt.Sprintf("https://raw.githubusercontent.com/twitter/twemoji/master/assets/svg/%s.svg", name)

	return url
}

var TokenCharset = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandomString(charset []rune, length int) string {
	str := make([]rune, 32)
	for i := range str {
		str[i] = charset[rand.Intn(len(charset))]
	}

	return string(str)
}

func EnsureFile(p string) (*os.File, error) {
	err := os.MkdirAll(path.Dir(p), os.ModePerm)
	if err != nil {
		return nil, errors.Wrap(err, "Error creating directory")
	}

	file, err := os.OpenFile(p, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, errors.Wrap(err, "Error opening file")
	}

	return file, nil
}

func ToGenericArray(arr ...interface{}) []interface{} {
	return arr
}

func FirstNRunes(s string, n int) string {
	i := 0
	for j := range s {
		if i == n {
			return s[:j]
		}
		i++
	}
	return s
}
