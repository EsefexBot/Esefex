package sounddb

import "esefexapi/types"

type ISoundDB interface {
	AddSound(guildID types.GuildID, name types.SoundName, icon Icon, pcm []int16) (SoundUID, error)
	DeleteSound(uid SoundUID) error
	GetSoundMeta(uid SoundUID) (SoundMeta, error)
	GetSoundPcm(uid SoundUID) (*[]int16, error)
	GetSoundUIDs(guildID types.GuildID) ([]SoundUID, error)
	GetGuildIDs() ([]types.GuildID, error)
	SoundExists(uid SoundUID) (bool, error)
	GetSoundNameByID(guildID types.GuildID, ID types.SoundID) (types.SoundName, error)
}

type SoundUID struct {
	GuildID   types.GuildID
	SoundName types.SoundName
}

type SoundMeta struct {
	SoundID types.SoundID   `json:"id"`
	GuildID types.GuildID   `json:"guildId"`
	Name    types.SoundName `json:"name"`
	Icon    Icon            `json:"icon"`
	Length  float32         `json:"length"`
	Created int64           `json:"created"`
}
