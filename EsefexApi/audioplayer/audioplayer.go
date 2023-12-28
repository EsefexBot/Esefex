package audioplayer

import "esefexapi/sounddb"

type IAudioPlayer interface {
	PlaySoundInsecure(uid sounddb.SoundUID, guildID string, userID string) error
	PlaySound(soundID string, userID string) error
}
