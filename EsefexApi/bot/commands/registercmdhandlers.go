package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

// call this before opening the session
func (c *CommandHandlers) RegisterComandHandlers() {
	log.Println("Registering command handlers...")
	c.ds.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := c.Handlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}
