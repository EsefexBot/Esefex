package discordplayer

import (
	"log"
	"time"
)

// TODO: better error handling
func (p *DiscordPlayer) afkKick() {
	for {
		time.Sleep(1 * time.Second)

		if !p.useTimeouts {
			continue
		}

		for _, vd := range p.vds {
			if vd.AfkTimeoutIn.After(time.Now()) {
				continue
			}

			if vd.vcon.IsPlaying() {
				vd.AfkTimeoutIn = time.Now().Add(p.timeout)
			} else {
				log.Printf("Kicking bot from %s", vd.ChannelID)
				err := p.UnregisterVcon(vd.ChannelID)
				if err != nil {
					log.Printf("Error unregistering vcon: %+v", err)
				}
			}
		}
	}
}
