package audioqueue

import (
	"esefexapi/appcontext"
	"esefexapi/audioprocessing"
	"esefexapi/filedb"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

type AudioQueue struct {
	playSound chan filedb.SoundUid
	stop      chan struct{}
	mixer     *audioprocessing.S16leMixReader
	enc       *audioprocessing.GopusEncoder
	vc        *discordgo.VoiceConnection
	ds        *discordgo.Session
	ctx       *appcontext.Context
}

func NewAudioQueue(ctx *appcontext.Context, guildID string, channelID string) (*AudioQueue, error) {
	vc, err := ctx.DiscordSession.ChannelVoiceJoin(guildID, channelID, false, true)
	if err != nil {
		return nil, err
	}

	mixer := audioprocessing.NewS16leMixReader()

	enc, err := audioprocessing.NewGopusEncoder(mixer)

	return &AudioQueue{
		playSound: make(chan filedb.SoundUid),
		stop:      make(chan struct{}),
		mixer:     mixer,
		enc:       enc,
		vc:        vc,
		ctx:       ctx,
	}, nil
}

// this is the main loop of the audio queue
func (a *AudioQueue) Run() {
	a.vc.Speaking(true)

	go func() {
		for {
			select {
			case sound := <-a.playSound:
				b, err := a.ctx.AudioCache.GetSound(sound)
				if err != nil {
					log.Println(err)
					continue
				}

				s := audioprocessing.NewS16leCacheReaderFromBytes(b)
				a.mixer.AddSource(s)
			case <-a.stop:
				return
			}

			ff := time.After(time.Duration(audioprocessing.FrameLengthMs))

			if !a.mixer.Empty() {
				buf, err := a.enc.EncodeNext()
				if err != nil {
					log.Panicln(err)
					return
				}
				a.vc.OpusSend <- buf
			}

			<-ff
		}
	}()
}

func (a *AudioQueue) Close() {
	close(a.stop)
	a.vc.Speaking(false)
	a.vc.Disconnect()
}
