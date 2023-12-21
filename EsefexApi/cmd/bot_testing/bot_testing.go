package main

import (
	"esefexapi/bot"
	"esefexapi/sounddb/filedb"

	"log"
	"os"
	"os/signal"
)

func main() {
	ds, err := bot.CreateSession()
	if err != nil {
		log.Fatal(err)
	}

	db := filedb.NewFileDB()

	bot := bot.NewDiscordBot(ds, db)

	<-bot.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	<-bot.Stop()
}
