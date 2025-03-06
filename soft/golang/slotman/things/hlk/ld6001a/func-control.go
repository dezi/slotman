package ld6001a

import (
	"fmt"
	"slotman/drivers/impl/uart"
	"slotman/things"
	"slotman/utils/log"
	"time"
)

func (se *LD6001a) SetHandler(handler Handler) {
	se.handler = handler
}

func (se *LD6001a) SetBaudRate(baudRate int) (err error) {

	if !se.IsOpen {
		err = things.ErrThingNotOpen
		log.Cerror(err)
		return
	}

	command := fmt.Sprintf(commandBaudrate, baudRate)

	err = se.writeWithOk(command)
	if err != nil {
		log.Cerror(err)
		return
	}

	//
	// Re-open serial port with new baudrate
	//

	_ = se.uart.Close()

	se.uart = uart.NewDevice(se.DevicePath, baudRate)
	err = se.uart.Open()
	if err != nil {
		se.IsOpen = false
		se.IsStarted = false
		return
	}

	_ = se.uart.SetReadTimeout(time.Millisecond * 100)

	return
}
