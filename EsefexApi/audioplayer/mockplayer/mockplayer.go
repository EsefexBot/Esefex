package mockplayer

import (
	"esefexapi/audioplayer"
	"esefexapi/sounddb"
	"esefexapi/types"
	"log"
)

var _ audioplayer.IAudioPlayer = &MockPlayer{}

// MockPlayer is a mock implementation of the audioplayer.AudioPlayer interface.
// MockPlayer implements audioplayer.AudioPlayer.
type MockPlayer struct {
}

func NewMockPlayer() *MockPlayer {
	return &MockPlayer{}
}

// PlaySoundInsecure implements audioplayer.AudioPlayer.
func (*MockPlayer) PlaySoundInsecure(uid sounddb.SoundURI, guildID types.GuildID, userID types.UserID) error {
	log.Printf("MockPlayer: Playing sound insecurely '%v' on guild '%v' for user '%v'", uid, guildID, userID)
	return nil
}

func (*MockPlayer) PlaySound(soundID types.SoundID, userID types.UserID) error {
	log.Printf("MockPlayer: Playing sound '%v' for user '%v'", soundID, userID)
	return nil
}
