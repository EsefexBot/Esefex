package discordplayer

import (
	"esefexapi/sounddb"
	"esefexapi/util/dcgoutil"
	"log"
)

func (c *DiscordPlayer) PlaySound(soundID string, userID string) error {
	log.Printf("Playing sound '%v' for user '%v'", soundID, userID)

	// usr, err := c.ds.User(userID)
	// if err != nil {
	// 	return err
	// }

	uservc, err := dcgoutil.UserVC(c.ds, userID)
	if err != nil {
		return err
	}

	vc, err := c.ensureVCon(uservc.GuildID, userID)
	if err != nil {
		return err
	}

	vc.PlaySound(sounddb.SuidFromStrings(uservc.GuildID, soundID))

	return nil
}
