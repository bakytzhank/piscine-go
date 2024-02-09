package piscine

func ReverseBits(oct byte) byte {
	var result byte
	for i := 0; i < 8; i++ {
		// Shift the result to the left
		result <<= 1
		// Extract the rightmost bit from oct and add it to the result
		result |= oct & 1
		// Shift oct to the right to get the next bit
		oct >>= 1
	}
	return result
}