package ld6001a

import (
	"encoding/json"
	"errors"
	"math"
	"slotman/things"
	"slotman/utils/log"
	"strings"
	"time"
)

func (se *LD6001a) readLoop() {

	defer se.loopGroup.Done()

	if !se.isProbe {
		log.Printf("LD6001a readLoop started...")
		defer log.Printf("LD6001a readLoop done.")
	}

	if !se.IsOpen {
		err := things.ErrThingNotOpen
		log.Cerror(err)
		return
	}

	parts := make([]byte, 1000)
	input := make([]byte, 0)

	//
	// Drain buffered junk from uart.
	//

	uart := se.uart
	if uart != nil {
		for {
			xfer, _ := uart.Read(parts)
			if xfer < len(parts) {
				break
			}
		}
	}

	for se.IsOpen {

		uart = se.uart
		if uart == nil {
			break
		}

		xfer, _ := uart.Read(parts)
		input = append(input, parts[:xfer]...)

		//log.Printf("###### read size=%d xfer=%d [ %02x ]", len(input), xfer, parts[:xfer])

		if len(input) > 1000 {

			//
			// Chaos recovery.
			//

			input = nil
			continue
		}

		//if xfer > 0 {
		//	log.Printf("###### read xfer=%d <%s>", xfer, string(parts[:xfer]))
		//}

		//if xfer == 0 {
		//	time.Sleep(time.Millisecond * 10)
		//	continue
		//}

		var ok bool
		var s1, s2, s3, s4, s5 bool

		for len(input) > 0 {

			//
			// Skip anything what is not a possible header.
			//

			if input[0] != 0x01 && input[0] != 0x55 && input[0] != 'A' &&
				input[0] != '{' && input[0] != '-' {
				input = input[1:]
				continue
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

			input, s4, ok = se.readParams(input)
			if ok {
				continue
			}

			input, s5, ok = se.readPositions(input)
			if ok {
				continue
			}

			if s1 || s2 || s3 || s4 || s5 {
				break
			}

			input = input[1:]
		}

		if len(input) == 0 {
			input = nil
		}

		time.Sleep(time.Millisecond * 2)
	}
}

func (se *LD6001a) readATOK(input []byte) (output []byte, short, ok bool) {

	output = input

	if len(output) < 6 {
		short = true
		return
	}

	if !strings.HasPrefix(string(output), "AT+") {
		return
	}

	var result []byte

	for len(output) > 0 && output[0] != '\n' && output[0] != '\r' {
		result = append(result, output[0])
		output = output[1:]
	}

	for len(output) > 0 && (output[0] == '\n' || output[0] == '\r') {
		output = output[1:]
	}

	//log.Printf("Response result=<%s>", string(result))

	se.results <- string(result)
	ok = true
	return
}

func (se *LD6001a) readPositions(input []byte) (output []byte, short, ok bool) {

	output = input

	if len(output) < 5 {
		short = true
		return
	}

	if !strings.HasPrefix(string(output), "-----") {
		return
	}

	if !strings.HasSuffix(string(output), "\r\n") {
		short = true
		return
	}

	log.Printf("################ positions=%s", string(output))

	ok = true
	return
}

func (se *LD6001a) readParams(input []byte) (output []byte, short, ok bool) {

	output = input

	if len(output) < 1 {
		short = true
		return
	}

	//
	// Fuck this shit.
	//
	// This is a real shitty implementation:
	//
	// 1. Developer starts response like JSON,
	// but fails to terminate it correctly.
	//
	// 2. Careless spelling like "PeopleCntSoftVerison".
	//
	// 3. Developer is too stupid to send CR/LF instead
	// he sends "\x09\x0a".
	//
	// 4. Developer messes up string with bogus "\xa3\xba"
	// character sequences.
	//
	// 5. Developer sends NO indication that
	// response is complete.
	//

	if output[0] != '{' {
		return
	}

	if !strings.Contains(string(output), "Target exit") {
		short = true
		return
	}

	for inx, ccc := range output {
		if ccc == 0x01 {
			output = output[:inx]
			break
		}
	}

	//
	// De-fuck output to make it work with JSON.
	//

	result := string(output)
	result = strings.ReplaceAll(result, "\x09\x0a", "\r\n")
	result = strings.ReplaceAll(result, "\xa3\xba", " ")
	result = strings.ReplaceAll(result, "Moving target", `"Moving target":`)
	result = strings.ReplaceAll(result, "Static target", `"Static target":`)
	result = strings.ReplaceAll(result, "Target exit", `"Target exit":`)
	result = strings.ReplaceAll(result, "s,", ",")
	result = strings.TrimSpace(result)
	result = strings.TrimSuffix(result, ",")
	result = result + "\r\n" + "}"

	output = nil

	//log.Printf("Response params=\n%s", result)

	err := json.Unmarshal([]byte(result), &se.Params)
	if err != nil {
		log.Printf("##### unmarshal result:\n%s", []byte(result))
		log.Cerror(err)
		return
	}

	//log.Printf("Parsed params=%+v", se.Params)

	//
	// Fake a result code to satisfy write ok.
	//

	se.results <- "AT+READ"
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

	if output[0] != 0x55 || output[1] != 0xaa || output[2] != 0x0a {
		return
	}

	if output[3] == 0x04 {

		people := output[8]

		log.Printf("People count=%d", people)

		// todo: inject result...

	}

	output = output[10:]

	//ok = true
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

	output = input

	if len(output) < 12 {
		short = true
		return
	}

	if output[0] != 0x01 || output[1] != 0x02 ||
		output[2] != 0x03 || output[3] != 0x04 ||
		output[4] != 0x05 || output[5] != 0x06 ||
		output[6] != 0x07 || output[7] != 0x08 {
		return
	}

	length := int(output[8])<<0 + int(output[9])<<8 +
		int(output[10])<<16 + int(output[11])<<24

	if length%32 != 0 || length >= 1024 {

		//
		// Recover from read chaos.
		//

		log.Printf("Flush corrupted (1)...")
		output = nil
		return
	}

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

	var xor byte

	//
	// Frame number.
	//

	xor ^= output[12]
	xor ^= output[13]
	xor ^= output[14]
	xor ^= output[15]

	trackLength := int(output[28])<<0 + int(output[29])<<8 +
		int(output[30])<<16 + int(output[31])<<24

	if trackLength%32 != 0 || trackLength >= 1024 {

		//
		// Recover from read chaos.
		//

		log.Printf("Flush corrupted (2)...")

		output = nil
		return
	}

	if len(input) < 32+trackLength+1 {
		short = true
		return
	}

	sets := trackLength / 32

	var ids []int
	var xps []float64
	var yps []float64
	var zps []float64
	var xvs []float64
	var yvs []float64
	var zvs []float64

	for index := 0; index < sets; index++ {

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

		ids = append(ids, id)
		xps = append(xps, float64(x))
		yps = append(yps, float64(y))
		zps = append(zps, float64(z))
		xvs = append(xvs, float64(vx))
		yvs = append(yvs, float64(vy))
		zvs = append(zvs, float64(vz))
	}

	for inx := 32; inx < length; inx++ {
		xor = xor ^ output[inx]
	}

	if xor != output[length] {
		log.Printf("Checksum error...")
		output = nil
		return
	}

	//for inx := range ids {
	//	log.Printf("Target id=%d %0.1f %0.1f %0.1f %0.1f %0.1f %0.1f",
	//		ids[inx], xps[inx], yps[inx], zps[inx], xvs[inx], yvs[inx], zvs[inx])
	//}

	handler := se.handler
	if handler != nil {
		go handler.OnHumanTracking(se, xps, yps)
		//go handler.OnHumanTracking3D(se, ids, xps, yps, zps, xvs, yvs, zvs)
	}

	_ = xvs
	_ = yvs
	_ = zvs

	//
	// Skip length + 1 (checksum).
	//

	output = output[length+1:]
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
