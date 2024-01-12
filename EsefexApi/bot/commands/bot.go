package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

var BotCommand = &discordgo.ApplicationCommand{
	Name:        "bot",
	Description: "All commands related to the bot.",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "invite",
			Description: "Get the invite link for the bot.",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
		},
		{
			Name:        "ping",
			Description: "Get the ping of the bot.",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
		},
		{
			Name:        "stats",
			Description: "Get the stats of the bot.",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
		},
		{
			Name:        "join",
			Description: "Join the voice channel you are in.",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
		},
		{
			Name:        "leave",
			Description: "Leave the voice channel you are in.",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
		},
		{
			Name:        "config",
			Description: "All commands related to the configuration of the bot.",
			Type:        discordgo.ApplicationCommandOptionSubCommandGroup,
		},
	},
}

func (c *CommandHandlers) Bot(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	switch i.ApplicationCommandData().Options[0].Name {
	case "invite":
		return c.BotInvite(s, i)
	case "ping":
		return c.BotPing(s, i)
	case "stats":
		return c.BotStats(s, i)
	case "join":
		return c.BotJoin(s, i)
	case "leave":
		return c.BotLeave(s, i)
	case "config":
		return c.BotConfig(s, i)
	default:
		return nil, errors.Wrap(fmt.Errorf("Not implemented"), "Bot")
	}
}

func (c *CommandHandlers) BotInvite(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	permissions := 8
	inviteUrl := fmt.Sprintf("https://discord.com/api/oauth2/authorize?client_id=%s&permissions=%d&scope=bot%%20applications.commands", s.State.User.ID, permissions)

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Invite me to your server with this link: \n%s", inviteUrl),
		},
	}, nil
}

// TODO: Add a way to get the ping of the bot (currently not updating)
func (c *CommandHandlers) BotPing(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Pong! %dms", s.HeartbeatLatency().Milliseconds()),
		},
	}, nil
}

// TODO: Implement BotStats
func (c *CommandHandlers) BotStats(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	return nil, errors.Wrap(fmt.Errorf("Not implemented"), "BotStats")
}

// TODO: Implement BotJoin
func (c *CommandHandlers) BotJoin(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	return nil, errors.Wrap(fmt.Errorf("Not implemented"), "BotJoin")
}

// TODO: Implement BotLeave
func (c *CommandHandlers) BotLeave(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	return nil, errors.Wrap(fmt.Errorf("Not implemented"), "BotLeave")
}

// TODO: Implement BotConfig
func (c *CommandHandlers) BotConfig(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	return nil, errors.Wrap(fmt.Errorf("Not implemented"), "BotConfig")
}
