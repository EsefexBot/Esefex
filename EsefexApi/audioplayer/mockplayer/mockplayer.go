package mockplayer

import (
	"esefexapi/audioplayer"
	"esefexapi/sounddb"
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

// PlaySound implements audioplayer.AudioPlayer.
func (*MockPlayer) PlaySound(uid sounddb.SoundUID, serverID string, userID string) error {
	log.Printf("MockPlayer: Playing sound '%v' on server '%v' for user '%v'", uid, serverID, userID)
	return nil
}
