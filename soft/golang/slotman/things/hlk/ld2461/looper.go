package ld2461

import (
	"errors"
	"fmt"
	"slotman/things"
	"slotman/utils/log"
	"time"
)

func (se *LD2461) readLoop() {

	defer se.loopGroup.Done()

	if !se.isProbe {
		log.Printf("LD2461 readLoop started...")
		defer log.Printf("LD2461 readLoop done.")
	}

	if !se.IsOpen {
		err := things.ErrThingNotOpen
		log.Cerror(err)
		return
	}

	var lastNumTargets []byte
	var lastNumTargetsTime int64
	var lastCoordinates []byte
	var lastCoordinatesTime int64

	parts := make([]byte, 100)
	input := make([]byte, 0)

	//
	// Drain buffered junk from port.
	//

	//port := se.uart
	//if port != nil {
	//	for {
	//		xfer, _ := port.Read(parts)
	//		if xfer == 0 {
	//			break
	//		}
	//	}
	//}

	for se.IsOpen {

		port := se.uart
		if port == nil {
			break
		}

		xfer, _ := port.Read(parts)
		input = append(input, parts[:xfer]...)

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

				if clen > 100 {

					//
					// Corrupted input.
					// Flush all.
					//

					input = nil
					break
				}

				//
				// Input not yet ready complete.
				//

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

			if data[0] == commandGetCoordinates {

				if se.IsStarted {

					if !bytesAreEqual(data, lastCoordinates) ||
						time.Now().Unix()-lastCoordinatesTime >= 60 {

						var xPos []float64
						var yPos []float64

						coords := ""

						if len(data) >= 3 {
							xPos = append(xPos, float64(int8(data[1]))/10)
							yPos = append(yPos, float64(int8(data[2]))/10)
							coords += fmt.Sprintf(" %0.1f/%0.1f",
								float64(int8(data[1]))/10,
								float64(int8(data[2]))/10)
						}

						if len(data) >= 5 {
							xPos = append(xPos, float64(int8(data[3]))/10)
							yPos = append(yPos, float64(int8(data[4]))/10)
							coords += fmt.Sprintf(" %0.1f/%0.1f",
								float64(int8(data[3]))/10,
								float64(int8(data[4]))/10)
						}

						if len(data) >= 7 {
							xPos = append(xPos, float64(int8(data[5]))/10)
							yPos = append(yPos, float64(int8(data[6]))/10)
							coords += fmt.Sprintf(" %0.1f/%0.1f",
								float64(int8(data[5]))/10,
								float64(int8(data[6]))/10)
						}

						//log.Printf("GetCoordinates data=[ %02x ] %s", data, coords)

						if se.handler != nil {
							go se.handler.OnHumanTracking(se, xPos, yPos)
						}

						lastCoordinates = data
						lastCoordinatesTime = time.Now().Unix()
					}
				}

			} else if data[0] == commandGetNumTargets {

				if se.IsStarted {

					if !bytesAreEqual(data, lastNumTargets) ||
						time.Now().Unix()-lastNumTargetsTime >= 60 {
						log.Printf("GetNumTargets data=[ %02x ]", data)

						lastNumTargets = data
						lastNumTargetsTime = time.Now().Unix()
					}
				}

			} else {
				se.results <- data
			}

			input = ti
		}

		if len(input) == 0 {
			input = nil
		}
	}
}

func (se *LD2461) readResult(command byte) (result []byte, err error) {

	for try := 1; try < 3; try++ {

		select {

		case result = <-se.results:
			//log.Printf("Serial result=[ %02x ]", result)
			err = nil

		case <-time.After(time.Millisecond * 500):
			err = errors.New("ld2461 serial timeout")
			continue
		}

		if result == nil || len(result) == 0 || result[0] != command {
			err = errors.New("ld2461 wrong response")
			continue
		}

		return
	}

	return
}

func (se *LD2461) write(command byte, data []byte) (err error) {

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
