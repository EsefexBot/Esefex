package actions

import "github.com/bwmarrin/discordgo"

var (
	channelID string = "777344211828604950"
)

func UnprovokedMessage(s *discordgo.Session) {
	s.ChannelMessageSend(channelID, "Hello, world! This is a message invoked by a channel message event.")
}
