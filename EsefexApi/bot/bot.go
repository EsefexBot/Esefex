package bot

import (
	"esefexapi/bot/commands"
	"esefexapi/clientnotifiy"
	"esefexapi/db"
	"esefexapi/service"

	"log"

	"github.com/bwmarrin/discordgo"
)

var _ service.IService = &DiscordBot{}

// DiscordBot implements Service
type DiscordBot struct {
	ds    *discordgo.Session
	cmdh  *commands.CommandHandlers
	cn    clientnotifiy.IClientNotifier
	stop  chan struct{}
	ready chan struct{}
}

func NewDiscordBot(ds *discordgo.Session, dbs *db.Databases, domain string, cn clientnotifiy.IClientNotifier) *DiscordBot {
	return &DiscordBot{
		ds:    ds,
		cmdh:  commands.NewCommandHandlers(ds, dbs, domain, cn),
		cn:    cn,
		stop:  make(chan struct{}, 1),
		ready: make(chan struct{}),
	}
}

func (b *DiscordBot) run() {
	defer close(b.stop)
	log.Println("Starting bot...")
	defer log.Println("Bot stopped")

	ds := b.ds
	ds.Identify.Intents = discordgo.IntentsAllWithoutPrivileged | discordgo.IntentsGuildMembers | discordgo.IntentsGuildPresences

	ready := b.WaitReady()

	b.cmdh.RegisterComandHandlers()
	b.RegisterClientUpdateHandlers()

	err := ds.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %+v", err)
	}
	defer ds.Close()

	<-ready

	err = b.cmdh.UpdateApplicationCommands()
	if err != nil {
		log.Printf("Cannot update command handlers: %+v", err)
	}

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
