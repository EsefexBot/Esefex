package filepermisssiondb

import (
	"esefexapi/opt"
	"esefexapi/permissions"
	"esefexapi/types"
	"esefexapi/util/dcgoutil"

	"github.com/pkg/errors"
)

// Query implements permissiondb.PermissionDB.
func (f *FilePermissionDB) Query(user types.UserID, guild types.GuildID) (permissions.Permissions, error) {
	f.rw.RLock()
	defer f.rw.RUnlock()

	m, err := f.ds.GuildMember(guild.String(), user.String())
	if err != nil {
		return permissions.Permissions{}, errors.Wrap(err, "Error getting guild member")
	}

	roles := make([]types.RoleID, len(m.Roles))
	for i, r := range m.Roles {
		roles[i] = types.RoleID(r)
	}

	userChannel, err := dcgoutil.UserGuildVC(f.ds, guild, user)
	if err != nil {
		return permissions.Permissions{}, errors.Wrap(err, "Error getting user channel")
	}

	if userChannel.IsNone() {
		return f.stack.Query(user, roles, opt.None[types.ChannelID]()), nil
	} else {
		chanID := types.ChannelID(userChannel.Unwrap().ChannelID)
		return f.stack.Query(user, roles, opt.Some(chanID)), nil
	}
}
