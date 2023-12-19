package actions

import (
	"esefexapi/ctx"
	"log"

	"github.com/bwmarrin/discordgo"
)

func ListenForApiRequests(s *discordgo.Session, c *ctx.Ctx) {
	for {
		msg := <-c.Channels.PlaySound
		PlaySound(s, msg)
		log.Printf("Received message: %v", msg)
	}
}
