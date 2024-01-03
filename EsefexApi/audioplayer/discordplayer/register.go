package discordplayer

import (
	"esefexapi/audioplayer/discordplayer/vcon"
	"esefexapi/timer"

	"github.com/pkg/errors"
)

func (c *DiscordPlayer) RegisterVcon(serverID string, channelID string) (*VconData, error) {
	vc, err := vcon.NewVCon(c.ds, c.dbs.SoundDB, serverID, channelID)
	if err != nil {
		return nil, errors.Wrap(err, "Error creating new VCon")
	}

	timer.MessageElapsed("Created new VCon")

	vd := &VconData{
		ChannelID: channelID,
		ServerID:  serverID,
		vcon:      vc,
	}

	c.vds[ChannelID(channelID)] = vd
	go vc.Run()

	return vd, nil
}

var VconNotFound = errors.New("VCon not found")

func (c *DiscordPlayer) UnregisterVcon(channelID string) error {
	vd, ok := c.vds[ChannelID(channelID)]
	if !ok {
		return VconNotFound
	}

	delete(c.vds, ChannelID(channelID))
	vd.vcon.Close()

	return nil
}
