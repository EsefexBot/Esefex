package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type DiscordBot struct {
	Session *discordgo.Session
	stop    chan struct{}
}

func NewDiscordBot(s *discordgo.Session) *DiscordBot {
	return &DiscordBot{
		Session: s,
		stop:    make(chan struct{}, 1),
	}
}

func (b *DiscordBot) run() {
	log.Println("Starting bot...")

	s := b.Session

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}
	defer s.Close()

	log.Println("Adding commands...")
	RegisterComands(s)

	log.Println("Bot Ready.")
	<-b.stop

	// defer actions.LeaveAllChannels(s)
	defer DeleteAllCommands(s)
}

func (b *DiscordBot) Start() {
	go b.run()
}

func (b *DiscordBot) Stop() {
	close(b.stop)
}
