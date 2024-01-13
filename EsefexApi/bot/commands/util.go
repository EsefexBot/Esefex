package commands

import (
	"esefexapi/permissions"
	"esefexapi/sounddb"
	"esefexapi/types"
	"fmt"
	"log"
	"regexp"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

func fmtMetaList(metas []sounddb.SoundMeta) string {
	// log.Printf("fmtMetaList: %v", metas)
	var str string
	for _, meta := range metas {
		str += fmt.Sprintf("- %s %s `%s`\n", meta.Icon.String(), meta.Name, meta.SoundID)
	}

	// log.Println(str)
	return str
}

// checks if its a user, role or channel
func extractTypeFromString(s *discordgo.Session, g types.GuildID, str string) (PermissionSet, error) {
	regex, err := regexp.Compile(`^<(@|#|@&)(\d+)>$|^(@everyone)$|^(\d+)$`)
	if err != nil {
		return PermissionSet{}, errors.Wrap(err, "Error compiling regex")
	}

	matches := regex.FindStringSubmatch(str)
	log.Printf("%d matches: %#v", len(matches), matches)

	if len(matches) != 5 {
		return PermissionSet{}, errors.Wrap(fmt.Errorf("Invalid id %s", str), "Error extracting type from string")
	}

	switch matches[1] {
	case "@":
		return PermissionSet{
			PermissionType: permissions.User,
			ID:             matches[2],
		}, nil
	case "#":
		return PermissionSet{
			PermissionType: permissions.Channel,
			ID:             matches[2],
		}, nil
	case "@&":
		return PermissionSet{
			PermissionType: permissions.Role,
			ID:             matches[2],
		}, nil
	}

	if matches[3] == "@everyone" {
		return PermissionSet{
			PermissionType: permissions.Role,
			ID:             "everyone",
		}, nil
	}

	id := matches[4]

	_, err = s.User(str)
	if err == nil {
		return PermissionSet{
			PermissionType: permissions.User,
			ID:             str,
		}, nil
	}

	roles, err := s.GuildRoles(g.String())
	if err != nil {
		return PermissionSet{}, errors.Wrap(err, "Error getting guild roles")
	}

	for _, r := range roles {
		if r.ID == id {
			return PermissionSet{
				PermissionType: permissions.Role,
				ID:             id,
			}, nil
		}
	}

	channels, err := s.GuildChannels(g.String())
	if err != nil {
		return PermissionSet{}, errors.Wrap(err, "Error getting guild channels")
	}

	for _, c := range channels {
		if c.ID == id {
			return PermissionSet{
				PermissionType: permissions.Channel,
				ID:             id,
			}, nil
		}
	}

	return PermissionSet{}, errors.Wrap(fmt.Errorf("Invalid id %s", str), "Error extracting type from string")
}

func formatPermissions(p permissions.Permissions) (string, error) {
	// b, err := toml.Marshal(p)
	// if err != nil {
	// 	return "", errors.Wrap(err, "Error marshalling permissions")
	// }

	// resp := string(b)
	// log.Println(resp)

	mdlang := "yaml"

	resp := "**Sound**\n"
	resp += fmt.Sprintf("```%s\n", mdlang)
	resp += fmt.Sprintf("Play:       %s\n", p.Sound.Play.String())
	resp += fmt.Sprintf("Upload:     %s\n", p.Sound.Upload.String())
	resp += fmt.Sprintf("Modify:     %s\n", p.Sound.Modify.String())
	resp += fmt.Sprintf("Delete:     %s\n", p.Sound.Delete.String())
	resp += "```"

	resp += "\n**Bot**\n"
	resp += fmt.Sprintf("```%s\n", mdlang)
	resp += fmt.Sprintf("Join:       %s\n", p.Bot.Join.String())
	resp += fmt.Sprintf("Leave:      %s\n", p.Bot.Leave.String())
	resp += "```"

	resp += "\n**Guild**\n"
	resp += fmt.Sprintf("```%s\n", mdlang)
	resp += fmt.Sprintf("ManageBot:  %s\n", p.Guild.ManageBot.String())
	resp += fmt.Sprintf("ManageUser: %s\n", p.Guild.ManageUser.String())
	resp += "```"

	return resp, nil
}

type PermissionSet struct {
	PermissionType permissions.PermissionType
	ID             string
}
