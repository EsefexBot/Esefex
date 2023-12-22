package main

import (
	"esefexapi/api"
	"esefexapi/audioplayer/discordplayer"
	"esefexapi/bot"
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
	ds, err := bot.CreateSession()
	if err != nil {
		log.Fatal(err)
	}

	db := dbcache.NewDBCache(filedb.NewFileDB())

	plr := discordplayer.NewDiscordPlayer(ds, db)

	api := api.NewHttpApi(db, plr, 8080, "http")
	bot := bot.NewDiscordBot(ds, db)

	<-api.Start()
	<-bot.Start()
	<-plr.Start()

	log.Println("Press Ctrl+C to exit")
	<-util.OsInterrupt()

	<-api.Stop()
	<-bot.Stop()
	<-plr.Stop()
}
