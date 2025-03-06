package ld6001a

import (
	"errors"
	"math"
	"slotman/things"
	"slotman/utils/log"
	"strings"
	"time"
)

func (se *LD6001a) readLoop() {

	if !se.isProbe {
		log.Printf("LD6001a readLoop started...")
		defer log.Printf("LD6001a readLoop done.")
	}

	if !se.IsOpen {
		err := things.ErrThingNotOpen
		log.Cerror(err)
		return
	}

	parts := make([]byte, 100)
	input := make([]byte, 0)

	//
	// Drain buffered junk from port.
	//

	port := se.uart
	if port != nil {
		for {
			xfer, _ := port.Read(parts)
			if xfer == 0 {
				break
			}
		}
	}

	for se.IsOpen {

		port = se.uart
		if port == nil {
			break
		}

		xfer, _ := port.Read(parts)
		input = append(input, parts[:xfer]...)

		log.Printf("###### read xfer=%d [ %02x ]", xfer, parts[:xfer])

		var ok bool
		var s1, s2, s3 bool

		for len(input) > 0 {

			//
			// Skip anything what is not a possible header.
			//

			for input[0] != 0x01 && input[0] != 0x55 && input[0] != 'A' {
				input = input[1:]
			}

			input, s1, ok = se.readLevel3(input)
			if ok {
				continue
			}

			input, s2, ok = se.readLevel0(input)
			if ok {
				continue
			}

			input, s3, ok = se.readATOK(input)
			if ok {
				continue
			}

			if s1 || s2 || s3 {
				break
			}

			input = input[1:]
		}

		if len(input) == 0 {
			input = nil
		}
	}
}

func (se *LD6001a) readATOK(input []byte) (output []byte, short, ok bool) {

	output = input

	if len(output) < 6 {
		short = true
		return
	}

	if !strings.HasPrefix(string(output), "AT+OK") {
		return
	}

	output = output[5:]

	for len(output) > 0 && (output[0] == '\n' || output[0] == '\r') {
		output = output[1:]
	}

	se.results <- "AT+OK"
	ok = true
	return
}

func (se *LD6001a) readLevel0(input []byte) (output []byte, short, ok bool) {

	//
	// 55 AA 0A 04 00 00 00 00 00 0E
	//  0  1  2  3  4  5  6  7  8  9
	//
	// 55 AA     : Frame header
	// 0A        : Line feed?
	// type=0x04 : People counting
	// 00 00     : Reserved
	// 00 00     : Reserved
	// 00        : Number of people counted
	// 0E        : XOR check (0A 04 00 00 00 00 00)
	//

	output = input

	if len(output) < 10 {
		short = true
		return
	}

	if output[0] != 0x55 && output[1] != 0xaa && output[2] != 0x0a {
		return
	}

	if output[3] == 0x04 {

		people := output[8]
		_ = people

		// todo: inject result...

	}

	output = output[10:]

	ok = true
	return
}

func (se *LD6001a) readLevel3(input []byte) (output []byte, short, ok bool) {

	//
	// HEAD        +  0: 01 02 03 04 05 06 07 08
	// LENGTH      +  8: 40 00 00 00 == 64 bytes incl. HEAD
	// FRAME       + 12: A3 01 00 00 == 419
	// TLVs        + 16: 01 00 00 00 == Always 1
	// POINTLENGTH + 20: 00 00 00 00 == Always 0
	// TLVs        + 24: 02 00 00 00 == Always 2
	// TRACKLENGTH + 28: 20 00 00 00 == 32 / 32 == 1 people
	//
	//   F  +  0: 00 00 00 00
	//   ID +  4: 00 00 00 00
	//   X  +  8: 21 28 96 BF
	//   Y  + 12: CB 85 20 40
	//   Z  + 16: 9A AB A3 3E
	//   Vx + 20: 8A BD C1 3D
	//   Vy + 24: 50 98 99 BD
	//   Vz + 28: 40 52 C3 3A
	//
	// CS: CC
	//

	if len(input) < 12 {
		short = true
		return
	}

	if output[0] != 0x01 && output[1] != 0x02 &&
		output[2] != 0x03 && output[3] != 0x04 &&
		output[4] != 0x05 && output[5] != 0x06 &&
		output[6] != 0x07 && output[7] != 0x08 {
		return
	}

	length := int(output[8])<<0 + int(output[9])<<8 +
		int(output[10])<<16 + int(output[11])<<24

	if length >= 32 && len(input) < length {
		short = true
		return
	}

	if length < 32 {

		//
		// Bad entry...
		//

		output = output[8+4:]
		return
	}

	trackLength := int(output[28])<<0 + int(output[29])<<8 +
		int(output[30])<<16 + int(output[31])<<24

	if len(input) < 32+trackLength {
		short = true
		return
	}

	for index := 0; index < trackLength/32; index++ {

		bi := 32 + index*32

		id := int(output[bi+4])<<0 + int(output[bi+5])<<8 +
			int(output[bi+6])<<16 + int(output[bi+7])<<24

		parsed := uint32(output[bi+8])<<0 + uint32(output[bi+9])<<8 +
			uint32(output[bi+10])<<16 + uint32(output[bi+11])<<24

		x := math.Float32frombits(parsed)

		parsed = uint32(output[bi+12])<<0 + uint32(output[bi+13])<<8 +
			uint32(output[bi+14])<<16 + uint32(output[bi+15])<<24

		y := math.Float32frombits(parsed)

		parsed = uint32(output[bi+16])<<0 + uint32(output[bi+17])<<8 +
			uint32(output[bi+18])<<16 + uint32(output[bi+19])<<24

		z := math.Float32frombits(parsed)

		parsed = uint32(output[bi+20])<<0 + uint32(output[bi+21])<<8 +
			uint32(output[bi+22])<<16 + uint32(output[bi+23])<<24

		vx := math.Float32frombits(parsed)

		parsed = uint32(output[bi+24])<<0 + uint32(output[bi+25])<<8 +
			uint32(output[bi+26])<<16 + uint32(output[bi+27])<<24

		vy := math.Float32frombits(parsed)

		parsed = uint32(output[bi+28])<<0 + uint32(output[bi+29])<<8 +
			uint32(output[bi+30])<<16 + uint32(output[bi+31])<<24

		vz := math.Float32frombits(parsed)

		_ = id
		_ = x
		_ = y
		_ = z
		_ = vx
		_ = vy
		_ = vz
	}

	// todo: inject result...

	output = output[length:]
	ok = true
	return
}

func (se *LD6001a) readResult() (result string, err error) {

	for try := 1; try < 3; try++ {

		select {

		case result = <-se.results:
			log.Printf("Serial result=<%s>", result)
			err = nil
			return

		case <-time.After(time.Millisecond * 500):
			err = errors.New("ld6001a serial timeout")
			continue
		}
	}

	return
}
