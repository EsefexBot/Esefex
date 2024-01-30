package filesounddb

import (
	"esefexapi/config"
	"esefexapi/sounddb"
	"esefexapi/types"
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileDB(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	icon := sounddb.Icon{
		Name: "icon1",
		ID:   "icon1",
		Url:  "https://raw.githubusercontent.com/EsefexBot/Esefex/main/EsefexApi/test/staticfiles/icon.webp",
	}

	guildID := types.GuildID("guild1")
	soundName := "sound1"
	soundPcm := []int16{115, 117, 115}

	config.InjectConfig(&config.Config{
		Database: config.Database{
			SounddbLocation: fmt.Sprintf("./dbtest_%d", rand.Intn(1000000)),
		},
	})

	db, err := NewFileDB()
	assert.Nil(t, err)

	// Test that we can add a sound
	uid, err := db.AddSound(guildID, types.SoundName(soundName), icon, soundPcm)
	assert.Nil(t, err)

	path := fmt.Sprintf("%s/%s/%s_meta.json", config.Get().Database.SounddbLocation, guildID, uid.SoundName.GetSoundID())
	_, err = os.Stat(path)
	assert.Nil(t, err)
	_, err = os.Stat(fmt.Sprintf("%s/%s/%s_sound.s16le", config.Get().Database.SounddbLocation, guildID, uid.SoundName.GetSoundID()))
	assert.Nil(t, err)

	// Test that the sound exists
	exists, err := db.SoundExists(uid)
	assert.Nil(t, err)
	assert.True(t, exists)

	// Test that we can get the sound
	sound, err := db.GetSoundMeta(uid)
	assert.Nil(t, err)
	assert.Equal(t, sounddb.SoundMeta{
		SoundID: uid.SoundName.GetSoundID(),
		GuildID: guildID,
		Name:    types.SoundName(soundName),
		Icon:    icon,
	}, sound)

	// Test that we can get the sound pcm
	soundPcm2, err := db.GetSoundPcm(uid)
	assert.Nil(t, err)
	assert.Equal(t, &soundPcm, soundPcm2)

	// Test that we can get the guild ids
	ids, err := db.GetGuildIDs()
	assert.Nil(t, err)
	assert.Equal(t, []types.GuildID{guildID}, ids)

	// Test that we can get the sound uids
	uids, err := db.GetSoundUIDs(guildID)
	assert.Nil(t, err)
	assert.Equal(t, []sounddb.SoundUID{uid}, uids)

	// Test that we can delete the sound
	err = db.DeleteSound(uid)
	assert.Nil(t, err)

	// Test that the sound doesn't exist
	exists, err = db.SoundExists(uid)
	assert.Nil(t, err)
	assert.False(t, exists)

	// Test that we can't get the sound
	_, err = db.GetSoundMeta(uid)
	assert.NotNil(t, err)

	// Test that we can't get the sound pcm
	_, err = db.GetSoundPcm(uid)
	assert.NotNil(t, err)

	// delete the db folder location
	err = os.RemoveAll(config.Get().Database.SounddbLocation)
	assert.Nil(t, err)
}
