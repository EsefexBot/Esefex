package commands

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strconv"

	"github.com/bwmarrin/discordgo"
	externalip "github.com/glendc/go-external-ip"
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
		{
			Name:        "debug",
			Description: "Prints debug information about the bot.",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
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
	case "debug":
		return c.BotDebug(s, i)
	default:
		return nil, errors.Wrap(fmt.Errorf("Unknown subcommand %s", i.ApplicationCommandData().Options[0].Name), "Error handling user command")
	}
}

func (c *CommandHandlers) BotInvite(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	inviteUrl := fmt.Sprintf("https://discord.com/api/oauth2/authorize?client_id=%s&permissions=%d&scope=bot%%20applications.commands", s.State.User.ID, c.permissionInteger)

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: "Invite to your server",
					URL:   inviteUrl,
				},
			},
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Style: discordgo.LinkButton,
							Label: "Invite",
							URL:   inviteUrl,
							// TODO: For some reason, you need to set the emoji to something, otherwise the request will fail
							// This is a bug in the discordgo library
							// I should probably make a PR to fix this
							Emoji: discordgo.ComponentEmoji{
								Name: "🤖",
							},
						},
					},
				},
			},
		},
	}, nil
}

// TODO: Add a way to get the ping of the bot (currently not updating)
func (c *CommandHandlers) BotPing(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: fmt.Sprintf("Pong! `%dms`", s.HeartbeatLatency().Milliseconds()),
				},
			},
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

func (c *CommandHandlers) BotDebug(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	resp := "Debug information:\n```js\n"

	resp += fmt.Sprintf("Domain: %s\n", c.domain)
	resp += fmt.Sprintf("Guilds: %d\n", len(s.State.Guilds))

	// log local ipLoc
	ipLoc, err := net.InterfaceAddrs()
	if err != nil {
		return nil, errors.Wrap(err, "Error getting local ip")
	}
	resp += fmt.Sprintf("Local IP: %s\n", ipLoc)

	// log public ip
	ipPub, err := externalip.DefaultConsensus(nil, nil).ExternalIP()
	if err != nil {
		return nil, errors.Wrap(err, "Error getting public ip")
	}
	resp += fmt.Sprintf("Public IP: %s\n", ipPub)

	// log PID
	resp += fmt.Sprintf("PID: %d\n", os.Getpid())

	// log CommandHandlersHash
	hash, err := c.ApplicationCommandsHash()
	if err != nil {
		return nil, errors.Wrap(err, "Error getting CommandHandlersHash")
	}
	resp += fmt.Sprintf("CommandHandlersHash: %s\n", hash)

	// bot ping
	resp += fmt.Sprintf("Bot ping: %dms\n", s.HeartbeatLatency().Milliseconds())

	// get ps output
	// using the following command:
	// ps -p <pid> -o uid,pid,ppid,c,systime,tty,time,stat,euid,ruid,tpgid,sess,pgrp,pcpu,comm,cmd
	out, err := exec.Command("ps", "-p", strconv.Itoa(os.Getpid()), "-o", "uid,pid,ppid,pcpu,c,time,pmem,nlwp,tname,stat,lstart,cmd").Output()
	if err != nil {
		return nil, errors.Wrap(err, "Error getting info of current process")
	}
	resp += fmt.Sprintf("Process info:\n%s\n", out)

	resp += "```"

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: resp,
		},
	}, nil
}
