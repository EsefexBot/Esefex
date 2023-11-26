package commands

import (
	// "esefexbot/util"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	UploadCommand = &discordgo.ApplicationCommand{
		Name:        "upload",
		Description: "Upload a sound effect to the bot",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionAttachment,
				Name:        "sound-file",
				Description: "The sound file to upload",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "name",
				Description: "The name of the sound effect",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "icon",
				Description: "The icon to use for the sound effect",
				Required:    true,
			},
		},
	}
)

func Upload(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options

	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	icon := optionMap["icon"]
	log.Println(icon.Value)

	// iconURL := util.GetIconURL(icon)

	// println(iconURL)

	soundFile := optionMap["sound-file"]
	// println(soundFile.Value)
	soundFileUrl := i.ApplicationCommandData().Resolved.Attachments[fmt.Sprint(soundFile.Value)].URL
	println(soundFileUrl)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Uploaded sound effect",
		},
	})
}
