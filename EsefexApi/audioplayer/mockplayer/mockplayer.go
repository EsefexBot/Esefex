package mockplayer

import (
	"esefexapi/audioplayer"
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

func (*MockPlayer) PlaySound(soundID types.SoundID, userID types.UserID) error {
	log.Printf("MockPlayer: Playing sound '%v' for user '%v'", soundID, userID)
	return nil
}
