package commands

import "github.com/bwmarrin/discordgo"

var (
	Commands = map[string]*discordgo.ApplicationCommand{}
	Handlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){}
)

func init() {
	Commands["upload"] = UploadCommand
	Handlers["upload"] = Upload

	Commands["session"] = SessionCommand
	Handlers["session"] = Session
}
