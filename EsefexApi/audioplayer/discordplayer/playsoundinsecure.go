package discordplayer

import (
	"esefexapi/sounddb"
	"esefexapi/types"
	"log"

	"github.com/pkg/errors"
)

func (c *DiscordPlayer) PlaySoundInsecure(uid sounddb.SoundUID, guildID types.GuildID, userID types.UserID) error {
	log.Printf("Playing sound %s\n", uid)

	vd, err := c.ensureVCon(guildID, userID)
	if err != nil {
		return errors.Wrap(err, "Error ensuring voice connection")
	}

	vd.vcon.PlaySound(uid)

	return nil
}
