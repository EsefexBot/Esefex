package types

type UserID string

func (u UserID) String() string {
	return string(u)
}

type RoleID string

func (r RoleID) String() string {
	return string(r)
}

type ChannelID string

func (c ChannelID) String() string {
	return string(c)
}

type GuildID string

func (g GuildID) String() string {
	return string(g)
}

type SoundID string

func (s SoundID) String() string {
	return string(s)
}
