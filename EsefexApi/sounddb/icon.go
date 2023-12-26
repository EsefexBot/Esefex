package sounddb

import (
	"esefexapi/util"
	"fmt"
	"regexp"
)

type Icon struct {
	RegularEmoji bool   `json:"regularEmoji"`
	Name         string `json:"name"`
	ID           string `json:"id"`
	Url          string `json:"url"`
}

func NewCustomIcon(name string, id string) Icon {
	return Icon{
		RegularEmoji: false,
		Name:         name,
		ID:           id,
		Url:          fmt.Sprintf("https://cdn.discordapp.com/emojis/%s.webp?quality=lossless", id),
	}
}

func NewEmojiIcon(emoji string) Icon {
	return Icon{
		RegularEmoji: true,
		Name:         emoji,
		Url:          util.GetEmojiURL(emoji),
	}
}

func (i *Icon) String() string {
	if i.RegularEmoji {
		return i.Name
	}
	return fmt.Sprintf("<:%s:%s>", i.Name, i.ID)
}

var iconRegex string = `<:([^:]+):(\d+)>`
var r *regexp.Regexp = regexp.MustCompile(fmt.Sprintf(`%s|(%s)`, iconRegex, util.EmojiRegex))

func ExtractIcon(str string) (Icon, error) {
	m := r.FindStringSubmatch(str)
	if m == nil {
		return Icon{}, fmt.Errorf("invalid icon, no match")
	}
	if len(m) != 4 {
		return Icon{}, fmt.Errorf("invalid icon, len(m) != 4")
	}

	if m[1] != "" && m[2] != "" {
		return NewCustomIcon(m[1], m[2]), nil
	}

	if m[3] != "" {
		return NewEmojiIcon(m[3]), nil
	}

	return Icon{}, fmt.Errorf("invalid icon")
}
