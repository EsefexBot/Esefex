package appctx

// import (
// 	"esefexapi/audioplayer"
// 	"esefexapi/sounddb"
// 	"log"
// 	"sync"

// 	"github.com/bwmarrin/discordgo"
// )

// type Context struct {
// 	CustomProtocol string
// 	ApiPort        string
// 	Channels       Channels
// 	DiscordSession *discordgo.Session
// 	AudioPlayer    *audioplayer.IAudioPlayer
// 	DB             *sounddb.ISoundDB
// }

// type Channels struct {
// 	// A2B chan msg.MessageA2B
// 	PlaySound chan msg.PlaySound
// 	// B2A  chan msg.MessageB2AS
// 	Stop chan struct{}
// }

// var lock = &sync.Mutex{}
// var instance *Context

// func GetInstance() *Context {
// 	if instance == nil {
// 		lock.Lock()
// 		defer lock.Unlock()
// 		if instance == nil {
// 			instance = &Context{
// 				Channels: Channels{
// 					// A2B:      make(chan msg.MessageA2B),
// 					PlaySound: make(chan msg.PlaySound),
// 					// B2A:      make(chan msg.MessageB2A),
// 					Stop: make(chan struct{}),
// 				},
// 			}
// 		}
// 	}
// 	return instance
// }

// func SetContext(ctx *Context) {
// 	if instance != nil {
// 		log.Println("Context should not be set more than once")
// 	}

// 	lock.Lock()
// 	defer lock.Unlock()
// 	instance = ctx
// }
