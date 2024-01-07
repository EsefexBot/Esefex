package commands

import (
	"esefexapi/types"
	"esefexapi/userdb"

	"github.com/bwmarrin/discordgo"
)

var (
	UnlinkCommand = &discordgo.ApplicationCommand{
		Name:        "unlink",
		Description: "Unlink your Discord account from Esefex. Useful if you think your account has been compromised.",
	}
)

func (c *CommandHandlers) Unlink(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
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
