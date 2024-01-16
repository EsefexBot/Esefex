package cmdhandler

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

type CommandHandlerWithErr func(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error)

func ErrorHandler(msg string) CommandHandlerWithErr {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
		return nil, errors.Wrap(fmt.Errorf(msg), "Error handling command")
	}
}
