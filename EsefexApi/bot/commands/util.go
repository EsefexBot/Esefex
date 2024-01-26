package commands

import (
	"esefexapi/db"
	"esefexapi/permissions"
	"esefexapi/sounddb"
	"esefexapi/types"
	"esefexapi/util/refl"
	"fmt"
	"regexp"
	"strings"

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
	regex, err := regexp.Compile(`^<(@|#|@&)(\d+)>$|^(@?everyone)$|^(\d+)$`)
	if err != nil {
		return PermissionSet{}, errors.Wrap(err, "Error compiling regex")
	}

	matches := regex.FindStringSubmatch(str)

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

	if matches[3] != "" {
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
	resp += fmt.Sprintf("Sound.Play:       %s\n", p.Sound.Play.String())
	resp += fmt.Sprintf("Sound.Upload:     %s\n", p.Sound.Upload.String())
	resp += fmt.Sprintf("Sound.Modify:     %s\n", p.Sound.Modify.String())
	resp += fmt.Sprintf("Sound.Delete:     %s\n", p.Sound.Delete.String())
	resp += "```"

	resp += "\n**Bot**\n"
	resp += fmt.Sprintf("```%s\n", mdlang)
	resp += fmt.Sprintf("Bot.Join:         %s\n", p.Bot.Join.String())
	resp += fmt.Sprintf("Bot.Leave:        %s\n", p.Bot.Leave.String())
	resp += "```"

	resp += "\n**Guild**\n"
	resp += fmt.Sprintf("```%s\n", mdlang)
	resp += fmt.Sprintf("Guild.BotManage:  %s\n", p.Guild.BotManage.String())
	resp += fmt.Sprintf("Guild.UserManage: %s\n", p.Guild.UserManage.String())
	resp += "```"

	return resp, nil
}

func formatPermissionsCompact(p permissions.Permissions) (string, error) {
	ppaths := refl.FindAllPaths(p)

	parts := make([]string, 0, len(ppaths))
	for _, ppath := range ppaths {
		ps, err := refl.GetNestedFieldValue(p, ppath)
		if err != nil {
			return "", errors.Wrap(err, "Error getting permission")
		}
		parts = append(parts, ps.(permissions.PermissionState).String())
	}

	return strings.Join(parts, "|"), nil
}

type PermissionSet struct {
	PermissionType permissions.PermissionType
	ID             string
}

func getPermissions(s *discordgo.Session, dbs *db.Databases, guildID types.GuildID, id string) (permissions.Permissions, error) {
	ty, err := extractTypeFromString(s, guildID, id)
	if err != nil {
		return permissions.Permissions{}, errors.Wrap(err, "Error extracting type from string")
	}
	var p permissions.Permissions

	switch ty.PermissionType {
	case permissions.User:
		p, err = dbs.PermissionDB.GetUser(guildID, types.UserID(ty.ID))
	case permissions.Role:
		p, err = dbs.PermissionDB.GetRole(guildID, types.RoleID(ty.ID))
	case permissions.Channel:
		p, err = dbs.PermissionDB.GetChannel(guildID, types.ChannelID(ty.ID))
	}

	if err != nil {
		return permissions.Permissions{}, errors.Wrap(err, "Error getting permissions")
	}

	return p, nil
}

func getPermission(p permissions.Permissions, key string) (permissions.PermissionState, error) {
	v, err := refl.GetNestedFieldValue(p, key)
	if err != nil {
		return permissions.Unset, errors.Wrap(err, "Error getting nested field value")
	}

	return v.(permissions.PermissionState), nil
}

func getPathOptions() []*discordgo.ApplicationCommandOptionChoice {
	util := refl.FindAllPaths(permissions.NewUnset())

	var options []*discordgo.ApplicationCommandOptionChoice

	for _, u := range util {
		options = append(options, &discordgo.ApplicationCommandOptionChoice{
			Name:  u,
			Value: u,
		})
	}

	return options
}
