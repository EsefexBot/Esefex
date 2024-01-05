package vcon

import (
	"esefexapi/audioprocessing"
	"esefexapi/sounddb"
	"esefexapi/timer"
	"esefexapi/types"
	"io"

	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

type VCon struct {
	playSound chan sounddb.SoundURI
	stop      chan struct{}
	mixer     *audioprocessing.S16leMixReader
	enc       *audioprocessing.GopusEncoder
	vc        *discordgo.VoiceConnection
	dc        *discordgo.Session
	db        sounddb.ISoundDB
}

func NewVCon(ds *discordgo.Session, db sounddb.ISoundDB, guildID types.GuildID, channelID types.ChannelID) (*VCon, error) {
	vc, err := ds.ChannelVoiceJoin(guildID.String(), channelID.String(), false, true)
	if err != nil {
		return nil, errors.Wrap(err, "Error joining voice channel")
	}

	mixer := audioprocessing.NewS16leMixReader()

	enc, err := audioprocessing.NewGopusEncoder(mixer)
	if err != nil {
		return nil, errors.Wrap(err, "Error creating encoder")
	}

	return &VCon{
		playSound: make(chan sounddb.SoundURI),
		stop:      make(chan struct{}),
		mixer:     mixer,
		enc:       enc,
		vc:        vc,
		dc:        ds,
		db:        db,
	}, nil
}

func (a *VCon) PlaySound(uid sounddb.SoundURI) {
	// log.Printf("channel: %s\n", a.vc.ChannelID)
	// log.Print(a.playSound)
	a.playSound <- uid
	timer.MessageElapsed("Sent sound to playSound")
}

// this is the main loop of the audio queue
func (a *VCon) Run() {
	log.Println("Running VCon")
	a.vc.Speaking(true)

	for {
		// log.Println("Looping...")
		select {
		case sound := <-a.playSound:
			timer.MessageElapsed("Got sound to play")
			// log.Printf("Playing sound %s\n", sound)
			ok, err := a.db.SoundExists(sound)
			if err != nil {
				log.Printf("Error checking if sound exists: %+v\n", err)
				continue
			} else if !ok {
				log.Printf("Sound does not exist: %+v\n", sound)
				continue
			}
			timer.MessageElapsed("Checked if sound exists")

			pcm, err := a.db.GetSoundPcm(sound)
			if err != nil {
				log.Printf("Error getting sound pcm: %+v\n", err)
				continue
			}
			timer.MessageElapsed("Got sound pcm")

			s := audioprocessing.NewS16leReferenceReaderFromRef(pcm)
			a.mixer.AddSource(s)
			// log.Println("Added sound to mixer")
			timer.MessageElapsed("Added sound to mixer")
		case <-a.stop:
			return
		default:
			ff := time.After(time.Duration(audioprocessing.FrameLengthMs))

			if !a.mixer.Empty() {
				buf, err := a.enc.EncodeNext()
				if err != nil && errors.Cause(err) != io.EOF {
					log.Printf("Error encoding next: %+v\n", err)
				}
				a.vc.OpusSend <- buf
			}

			<-ff
		}

	}
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
