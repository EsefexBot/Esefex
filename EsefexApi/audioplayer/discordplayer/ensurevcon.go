package discordplayer

import (
	"esefexapi/audioplayer/discordplayer/vcon"
	"esefexapi/util/dcgoutil"

	"github.com/pkg/errors"
)

func (c *DiscordPlayer) ensureVCon(serverID, userID string) (*vcon.VCon, error) {
	usrChanID, err := dcgoutil.UserServerVC(c.ds, serverID, userID)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting user voice channel")
	}

	botChan, err := dcgoutil.GetBotVC(c.ds, serverID)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting bot voice channel")
	}

	// if the bot is in the server and in the user's channel, return the VCon for that channel (if it exists) or create a new one

	if usrChanID == botChan.ChannelID && c.vcs[usrChanID] != nil {
		return c.vcs[usrChanID], nil
	}

	if usrChanID == botChan.ChannelID && c.vcs[usrChanID] == nil {
		vc, err := vcon.NewVCon(c.ds, c.dbs.SoundDB, serverID, usrChanID)
		if err != nil {
			return nil, errors.Wrap(err, "Error creating new VCon")
		}

		c.vcs[usrChanID] = vc
		vc.Run()

		return vc, nil
	}

	// if the bot is not in the server, join the user's channel by creating a new VCon

	if botChan == nil {
		vc, err := vcon.NewVCon(c.ds, c.dbs.SoundDB, serverID, usrChanID)
		if err != nil {
			return nil, errors.Wrap(err, "Error creating new VCon")
		}

		c.vcs[usrChanID] = vc
		vc.Run()

		return vc, nil
	}

	// if the bot is in the server but not in the user's channel, delete the VCon for the bot's current channel and create a new one for the user's channel

	if c.vcs[botChan.ChannelID] != nil {
		c.vcs[botChan.ChannelID].Close()
		delete(c.vcs, botChan.ChannelID)
	}

	vc, err := vcon.NewVCon(c.ds, c.dbs.SoundDB, serverID, usrChanID)
	if err != nil {
		return nil, errors.Wrap(err, "Error creating new VCon")
	}

	c.vcs[usrChanID] = vc
	vc.Run()

	return vc, nil
}
