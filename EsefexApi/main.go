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

	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	log.Printf("Starting Esefex API with PID: %d", os.Getpid())

	cfg, err := config.LoadConfig("config.toml")
	if err != nil {
		log.Fatal(err)
	}

	ds, err := bot.CreateSession()
	if err != nil {
		log.Fatal(err)
	}

	sdb, err := filesounddb.NewFileDB(cfg.FileSoundDB.Location)
	if err != nil {
		log.Fatal(err)
	}

	sdbc := dbcache.NewSoundDBCache(sdb)

	udb, err := fileuserdb.NewFileUserDB(cfg.FileUserDB.Location)
	if err != nil {
		log.Fatal(err)
	}

	ldb := memorylinktokenstore.NewMemoryLinkTokenStore(time.Minute * 5)

	dbs := db.Databases{
		SoundDB:        sdbc,
		UserDB:         udb,
		LinkTokenStore: ldb,
	}

	plr := discordplayer.NewDiscordPlayer(ds, dbs, cfg.Bot.UseTimeouts, time.Duration(cfg.Bot.Timeout)*time.Second)

	api := api.NewHttpApi(dbs, plr, cfg.HttpApi.Port, cfg.HttpApi.CustomProtocol)
	bot := bot.NewDiscordBot(ds, dbs, cfg.HttpApi.Domain)

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
	println()
	log.Println("Gracefully shutting down...")
}
