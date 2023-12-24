package main

import (
	"esefexapi/bot"
	"esefexapi/config"
	"esefexapi/sounddb/dbcache"
	"esefexapi/sounddb/filedb"

	"log"
	"os"
	"os/signal"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
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

	bot := bot.NewDiscordBot(ds, db, cfg.HttpApi.Domain)

	<-bot.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	<-bot.Stop()
}
