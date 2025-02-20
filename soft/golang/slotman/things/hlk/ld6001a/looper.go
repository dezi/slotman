package ld6001a

import (
	"errors"
	"slotman/things"
	"slotman/utils/log"
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

		//log.Printf("###### read xfer=%d [ %02x ]", xfer, parts[:xfer])

		var tb byte

		for len(input) >= 11 {

			ti := input

			if ti[0] != 0xFF || ti[1] != 0xEE || ti[2] != 0xDD {
				input = ti[1:]
				continue
			}

			ti = ti[3:]

			clen := int(ti[0])<<8 + int(ti[1])
			ti = ti[2:]

			if len(ti) < clen+4 {

				//
				// Input not yet ready complete.
				//

				input = append([]byte{0xFF, 0xEE, 0xDD}, input...)

				break
			}

			sumA := byte(0)

			var data []byte
			for clen > 0 {
				sumA += ti[0]
				data = append(data, ti[0])
				ti = ti[1:]
				clen--
			}

			sumB := ti[0]
			ti = ti[1:]

			if sumA != sumB {
				// Discard input until here.
				input = ti
				continue
			}

			tb = ti[0]
			ti = ti[1:]

			if tb != 0xDD {
				// Discard input until here.
				input = ti
				continue
			}

			tb = ti[0]
			ti = ti[1:]

			if tb != 0xEE {
				// Discard input until here.
				input = ti
				continue
			}

			tb = ti[0]
			ti = ti[1:]

			if tb != 0xFF {
				// Discard input until here.
				input = ti
				continue
			}

			//
			// Valid message received.
			//

			input = ti
			break
		}

		if len(input) == 0 {
			input = nil
		}
	}
}

func (se *LD6001a) readResult(command byte) (result []byte, err error) {

	for try := 1; try < 3; try++ {

		select {

		case result = <-se.results:
			//log.Printf("Serial result=[ %02x ]", result)
			err = nil

		case <-time.After(time.Millisecond * 500):
			err = errors.New("ld6001a serial timeout")
			continue
		}

		if result == nil || len(result) == 0 || result[0] != command {
			err = errors.New("ld6001a wrong response")
			continue
		}

		return
	}

	return
}

func (se *LD6001a) write(command byte, data []byte) (err error) {

	if !se.IsOpen {
		err = things.ErrThingNotOpen
		log.Cerror(err)
		return
	}

	var cmdBytes []byte

	cmdBytes = append(cmdBytes, 0xFF)
	cmdBytes = append(cmdBytes, 0xEE)
	cmdBytes = append(cmdBytes, 0xDD)

	cmdSize := 1 + len(data)
	cmdBytes = append(cmdBytes, byte(cmdSize>>8))
	cmdBytes = append(cmdBytes, byte(cmdSize))

	cmdBytes = append(cmdBytes, command)
	cmdBytes = append(cmdBytes, data...)

	sum := command
	for _, byt := range data {
		sum += byt
	}

	cmdBytes = append(cmdBytes, sum)

	cmdBytes = append(cmdBytes, 0xDD)
	cmdBytes = append(cmdBytes, 0xEE)
	cmdBytes = append(cmdBytes, 0xFF)

	//log.Printf("Write command=%02x data=[ %02x ]", command, data)
	//log.Printf("Write cmdBytes=[ %02x ]", cmdBytes)

	_, err = se.uart.Write(cmdBytes)
	log.Cerror(err)

	return
}
