package actions

import (
	"esefexbot/appcontext"
	"log"

	"github.com/bwmarrin/discordgo"
)

func ListenForApiRequests(s *discordgo.Session, c *appcontext.Context) {
	for {
		msg := <-c.Channels.PlaySound
		PlaySound(s, msg)
		log.Printf("Received message: %v", msg)
	}
}
