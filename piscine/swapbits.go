package piscine

func SwapBits(octet byte) byte {
	firsthalf := octet << 4
	secondhalf := octet >> 4
	return secondhalf | firsthalf
}