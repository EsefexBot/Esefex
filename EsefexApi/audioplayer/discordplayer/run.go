package discordplayer

func (p *DiscordPlayer) run() {
	defer close(p.stop)

	close(p.ready)
	<-p.stop
}
