package discordplayer

import (
	"esefexapi/audioplayer"
	"esefexapi/sounddb"
	"esefexapi/util/dcgoutil"
	"log"

	"github.com/pkg/errors"
)

func (c *DiscordPlayer) PlaySound(soundID string, userID string) error {
	log.Printf("Playing sound '%v' for user '%v'", soundID, userID)

	OuserVc, err := dcgoutil.UserVCAny(c.ds, userID)
	if err != nil {
		return errors.Wrap(err, "Error getting user's voice channel")
	} else if OuserVc.IsNone() {
		return audioplayer.UserNotInVC
	}
	userVC := OuserVc.Unwrap()

	vd, err := c.ensureVCon(userVC.GuildID, userID)
	if err != nil {
		return errors.Wrap(err, "Error ensuring voice connection")
	}

	vd.vcon.PlaySound(sounddb.SuidFromStrings(userVC.GuildID, soundID))
	vd.AfkTimeoutIn = vd.AfkTimeoutIn.Add(c.timeout)

	return nil
}
