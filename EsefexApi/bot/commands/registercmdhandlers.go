package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

// call this before opening the session
func (c *CommandHandlers) RegisterApplicationComandHandlers() {
	log.Println("Registering command handlers...")
	c.ds.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionApplicationCommand {
			return
		}

		if h, ok := c.Handlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}
