package commands

import (
	"crypto/sha256"
	"encoding/json"
	"esefexapi/bot/commands/cmdhandler"
	"esefexapi/bot/commands/middleware"
	"esefexapi/clientnotifiy"
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
	ds                *discordgo.Session
	dbs               *db.Databases
	domain            string
	permissionInteger int64
	mw                *middleware.CommandMiddleware
	cn                clientnotifiy.IClientNotifier
	Commands          map[string]*discordgo.ApplicationCommand
	Handlers          map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

func NewCommandHandlers(ds *discordgo.Session, dbs *db.Databases, domain string, cn clientnotifiy.IClientNotifier, permissionInteger int64) *CommandHandlers {
	c := &CommandHandlers{
		ds:                ds,
		dbs:               dbs,
		domain:            domain,
		permissionInteger: permissionInteger,
		mw:                middleware.NewCommandMiddleware(dbs),
		cn:                cn,
		Commands:          map[string]*discordgo.ApplicationCommand{},
		Handlers:          map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){},
	}

	c.Commands["bot"] = BotCommand
	c.Handlers["bot"] = WithErrorHandling(c.mw.CheckPerms(c.Bot, "Guild.UseSlashCommands"))

	c.Commands["help"] = HelpCommand
	c.Handlers["help"] = WithErrorHandling(c.mw.CheckPerms(c.Help, "Guild.UseSlashCommands"))

	c.Commands["permission"] = PermissionCommand
	c.Handlers["permission"] = WithErrorHandling(c.mw.CheckPerms(c.Permission, "Guild.UseSlashCommands", "Guild.ManageUser"))

	c.Commands["sound"] = SoundCommand
	c.Handlers["sound"] = WithErrorHandling(c.mw.CheckPerms(c.Sound, "Guild.UseSlashCommands"))

	c.Commands["user"] = UserCommand
	c.Handlers["user"] = WithErrorHandling(c.mw.CheckPerms(c.User, "Guild.UseSlashCommands"))

	c.Commands["webui"] = WebUICommand
	c.Handlers["webui"] = WithErrorHandling(c.mw.CheckPerms(c.WebUI, "Guild.UseSlashCommands"))

	return c
}

func WithErrorHandling(h cmdhandler.CommandHandlerWithErr) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
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

// ApplicationCommandsHash returns a hash of the application commands
// This is used to determine if the commands need to be updated
// The hash is based on the CommandHandlers.Commands field
func (c *CommandHandlers) ApplicationCommandsHash() (string, error) {
	data, err := json.Marshal(c.Commands)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash), nil
}
