package apimockdb

import (
	"esefexapi/sounddb"
	"esefexapi/types"
)

var mockData = map[string]map[string]sounddb.SoundMeta{
	"guild1": {
		"sound1": {
			SoundID: "sound1",
			GuildID: "guild1",
			Name:    "sound1Name",
			Icon: sounddb.Icon{
				Name: "icon1",
				ID:   "icon1ID",
				Url:  "https://icon1Url.webp",
			},
		},
		"sound2": {
			SoundID: "sound2",
			GuildID: "guild1",
			Name:    "sound2Name",
			Icon: sounddb.Icon{
				Name: "icon2",
				ID:   "icon2ID",
				Url:  "https://icon2Url.webp",
			},
		},
	},
	"guild2": {
		"sound3": {
			SoundID: "sound3",
			GuildID: "guild2",
			Name:    "sound3Name",
			Icon: sounddb.Icon{
				Name: "icon3",
				ID:   "icon3ID",
				Url:  "https://icon3Url.webp",
			},
		},
		"sound4": {
			SoundID: "sound4",
			GuildID: "guild2",
			Name:    "sound4Name",
			Icon: sounddb.Icon{
				Name: "icon4",
				ID:   "icon4ID",
				Url:  "https://icon4Url.webp",
			},
		},
	},
}

var _ sounddb.ISoundDB = &ApiMockDB{}

// ApiMockDB implements ISoundDB
type ApiMockDB struct{}

func NewApiMockDB() *ApiMockDB {
	return &ApiMockDB{}
}

// AddSound implements sounddb.ISoundDB.
func (*ApiMockDB) AddSound(guildID types.GuildID, name string, icon sounddb.Icon, pcm []int16) (sounddb.SoundURI, error) {
	panic("unimplemented")
}

// DeleteSound implements sounddb.ISoundDB.
func (*ApiMockDB) DeleteSound(uid sounddb.SoundURI) error {
	panic("unimplemented")
}

// GetGuildIDs implements sounddb.ISoundDB.
func (*ApiMockDB) GetGuildIDs() ([]types.GuildID, error) {
	ids := make([]types.GuildID, 0, len(mockData))
	for id := range mockData {
		ids = append(ids, types.GuildID(id))
	}
	return ids, nil
}

// GetSoundMeta implements sounddb.ISoundDB.
func (*ApiMockDB) GetSoundMeta(uid sounddb.SoundURI) (sounddb.SoundMeta, error) {
	return mockData[uid.GuildID.String()][uid.SoundID.String()], nil
}

// GetSoundPcm implements sounddb.ISoundDB.
func (*ApiMockDB) GetSoundPcm(uid sounddb.SoundURI) (*[]int16, error) {
	panic("unimplemented")
}

// GetSoundUIDs implements sounddb.ISoundDB.
func (*ApiMockDB) GetSoundUIDs(guildID types.GuildID) ([]sounddb.SoundURI, error) {
	uids := make([]sounddb.SoundURI, 0, len(mockData[guildID.String()]))
	for id := range mockData[guildID.String()] {
		uids = append(uids, sounddb.SoundURI{
			GuildID: guildID,
			SoundID: types.SoundID(id),
		})
	}
	return uids, nil
}

// SoundExists implements sounddb.ISoundDB.
func (*ApiMockDB) SoundExists(uid sounddb.SoundURI) (bool, error) {
	_, ok := mockData[uid.GuildID.String()][uid.SoundID.String()]
	return ok, nil
}
