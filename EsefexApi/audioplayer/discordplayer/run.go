package discordplayer

import (
	"log"
)

func (p *DiscordPlayer) run() {
	defer close(p.stop)
	log.Println("Starting discord audio player...")
	defer log.Println("Discord audio player stopped")

	go p.afkKick()

	close(p.ready)
	<-p.stop
}
