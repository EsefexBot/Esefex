package main

import (
	"esefexapi/api"
	"esefexapi/audioplayer/discordplayer"
	"esefexapi/bot"
	"esefexapi/config"
	"esefexapi/db"
	"esefexapi/linktokenstore/memorylinktokenstore"
	"esefexapi/sounddb/dbcache"
	"esefexapi/sounddb/filesounddb"
	"esefexapi/userdb/fileuserdb"
	"esefexapi/util"
	. "esefexapi/util/must"

	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	log.Printf("Starting Esefex API with PID: %d", os.Getpid())

	Must(godotenv.Load())

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	cfg, err := config.LoadConfig("config.toml")
	Must(err)

	domain := os.Getenv("DOMAIN")

	ds, err := bot.CreateSession()
	Must(err)

	sdb, err := filesounddb.NewFileDB(cfg.FileSoundDB.Location)
	Must(err)

	sdbc, err := dbcache.NewSoundDBCache(sdb)
	Must(err)

	udb, err := fileuserdb.NewFileUserDB(cfg.FileUserDB.Location)
	Must(err)

	verT := time.Duration(cfg.VerificationExpiry * float32(time.Minute))
	ldb := memorylinktokenstore.NewMemoryLinkTokenStore(verT)

	dbs := &db.Databases{
		SoundDB:        sdbc,
		UserDB:         udb,
		LinkTokenStore: ldb,
	}

	botT := time.Duration(cfg.Bot.Timeout * float32(time.Minute))
	plr := discordplayer.NewDiscordPlayer(ds, dbs, cfg.Bot.UseTimeouts, botT)

	api := api.NewHttpApi(dbs, plr, ds, cfg.HttpApi.Port, cfg.HttpApi.CustomProtocol)
	bot := bot.NewDiscordBot(ds, dbs, domain)

	log.Println("Components bootstraped, starting...")

	<-api.Start()
	<-bot.Start()
	<-plr.Start()

	defer func() {
		<-api.Stop()
		<-bot.Stop()
		<-plr.Stop()

		udb.Close()

		log.Println("All components stopped, exiting...")
	}()

	log.Println("All components started successfully :)")
	log.Println("Press Ctrl+C to exit")
	<-util.Interrupt()
	fmt.Println()
	log.Println("Gracefully shutting down...")
}
