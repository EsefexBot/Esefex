package commands

import (
	"esefexapi/types"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

var (
	LinkCommand = &discordgo.ApplicationCommand{
		Name:        "link",
		Description: "Link your Discord account to Esefex",
	}
)

func (c *CommandHandlers) Link(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
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
