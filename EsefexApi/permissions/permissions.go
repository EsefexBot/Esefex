package permissions

type PermissionType int

const (
	User PermissionType = iota
	Role
	Channel
)

func (pt PermissionType) String() string {
	switch pt {
	case User:
		return "user"
	case Role:
		return "role"
	case Channel:
		return "channel"
	default:
		return "unknown"
	}
}

type PermissionState int

const (
	Allow PermissionState = iota
	Deny
	Unset
)

func (ps PermissionState) String() string {
	switch ps {
	case Allow:
		return "allow"
	case Deny:
		return "deny"
	case Unset:
		return "unset"
	default:
		return "unknown"
	}
}

func (ps PermissionState) Allowed() bool {
	return ps == Allow
}

type Permissions struct {
	Sound SoundPermissions
	Bot   BotPermissions
	Guild GuildPermissions
}

type SoundPermissions struct {
	Play   PermissionState
	Upload PermissionState
	Modify PermissionState
	Delete PermissionState
}

type BotPermissions struct {
	Join  PermissionState
	Leave PermissionState
}

type GuildPermissions struct {
	ManageBot  PermissionState
	ManageUser PermissionState
}

// Default returns a Permissions struct with all permissions set to Allow.
func NewAllow() Permissions {
	return Permissions{
		Sound: SoundPermissions{
			Play:   Allow,
			Upload: Allow,
			Modify: Allow,
			Delete: Allow,
		},
		Bot: BotPermissions{
			Join:  Allow,
			Leave: Allow,
		},
		Guild: GuildPermissions{
			ManageBot:  Allow,
			ManageUser: Allow,
		},
	}
}

func NewUnset() Permissions {
	return Permissions{
		Sound: SoundPermissions{
			Play:   Unset,
			Upload: Unset,
			Modify: Unset,
			Delete: Unset,
		},
		Bot: BotPermissions{
			Join:  Unset,
			Leave: Unset,
		},
		Guild: GuildPermissions{
			ManageBot:  Unset,
			ManageUser: Unset,
		},
	}
}

func NewDeny() Permissions {
	return Permissions{
		Sound: SoundPermissions{
			Play:   Deny,
			Upload: Deny,
			Modify: Deny,
			Delete: Deny,
		},
		Bot: BotPermissions{
			Join:  Deny,
			Leave: Deny,
		},
		Guild: GuildPermissions{
			ManageBot:  Deny,
			ManageUser: Deny,
		},
	}
}
