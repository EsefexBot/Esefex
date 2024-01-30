package audioplayer

import (
	"esefexapi/sounddb"

	"github.com/pkg/errors"
)

type IAudioPlayer interface {
	PlaySoundInsecure(uid sounddb.SoundUID, guildID string, userID string) error
	PlaySound(soundID string, userID string) error
}

var UserNotInVC = errors.New("User is not in a voice channel")
