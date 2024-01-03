package discordplayer

import (
	"esefexapi/audioplayer"
	"esefexapi/timer"
	"esefexapi/util/dcgoutil"

	"github.com/pkg/errors"
)

func (c *DiscordPlayer) ensureVCon(serverID, userID string) (*VconData, error) {
	OusrChan, err := dcgoutil.UserServerVC(c.ds, serverID, userID)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting user voice channel")
	} else if OusrChan.IsNone() {
		return nil, audioplayer.UserNotInVC
	}
	usrChan := OusrChan.Unwrap()

	timer.MessageElapsed("Got user's voice channel")

	ObotChan, err := dcgoutil.GetBotVC(c.ds, serverID)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting bot voice channel")
	}

	timer.MessageElapsed("Got bot's voice channel")

	botInGuild := ObotChan.IsSome()
	if botInGuild && ObotChan.Unwrap().ChannelID == usrChan.ChannelID {
		return c.vds[ChannelID(usrChan.ChannelID)], nil
	}

	timer.MessageElapsed("Checked if bot is in guild")

	vc, err := c.RegisterVcon(serverID, usrChan.ChannelID)
	if err != nil {
		return nil, errors.Wrap(err, "Error registering VCon")
	}

	timer.MessageElapsed("Registered VCon")

	return vc, nil
}
