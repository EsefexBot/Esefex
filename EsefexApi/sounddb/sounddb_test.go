package sounddb

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func ExtractIconTest(t *testing.T) {
	text := "<:emoji:630819109726191617>🀄🆘🧌🤡🆘"
	icon, err := ExtractIcon(text)
	assert.Nil(t, err)
	assert.Equal(t, Icon{
		RegularEmoji: false,
		Name:         "emoji",
		ID:           "630819109726191617",
		Url:          "https://cdn.discordapp.com/emojis/630819109726191617.webp?quality=lossless",
	}, icon)

	text = "👨‍👨‍👧‍👦🀄🆘🧌🤡🆘"
	icon, err = ExtractIcon(text)
	assert.Nil(t, err)
	assert.Equal(t, Icon{
		RegularEmoji: true,
		Name:         "👨‍👨‍👧‍👦",
		Url:          "https://raw.githubusercontent.com/twitter/twemoji/master/assets/svg/1f468-200d-1f468-200d-1f467-200d-1f466.svg",
	}, icon)

	text = "<:emoji:630819109726191617>"
	icon, err = ExtractIcon(text)
	assert.Nil(t, err)
	assert.Equal(t, Icon{
		RegularEmoji: false,
		Name:         "emoji",
		ID:           "630819109726191617",
		Url:          "https://cdn.discordapp.com/emojis/630819109726191617.webp?quality=lossless",
	}, icon)

	text = "asdasc<:emoji:630819109726191617>asdasc"
	icon, err = ExtractIcon(text)
	assert.Nil(t, err)
	assert.Equal(t, Icon{
		RegularEmoji: false,
		Name:         "emoji",
		ID:           "630819109726191617",
		Url:          "https://cdn.discordapp.com/emojis/630819109726191617.webp?quality=lossless",
	}, icon)

	text = "asdasc<:emoji1:111>asdasc<:emoji2:222>asdasc"
	icon, err = ExtractIcon(text)
	assert.Nil(t, err)
	assert.Equal(t, Icon{
		RegularEmoji: false,
		Name:         "emoji1",
		ID:           "111",
		Url:          "https://cdn.discordapp.com/emojis/111.webp?quality=lossless",
	}, icon)

	text = "asdas<><<s;;::c🤡"
	icon, err = ExtractIcon(text)
	assert.Nil(t, err)
	assert.Equal(t, Icon{
		RegularEmoji: true,
		Name:         "🤡",
		Url:          "https://raw.githubusercontent.com/twitter/twemoji/master/assets/svg/1f921.svg",
	}, icon)
}
