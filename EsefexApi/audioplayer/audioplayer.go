package audioplayer

import (
	"esefexapi/sounddb"
	"fmt"
)

type IAudioPlayer interface {
	PlaySoundInsecure(uid sounddb.SoundUID, guildID string, userID string) error
	PlaySound(soundID string, userID string) error
}

var UserNotInVC = fmt.Errorf("User is not in a voice channel")
