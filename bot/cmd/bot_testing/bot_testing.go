package main

import (
	"esefexbot/appcontext"
	"esefexbot/bot"
	"esefexbot/msg"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
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

	context := appcontext.Context{
		Channels: appcontext.Channels{
			A2B: make(chan msg.MessageA2B),
			B2A: make(chan msg.MessageB2A),
		},
		DiscordSession: s,
	}

	bot.Run(&context)
}
