package audioprocessing

import (
	"esefexapi/audioprocessing/pcmutil"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestS16leReferenceReader(t *testing.T) {
	data := []int16{1, 2, 3, 4, 5, 6}

	reader := NewS16leReferenceReaderFromRef(&data)

	assert.Equal(t, data, *reader.data)

	// assert.Equal(t, byte(1), reader.getByte(0))
	// assert.Equal(t, byte(0), reader.getByte(1))
	// assert.Equal(t, byte(2), reader.getByte(2))
	// assert.Equal(t, byte(0), reader.getByte(3))

	int16buf := make([]int16, len(data))
	n, err := pcmutil.ReadPCM(reader, &int16buf)
	assert.Nil(t, err)
	log.Printf("n: %d", n)

	assert.Equal(t, data, int16buf)
}
