package util

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
)

func TestClampInt(t *testing.T) {
	assert.Equal(t, 0, ClampInt(-1, 0, 10))
	assert.Equal(t, 10, ClampInt(11, 0, 10))
	assert.Equal(t, 5, ClampInt(5, 0, 10))
}

func TestExtractIconUrl(t *testing.T) {
	assert.Equal(t, "https://cdn.discordapp.com/emojis/123.webp", ExtractIconUrl(&discordgo.ApplicationCommandInteractionDataOption{
		Value: "<:test:123>",
	}))
	assert.Equal(t, "https://cdn.discordapp.com/emojis/123.webp", ExtractIconUrl(&discordgo.ApplicationCommandInteractionDataOption{
		Value: "<:test:123> ",
	}))
	assert.Equal(t, "https://cdn.discordapp.com/emojis/123.webp", ExtractIconUrl(&discordgo.ApplicationCommandInteractionDataOption{
		Value: " <:test:123>",
	}))
	assert.Equal(t, "https://cdn.discordapp.com/emojis/123.webp", ExtractIconUrl(&discordgo.ApplicationCommandInteractionDataOption{
		Value: " <:test:123> ",
	}))
}
