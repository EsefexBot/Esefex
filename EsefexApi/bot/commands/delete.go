package commands

import (
	"esefexapi/sounddb"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

var DeleteCommand = &discordgo.ApplicationCommand{
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

func (c *CommandHandlers) Delete(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	options := OptionsMap(i)
	soundID := options["sound-id"]

	uid := sounddb.SuidFromStrings(i.GuildID, fmt.Sprint(soundID.Value))

	exists, err := c.dbs.SoundDB.SoundExists(uid)
	if err != nil {
		return nil, errors.Wrap(err, "Error checking if sound exists")
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

	err = c.dbs.SoundDB.DeleteSound(uid)
	if err != nil {
		return nil, errors.Wrap(err, "Error deleting sound")
	}

	log.Printf("Deleted sound effect %v from guild %v", soundID.Value, i.GuildID)

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Deleted sound effect `%s`", soundID.Value),
		},
	}, nil
}
