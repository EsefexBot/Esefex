package discordgo

import "github.com/bwmarrin/discordgo"

// util functions for discordgo

func BotInVC(ds *discordgo.Session, serverID, channelID string) bool {
	vs, err := ds.State.VoiceState(serverID, ds.State.User.ID)
	if err != nil {
		return false
	}
	if vs == nil {
		return false
	}
	return vs.ChannelID == channelID
}

func GetBotVC(ds *discordgo.Session, serverID string) (*discordgo.VoiceState, error) {
	return ds.State.VoiceState(serverID, ds.State.User.ID)
}
