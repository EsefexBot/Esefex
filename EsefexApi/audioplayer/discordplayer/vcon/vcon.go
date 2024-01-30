package vcon

import (
	"esefexapi/audioprocessing"
	"esefexapi/sounddb"

	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

type VCon struct {
	playSound chan sounddb.SoundUID
	stop      chan struct{}
	mixer     *audioprocessing.S16leMixReader
	enc       *audioprocessing.GopusEncoder
	vc        *discordgo.VoiceConnection
	dc        *discordgo.Session
	db        sounddb.ISoundDB
}

func NewVCon(dc *discordgo.Session, db sounddb.ISoundDB, guildID string, channelID string) (*VCon, error) {
	vc, err := dc.ChannelVoiceJoin(guildID, channelID, false, true)
	if err != nil {
		return nil, errors.Wrap(err, "Error joining voice channel")
	}

	mixer := audioprocessing.NewS16leMixReader()

	enc, err := audioprocessing.NewGopusEncoder(mixer)
	if err != nil {
		return nil, errors.Wrap(err, "Error creating encoder")
	}

	return &VCon{
		playSound: make(chan sounddb.SoundUID),
		stop:      make(chan struct{}),
		mixer:     mixer,
		enc:       enc,
		vc:        vc,
		dc:        dc,
		db:        db,
	}, nil
}

func (a *VCon) PlaySound(uid sounddb.SoundUID) {
	log.Printf("channel: %s\n", a.vc.ChannelID)
	log.Print(a.playSound)
	a.playSound <- uid
}

// this is the main loop of the audio queue
func (a *VCon) Run() {
	log.Println("Running VCon")
	a.vc.Speaking(true)

	go func() {
		for {
			// log.Println("Looping...")
			select {
			case sound := <-a.playSound:
				// log.Printf("Playing sound %s\n", sound)
				ok, err := a.db.SoundExists(sound)
				if err != nil {
					log.Printf("Error checking if sound exists: %+v\n", err)
					continue
				} else if !ok {
					log.Printf("Sound does not exist: %+v\n", sound)
					continue
				}
				pcm, err := a.db.GetSoundPcm(sound)
				if err != nil {
					log.Printf("Error getting sound pcm: %+v\n", err)
					continue
				}

				s := audioprocessing.NewS16leReferenceReaderFromRef(pcm)
				a.mixer.AddSource(s)
				// log.Println("Added sound to mixer")
			case <-a.stop:
				return
			default:
				ff := time.After(time.Duration(audioprocessing.FrameLengthMs))

				if !a.mixer.Empty() {
					buf, err := a.enc.EncodeNext()
					if err != nil {
						log.Printf("Error encoding next: %+v\n", err)
					}
					a.vc.OpusSend <- buf
				}

				<-ff
			}
		}
	}()
}

func (a *VCon) Close() {
	log.Println("Closing VCon")
	close(a.stop)
	a.vc.Speaking(false)
	a.vc.Disconnect()
}

func (a *VCon) IsPlaying() bool {
	return !a.mixer.Empty()
}
