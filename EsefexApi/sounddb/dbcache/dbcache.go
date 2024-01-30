package dbcache

import (
	"esefexapi/sounddb"
	"esefexapi/types"
	"fmt"
	"sync"

	"github.com/pkg/errors"
)

var _ sounddb.ISoundDB = &SoundDBCache{}

// DB Cache loads all sounds into memory and caches them.
// SoundDBCache implements db.SoundDB
type SoundDBCache struct {
	sounds map[sounddb.SoundUID]*CachedSound
	db     sounddb.ISoundDB
	rw     sync.RWMutex
}

// GetSoundNameByID implements sounddb.ISoundDB.
func (c *SoundDBCache) GetSoundNameByID(guildID types.GuildID, ID types.SoundID) (types.SoundName, error) {
	c.rw.RLock()
	defer c.rw.RUnlock()

	for _, sound := range c.sounds {
		if sound.Meta.SoundID == ID && sound.Meta.GuildID == guildID {
			return sound.Meta.Name, nil
		}
	}

	return "", errors.Wrap(fmt.Errorf("Sound with ID '%s' not found", ID), "Error getting sound name by ID")
}

type CachedSound struct {
	Data *[]int16
	Meta sounddb.SoundMeta
}

// NewSoundDBCache creates a new DBCache.
func NewSoundDBCache(db sounddb.ISoundDB) (*SoundDBCache, error) {
	c := &SoundDBCache{
		sounds: make(map[sounddb.SoundUID]*CachedSound),
		db:     db,
	}
	err := c.CacheAll()
	if err != nil {
		return nil, errors.Wrap(err, "Error caching all sounds")
	}

	return c, nil
}

// AddSound implements db.SoundDB.
func (c *SoundDBCache) AddSound(guildID types.GuildID, name types.SoundName, icon sounddb.Icon, pcm []int16) (sounddb.SoundUID, error) {
	c.rw.Lock()
	defer c.rw.Unlock()

	uid, err := c.db.AddSound(guildID, name, icon, pcm)
	if err != nil {
		return sounddb.SoundUID{}, errors.Wrap(err, "Error adding sound")
	}

	bitrate := 48000
	channels := 2

	c.sounds[uid] = &CachedSound{
		Data: &pcm,
		Meta: sounddb.SoundMeta{
			SoundID: uid.SoundName.GetSoundID(),
			GuildID: guildID,
			Name:    name,
			Icon:    icon,
			Length:  float32(len(pcm)) / float32(bitrate*channels),
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
		return errors.Wrap(err, "Error deleting sound")
	}

	delete(c.sounds, uid)

	return nil
}

// GetGuildIDs implements db.SoundDB.
func (c *SoundDBCache) GetGuildIDs() ([]types.GuildID, error) {
	c.rw.RLock()
	defer c.rw.RUnlock()

	uniqueGuildIDs := make(map[types.GuildID]struct{})

	for uid := range c.sounds {
		uniqueGuildIDs[uid.GuildID] = struct{}{}
	}

	guildIDs := make([]types.GuildID, 0, len(uniqueGuildIDs))
	for guildID := range uniqueGuildIDs {
		guildIDs = append(guildIDs, guildID)
	}

	return guildIDs, nil
}

// GetSoundMeta implements db.SoundDB.
func (c *SoundDBCache) GetSoundMeta(uid sounddb.SoundUID) (sounddb.SoundMeta, error) {
	c.rw.RLock()

	if sound, ok := c.sounds[uid]; ok {
		c.rw.RUnlock()
		return sound.Meta, nil
	}
	c.rw.RUnlock()
	s, err := c.LoadSound(uid)

	c.rw.RLock()
	defer c.rw.RUnlock()
	if err != nil {
		return sounddb.SoundMeta{}, errors.Wrap(err, "Error loading sound")
	}

	return s.Meta, nil
}

// GetSoundPcm implements db.SoundDB.
func (c *SoundDBCache) GetSoundPcm(uid sounddb.SoundUID) (*[]int16, error) {
	c.rw.RLock()

	if sound, ok := c.sounds[uid]; ok {
		c.rw.RUnlock()
		return sound.Data, nil
	}

	c.rw.RUnlock()

	s, err := c.LoadSound(uid)

	c.rw.RLock()
	defer c.rw.RUnlock()
	if err != nil {
		return nil, errors.Wrap(err, "Error loading sound")
	}

	return s.Data, nil
}

// GetSoundUIDs implements db.SoundDB.
func (c *SoundDBCache) GetSoundUIDs(guildID types.GuildID) ([]sounddb.SoundUID, error) {
	c.rw.RLock()
	defer c.rw.RUnlock()

	uids := make([]sounddb.SoundUID, 0)

	for uid := range c.sounds {
		if uid.GuildID == guildID {
			uids = append(uids, uid)
		}
	}

	return uids, nil
}

func (c *SoundDBCache) CacheAll() error {
	guilds, err := c.db.GetGuildIDs()
	if err != nil {
		return errors.Wrap(err, "Error getting guild ids")
	}

	for _, guildID := range guilds {
		uids, err := c.db.GetSoundUIDs(guildID)
		if err != nil {
			return errors.Wrap(err, "Error getting sound uids")
		}

		for _, uid := range uids {
			_, err := c.LoadSound(uid)
			if err != nil {
				return errors.Wrap(err, "Error loading sound")
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
		return nil, errors.Wrap(err, "Error getting sound pcm")
	}

	meta, err := c.db.GetSoundMeta(uid)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting sound meta")
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
