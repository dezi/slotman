package ld6001a

import (
	"slotman/things"
	"slotman/utils/log"
)

func (se *LD6001a) writeWithOk(command string) (err error) {

	if !se.IsOpen {
		err = things.ErrThingNotOpen
		log.Cerror(err)
		return
	}

	log.Printf("Write command=%s", command)

	_, err = se.uart.Write([]byte(command + "\n"))
	log.Cerror(err)

	return
}
