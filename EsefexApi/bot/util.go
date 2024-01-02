package bot

import (
	"github.com/pkg/errors"
	"log"
	"os"

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
func (b *DiscordBot) RegisterComands() {
	ds := b.Session

	for _, v := range b.cmdh.Commands {
		_, err := ds.ApplicationCommandCreate(ds.State.User.ID, "", v)
		if err != nil {
			log.Printf("Cannot create '%v' command: %v", v.Name, err)
		}

		log.Printf("Registered '%v' command", v.Name)
	}
}

func (b *DiscordBot) DeleteAllCommands() {
	ds := b.Session

	log.Println("Deleting all commands...")

	for _, g := range ds.State.Guilds {
		b.DeleteGuildCommands(g.ID)
	}

	b.DeleteGuildCommands("")

	log.Println("Deleted all commands")
}

func (b *DiscordBot) DeleteGuildCommands(guildID string) {
	ds := b.Session

	cmds, err := ds.ApplicationCommands(ds.State.User.ID, guildID)
	if err != nil {
		log.Printf("Cannot get commands for guild '%v': %v", guildID, err)
	}

	for _, v := range cmds {
		ds.ApplicationCommandDelete(ds.State.User.ID, guildID, v.ID)
		log.Printf("Deleted '%v' command", v.Name)
	}
}

var BotTokenNotSet = errors.New("BOT_TOKEN is not set")

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
