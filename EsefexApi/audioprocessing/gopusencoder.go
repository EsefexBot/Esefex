package audioprocessing

import (
	"esefexapi/audioprocessing/pcmutil"
	"io"

	"github.com/pkg/errors"
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
		return nil, errors.Wrap(err, "Error creating gopus encoder")
	}

	return &GopusEncoder{
		source:  &s,
		encoder: enc,
	}, nil
}

// Returns the next encoded opus frame (20ms)
func (e *GopusEncoder) EncodeNext() ([]byte, error) {
	pcm := make([]int16, 960*2)

	_, err := pcmutil.ReadPCM(*e.source, &pcm)
	// Read from the source
	if err != nil && err != io.EOF {
		return nil, errors.Wrap(err, "Error reading from source")
	}

	// Encode
	return e.encoder.Encode(pcm, 960, 960*2*2)
}
