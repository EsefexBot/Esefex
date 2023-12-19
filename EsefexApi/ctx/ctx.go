package ctx

import (
	"esefexapi/audioprocessing"
	"esefexapi/msg"
	"esefexapi/vchandler"

	"log"
	"sync"

	"github.com/bwmarrin/discordgo"
)

type Ctx struct {
	CustomProtocol    string
	ApiPort           string
	Channels          Channels
	DiscordSession    *discordgo.Session
	AudioCache        *audioprocessing.AudioCache
	ConnectionHandler *vchandler.ConnectionHandler
}

type Channels struct {
	// A2B chan msg.MessageA2B
	PlaySound chan msg.PlaySound
	// B2A  chan msg.MessageB2A
	Stop chan struct{}
}

var lock = &sync.Mutex{}
var instance *Ctx

func GetInstance() *Ctx {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &Ctx{
				Channels: Channels{
					// A2B:      make(chan msg.MessageA2B),
					PlaySound: make(chan msg.PlaySound),
					// B2A:      make(chan msg.MessageB2A),
					Stop: make(chan struct{}),
				},
			}
		}
	}
	return instance
}

func SetContext(ctx *Ctx) {
	if instance != nil {
		log.Println("Context should not be set more than once")
	}

	lock.Lock()
	defer lock.Unlock()
	instance = ctx
}
