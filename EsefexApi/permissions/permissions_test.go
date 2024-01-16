package permissions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	assert.Equal(t, Allow, Deny.MergeParent(Allow))
	assert.Equal(t, Deny, Allow.MergeParent(Deny))
	assert.Equal(t, Allow, Allow.MergeParent(Allow))
	assert.Equal(t, Allow, Allow.MergeParent(Unset))

	assert.Equal(t, NewAllow(), NewAllow().MergeParent(NewAllow()))
	assert.Equal(t, NewAllow(), NewAllow().MergeParent(NewUnset()))
	assert.Equal(t, NewAllow(), NewUnset().MergeParent(NewAllow()))
	assert.Equal(t, NewDeny(), NewAllow().MergeParent(NewDeny()))
	assert.Equal(t, NewAllow(), NewDeny().MergeParent(NewAllow()))
	assert.Equal(t, NewDeny(), NewDeny().MergeParent(NewDeny()))
	assert.Equal(t, NewDeny(), NewDeny().MergeParent(NewUnset()))
	assert.Equal(t, NewUnset(), NewUnset().MergeParent(NewUnset()))

	bp1 := BotPermissions{
		Join:  Allow,
		Leave: Allow,
	}

	bp2 := BotPermissions{
		Join:  Deny,
		Leave: Deny,
	}

	assert.Equal(t, bp2, bp1.MergeParent(bp2))
	assert.Equal(t, bp1, bp2.MergeParent(bp1))

	assert.Equal(t, BotPermissions{
		Join:  Allow,
		Leave: Deny,
	}, BotPermissions{
		Join:  Allow,
		Leave: Unset,
	}.MergeParent(BotPermissions{
		Join:  Unset,
		Leave: Deny,
	}))

	p1 := NewUnset()
	p2 := NewUnset()

	p1.Sound.Play = Deny
	p1.Guild.ManageBot = Deny

	p2.Sound.Play = Allow

	merged := Permissions{
		Sound: SoundPermissions{
			Play:   Allow,
			Upload: Unset,
			Modify: Unset,
			Delete: Unset,
		},
		Bot: BotPermissions{
			Join:  Unset,
			Leave: Unset,
		},
		Guild: GuildPermissions{
			UseSlashCommands: Unset,
			ManageBot:        Deny,
			ManageUser:       Unset,
		},
	}

	assert.Equal(t, merged, p1.MergeParent(p2))
}
