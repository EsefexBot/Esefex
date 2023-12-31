package commands

import (
	"esefexapi/sounddb"
	"esefexapi/util"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
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

	iconOption := options["icon"]
	icon, err := sounddb.ExtractIcon(fmt.Sprint(iconOption.Value))
	if err != nil {
		return nil, errors.Wrap(err, "Error extracting icon")
	}

	soundFile := options["sound-file"]
	soundFileUrl := i.ApplicationCommandData().Resolved.Attachments[fmt.Sprint(soundFile.Value)].URL

	pcm, err := util.Download2PCM(soundFileUrl)
	if err != nil {
		return nil, errors.Wrap(err, "Error downloading sound file")
	}

	uid, err := c.dbs.SoundDB.AddSound(i.GuildID, fmt.Sprint(options["name"].Value), icon, pcm)
	if err != nil {
		return nil, errors.Wrap(err, "Error adding sound")
	}

	log.Printf("Uploaded sound effect %v to server %v", uid.SoundID, i.GuildID)
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Uploaded sound effect %s %s", uid.SoundID, icon.Name),
		},
	}, nil
}
