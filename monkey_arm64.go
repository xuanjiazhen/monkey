package monkey

// Assembles a jump to a function value
func jmpToFunctionValue(to uintptr) []byte {
	return []byte{
		0x49, 0x00,0x00,0x58,	// LDR x10, .+8
		0x20, 0x01,0x1F,0xD6,	// BR x10
		byte(to),
		byte(to >> 8),
		byte(to >> 16),
		byte(to >> 24),
		byte(to >> 32),
		byte(to >> 40),
		byte(to >> 48),
		byte(to >> 56), 
	}
}
