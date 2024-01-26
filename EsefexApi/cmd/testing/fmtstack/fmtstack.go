package main

import (
	"esefexapi/bot"
	"esefexapi/permissions"
	"esefexapi/types"
	. "esefexapi/util/must"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	guild := types.GuildID("777344211246120991")

	ps := permissions.NewPermissionStack()
	ps.SetChannel(types.ChannelID("777344211828604950"), permissions.NewEveryoneDefault())
	ps.SetRole(types.RoleID("1177022771260297278"), permissions.NewAllow())
	ps.SetUser(types.UserID("246941577904128000"), permissions.NewDeny())
	ps.SetUser(types.UserID("247763762298355712"), permissions.NewUnset())

	ds, err := bot.CreateSession()
	Must(err)

	ds.Open()
	defer ds.Close()

	fmt.Println(ps.FmtStack(ds, guild))
}
