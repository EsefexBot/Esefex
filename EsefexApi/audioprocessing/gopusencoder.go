package audioprocessing

import (
	"encoding/binary"
	"io"

	"layeh.com/gopus"
)

var FrameSize = 960
var FrameLengthMs = 20

type GopusEncoder struct {
	source  *io.Reader
	encoder *gopus.Encoder
}

func NewGopusEncoder(s io.Reader) (*GopusEncoder, error) {
	enc, err := gopus.NewEncoder(48000, 2, gopus.Voip)
	if err != nil {
		return nil, err
	}

	return &GopusEncoder{
		source:  &s,
		encoder: enc,
	}, nil
}

// Returns the next encoded opus frame (20ms)
func (e *GopusEncoder) EncodeNext() ([]byte, error) {
	pcm := make([]int16, 960*2)

	err := binary.Read(*e.source, binary.LittleEndian, &pcm)
	// Read from the source
	if err != nil {
		return nil, err
	}

	// Encode
	return e.encoder.Encode(pcm, 960, 960*2*2)
}
