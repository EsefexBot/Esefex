package commands

import (
	"esefexapi/sounddb"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	DeleteCommand = &discordgo.ApplicationCommand{
		Name:        "delete",
		Description: "Delete a sound effect",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "sound-id",
				Description: "The sound effect to delete",
				Required:    true,
			},
		},
	}
)

func (c *CommandHandlers) Delete(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	options := OptionsMap(i)
	soundID := options["sound-id"]

	uid := sounddb.SuidFromStrings(i.GuildID, fmt.Sprint(soundID.Value))

	exists, err := c.db.SoundExists(uid)
	if err != nil {
		return nil, err
	}
	if !exists {
		return &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("Sound effect `%s` does not exist", soundID.Value),
			},
		}, nil
	}

	log.Print("a")

	err = c.db.DeleteSound(uid)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Printf("Deleted sound effect %v from server %v", soundID.Value, i.GuildID)

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Deleted sound effect `%s`", soundID.Value),
		},
	}, nil
}
