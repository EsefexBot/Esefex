package discordplayer

import (
	"esefexapi/sounddb"
	"esefexapi/util/dcgoutil"
	"log"

	"github.com/pkg/errors"
)

func (c *DiscordPlayer) PlaySound(soundID string, userID string) error {
	log.Printf("Playing sound '%v' for user '%v'", soundID, userID)

	// usr, err := c.ds.User(userID)
	// if err != nil {
	// 	return err
	// }

	uservc, err := dcgoutil.UserVC(c.ds, userID)
	if err != nil {
		return errors.Wrap(err, "Error getting user's voice channel")
	}

	vc, err := c.ensureVCon(uservc.GuildID, userID)
	if err != nil {
		return errors.Wrap(err, "Error ensuring voice connection")
	}

	vc.PlaySound(sounddb.SuidFromStrings(uservc.GuildID, soundID))

	return nil
}
