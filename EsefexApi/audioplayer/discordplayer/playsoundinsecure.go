package discordplayer

import (
	"esefexapi/sounddb"
	"log"
)

func (c *DiscordPlayer) PlaySoundInsecure(uid sounddb.SoundUID, serverID, userID string) error {
	log.Printf("Playing sound %s\n", uid)

	vc, err := c.ensureVCon(serverID, userID)
	if err != nil {
		return err
	}

	vc.PlaySound(uid)

	return nil
}
