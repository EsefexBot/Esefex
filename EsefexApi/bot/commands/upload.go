package commands

import (
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

func (c *CommandHandlers) Upload(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	options := OptionsMap(i)

	icon := options["icon"]
	iconURL := util.ExtractIconUrl(icon)

	soundFile := options["sound-file"]
	soundFileUrl := i.ApplicationCommandData().Resolved.Attachments[fmt.Sprint(soundFile.Value)].URL

	pcm, err := util.Download2PCM(soundFileUrl)
	if err != nil {
		return nil, err
	}

	c.db.AddSound(i.GuildID, fmt.Sprint(options["name"].Value), iconURL, pcm)

	log.Printf("Uploaded sound effect %v to server %v", options["name"].Value, i.GuildID)
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Uploaded sound effect",
		},
	}, nil
}
