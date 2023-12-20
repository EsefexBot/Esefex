package actions

// import (
// 	"esefexapi/util"
// 	"log"

// 	"github.com/bwmarrin/discordgo"
// )

// func PlaySound(s *discordgo.Session, ps msg.PlaySound) {
// 	g, err := s.State.Guild(ps.GuildID)
// 	if err != nil {
// 		log.Printf("Error getting guild: %v", err)
// 		return
// 	}

// 	log.Println(g.VoiceStates)

// 	vs, err := s.State.VoiceState(ps.GuildID, s.State.User.ID)
// 	if err != nil {
// 		log.Printf("Error getting voice state: %v", err)
// 	}

// 	if vs != nil {
// 		go PlayDebugSound(s, ps.GuildID, vs.ChannelID)
// 	}

// 	log.Println(vs)
// }

// func PlayDebugSound(s *discordgo.Session, gID string, cID string) {
// 	// sound, err := util.LoadDcaSound("debug.dca")
// 	// if err != nil {
// 	// 	log.Printf("Error loading sound: %v", err)
// 	// }

// 	sound := util.HallucinateDcaData(100, 150)

// 	vc, err := s.ChannelVoiceJoin(gID, cID, false, false)
// 	if err != nil {
// 		log.Printf("Error joining voice channel: %v", err)
// 	}

// 	vc.Speaking(true)
// 	defer vc.Speaking(false)

// 	log.Printf("Sending %v bytes", len(sound))

// 	for _, buff := range sound {
// 		vc.OpusSend <- buff
// 	}
// }
