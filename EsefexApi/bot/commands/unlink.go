package commands

import (
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
	c.dbs.UserDB.DeleteUser(i.Member.User.ID)
	c.dbs.UserDB.SetUser(userdb.User{
		ID:     i.Member.User.ID,
		Tokens: []userdb.Token{},
	})

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Your account has been unlinked from Esefex. You can now link your account again.",
		},
	}, nil
}
