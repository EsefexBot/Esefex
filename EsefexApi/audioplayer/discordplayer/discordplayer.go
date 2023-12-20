package discordplayer

import (
	"esefexapi/audioplayer"
	"esefexapi/sounddb"

	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	useTimeouts    = false
	sessionTimeout = 5 * time.Minute
)

var _ audioplayer.IAudioPlayer = &DiscordPlayer{}

// DiscordPlayer implements PlaybackManager
type DiscordPlayer struct {
	s  map[string]*SfxPlayer
	ds *discordgo.Session
	db sounddb.ISoundDB
}

func NewDiscordPlayer(ds *discordgo.Session, db sounddb.ISoundDB) *DiscordPlayer {
	return &DiscordPlayer{
		s:  make(map[string]*SfxPlayer),
		ds: ds,
		db: db,
	}
}

func (c *DiscordPlayer) PlaySound(uid sounddb.SoundUID, userID string) error {
	panic("not implemented")
	return nil
}
