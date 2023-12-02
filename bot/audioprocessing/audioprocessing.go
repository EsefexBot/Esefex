package audioprocessing

func MixPCMs16le(bytes []int16) int16 {
	var mix int16 = 0

	for _, b := range bytes {
		mix += b
	}

	return mix
}

func AsPCMs16le(bytes []byte) []int16 {
	var shorts []int16
	for i := 0; i < len(bytes); i += 2 {
		shorts = append(shorts, int16(bytes[i])|int16(bytes[i+1])<<8)
	}
	return shorts
}
