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

	log.Printf("Write inp command=%s", command)

	_, err = se.uart.Write([]byte(command + "\n"))
	log.Cerror(err)

	select {
	case result := <-se.results:
		if result == command {
			log.Printf("Write out command=%s success", command)
			return
		}

		if strings.HasPrefix(result, "AT+OK=") {
			log.Printf("Write out command=%s result=%s", command, result)
			return
		}

		err = errors.New(fmt.Sprintf("serial fail <%s>", result))

	case <-time.After(time.Millisecond * 250):
		err = ErrSerialTimeout
	}

	return
}
