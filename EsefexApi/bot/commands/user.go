package commands

import (
	"esefexapi/types"
	"esefexapi/userdb"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

var UserCommand = &discordgo.ApplicationCommand{
	Name:        "user",
	Description: "All commands related to users.",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "stats",
			Description: "Get the stats of a user.",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "user",
					Description: "The user to get the stats of. (Leave empty to get your own stats)",
					Type:        discordgo.ApplicationCommandOptionUser,
					Required:    false,
				},
			},
		},
		{
			Name:        "link",
			Description: "Link your Discord account to Esefex",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
		},
		{
			Name:        "unlink",
			Description: "Unlink your Discord account from Esefex. Useful if you think your account has been compromised.",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
		},
	},
}

func (c *CommandHandlers) User(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	switch i.ApplicationCommandData().Options[0].Name {
	case "stats":
		return c.UserStats(s, i)
	case "link":
		return c.UserLink(s, i)
	case "unlink":
		return c.UserUnlink(s, i)
	default:
		return nil, errors.Wrap(fmt.Errorf("Unknown subcommand %s", i.ApplicationCommandData().Options[0].Name), "Error handling user command")
	}
}

func (c *CommandHandlers) UserStats(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	// TODO: index out of range error when no user is specified
	var user *discordgo.User

	if len(i.ApplicationCommandData().Options[0].Options) == 0 {
		user = i.Member.User
	} else {
		user = i.ApplicationCommandData().Options[0].Options[0].UserValue(s)
	}

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: fmt.Sprintf("Stats for @%s", user.Username),
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:   "ID",
							Value:  user.ID,
							Inline: true,
						},
					},
					Footer: &discordgo.MessageEmbedFooter{
						Text: "No more stats for now :(",
					},
				},
			},
		},
	}, nil
}

func (c *CommandHandlers) UserLink(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	linkToken, err := c.dbs.LinkTokenStore.CreateToken(types.UserID(i.Member.User.ID))
	if err != nil {
		return nil, errors.Wrap(err, "Error creating link token")
	}

	// https://esefex.com/link?<linktoken>
	linkUrl := fmt.Sprintf("%s/link?t=%s", c.domain, linkToken.Token)
	expiresIn := time.Until(linkToken.Expiry)

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: "Open the link to connect your Discord account to Esefex",
					URL:   linkUrl,
					Footer: &discordgo.MessageEmbedFooter{
						Text: fmt.Sprintf("Expires in %s", expiresIn.String()),
					},
				},
			},
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Style: discordgo.LinkButton,
							Label: "Link Your Account",
							URL:   linkUrl,
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

func (c *CommandHandlers) UserUnlink(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	err := c.dbs.UserDB.DeleteUser(types.UserID(i.Member.User.ID))
	if err != nil {
		return nil, err
	}

	err = c.dbs.UserDB.SetUser(userdb.User{
		ID:     types.UserID(i.Member.User.ID),
		Tokens: []userdb.Token{},
	})
	if err != nil {
		return nil, err
	}

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "Unlinked your account from Esefex",
					Description: "Your account has been unlinked from Esefex. You can now link your account again.",
				},
			},
		},
	}, nil
}
