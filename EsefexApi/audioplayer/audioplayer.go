package audioplayer

import "esefexapi/sounddb"

type IAudioPlayer interface {
	PlaySound(uid sounddb.SoundUID, userID string) error
}
