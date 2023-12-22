package discordplayer

import "esefexapi/sounddb"

func (c *DiscordPlayer) PlaySound(uid sounddb.SoundUID, serverID, userID string) error {
	// check if bot is joined to server but in a different channel than the user
	// if so, move bot to user's channel if no other users are in the original channel

	// if the bot is not in the server, join the user's channel

	// if the bot is in the server and in the user's channel, play the sound

	// if the sound sound played reset the auto-disconnect (afk) timer

	return nil
}
