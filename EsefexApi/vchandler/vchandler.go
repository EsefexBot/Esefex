package vchandler

import (
	"esefexapi/sfxplayer"
	"time"
)

var (
	useTimeouts    = false
	sessionTimeout = 5 * time.Minute
)

type ConnectionHandler struct {
	s []*sfxplayer.SfxPlayer
}

func NewConnectionHandler() *ConnectionHandler {
	return &ConnectionHandler{
		s: make([]*sfxplayer.SfxPlayer, 0),
	}
}

func (c *ConnectionHandler) PlaySound()
