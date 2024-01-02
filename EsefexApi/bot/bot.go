package bot

import (
	"esefexapi/bot/commands"
	"esefexapi/db"
	"esefexapi/service"

	"log"

	"github.com/bwmarrin/discordgo"
)

var _ service.IService = &DiscordBot{}

// DiscordBot implements Service
type DiscordBot struct {
	Session *discordgo.Session
	cmdh    *commands.CommandHandlers
	stop    chan struct{}
	ready   chan struct{}
}

func NewDiscordBot(s *discordgo.Session, dbs *db.Databases, domain string) *DiscordBot {
	return &DiscordBot{
		Session: s,
		cmdh:    commands.NewCommandHandlers(dbs, domain),
		stop:    make(chan struct{}, 1),
		ready:   make(chan struct{}),
	}
}

func (b *DiscordBot) run() {
	defer close(b.stop)
	log.Println("Starting bot...")
	defer log.Println("Bot stopped")

	ds := b.Session
	ds.Identify.Intents = discordgo.IntentsAllWithoutPrivileged | discordgo.IntentsGuildMembers | discordgo.IntentsGuildPresences

	ready := b.WaitReady()
	log.Println("Registering command handlers...")
	b.RegisterComandHandlers()

	err := ds.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %+v", err)
	}
	defer ds.Close()

	<-ready

	log.Println("Registering commands...")
	b.RegisterComands()
	defer b.DeleteAllCommands()

	log.Println("Bot Ready.")
	close(b.ready)
	<-b.stop
}

func (b *DiscordBot) Start() <-chan struct{} {
	go b.run()
	return b.ready
}

func (b *DiscordBot) Stop() <-chan struct{} {
	b.stop <- struct{}{}
	return b.stop
}
