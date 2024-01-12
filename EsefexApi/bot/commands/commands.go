package commands

import (
	"esefexapi/db"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

type Command struct {
	ApplicationCommand discordgo.ApplicationCommand
	Handler            func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

type SubcommandGroup struct {
	Name        string
	Description string
	Commands    []*Command
}

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

	ch.Commands["bot"] = BotCommand
	ch.Handlers["bot"] = WithErrorHandling(ch.Bot)

	ch.Commands["help"] = HelpCommand
	ch.Handlers["help"] = WithErrorHandling(ch.Help)

	ch.Commands["permission"] = PermissionCommand
	ch.Handlers["permission"] = WithErrorHandling(ch.Permission)

	ch.Commands["sound"] = SoundCommand
	ch.Handlers["sound"] = WithErrorHandling(ch.Sound)

	ch.Commands["user"] = UserCommand
	ch.Handlers["user"] = WithErrorHandling(ch.User)

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

func OptionsMap(options []*discordgo.ApplicationCommandInteractionDataOption) map[string]*discordgo.ApplicationCommandInteractionDataOption {
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	return optionMap
}
