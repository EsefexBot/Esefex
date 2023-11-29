package main

import (
	"esefexbot/appcontext"
	"esefexbot/bot"
	"esefexbot/msg"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	godotenv.Load()
	token := os.Getenv("BOT_TOKEN")

	if token == "" {
		log.Fatal("BOT_TOKEN is not set")
		return
	}

	s, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}

	c := appcontext.Context{
		Channels: appcontext.Channels{
			A2B:  make(chan msg.MessageA2B),
			B2A:  make(chan msg.MessageB2A),
			Stop: make(chan bool, 1),
		},
		DiscordSession: s,
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		bot.Run(&c)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	close(c.Channels.Stop)
	wg.Wait()
}
