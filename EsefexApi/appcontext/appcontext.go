package appcontext

import (
	// "net/http"

	"esefexapi/audioprocessing"
	"esefexapi/msg"

	"github.com/bwmarrin/discordgo"
)

type Context struct {
	CustomProtocol string
	ApiPort        string
	Channels       Channels
	DiscordSession *discordgo.Session
	AudioCache     *audioprocessing.AudioCache
}

type Channels struct {
	// A2B chan msg.MessageA2B
	PlaySound chan msg.PlaySound
	// B2A  chan msg.MessageB2A
	Stop chan struct{}
}
