package permissions

import (
	"esefexapi/opt"
	"esefexapi/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := NewPermissionStack()

	u1 := types.UserID(rune(1))

	s.SetUser(u1, NewAllow())
	assert.Equal(t, NewAllow(), s.GetUser(u1))

	u2 := types.UserID(rune(2))
	up2 := Permissions{
		Sound: SoundPermissions{
			Play: Allow,
		},
	}
	s.SetUser(u2, up2)

	assert.Equal(t, up2, s.GetUser(u2))

	r1 := types.RoleID(rune(1))
	rp1 := Permissions{
		Sound: SoundPermissions{
			Upload: Allow,
			Delete: Allow,
		},
	}

	s.SetRole(r1, NewAllow())
	assert.Equal(t, NewAllow(), s.GetRole(r1))

	assert.Equal(t, NewUnset().MergeParent(rp1).MergeParent(up2), s.Query(u2, []types.RoleID{r1}, opt.None[types.ChannelID]()))
}
