package audioplayer

import "esefexapi/sounddb"

type IAudioPlayer interface {
	PlaySound(uid sounddb.SoundUID, guildID string, userID string) error
}
