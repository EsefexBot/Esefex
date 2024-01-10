package discordplayer

import (
	"esefexapi/audioplayer/discordplayer/vcon"
	"esefexapi/timer"
	"esefexapi/types"
	"fmt"

	"github.com/pkg/errors"
)

func (c *DiscordPlayer) RegisterVcon(guildID types.GuildID, channelID types.ChannelID) (*VconData, error) {
	vc, err := vcon.NewVCon(c.ds, c.dbs.SoundDB, guildID, channelID)
	if err != nil {
		return nil, errors.Wrap(err, "Error creating new VCon")
	}

	timer.MessageElapsed("Created new VCon")

	vd := &VconData{
		ChannelID: channelID,
		GuildID:   guildID,
		vcon:      vc,
	}

	c.vds[types.ChannelID(channelID)] = vd
	go vc.Run()

	return vd, nil
}

var VconNotFound = fmt.Errorf("VCon not found")

func (c *DiscordPlayer) UnregisterVcon(channelID types.ChannelID) error {
	vd, ok := c.vds[types.ChannelID(channelID)]
	if !ok {
		return VconNotFound
	}

	delete(c.vds, types.ChannelID(channelID))
	vd.vcon.Close()

	return nil
}
