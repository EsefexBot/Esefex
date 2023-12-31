package discordplayer

import (
	"esefexapi/sounddb"
	"log"

	"github.com/pkg/errors"
)

func (c *DiscordPlayer) PlaySoundInsecure(uid sounddb.SoundUID, serverID, userID string) error {
	log.Printf("Playing sound %s\n", uid)

	vc, err := c.ensureVCon(serverID, userID)
	if err != nil {
		return errors.Wrap(err, "Error ensuring voice connection")
	}

	vc.PlaySound(uid)

	return nil
}
