package discordplayer

import (
	"esefexapi/audioplayer"
	"esefexapi/sounddb"
	"esefexapi/timer"
	"esefexapi/types"
	"esefexapi/util/dcgoutil"
	"log"

	"github.com/pkg/errors"
)

func (c *DiscordPlayer) PlaySound(soundID types.SoundID, userID types.UserID) error {
	log.Printf("Playing sound '%v' for user '%v'", soundID, userID)

	OuserVc, err := dcgoutil.UserVCAny(c.ds, userID)
	if err != nil {
		return errors.Wrap(err, "Error getting user's voice channel")
	} else if OuserVc.IsNone() {
		return audioplayer.UserNotInVC
	}
	userVC := OuserVc.Unwrap()

	timer.MessageElapsed("Got user's voice channel")

	vd, err := c.ensureVCon(types.GuildID(userVC.GuildID), userID)
	if err != nil {
		return errors.Wrap(err, "Error ensuring voice connection")
	}

	timer.MessageElapsed("Got voice connection")

	vd.vcon.PlaySound(sounddb.SuidFromStrings(userVC.GuildID, soundID.String()))
	vd.AfkTimeoutIn = vd.AfkTimeoutIn.Add(c.timeout)

	return nil
}
