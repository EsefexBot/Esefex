package main

import (
	"esefexapi/bot"
	. "esefexapi/util/must"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

func main() {
	Must(godotenv.Load())

	ds, err := bot.CreateSession()
	Must(err)

	err = ds.Open()
	Must(err)
	defer ds.Close()

	guildID := "489017101894418444"

	cmds, err := ds.ApplicationCommands(ds.State.User.ID, guildID)
	Must(err)

	spew.Dump(cmds)
}
