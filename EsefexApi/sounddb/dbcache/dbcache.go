package dbcache

import (
	"esefexapi/sounddb"
	"sync"
)

var _ sounddb.ISoundDB = &DBCache{}

// DB Cache loads all sounds into memory and caches them.
// DBCache implements db.SoundDB
type DBCache struct {
	sounds map[sounddb.SoundUID]*CachedSound
	db     sounddb.ISoundDB
	rw     sync.RWMutex
}

type CachedSound struct {
	Data []int16
	Meta sounddb.SoundMeta
}

// NewDBCache creates a new DBCache.
func NewDBCache(db sounddb.ISoundDB) *DBCache {
	c := &DBCache{
		sounds: make(map[sounddb.SoundUID]*CachedSound),
		db:     db,
	}
	c.CacheAll()
	return c
}

// AddSound implements db.SoundDB.
func (c *DBCache) AddSound(serverID string, name string, icon string, pcm []int16) (sounddb.SoundUID, error) {
	uid, err := c.db.AddSound(serverID, name, icon, pcm)
	if err != nil {
		return sounddb.SoundUID{}, err
	}

	c.rw.Lock()
	defer c.rw.Unlock()

	c.sounds[uid] = &CachedSound{
		Data: pcm,
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
func (c *DBCache) DeleteSound(uid sounddb.SoundUID) error {
	c.rw.RLock()
	defer c.rw.RUnlock()

	err := c.db.DeleteSound(uid)
	if err != nil {
		return err
	}

	c.rw.Lock()
	defer c.rw.Unlock()

	delete(c.sounds, uid)

	return nil
}

// GetServerIDs implements db.SoundDB.
func (c *DBCache) GetServerIDs() ([]string, error) {
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
func (c *DBCache) GetSoundMeta(uid sounddb.SoundUID) (sounddb.SoundMeta, error) {
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
func (c *DBCache) GetSoundPcm(uid sounddb.SoundUID) ([]int16, error) {
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
func (c *DBCache) GetSoundUIDs(serverID string) ([]sounddb.SoundUID, error) {
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

func (c *DBCache) CacheAll() error {
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

func (c *DBCache) LoadSound(uid sounddb.SoundUID) (*CachedSound, error) {
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