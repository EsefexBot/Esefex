package sounddb

import "esefexapi/types"

type ISoundDB interface {
	AddSound(guildID types.GuildID, name string, icon Icon, pcm []int16) (SoundURI, error)
	DeleteSound(uid SoundURI) error
	GetSoundMeta(uid SoundURI) (SoundMeta, error)
	GetSoundPcm(uid SoundURI) (*[]int16, error)
	GetSoundUIDs(guildID types.GuildID) ([]SoundURI, error)
	GetGuildIDs() ([]types.GuildID, error)
	SoundExists(uid SoundURI) (bool, error)
}

type SoundURI struct {
	GuildID types.GuildID
	SoundID types.SoundID
}

type SoundMeta struct {
	SoundID types.SoundID `json:"id"`
	GuildID types.GuildID `json:"guildId"`
	Name    string        `json:"name"`
	Icon    Icon          `json:"icon"`
}
