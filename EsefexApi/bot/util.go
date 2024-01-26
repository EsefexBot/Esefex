package bot

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"

	"github.com/bwmarrin/discordgo"
)

var BotTokenNotSet = fmt.Errorf("BOT_TOKEN is not set")

func CreateSession() (*discordgo.Session, error) {
	token := os.Getenv("BOT_TOKEN")

	if token == "" {
		return nil, BotTokenNotSet
	}

	s, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, errors.Wrap(err, "Error creating session")
	}

	return s, nil
}

func (ds *DiscordBot) WaitReady() chan struct{} {
	ready := make(chan struct{}, 1)

	ds.ds.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
		close(ready)
	})

	return ready
}
