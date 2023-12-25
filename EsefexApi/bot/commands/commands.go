package commands

import (
	"esefexapi/sounddb"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

type CommandHandlers struct {
	db       sounddb.ISoundDB
	domain   string
	Commands map[string]*discordgo.ApplicationCommand
	Handlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

func NewCommandHandlers(db sounddb.ISoundDB, domain string) *CommandHandlers {
	ch := &CommandHandlers{
		db:       db,
		domain:   domain,
		Commands: map[string]*discordgo.ApplicationCommand{},
		Handlers: map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){},
	}

	ch.Commands["upload"] = UploadCommand
	ch.Handlers["upload"] = WithErrorHandling(ch.Upload)

	ch.Commands["session"] = SessionCommand
	ch.Handlers["session"] = WithErrorHandling(ch.Session)

	ch.Commands["list"] = ListCommand
	ch.Handlers["list"] = WithErrorHandling(ch.List)

	ch.Commands["delete"] = DeleteCommand
	ch.Handlers["delete"] = WithErrorHandling(ch.Delete)

	return ch
}

func WithErrorHandling(h func(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error)) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		r, err := h(s, i)
		if err != nil {
			log.Printf("Cannot execute command: %v", err)

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("An error has occurred while executing the command: %v", err),
				},
			})
		}

		if r != nil {
			s.InteractionRespond(i.Interaction, r)
		}
	}
}

func OptionsMap(i *discordgo.InteractionCreate) map[string]*discordgo.ApplicationCommandInteractionDataOption {
	options := i.ApplicationCommandData().Options

	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	return optionMap
}
