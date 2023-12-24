package main

import (
	"esefexapi/api"
	"esefexapi/audioplayer/discordplayer"
	"esefexapi/bot"
	"esefexapi/config"
	"esefexapi/sounddb/dbcache"
	"esefexapi/sounddb/filedb"
	"esefexapi/util"

	"log"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	log.Println("Starting Esefex API")

	cfg, err := config.LoadConfig("config.toml")
	if err != nil {
		log.Fatal(err)
	}

	ds, err := bot.CreateSession()
	if err != nil {
		log.Fatal(err)
	}

	fdb, err := filedb.NewFileDB(cfg.FileDatabase.Location)
	if err != nil {
		log.Fatal(err)
	}

	db := dbcache.NewDBCache(fdb)

	plr := discordplayer.NewDiscordPlayer(ds, db)

	api := api.NewHttpApi(db, plr, cfg.HttpApi.Port, cfg.HttpApi.CustomProtocol)
	bot := bot.NewDiscordBot(ds, db, cfg.HttpApi.Domain)

	log.Println("Components bootstraped, starting...")

	<-api.Start()
	<-bot.Start()
	<-plr.Start()

	log.Println("All components started successfully :)")
	log.Println("Press Ctrl+C to exit")
	<-util.Interrupt()
	log.Println("Gracefully shutting down...")

	<-api.Stop()
	<-bot.Stop()
	<-plr.Stop()

	log.Println("All components stopped, exiting...")
}
