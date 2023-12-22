package apimockdb

import "esefexapi/sounddb"

var mockData = map[string]map[string]sounddb.SoundMeta{
	"server1": {
		"sound1": {
			SoundID:  "sound1",
			ServerID: "server1",
			Name:     "sound1Name",
			Icon:     "icon1",
		},
		"sound2": {
			SoundID:  "sound2",
			ServerID: "server1",
			Name:     "sound2Name",
			Icon:     "icon2",
		},
	},
	"server2": {
		"sound3": {
			SoundID:  "sound3",
			ServerID: "server2",
			Name:     "sound3Name",
			Icon:     "icon3",
		},
		"sound4": {
			SoundID:  "sound4",
			ServerID: "server2",
			Name:     "sound4Name",
			Icon:     "icon4",
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
func (*ApiMockDB) AddSound(serverID string, name string, icon string, pcm []int16) (sounddb.SoundUID, error) {
	panic("unimplemented")
}

// DeleteSound implements sounddb.ISoundDB.
func (*ApiMockDB) DeleteSound(uid sounddb.SoundUID) error {
	panic("unimplemented")
}

// GetServerIDs implements sounddb.ISoundDB.
func (*ApiMockDB) GetServerIDs() ([]string, error) {
	ids := make([]string, 0, len(mockData))
	for id := range mockData {
		ids = append(ids, id)
	}
	return ids, nil
}

// GetSoundMeta implements sounddb.ISoundDB.
func (*ApiMockDB) GetSoundMeta(uid sounddb.SoundUID) (sounddb.SoundMeta, error) {
	return mockData[uid.ServerID][uid.SoundID], nil
}

// GetSoundPcm implements sounddb.ISoundDB.
func (*ApiMockDB) GetSoundPcm(uid sounddb.SoundUID) ([]int16, error) {
	panic("unimplemented")
}

// GetSoundUIDs implements sounddb.ISoundDB.
func (*ApiMockDB) GetSoundUIDs(serverID string) ([]sounddb.SoundUID, error) {
	uids := make([]sounddb.SoundUID, 0, len(mockData[serverID]))
	for id := range mockData[serverID] {
		uids = append(uids, sounddb.SoundUID{
			ServerID: serverID,
			SoundID:  id,
		})
	}
	return uids, nil
}
