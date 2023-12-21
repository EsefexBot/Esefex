package discordplayer

import (
	"esefexapi/audioprocessing"
	"esefexapi/sounddb"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

type SfxPlayer struct {
	playSound chan sounddb.SoundUID
	stop      chan struct{}
	mixer     *audioprocessing.S16leMixReader
	enc       *audioprocessing.GopusEncoder
	vc        *discordgo.VoiceConnection
	dc        *discordgo.Session
	db        sounddb.ISoundDB
}

func NewSfxPlayer(dc *discordgo.Session, db sounddb.ISoundDB, guildID string, channelID string) (*SfxPlayer, error) {
	vc, err := dc.ChannelVoiceJoin(guildID, channelID, false, true)
	if err != nil {
		return nil, err
	}

	mixer := audioprocessing.NewS16leMixReader()

	enc, err := audioprocessing.NewGopusEncoder(mixer)

	return &SfxPlayer{
		playSound: make(chan sounddb.SoundUID),
		stop:      make(chan struct{}),
		mixer:     mixer,
		enc:       enc,
		vc:        vc,
		dc:        dc,
		db:        db,
	}, nil
}

// this is the main loop of the audio queue
func (a *SfxPlayer) Run() {
	a.vc.Speaking(true)

	go func() {
		for {
			select {
			case sound := <-a.playSound:
				pcm, err := a.db.GetSoundPcm(sound)
				if err != nil {
					log.Println(err)
					continue
				}

				s := audioprocessing.NewS16leCacheReaderFromPCM(pcm)
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

func (a *SfxPlayer) Close() {
	close(a.stop)
	a.vc.Speaking(false)
	a.vc.Disconnect()
}
