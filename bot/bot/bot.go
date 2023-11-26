package bot

import (
	"esefexbot/bot/commands"
	"log"

	"github.com/bwmarrin/discordgo"
)

func RegisterComands(s *discordgo.Session) {
	log.Println("Registering commands...")

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commands.Handlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	for _, v := range commands.Commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
		if err != nil {
			log.Printf("Cannot create '%v' command: %v", v.Name, err)
		}

		log.Printf("Registered '%v' command", v.Name)
	}
}

func DeleteAllCommands(s *discordgo.Session) {
	log.Println("Deleting all commands...")

	for _, g := range s.State.Guilds {
		DeleteGuildCommands(s, g.ID)
	}

	DeleteGuildCommands(s, "")

	log.Println("Deleted all commands")
}

func DeleteGuildCommands(s *discordgo.Session, guildID string) {
	cmds, err := s.ApplicationCommands(s.State.User.ID, guildID)
	if err != nil {
		log.Println(err)
	}

	for _, v := range cmds {
		s.ApplicationCommandDelete(s.State.User.ID, guildID, v.ID)
		log.Printf("Deleted '%v' command", v.Name)
	}
}
