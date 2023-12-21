package main

import (
	"esefexapi/api"
	"esefexapi/audioplayer"
	"esefexapi/audioplayer/mockplayer"
	"esefexapi/sounddb"
	"esefexapi/sounddb/dbcache"
	"esefexapi/sounddb/filedb"
	"log"
	"os"
	"os/signal"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var db sounddb.ISoundDB = dbcache.NewDBCache(filedb.NewFileDB())
	var player audioplayer.IAudioPlayer = mockplayer.NewMockPlayer()

	api := api.NewHttpApi(db, player, 8080, "esefexapi")

	<-api.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	<-api.Stop()
}
