package audioprocessing

import (
	"esefexapi/db"
	"esefexapi/db/filedb"
	"log"
	"sync"
)

type AudioCache struct {
	sounds map[db.SoundUID][]byte
	rw     sync.RWMutex
}

func NewAudioCache() *AudioCache {
	return &AudioCache{
		sounds: make(map[db.SoundUID][]byte),
	}
}

func (a *AudioCache) GetSound(uid db.SoundUID) ([]byte, error) {
	a.rw.RLock()
	defer a.rw.RUnlock()

	if sound, ok := a.sounds[uid]; ok {
		return sound, nil
	}

	err := a.LoadSound(uid)
	if err != nil {
		return nil, err
	}

	return a.sounds[uid], nil
}

func (a *AudioCache) LoadSound(uid db.SoundUID) error {
	a.rw.Lock()
	defer a.rw.Unlock()

	sb, err := filedb.LoadSoundBytes(uid)
	if err != nil {
		log.Println(err)
		return err
	}

	a.sounds[uid] = sb
	return nil
}

func (a *AudioCache) LoadAll() {
	a.rw.Lock()
	defer a.rw.Unlock()

	for _, uid := range filedb.GetAllSoundUids() {
		a.LoadSound(uid)
	}
}
