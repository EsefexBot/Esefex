package bot

import (
	"errors"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

func (b *DiscordBot) RegisterComands(s *discordgo.Session) {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := b.cmdh.Handlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	for _, v := range b.cmdh.Commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
		if err != nil {
			log.Printf("Cannot create '%v' command: %v", v.Name, err)
		}

		log.Printf("Registered '%v' command", v.Name)
	}
}

func (b *DiscordBot) DeleteAllCommands(s *discordgo.Session) {
	log.Println("Deleting all commands...")

	for _, g := range s.State.Guilds {
		b.DeleteGuildCommands(s, g.ID)
	}

	b.DeleteGuildCommands(s, "")

	log.Println("Deleted all commands")
}

func (b *DiscordBot) DeleteGuildCommands(s *discordgo.Session, guildID string) {
	cmds, err := s.ApplicationCommands(s.State.User.ID, guildID)
	if err != nil {
		log.Println(err)
	}

	for _, v := range cmds {
		s.ApplicationCommandDelete(s.State.User.ID, guildID, v.ID)
		log.Printf("Deleted '%v' command", v.Name)
	}
}

func CreateSession() (*discordgo.Session, error) {
	token := os.Getenv("BOT_TOKEN")

	if token == "" {
		return nil, errors.New("BOT_TOKEN is not set")
	}

	s, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	return s, nil
}
