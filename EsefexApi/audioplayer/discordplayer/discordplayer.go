package discordplayer

import (
	"esefexapi/audioplayer/discordplayer/vcon"
	"esefexapi/db"
	"esefexapi/types"
	"esefexapi/util/dcgoutil"
	"log"

	"esefexapi/audioplayer"
	"esefexapi/service"

	"time"

	"github.com/bwmarrin/discordgo"
)

var _ service.IService = &DiscordPlayer{}
var _ audioplayer.IAudioPlayer = &DiscordPlayer{}

// DiscordPlayer implements PlaybackManager
type DiscordPlayer struct {
	vds         map[types.ChannelID]*VconData
	ds          *discordgo.Session
	dbs         db.Databases
	stop        chan struct{}
	ready       chan struct{}
	useTimeouts bool
	timeout     time.Duration
}

type VconData struct {
	ChannelID    types.ChannelID
	GuildID      types.GuildID
	AfkTimeoutIn time.Time
	vcon         *vcon.VCon
}

func NewDiscordPlayer(ds *discordgo.Session, dbs *db.Databases, useTimeouts bool, timeout time.Duration) *DiscordPlayer {
	dp := &DiscordPlayer{
		vds:         make(map[types.ChannelID]*VconData),
		ds:          ds,
		dbs:         *dbs,
		stop:        make(chan struct{}),
		ready:       make(chan struct{}),
		useTimeouts: useTimeouts,
		timeout:     timeout,
	}

	ds.AddHandler(func(s *discordgo.Session, e *discordgo.VoiceStateUpdate) {
		// check if previous state has a vcon associated with it and close it, make sure that it is not closed twice
		if e.BeforeUpdate == nil || e.ChannelID != "" || e.UserID != s.State.User.ID {
			return
		}

		if _, ok := dp.vds[types.ChannelID(e.BeforeUpdate.ChannelID)]; ok {
			log.Printf("Closing VCon: %s", e.BeforeUpdate.ChannelID)
			err := dp.UnregisterVcon(types.ChannelID(e.BeforeUpdate.ChannelID))
			if err != nil {
				log.Printf("Error unregistering vcon: %+v", err)
			}
		}
	})

	// check if there are still users in a channel that the bot is in, if not close the vcon
	ds.AddHandler(func(s *discordgo.Session, e *discordgo.VoiceStateUpdate) {
		if e.BeforeUpdate == nil {
			return
		}

		if _, ok := dp.vds[types.ChannelID(e.BeforeUpdate.ChannelID)]; !ok {
			return
		}

		users, err := dcgoutil.GetVCUsers(s, e.GuildID, e.BeforeUpdate.ChannelID)
		if err != nil {
			log.Printf("Error getting users in channel: %s", err)
			return
		}

		// log.Printf("Users in channel: %d", len(users))

		if len(users) == 1 {
			log.Printf("Channel empty, closing vcon: %s", e.BeforeUpdate.ChannelID)
			err := dp.UnregisterVcon(types.ChannelID(e.BeforeUpdate.ChannelID))
			if err != nil {
				log.Printf("Error unregistering vcon: %+v", err)
			}
		}
	})

	return dp
}
