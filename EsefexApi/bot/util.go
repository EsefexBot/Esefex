package bot

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"

	"github.com/bwmarrin/discordgo"
)

// call this before opening the session
func (b *DiscordBot) RegisterComandHandlers() {
	ds := b.Session

	ds.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := b.cmdh.Handlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

// Call after opening the session
func (b *DiscordBot) RegisterComands() error {
	ds := b.Session

	for _, v := range b.cmdh.Commands {
		_, err := ds.ApplicationCommandCreate(ds.State.User.ID, "", v)
		if err != nil {
			return errors.Wrapf(err, "Cannot create '%v' command", v.Name)
		}

		log.Printf("Registered '%v' command", v.Name)
	}

	return nil
}

func (b *DiscordBot) DeleteAllCommands() error {
	ds := b.Session

	log.Println("Deleting all commands...")

	for _, g := range ds.State.Guilds {
		err := b.DeleteGuildCommands(g.ID)
		if err != nil {
			return errors.Wrapf(err, "Cannot delete commands for guild '%v'", g.ID)
		}
	}

	err := b.DeleteGuildCommands("")
	if err != nil {
		return errors.Wrap(err, "Cannot delete global commands")
	}

	log.Println("Deleted all commands")

	return nil
}

func (b *DiscordBot) DeleteGuildCommands(guildID string) error {
	ds := b.Session

	cmds, err := ds.ApplicationCommands(ds.State.User.ID, guildID)
	if err != nil {
		return errors.Wrapf(err, "Cannot get commands for guild '%v'", guildID)
	}

	for _, v := range cmds {
		err = ds.ApplicationCommandDelete(ds.State.User.ID, guildID, v.ID)
		if err != nil {
			return errors.Wrapf(err, "Cannot delete '%v' command", v.Name)
		}
		log.Printf("Deleted '%v' command", v.Name)
	}

	return nil
}

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

	ds.Session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
		close(ready)
	})

	return ready
}
