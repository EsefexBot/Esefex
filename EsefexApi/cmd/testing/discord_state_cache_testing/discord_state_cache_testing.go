package main

import (
	"esefexapi/bot"
	// "esefexapi/util"
	// "reflect"
	"time"

	"log"

	// "github.com/bwmarrin/discordgo"
	// "github.com/davecgh/go-spew/spew"
	// "github.com/bwmarrin/discordgo"
	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	spew.Config.MaxDepth = 2
}

func main() {
	gid := "777344211246120991"

	s, err := bot.CreateSession()
	if err != nil {
		panic(err)
	}

	// s.Debug = true

	log.Printf("state tracking: %t", s.StateEnabled)
	log.Printf("Sync Events: %t", s.SyncEvents)

	// s.AddHandler(func(s *discordgo.Session, e interface{}) {
	// 	eventName := reflect.TypeOf(e)
	// 	log.Printf("Event Received: %v", eventName)
	// })

	err = s.Open()
	if err != nil {
		panic(err)
	}

	log.Println("waiting for data ready")
	for s.DataReady == false {
		print(".")
	}
	print("\n")
	log.Println("data ready")

	// g, err := s.Guild(gid)
	// if err != nil {
	// 	panic(err)
	// }

	for {
		// g, err := s.Guild(gid)
		g, err := s.State.Guild(gid)
		if err != nil {
			panic(err)
		}

		spew.Dump(g.VoiceStates)

		time.Sleep(1 * time.Second)
	}

	// <-util.Interrupt()
}
