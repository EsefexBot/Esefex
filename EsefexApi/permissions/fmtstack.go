package permissions

import (
	"bytes"
	"esefexapi/types"
	"esefexapi/util"
	"esefexapi/util/dcgoutil"
	"esefexapi/util/refl"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/jedib0t/go-pretty/v6/table"
)

func (ps *PermissionStack) FmtStack(ds *discordgo.Session, guildID types.GuildID) (string, error) {
	t := table.NewWriter()
	buf := new(bytes.Buffer)
	t.SetOutputMirror(buf)

	paths := refl.FindAllPaths(Permissions{})

	pps := []interface{}{}
	for _, p := range shortenPaths(paths) {
		pps = append(pps, p)
	}

	header := append(table.Row{"#"}, pps...)
	t.AppendHeader(header)

	t.AppendSeparator()
	t.AppendRow(table.Row{"USERS"})
	for k, v := range ps.Users {
		uname, err := dcgoutil.UserIDName(ds, k)
		if err != nil {
			return "", err
		}
		row := table.Row{uname}
		for _, p := range paths {
			ps, err := refl.GetNestedFieldValue(v, p)
			if err != nil {
				return "", err
			}

			row = append(row, ps.(PermissionState).String())
		}
		t.AppendRow(row)
	}

	t.AppendSeparator()
	t.AppendRow(table.Row{"ROLES"})
	for k, v := range ps.Roles {
		rname, err := dcgoutil.RoleIDName(ds, guildID, k)
		if err != nil {
			return "", err
		}
		row := table.Row{rname}
		for _, p := range paths {
			ps, err := refl.GetNestedFieldValue(v, p)
			if err != nil {
				return "", err
			}

			row = append(row, ps.(PermissionState).String())
		}
		t.AppendRow(row)
	}

	t.AppendSeparator()
	t.AppendRow(table.Row{"CHANNELS"})
	for k, v := range ps.Channels {
		cname, err := dcgoutil.ChannelIDName(ds, guildID, k)
		if err != nil {
			return "", err
		}
		row := table.Row{cname}
		for _, p := range paths {
			ps, err := refl.GetNestedFieldValue(v, p)
			if err != nil {
				return "", err
			}

			row = append(row, ps.(PermissionState).String())
		}
		t.AppendRow(row)
	}

	t.SetStyle(table.StyleRounded)
	t.Render()
	return buf.String(), nil
}

var segLens []int = []int{1, 15}
var segLenDefault = 3

func shortenPaths(paths []string) []string {
	shortened := []string{}
	for _, p := range paths {
		pathSegs := []string{}

		for i, spl := range strings.Split(p, ".") {
			if i >= len(segLens) {
				pathSegs = append(pathSegs, util.FirstNRunes(spl, segLenDefault))
				continue
			}

			pathSegs = append(pathSegs, util.FirstNRunes(spl, segLens[i]))
		}

		shortened = append(shortened, strings.Join(pathSegs, "."))
	}
	return shortened
}
