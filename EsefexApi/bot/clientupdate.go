package bot

import (
	"esefexapi/types"
	"log"

	"github.com/bwmarrin/discordgo"
)

func (b *DiscordBot) RegisterClientUpdateHandlers() {
	b.ds.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		err := b.cn.UpdateNotificationUsers()
		if err != nil {
			log.Printf("Error notifying clients: %+v", err)
		}
	})

	b.ds.AddHandler(func(s *discordgo.Session, r *discordgo.VoiceStateUpdate) {
		userID := types.UserID(r.UserID)

		err := b.cn.UpdateNotificationUsers(userID)
		if err != nil {
			log.Printf("Error notifying clients: %+v", err)
		}
	})

}
