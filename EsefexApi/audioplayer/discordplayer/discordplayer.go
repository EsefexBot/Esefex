package discordplayer

import (
	"esefexapi/audioplayer/discordplayer/vcon"
	"esefexapi/db"

	"esefexapi/audioplayer"
	"esefexapi/service"

	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	useTimeouts    = false
	sessionTimeout = 5 * time.Minute
)

var _ service.IService = &DiscordPlayer{}
var _ audioplayer.IAudioPlayer = &DiscordPlayer{}

// DiscordPlayer implements PlaybackManager
type DiscordPlayer struct {
	vcs   map[string]*vcon.VCon
	ds    *discordgo.Session
	dbs   db.Databases
	stop  chan struct{}
	ready chan struct{}
}

func NewDiscordPlayer(ds *discordgo.Session, dbs db.Databases) *DiscordPlayer {
	return &DiscordPlayer{
		vcs:   make(map[string]*vcon.VCon),
		ds:    ds,
		dbs:   dbs,
		stop:  make(chan struct{}),
		ready: make(chan struct{}),
	}
}
