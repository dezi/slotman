package ld2451

import (
	"errors"
	"math"
	"slotman/things"
	"slotman/utils/log"
	"time"
)

func (se *LD2451) readLoop() {

	if !se.isProbe {
		log.Printf("LD2451 readLoop started...")
		defer log.Printf("LD2451 readLoop done.")
	}

	if !se.IsOpen {
		err := things.ErrThingNotOpen
		log.Cerror(err)
		return
	}

	//var lastNumTargets []byte
	//var lastNumTargetsTime int64
	//var lastCoordinates []byte
	//var lastCoordinatesTime int64

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

		xfer, err := port.Read(parts)
		input = append(input, parts[:xfer]...)

		if err != nil {
			time.Sleep(time.Millisecond * 100)
			continue
		}

		//if xfer > 0 {
		//	log.Printf("###### read xfer=%d [ %02x ]", xfer, parts[:xfer])
		//}

		var tb byte
		var isConfig bool
		var isReport bool

		for len(input) >= 10 {

			ti := input

			if (ti[0] != 0xfd || ti[1] != 0xfc || ti[2] != 0xfb || ti[3] != 0xfa) &&
				(ti[0] != 0xf4 || ti[1] != 0xf3 || ti[2] != 0xf2 || ti[3] != 0xf1) {
				input = ti[1:]
				continue
			}

			isConfig = ti[0] == 0xfd
			isReport = ti[0] == 0xf4

			ti = ti[4:]

			clen := int(ti[0]) + int(ti[1])<<8
			ti = ti[2:]

			if len(ti) < clen+4 {

				//
				// Input not yet ready complete.
				//

				break
			}

			var data []byte
			for clen > 0 {
				data = append(data, ti[0])
				ti = ti[1:]
				clen--
			}

			tb = ti[0]
			ti = ti[1:]

			if isConfig {
				if tb != 0x04 {
					// Discard input until here.
					input = ti
					continue
				}

				tb = ti[0]
				ti = ti[1:]

				if tb != 0x03 {
					// Discard input until here.
					input = ti
					continue
				}

				tb = ti[0]
				ti = ti[1:]

				if tb != 0x02 {
					// Discard input until here.
					input = ti
					continue
				}

				tb = ti[0]
				ti = ti[1:]

				if tb != 0x01 {
					// Discard input until here.
					input = ti
					continue
				}
			}

			if isReport {
				if tb != 0xf8 {
					// Discard input until here.
					input = ti
					continue
				}

				tb = ti[0]
				ti = ti[1:]

				if tb != 0x0f7 {
					// Discard input until here.
					input = ti
					continue
				}

				tb = ti[0]
				ti = ti[1:]

				if tb != 0x0f6 {
					// Discard input until here.
					input = ti
					continue
				}

				tb = ti[0]
				ti = ti[1:]

				if tb != 0xf5 {
					// Discard input until here.
					input = ti
					continue
				}
			}

			//
			// Valid message received.
			//

			if isReport {
				log.Printf("###### report xfer=%d [ %02x ]", len(data), data)
				se.evalReport(data)
			}

			if isConfig {
				//log.Printf("###### config xfer=%d [ %02x ]", len(data), data)
				se.results <- data
			}

			input = ti
			break
		}

		if len(input) == 0 {
			input = nil
		}
	}
}

func (se *LD2451) evalReport(data []byte) {

	xPos := make([]float64, 0)
	yPos := make([]float64, 0)

	if len(data) >= 2 {

		targets := data[0]
		approach := data[1]

		_ = targets
		_ = approach

		data = data[2:]

		for len(data) >= 5 {

			degree := float64(int(data[0]) - 0x80)
			distance := float64(data[1])
			direction := data[2]
			speed := int(data[3])
			noise := int(data[4])

			_ = direction
			_ = speed
			_ = noise

			//log.Printf("########### degree=%0f distance=%0f direction=%d speed=%d noise=%d",
			//	degree, distance, direction, speed, noise)

			xP := distance * math.Cos(degree/180*math.Pi)
			yP := distance * math.Sin(degree/180*math.Pi)

			log.Printf("########### xp=%0f xp=%0f", xP, yP)

			xPos = append(xPos, xP)
			yPos = append(yPos, yP)

			data = data[5:]
		}
	}

	handler := se.handler
	if handler != nil {
		go handler.OnHumanTracking(se, xPos, yPos)
	}
}

func (se *LD2451) readResult(command byte) (result []byte, err error) {

	for try := 1; try < 3; try++ {

		select {

		case result = <-se.results:
			log.Printf("Serial result=[ %02x ]", result)
			err = nil

		case <-time.After(time.Millisecond * 500):
			err = errors.New("ld2451 serial timeout")
			continue
		}

		if result == nil || len(result) == 0 || result[0] != command {
			err = errors.New("ld2451 wrong response")
			continue
		}

		return
	}

	return
}
