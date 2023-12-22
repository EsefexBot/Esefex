package discordplayer

import (
	"esefexapi/audioplayer/discordplayer/vcon"

	"esefexapi/audioplayer"
	"esefexapi/service"
	"esefexapi/sounddb"

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
	db    sounddb.ISoundDB
	stop  chan struct{}
	ready chan struct{}
}

func NewDiscordPlayer(ds *discordgo.Session, db sounddb.ISoundDB) *DiscordPlayer {
	return &DiscordPlayer{
		vcs:   make(map[string]*vcon.VCon),
		ds:    ds,
		db:    db,
		stop:  make(chan struct{}),
		ready: make(chan struct{}),
	}
}
