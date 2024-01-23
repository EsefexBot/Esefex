package discordplayer

func (p *DiscordPlayer) Start() <-chan struct{} {
	go p.run()
	return p.ready
}

func (p *DiscordPlayer) Stop() <-chan struct{} {
	p.stop <- struct{}{}
	return p.stop
}
