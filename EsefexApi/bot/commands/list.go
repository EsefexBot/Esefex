package commands

import (
	"esefexapi/sounddb"
	"esefexapi/types"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

var ListCommand = &discordgo.ApplicationCommand{
	Name:        "list",
	Description: "List all sound effects in the guild",
}

func (c *CommandHandlers) List(s *discordgo.Session, i *discordgo.InteractionCreate) (*discordgo.InteractionResponse, error) {
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

	if len(metas) == 0 {
		return &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "There are no sounds in this guild",
			},
		}, nil
	}

	if len(metas) == 1 {
		return &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("There is 1 sound in this guild: \n%s", fmtMetaList(metas)),
			},
		}, nil
	}

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("There are %d sounds in this guild: \n%s", len(metas), fmtMetaList(metas)),
		},
	}, nil
}

func fmtMetaList(metas []sounddb.SoundMeta) string {
	// log.Printf("fmtMetaList: %v", metas)
	var str string
	for _, meta := range metas {
		str += fmt.Sprintf("- %s %s `%s`\n", meta.Icon.String(), meta.Name, meta.SoundID)
	}

	// log.Println(str)
	return str
}
