package packet

type Flags struct {
	Bit0 bool // 1
	Bit1 bool // 2
	Bit2 bool // 4
	Bit3 bool // 8
	Bit4 bool // 16
	Bit5 bool // 32
	Bit6 bool // 64
	Bit7 bool // 128
}

func DefaultFlags() *Flags {
	return &Flags{
		Bit0: false,
		Bit1: false,
		Bit2: false,
		Bit3: false,
		Bit4: false,
		Bit5: false,
		Bit6: false,
		Bit7: false,
	}
}

// ParseFlags parses a byte into a Flags struct
func parseFlags(flagsByte byte) Flags {
	return Flags{
		//! ENDIANESS !
		Bit0: flagsByte&0x01 != 0, // Check if the least significant bit is set
		Bit1: flagsByte&0x02 != 0, // Check if the second bit is set
		Bit2: flagsByte&0x04 != 0, // Check if the third bit is set
		Bit3: flagsByte&0x08 != 0, // Check if the fourth bit is set
		Bit4: flagsByte&0x10 != 0, // Check if the fifth bit is set
		Bit5: flagsByte&0x20 != 0, // Check if the sixth bit is set
		Bit6: flagsByte&0x40 != 0, // Check if the seventh bit is set
		Bit7: flagsByte&0x80 != 0, // Check if the eighth (most significant) bit is set
	}
}

// ToByte converts a Flags struct back into a single byte
func (f Flags) ToByte() byte {
	var result byte = 0

	if f.Bit0 {
		result |= 0x01 // Set the least significant bit
	}
	if f.Bit1 {
		result |= 0x02 // Set the second bit
	}
	if f.Bit2 {
		result |= 0x04 // Set the third bit
	}
	if f.Bit3 {
		result |= 0x08 // Set the fourth bit
	}
	if f.Bit4 {
		result |= 0x10 // Set the fifth bit
	}
	if f.Bit5 {
		result |= 0x20 // Set the sixth bit
	}
	if f.Bit6 {
		result |= 0x40 // Set the seventh bit
	}
	if f.Bit7 {
		result |= 0x80 // Set the most significant bit
	}

	return result
}
