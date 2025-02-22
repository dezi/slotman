package sgp40

import (
	"errors"
	"slotman/things"
	"slotman/utils/log"
	"time"
)

func (se *SGP40) writeCommandAndRead(
	command []byte, waitTime time.Duration, resultWords int) (
	result []byte, err error) {

	if !se.IsOpen {
		err = things.ErrThingNotOpen
		return
	}

	err = se.i2cDev.TransLock()
	if err != nil {
		log.Cerror(err)
		return
	}

	defer func() {
		drr := se.i2cDev.TransUnlock()
		log.Cerror(drr)
	}()

	xfer, err := se.i2cDev.WriteBytes(command)
	if err != nil {
		return
	}

	if xfer != len(command) {
		err = errors.New("short write")
		return
	}

	time.Sleep(waitTime)

	resbuf := make([]byte, 3)

	for rw := 0; rw < resultWords; rw++ {

		xfer, err = se.i2cDev.ReadBytes(resbuf)
		if err != nil {
			return
		}

		if xfer != 3 {
			err = errors.New("short read")
			return
		}

		if resbuf[2] != calculateCrc(resbuf[0:2]) {
			err = errors.New("checksum error")
			log.Cerror(err)
			return
		}

		result = append(result, resbuf[0], resbuf[1])
	}

	return
}

func calculateCrc(data []uint8) (crc uint8) {

	crc = 0xff

	for _, byt := range data {
		crc ^= byt
		for b := 0; b < 8; b++ {
			if (crc & 0x80) != 0 {
				crc = (crc << 1) ^ 0x31
			} else {
				crc <<= 1
			}
		}
	}

	return
}
