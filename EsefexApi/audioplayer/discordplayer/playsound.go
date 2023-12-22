package discordplayer

import (
	"esefexapi/audioplayer/discordplayer/vcon"
	"esefexapi/sounddb"
	"esefexapi/util/dcgoutil"
	"log"
)

func (c *DiscordPlayer) PlaySound(uid sounddb.SoundUID, serverID, userID string) error {
	// check if bot is joined to server but in a different channel than the user
	// if so, move bot to user's channel

	// if the bot is not in the server, join the user's channel

	// if the bot is in the server and in the user's channel, play the sound

	// if the sound sound played reset the auto-disconnect (afk) timer

	log.Printf("Playing sound %s\n", uid)

	vc, err := c.ensureVCon(serverID, userID)
	if err != nil {
		return err
	}

	vc.PlaySound(uid)

	return nil
}

func (c *DiscordPlayer) ensureVCon(serverID, userID string) (*vcon.VCon, error) {
	usrChanID, err := dcgoutil.UserVC(c.ds, serverID, userID)
	if err != nil {
		return nil, err
	}

	botChan, err := dcgoutil.GetBotVC(c.ds, serverID)
	if err != nil {
		return nil, err
	}

	// if the bot is in the server and in the user's channel, return the VCon for that channel (if it exists) or create a new one

	if usrChanID == botChan.ChannelID && c.vcs[usrChanID] != nil {
		return c.vcs[usrChanID], nil
	}

	if usrChanID == botChan.ChannelID && c.vcs[usrChanID] == nil {
		vc, err := vcon.NewVCon(c.ds, c.db, serverID, usrChanID)
		if err != nil {
			return nil, err
		}

		c.vcs[usrChanID] = vc
		vc.Run()

		return vc, nil
	}

	// if the bot is not in the server, join the user's channel by creating a new VCon

	if botChan == nil {
		vc, err := vcon.NewVCon(c.ds, c.db, serverID, usrChanID)
		if err != nil {
			return nil, err
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

	vc, err := vcon.NewVCon(c.ds, c.db, serverID, usrChanID)
	if err != nil {
		return nil, err
	}

	c.vcs[usrChanID] = vc
	vc.Run()

	return vc, nil
}
