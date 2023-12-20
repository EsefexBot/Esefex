package main

import (
	"esefexapi/bot"

	"log"
	"os"
	"os/signal"

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

	bot := bot.NewDiscordBot(s)

	bot.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	bot.Stop()
}
