package main

import (
	"esefexapi/api"
	"esefexapi/audioplayer/discordplayer"
	"esefexapi/bot"
	"esefexapi/bot/commands/cmdhashstore"
	"esefexapi/clientnotifiy"
	"esefexapi/config"
	"esefexapi/db"
	"esefexapi/linktokenstore/memorylinktokenstore"
	"esefexapi/permissiondb/filepermisssiondb"
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

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, assuming all variables are set in the environment")
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	domain := os.Getenv("DOMAIN")

	ds, err := bot.CreateSession()
	Must(err)

	sdb, err := filesounddb.NewFileDB()
	Must(err)

	sdbc, err := dbcache.NewSoundDBCache(sdb)
	Must(err)

	udb, err := fileuserdb.NewFileUserDB()
	Must(err)

	fpdb, err := filepermisssiondb.NewFilePermissionDB(ds)
	Must(err)

	verT := time.Duration(config.Get().VerificationExpiry * float32(time.Minute))
	ldb := memorylinktokenstore.NewMemoryLinkTokenStore(verT)

	fcmhs := cmdhashstore.NewFileCmdHashStore()

	dbs := &db.Databases{
		SoundDB:        sdbc,
		UserDB:         udb,
		LinkTokenStore: ldb,
		PermissionDB:   fpdb,
		CmdHashStore:   fcmhs,
	}

	botT := time.Duration(config.Get().Bot.Timeout * float32(time.Minute))
	plr := discordplayer.NewDiscordPlayer(ds, dbs, botT)

	wsCN := clientnotifiy.NewWsClientNotifier(ds)

	api := api.NewHttpApi(dbs, plr, ds, wsCN, domain)
	bot := bot.NewDiscordBot(ds, dbs, domain, wsCN)

	log.Println("Components bootstraped, starting...")

	<-api.Start()
	<-bot.Start()
	<-plr.Start()

	defer func() {
		<-api.Stop()
		<-bot.Stop()
		<-plr.Stop()

		udb.Close()
		fpdb.Close()

		log.Println("All components stopped, exiting...")
	}()

	log.Println("All components started successfully :)")
	log.Println("Press Ctrl+C to exit")
	<-util.Interrupt()
	fmt.Println()
	log.Println("Gracefully shutting down...")
}
