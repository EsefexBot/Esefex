package main

import (
	"esefexapi/api"
	"esefexapi/audioplayer"
	"esefexapi/audioplayer/mockplayer"
	"esefexapi/config"
	"esefexapi/sounddb"
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

	fdb, err := filedb.NewFileDB(cfg.FileDatabase.Location)
	if err != nil {
		log.Fatal(err)
	}

	var db sounddb.ISoundDB = dbcache.NewDBCache(fdb)
	var player audioplayer.IAudioPlayer = mockplayer.NewMockPlayer()

	api := api.NewHttpApi(db, player, cfg.HttpApi.Port, cfg.HttpApi.CustomProtocol)

	<-api.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	<-api.Stop()
}
