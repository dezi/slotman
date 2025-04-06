package ld6001a

import (
	"errors"
	"fmt"
	"slotman/things"
	"slotman/utils/log"
	"strings"
	"time"
)

func (se *LD6001a) writeWithOk(command string) (err error) {

	if !se.IsOpen {
		err = things.ErrThingNotOpen
		log.Cerror(err)
		return
	}

	//
	// Empty any leftover commend results.
	//

	for len(se.results) > 0 {
		<-se.results
	}

	devAddr := se.DevicePath

	parts := strings.Split(devAddr, ":")
	if len(parts) == 2 {
		devAddr = parts[1]
	}

	log.Printf("Write inp dev=%s command=%s", devAddr, command)

	_, err = se.uart.Write([]byte(command + "\n"))
	log.Cerror(err)

	select {

	case result := <-se.results:
		if result == command {
			log.Printf("Write out dev=%s command=%s success", devAddr, command)
			return
		}

		if strings.HasPrefix(result, "AT+OK=") {
			log.Printf("Write out dev=%s command=%s result=%s", devAddr, command, result)
			return
		}

		err = errors.New(fmt.Sprintf("dev=%s serial fail <%s>", devAddr, result))

	case <-time.After(time.Millisecond * 2000):
		err = ErrSerialTimeout
	}

	log.Cerror(err)

	return
}
