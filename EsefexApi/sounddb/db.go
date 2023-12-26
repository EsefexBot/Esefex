package sounddb

type ISoundDB interface {
	AddSound(serverID string, name string, icon Icon, pcm []int16) (SoundUID, error)
	DeleteSound(uid SoundUID) error
	GetSoundMeta(uid SoundUID) (SoundMeta, error)
	GetSoundPcm(uid SoundUID) ([]int16, error)
	GetSoundUIDs(serverID string) ([]SoundUID, error)
	GetServerIDs() ([]string, error)
	SoundExists(uid SoundUID) (bool, error)
}

type SoundUID struct {
	ServerID string
	SoundID  string
}

type SoundMeta struct {
	SoundID  string `json:"id"`
	ServerID string `json:"serverId"`
	Name     string `json:"name"`
	Icon     Icon   `json:"icon"`
}
