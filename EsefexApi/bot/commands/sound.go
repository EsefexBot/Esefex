package commands

import (
	"esefexapi/sounddb"
	"esefexapi/types"
	"esefexapi/util"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

var SoundCommand = &discordgo.ApplicationCommand{
	Name:        "sound",
	Description: "All commands related to sounds.",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "upload",
			Description: "Upload a sound effect to the bot",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
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
		},
		{
			Name:        "delete",
			Description: "Delete a sound effect",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "sound-name",
					Description: "The sound effect to delete",
					Required:    true,
				},
			},
		},
		{
			Name:        "list",
			Description: "List all sound effects in the guild",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
		},
		{
			Name:        "play",
			Description: "Play a sound effect",
			Type:        discordgo.ApplicationCommandOptionSubCommand,
		},
	},
}

func (c *CommandHandlers) Sound(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	switch i.ApplicationCommandData().Options[0].Name {
	case "upload":
		return c.mw.CheckPerms(c.SoundUpload, "Sound.Upload")(s, i)
	case "delete":
		return c.mw.CheckPerms(c.SoundDelete, "Sound.Delete")(s, i)
	case "list":
		return c.SoundList(s, i)
	case "play":
		return c.mw.CheckPerms(c.SoundPlay, "Sound.Play")(s, i)
	default:
		return nil, errors.Wrap(fmt.Errorf("Unknown subcommand %s", i.ApplicationCommandData().Options[0].Name), "Error handling user command")
	}
}

func (c *CommandHandlers) SoundUpload(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	options := OptionsMap(i.ApplicationCommandData().Options[0].Options)

	iconOption := options["icon"]
	icon, err := sounddb.ExtractIcon(fmt.Sprint(iconOption.Value))
	if err != nil {
		return nil, errors.Wrap(err, "Error extracting icon")
	}

	soundFile := fmt.Sprint(options["sound-file"].Value)
	soundFileUrl := i.ApplicationCommandData().Resolved.Attachments[soundFile].URL

	pcm, err := util.Download2PCM(soundFileUrl)
	if err != nil {
		return nil, errors.Wrap(err, "Error downloading sound file")
	}

	soundName := types.SoundName(fmt.Sprint(options["name"].Value))

	uid, err := c.dbs.SoundDB.AddSound(types.GuildID(i.GuildID), soundName, icon, pcm)
	if err != nil {
		return nil, errors.Wrap(err, "Error adding sound")
	}

	guildID := types.GuildID(i.GuildID)
	c.cn.UpdateNotificationGuilds(guildID)

	log.Printf("Uploaded sound effect %v to guild %v", uid.SoundName, i.GuildID)
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: "Sound Effect Uploaded",
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:   "Name",
							Value:  fmt.Sprintf("%s %s", options["name"].Value, icon.String()),
							Inline: true,
						},
					},
				},
			},
		},
	}, nil
}

func (c *CommandHandlers) SoundDelete(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	options := OptionsMap(i.ApplicationCommandData().Options[0].Options)
	soundName := types.SoundName(fmt.Sprint(options["sound-name"].Value))

	uid := sounddb.SoundUID{
		GuildID:   types.GuildID(i.GuildID),
		SoundName: soundName,
	}

	exists, err := c.dbs.SoundDB.SoundExists(uid)
	if err != nil {
		return nil, errors.Wrap(err, "Error checking if sound exists")
	}
	if !exists {
		return nil, errors.Wrap(fmt.Errorf("Sound effect %s does not exist", soundName), "Error deleting sound")
	}

	log.Print("a")

	err = c.dbs.SoundDB.DeleteSound(uid)
	if err != nil {
		return nil, errors.Wrap(err, "Error deleting sound")
	}

	log.Printf("Deleted sound effect %v from guild %v", soundName, i.GuildID)

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: fmt.Sprintf("Deleted sound effect `%s`", soundName),
				},
			},
		},
	}, nil
}

func (c *CommandHandlers) SoundList(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	uids, err := c.dbs.SoundDB.GetSoundUIDs(types.GuildID(i.GuildID))
	if err != nil {
		return nil, errors.Wrap(err, "Error getting sound UIDs")
	}
	// log.Printf("List: %v", uids)

	var metas = make([]sounddb.SoundMeta, 0, len(uids))
	for _, uid := range uids {
		meta, err := c.dbs.SoundDB.GetSoundMeta(uid)
		if err != nil {
			return nil, errors.Wrap(err, "Error getting sound meta")
		}
		metas = append(metas, meta)
	}

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				fmtMetaListAsEmbed(metas),
			},
		},
	}, nil
}

// TODO: Implement SoundPlay
func (c *CommandHandlers) SoundPlay(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
	return nil, errors.Wrap(fmt.Errorf("Not implemented"), "Sound")
}
