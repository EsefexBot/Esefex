package main

import (
	"esefexapi/api"
	"esefexapi/audioprocessing"
	"esefexapi/bot"
	"esefexapi/ctx"
	"esefexapi/db"
	"esefexapi/db/filedb"
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
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	var db db.SoundDB = filedb.NewFileDb()

	c := ctx.Ctx{
		Channels: ctx.Channels{
			// A2B:  make(chan msg.MessageA2B),
			// B2A:  make(chan msg.MessageB2A),
			PlaySound: make(chan msg.PlaySound),
			Stop:      make(chan struct{}, 1),
		},
		DiscordSession: bot.CreateSession(),
		CustomProtocol: "esefexapi",
		ApiPort:        "8080",
		AudioCache:     audioprocessing.NewAudioCache(),
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

	print("\n")
	log.Println("Stopping...")

	close(c.Channels.Stop)
	wg.Wait()
}
