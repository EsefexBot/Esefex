package discordplayer

import (
	"esefexapi/util/dcgoutil"
	"github.com/pkg/errors"
)

var BotNotInVC = errors.New("Bot is not in a voice channel")
var UserNotInVC = errors.New("User is not in a voice channel")

func (c *DiscordPlayer) ensureVCon(serverID, userID string) (*VconData, error) {
	OusrChan, err := dcgoutil.UserServerVC(c.ds, serverID, userID)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting user voice channel")
	} else if OusrChan.IsNone() {
		return nil, UserNotInVC
	}
	usrChan := OusrChan.Unwrap()

	ObotChan, err := dcgoutil.GetBotVC(c.ds, serverID)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting bot voice channel")
	}

	botInGuild := ObotChan.IsSome()
	if botInGuild && ObotChan.Unwrap().ChannelID == usrChan.ChannelID {
		return c.vds[ChannelID(usrChan.ChannelID)], nil
	}

	vc, err := c.RegisterVcon(serverID, usrChan.ChannelID)
	if err != nil {
		return nil, errors.Wrap(err, "Error registering VCon")
	}

	return vc, nil
}
