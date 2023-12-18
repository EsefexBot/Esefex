package main

import (
	"esefexapi/api"
	"esefexapi/appcontext"
	"esefexapi/bot"
	"esefexapi/msg"

	// "esefexapi/msg"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	c := appcontext.Context{
		Channels: appcontext.Channels{
			// A2B:  make(chan msg.MessageA2B),
			// B2A:  make(chan msg.MessageB2A),
			PlaySound: make(chan msg.PlaySound),
			Stop:      make(chan bool, 1),
		},
		DiscordSession: bot.CreateSession(),
		CustomProtocol: "esefexapi",
		ApiPort:        "8080",
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		bot.Run(&c)
	}()

	go func() {
		defer wg.Done()
		api.Run(&c)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	close(c.Channels.Stop)
	wg.Wait()
}
