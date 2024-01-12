package commands

import (
	"esefexapi/types"
	"esefexapi/userdb"
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/davecgh/go-spew/spew"
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
	}

	log.Println("User command called with options:")
	spew.Dump(i.ApplicationCommandData().Options)

	return nil, errors.Wrap(fmt.Errorf("Not implemented"), "Sound")
}

func (c *CommandHandlers) UserStats(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	// TODO: index out of range error when no user is specified
	var user *discordgo.User

	if len(i.ApplicationCommandData().Options[0].Options) == 0 {
		user = i.Member.User
	} else {
		user = i.ApplicationCommandData().Options[0].Options[0].UserValue(s)
	}

	stats := fmt.Sprintf("Stats for %s:\n", user.Username)
	stats += fmt.Sprintf("ID: %s\n", user.ID)
	stats += "No more stats for now... :("

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: stats,
		},
	}, nil
}

func (c *CommandHandlers) UserLink(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	linkToken, err := c.dbs.LinkTokenStore.CreateToken(types.UserID(i.Member.User.ID))
	if err != nil {
		return nil, errors.Wrap(err, "Error creating link token")
	}

	channel, err := s.UserChannelCreate(i.Member.User.ID)
	if err != nil {
		return nil, errors.Wrap(err, "Error creating DM channel")
	}

	// https://esefex.com/link?<linktoken>
	linkUrl := fmt.Sprintf("%s/link?t=%s", c.domain, linkToken.Token)
	expiresIn := time.Until(linkToken.Expiry)
	_, err = s.ChannelMessageSend(channel.ID, fmt.Sprintf("Click this link to link your Discord account to Esefex (expires in %d Minutes): \n%s", int(expiresIn.Minutes()), linkUrl))
	if err != nil {
		return nil, errors.Wrap(err, "Error sending DM message")
	}

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Please check your DMs for a link to link your Discord account to Esefex",
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
			Content: "Your account has been unlinked from Esefex. You can now link your account again.",
		},
	}, nil
}
