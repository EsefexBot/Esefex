package main

import (
	"esefexapi/api"
	"esefexapi/audioplayer/discordplayer"
	"esefexapi/bot"
	"esefexapi/sounddb/dbcache"
	"esefexapi/sounddb/filedb"

	// "esefexapi/msg"
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

	a := discordplayer.NewDiscordPlayer(ds, db)

	api := api.NewHttpApi(db, a, 8080, "http")
	bot := bot.NewDiscordBot(ds, db)

	<-api.Start()
	<-bot.Start()

	<-api.Stop()
	<-api.Stop()
}
