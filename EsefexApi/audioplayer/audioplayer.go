package audioplayer

import (
	"esefexapi/sounddb"
	"esefexapi/types"
	"fmt"
)

type IAudioPlayer interface {
	PlaySoundInsecure(uid sounddb.SoundURI, guildID types.GuildID, userID types.UserID) error
	PlaySound(soundID types.SoundID, userID types.UserID) error
}

var UserNotInVC = fmt.Errorf("User is not in a voice channel")
