package permissions

// Merge returns the permission state that results from merging two permission states.
// otherPS has precedence over ps.
func (ps PermissionState) Merge(otherPS PermissionState) PermissionState {
	if otherPS == Allow {
		return Allow
	} else if otherPS == Deny {
		return Deny
	} else {
		return ps
	}
}

func (p Permissions) Merge(otherP *Permissions) Permissions {
	return Permissions{
		Sound: p.Sound.Merge(otherP.Sound),
		Bot:   p.Bot.Merge(otherP.Bot),
		Guild: p.Guild.Merge(otherP.Guild),
	}
}

func (p SoundPermissions) Merge(otherP SoundPermissions) SoundPermissions {
	return SoundPermissions{
		Play:   p.Play.Merge(otherP.Play),
		Upload: p.Upload.Merge(otherP.Upload),
		Modify: p.Modify.Merge(otherP.Modify),
		Delete: p.Delete.Merge(otherP.Delete),
	}
}

func (p BotPermissions) Merge(otherP BotPermissions) BotPermissions {
	return BotPermissions{
		Join:  p.Join.Merge(otherP.Join),
		Leave: p.Leave.Merge(otherP.Leave),
	}
}

func (p GuildPermissions) Merge(otherP GuildPermissions) GuildPermissions {
	return GuildPermissions{
		ManageBot:  p.ManageBot.Merge(otherP.ManageBot),
		ManageUser: p.ManageUser.Merge(otherP.ManageUser),
	}
}
