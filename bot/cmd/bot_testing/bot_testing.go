package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	token := os.Getenv("TOKEN")
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Println("error creating Discord session,", err)
	}

	applicationCommand := discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "ping pong",
	}

	cmd, err := dg.ApplicationCommandCreate(dg.State.User.ID, "", &applicationCommand)
	if err != nil {
		log.Println("error creating ApplicationCommand,", err)
	}

	log.Println(cmd.ID)

	// append(registeredCommands, a)

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		log.Println("error opening connection,", err)
		return
	}

	// dg.ChannelMessageSend("777344211828604950", "https://media.tenor.com/sAhYu4Wd7IcAAAAd/blm.gif")

	// Wait here until CTRL-C or other term signal is received.
	log.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	s.ChannelMessageSend(m.ChannelID, "https://media.tenor.com/sAhYu4Wd7IcAAAAd/blm.gif")
}
