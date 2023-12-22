package audioprocessing

import "esefexapi/util"

func MixPCMs16leClip(bytes []int16) int16 {
	var mix int = 0

	for _, b := range bytes {
		mix += int(b)
	}

	return int16(util.ClampInt(mix, -32768, 32767))
}

func MixPCMs16leSum(bytes []int16) int16 {
	var mix int16 = 0

	for _, b := range bytes {
		mix += b
	}

	return mix
}

func MixPCMs16leAverage(bytes []int16) int16 {
	if len(bytes) == 0 {
		return 0
	}

	var mix int = 0

	for _, b := range bytes {
		mix += int(b)
	}

	mix /= len(bytes)

	return int16(mix)
}

func AsPCMs16le(bytes []byte) []int16 {
	var shorts []int16
	for i := 0; i < len(bytes); i += 2 {
		shorts = append(shorts, int16(bytes[i])|int16(bytes[i+1])<<8)
	}
	return shorts
}
