package commands

import (
	"esefexapi/db"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

type CommandHandlers struct {
	dbs      *db.Databases
	domain   string
	Commands map[string]*discordgo.ApplicationCommand
	Handlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

func NewCommandHandlers(dbs *db.Databases, domain string) *CommandHandlers {
	ch := &CommandHandlers{
		dbs:      dbs,
		domain:   domain,
		Commands: map[string]*discordgo.ApplicationCommand{},
		Handlers: map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){},
	}

	ch.Commands["upload"] = UploadCommand
	ch.Handlers["upload"] = WithErrorHandling(ch.Upload)

	ch.Commands["list"] = ListCommand
	ch.Handlers["list"] = WithErrorHandling(ch.List)

	ch.Commands["delete"] = DeleteCommand
	ch.Handlers["delete"] = WithErrorHandling(ch.Delete)

	ch.Commands["link"] = LinkCommand
	ch.Handlers["link"] = WithErrorHandling(ch.Link)

	ch.Commands["unlink"] = UnlinkCommand
	ch.Handlers["unlink"] = WithErrorHandling(ch.Unlink)

	return ch
}

func WithErrorHandling(h func(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error)) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		r, err := h(s, i)
		if err != nil {
			log.Printf("Cannot execute command: %+v", err)

			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("An error has occurred while executing the command: \n```%+v```", errors.Cause(err)),
				},
			})
			if err != nil {
				log.Printf("Cannot respond to interaction: %+v", err)
			}
		}

		if r != nil {
			err = s.InteractionRespond(i.Interaction, r)
			if err != nil {
				log.Printf("Cannot respond to interaction: %+v", err)
			}
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
