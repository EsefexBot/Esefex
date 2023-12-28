package dbcache

import (
	"esefexapi/sounddb"
	"sync"
)

var _ sounddb.ISoundDB = &SoundDBCache{}

// DB Cache loads all sounds into memory and caches them.
// SoundDBCache implements db.SoundDB
type SoundDBCache struct {
	sounds map[sounddb.SoundUID]*CachedSound
	db     sounddb.ISoundDB
	rw     sync.RWMutex
}

type CachedSound struct {
	Data *[]int16
	Meta sounddb.SoundMeta
}

// NewSoundDBCache creates a new DBCache.
func NewSoundDBCache(db sounddb.ISoundDB) *SoundDBCache {
	c := &SoundDBCache{
		sounds: make(map[sounddb.SoundUID]*CachedSound),
		db:     db,
	}
	c.CacheAll()
	return c
}

// AddSound implements db.SoundDB.
func (c *SoundDBCache) AddSound(serverID string, name string, icon sounddb.Icon, pcm []int16) (sounddb.SoundUID, error) {
	c.rw.Lock()
	defer c.rw.Unlock()

	uid, err := c.db.AddSound(serverID, name, icon, pcm)
	if err != nil {
		return sounddb.SoundUID{}, err
	}

	c.sounds[uid] = &CachedSound{
		Data: &pcm,
		Meta: sounddb.SoundMeta{
			SoundID:  uid.SoundID,
			ServerID: serverID,
			Name:     name,
			Icon:     icon,
		},
	}

	return uid, nil
}

// DeleteSound implements db.SoundDB.
func (c *SoundDBCache) DeleteSound(uid sounddb.SoundUID) error {
	c.rw.Lock()
	defer c.rw.Unlock()

	err := c.db.DeleteSound(uid)
	if err != nil {
		return err
	}

	delete(c.sounds, uid)

	return nil
}

// GetServerIDs implements db.SoundDB.
func (c *SoundDBCache) GetServerIDs() ([]string, error) {
	c.rw.RLock()
	defer c.rw.RUnlock()

	uniqueServerIDs := make(map[string]struct{})

	for uid := range c.sounds {
		uniqueServerIDs[uid.ServerID] = struct{}{}
	}

	serverIDs := make([]string, 0, len(uniqueServerIDs))
	for serverID := range uniqueServerIDs {
		serverIDs = append(serverIDs, serverID)
	}

	return serverIDs, nil
}

// GetSoundMeta implements db.SoundDB.
func (c *SoundDBCache) GetSoundMeta(uid sounddb.SoundUID) (sounddb.SoundMeta, error) {
	c.rw.RLock()
	defer c.rw.RUnlock()

	if sound, ok := c.sounds[uid]; ok {
		return sound.Meta, nil
	}

	c.rw.RUnlock()
	s, err := c.LoadSound(uid)
	if err != nil {
		return sounddb.SoundMeta{}, err
	}

	return s.Meta, nil
}

// GetSoundPcm implements db.SoundDB.
func (c *SoundDBCache) GetSoundPcm(uid sounddb.SoundUID) (*[]int16, error) {
	c.rw.RLock()
	defer c.rw.RUnlock()

	if sound, ok := c.sounds[uid]; ok {
		return sound.Data, nil
	}

	c.rw.RUnlock()
	s, err := c.LoadSound(uid)
	if err != nil {
		return nil, err
	}

	return s.Data, nil
}

// GetSoundUIDs implements db.SoundDB.
func (c *SoundDBCache) GetSoundUIDs(serverID string) ([]sounddb.SoundUID, error) {
	c.rw.RLock()
	defer c.rw.RUnlock()

	uids := make([]sounddb.SoundUID, 0)

	for uid := range c.sounds {
		if uid.ServerID == serverID {
			uids = append(uids, uid)
		}
	}

	return uids, nil
}

func (c *SoundDBCache) CacheAll() error {
	servers, err := c.db.GetServerIDs()
	if err != nil {
		return err
	}

	for _, serverID := range servers {
		uids, err := c.db.GetSoundUIDs(serverID)
		if err != nil {
			return err
		}

		for _, uid := range uids {
			_, err := c.LoadSound(uid)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *SoundDBCache) LoadSound(uid sounddb.SoundUID) (*CachedSound, error) {
	c.rw.Lock()
	defer c.rw.Unlock()

	if sound, ok := c.sounds[uid]; ok {
		return sound, nil
	}

	pcm, err := c.db.GetSoundPcm(uid)
	if err != nil {
		return nil, err
	}

	meta, err := c.db.GetSoundMeta(uid)
	if err != nil {
		return nil, err
	}

	s := CachedSound{
		Data: pcm,
		Meta: meta,
	}

	c.sounds[uid] = &s

	return &s, nil

}

func (c *SoundDBCache) SoundExists(uid sounddb.SoundUID) (bool, error) {
	c.rw.RLock()
	defer c.rw.RUnlock()

	if _, ok := c.sounds[uid]; ok {
		return true, nil
	}

	return c.db.SoundExists(uid)
}
