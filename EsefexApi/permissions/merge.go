package permissions

// MergeParent returns the permission state that results from merging two permission states.
// otherPS has precedence over ps.
func (ps PermissionState) MergeParent(otherPS PermissionState) PermissionState {
	if otherPS == Allow {
		return Allow
	} else if otherPS == Deny {
		return Deny
	} else {
		return ps
	}
}

func (p Permissions) MergeParent(otherP Permissions) Permissions {
	return Permissions{
		Sound: p.Sound.MergeParent(otherP.Sound),
		Bot:   p.Bot.MergeParent(otherP.Bot),
		Guild: p.Guild.MergeParent(otherP.Guild),
	}
}

func (p SoundPermissions) MergeParent(otherP SoundPermissions) SoundPermissions {
	return SoundPermissions{
		Play:   p.Play.MergeParent(otherP.Play),
		Upload: p.Upload.MergeParent(otherP.Upload),
		Modify: p.Modify.MergeParent(otherP.Modify),
		Delete: p.Delete.MergeParent(otherP.Delete),
	}
}

func (p BotPermissions) MergeParent(otherP BotPermissions) BotPermissions {
	return BotPermissions{
		Join:  p.Join.MergeParent(otherP.Join),
		Leave: p.Leave.MergeParent(otherP.Leave),
	}
}

func (p GuildPermissions) MergeParent(otherP GuildPermissions) GuildPermissions {
	return GuildPermissions{
		ManageBot:  p.ManageBot.MergeParent(otherP.ManageBot),
		ManageUser: p.ManageUser.MergeParent(otherP.ManageUser),
	}
}
