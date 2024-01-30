package sounddb

import "esefexapi/types"

func (sMeta SoundMeta) GetUID() SoundUID {
	return SoundUID{
		GuildID:   sMeta.GuildID,
		SoundName: types.SoundName(sMeta.Name),
	}
}
