package sounddb

func SuidFromStrings(serverId string, soundId string) SoundUID {
	return SoundUID{
		ServerID: serverId,
		SoundID:  soundId,
	}
}

func (sMeta SoundMeta) GetUID() SoundUID {
	return SoundUID{
		ServerID: sMeta.ServerID,
		SoundID:  sMeta.SoundID,
	}
}
