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

func NewDiscordBot(s *discordgo.Session, dbs db.Databases, domain string) *DiscordBot {
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

	s := b.Session

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	s.Identify.Intents = discordgo.IntentsAllWithoutPrivileged | discordgo.IntentsGuildMembers | discordgo.IntentsGuildPresences

	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}
	defer s.Close()

	log.Println("Registering commands...")
	b.RegisterComands(s)
	defer b.DeleteAllCommands(s)
	// defer actions.LeaveAllChannels(s)

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
