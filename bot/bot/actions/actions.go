package actions

import (
	// "log"

	"github.com/bwmarrin/discordgo"
)

var (
	channelID string = "777344211828604950"
)

func UnprovokedMessage(s *discordgo.Session) {
	s.ChannelMessageSend(channelID, "Hello, world! This is a message invoked by a channel message event.")
}

func JoinChannelVoice(s *discordgo.Session, gID string, cID string) {
	s.ChannelVoiceJoin(gID, cID, false, false)
}

// func LeaveChannelVoice(s *discordgo.Session, gID string) {
// 	// //get guild
// 	// g, err := s.State.Guild(i.GuildID)
// 	// if err != nil {
// 	// 	log.Printf("Cannot get guild: %v", err)
// 	// }

// 	// // get voice connection
// 	// vc, err := s.State.Vo
// 	log.Printf("Leaving voice channel in guild %v", gID)
// 	log.Println("Unfortunately this is not implemented yet. Sucks to be you :P")
// }

// func LeaveAllChannels(s *discordgo.Session) {
// 	for _, g := range s.State.Guilds {
// 		LeaveChannelVoice(s, g.ID)
// 	}
// }

func PlaySound(s *discordgo.Session, gID string, cID string) {

}
