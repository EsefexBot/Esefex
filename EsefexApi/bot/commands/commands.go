package commands

import (
	"crypto/sha256"
	"encoding/json"
	"esefexapi/bot/commands/middleware"
	"esefexapi/clientnotifiy"
	"esefexapi/config"
	"esefexapi/db"
	"fmt"

	"github.com/bwmarrin/discordgo"
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

func NewCommandHandlers(ds *discordgo.Session, dbs *db.Databases, domain string, cn clientnotifiy.IClientNotifier) *CommandHandlers {
	c := &CommandHandlers{
		ds:                ds,
		dbs:               dbs,
		domain:            domain,
		permissionInteger: config.Get().Bot.PermissionsInteger,
		mw:                middleware.NewCommandMiddleware(dbs),
		cn:                cn,
		Commands:          map[string]*discordgo.ApplicationCommand{},
		Handlers:          map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){},
	}

	eh := c.mw.WithErrorHandling
	rdm := c.mw.RejectDMs
	cp := c.mw.CheckPerms

	c.Commands["bot"] = BotCommand
	c.Handlers["bot"] = eh(rdm(cp(c.Bot, "Guild.UseSlashCommands")))

	c.Commands["help"] = HelpCommand
	c.Handlers["help"] = eh(rdm(cp(c.Help, "Guild.UseSlashCommands")))

	c.Commands["permission"] = PermissionCommand
	c.Handlers["permission"] = eh(rdm(cp(c.Permission, "Guild.UseSlashCommands", "Guild.ManageUser")))

	c.Commands["sound"] = SoundCommand
	c.Handlers["sound"] = eh(rdm(cp(c.Sound, "Guild.UseSlashCommands")))

	c.Commands["user"] = UserCommand
	c.Handlers["user"] = eh(rdm(cp(c.User, "Guild.UseSlashCommands")))

	c.Commands["webui"] = WebUICommand
	c.Handlers["webui"] = eh(rdm(cp(c.WebUI, "Guild.UseSlashCommands")))

	return c
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
