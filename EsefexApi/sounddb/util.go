package sounddb

import "esefexapi/types"

func SuidFromStrings(guildID string, soundID string) SoundURI {
	return SoundURI{
		GuildID: types.GuildID(guildID),
		SoundID: types.SoundID(soundID),
	}
}

func New(guildID types.GuildID, soundID types.SoundID) SoundURI {
	return SoundURI{
		GuildID: guildID,
		SoundID: soundID,
	}
}

func (sMeta SoundMeta) GetUID() SoundURI {
	return SoundURI{
		GuildID: sMeta.GuildID,
		SoundID: sMeta.SoundID,
	}
}
