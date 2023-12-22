package commands

import (
	"esefexapi/sounddb"

	"github.com/bwmarrin/discordgo"
)

type CommandHandlers struct {
	db       sounddb.ISoundDB
	Commands map[string]*discordgo.ApplicationCommand
	Handlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

func NewCommandHandlers(db sounddb.ISoundDB) *CommandHandlers {
	ch := &CommandHandlers{
		db:       db,
		Commands: map[string]*discordgo.ApplicationCommand{},
		Handlers: map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){},
	}

	ch.Commands["upload"] = UploadCommand
	ch.Handlers["upload"] = ch.Upload

	ch.Commands["session"] = SessionCommand
	ch.Handlers["session"] = ch.Session

	return ch
}
