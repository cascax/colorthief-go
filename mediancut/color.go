package mediancut

// RGB convert RGB color to 16 bit value
func RGB(r uint8, g uint8, b uint8) uint16 {
	return (uint16(b&^7) << 7) | (uint16((g)&^7) << 2) | (uint16(r) >> 3)
}

// RedColor return red value from uint16 color
func RedColor(color uint16) uint8 {
	return uint8((color & 0x1F) << 3)
}

// GreenColor return green value from uint16 color
func GreenColor(color uint16) uint8 {
	return uint8((color & 0x3E0) >> 2)
}

// BlueColor return blue value from uint16 color
func BlueColor(color uint16) uint8 {
	return uint8((color & 0x7C00) >> 7)
}

func GetRGB(color uint16) (uint8, uint8, uint8) {
	return RedColor(color), GreenColor(color), BlueColor(color)
}
