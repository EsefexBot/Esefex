package audioplayer

import (
	"esefexapi/types"
	"fmt"
)

type IAudioPlayer interface {
	PlaySound(soundID types.SoundID, userID types.UserID) error
}

var UserNotInVC = fmt.Errorf("User is not in a voice channel")
