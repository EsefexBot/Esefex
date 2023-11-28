package appcontext

import (
	"esefexbot/msg"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

type Context struct {
	CustomProtocol string
	Channels       Channels
	DiscordSession *discordgo.Session
	ApiPort        string
}

type Channels struct {
	A2B  chan msg.MessageA2B
	B2A  chan msg.MessageB2A
	Stop chan bool
}

func Wrap(fn func(http.ResponseWriter, *http.Request, *Context), c *Context) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, c)
	}
}
