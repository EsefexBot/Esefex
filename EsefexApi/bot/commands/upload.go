package commands

import (
	// "esefexapi/util"
	"esefexapi/filedb"
	"esefexapi/util"
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
	iconURL := util.ExtractIconUrl(icon)

	soundFile := optionMap["sound-file"]
	soundFileUrl := i.ApplicationCommandData().Resolved.Attachments[fmt.Sprint(soundFile.Value)].URL

	filedb.AddSound(i.GuildID, fmt.Sprint(optionMap["name"].Value), iconURL, soundFileUrl)

	log.Printf("Uploaded sound effect %v to server %v", optionMap["name"].Value, i.GuildID)
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Uploaded sound effect",
		},
	})
}
