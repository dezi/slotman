package ld2451

import (
	"errors"
	"slotman/things"
	"slotman/utils/log"
)

func (se *LD2451) writeAndRead(
	command byte, data []byte, size int) (
	result []byte, err error) {

	err = se.writeCommand(command, data)
	if err != nil {
		return
	}

	result, err = se.readResult(command)
	if err != nil {
		return
	}

	if len(result) != size || size < 2 {
		err = errors.New("invalid result length")
		log.Cerror(err)
		return
	}

	ack := int(result[2]) | int(result[3])<<8

	if ack != 0x0000 {
		err = errors.New("invalid ack")
		return
	}

	return
}

func (se *LD2451) writeCommand(command byte, data []byte) (err error) {

	if !se.IsOpen {
		err = things.ErrThingNotOpen
		log.Cerror(err)
		return
	}

	log.Printf("Write command=%02x data=[ %02x ]", command, data)

	se.lock.Lock()
	defer se.lock.Unlock()

	size := 2 + len(data)

	var buffer []byte

	buffer = append(buffer, 0xfd, 0xfc, 0xfb, 0xfa)
	buffer = append(buffer, byte(size), byte(size>>8))
	buffer = append(buffer, command, 0x00)
	buffer = append(buffer, data...)
	buffer = append(buffer, 0x04, 0x03, 0x02, 0x01)

	//log.Printf("#### write command=%02x data=[ %02x ]", command, buffer)

	_, err = se.uart.Write(buffer)
	log.Cerror(err)

	return
}
