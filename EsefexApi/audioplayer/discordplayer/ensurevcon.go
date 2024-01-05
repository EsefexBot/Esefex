package discordplayer

import (
	"esefexapi/audioplayer"
	"esefexapi/timer"
	"esefexapi/types"
	"esefexapi/util/dcgoutil"

	"github.com/pkg/errors"
)

func (c *DiscordPlayer) ensureVCon(guildID types.GuildID, userID types.UserID) (*VconData, error) {
	OusrChan, err := dcgoutil.UserGuildVC(c.ds, guildID, userID)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting user voice channel")
	} else if OusrChan.IsNone() {
		return nil, audioplayer.UserNotInVC
	}
	usrChan := OusrChan.Unwrap()

	timer.MessageElapsed("Got user's voice channel")

	ObotChan, err := dcgoutil.GetBotVC(c.ds, guildID)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting bot voice channel")
	}

	timer.MessageElapsed("Got bot's voice channel")

	botInGuild := ObotChan.IsSome()
	if botInGuild && ObotChan.Unwrap().ChannelID == usrChan.ChannelID {
		return c.vds[types.ChannelID(usrChan.ChannelID)], nil
	}

	timer.MessageElapsed("Checked if bot is in guild")

	vc, err := c.RegisterVcon(guildID, types.ChannelID(usrChan.ChannelID))
	if err != nil {
		return nil, errors.Wrap(err, "Error registering VCon")
	}

	timer.MessageElapsed("Registered VCon")

	return vc, nil
}
