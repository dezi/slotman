package mxt550

import (
	"errors"
	"fmt"
	"slotman/things"
	"slotman/utils/log"
)

func (se *MXT550) sendCommand(cmd []byte) (err error) {

	if !se.IsOpen {
		err = things.ErrThingNotOpen
		return
	}

	sendCrc := (se.protocolOptions & 1 << MotoronProtocolOptionCrcForCommands) != 0
	err = se.sendCommandCore(cmd, sendCrc)

	return
}

func (se *MXT550) sendCommandCrc(cmd []byte) (err error) {

	if !se.IsOpen {
		err = things.ErrThingNotOpen
		return
	}

	err = se.sendCommandCore(cmd, true)

	return
}

func (se *MXT550) sendCommandAndReadResponse(cmd []byte, length byte) (response []byte, err error) {

	if !se.IsOpen {
		err = things.ErrThingNotOpen
		return
	}

	sendCrc := (se.protocolOptions & 1 << MotoronProtocolOptionCrcForCommands) != 0
	err = se.sendCommandCore(cmd, sendCrc)
	if err != nil {
		return
	}

	response, err = se.readResponse(cmd[0], length)
	return
}

func (se *MXT550) sendCommandCore(cmd []byte, sendCrc bool) (err error) {

	if sendCrc {
		crc := se.calculateCrc(cmd)
		cmd = append(cmd, crc)
	}

	_, err = se.i2cDev.WriteBytes(cmd)

	if se.debug {
		log.Printf("Write cmd=%s cmd=[ %02x ]", MotoronCmd2Str[MotoronCmd(cmd[0])], cmd)
	}

	if err != nil {
		err = fmt.Errorf("%s @wb=%d", err.Error(), len(cmd))
		return
	}

	return
}

func (se *MXT550) readResponse(cmdByte byte, length byte) (response []byte, err error) {

	crcEnabled := (se.protocolOptions & 1 << MotoronProtocolOptionCrcForResponses) != 0

	if crcEnabled {
		response = make([]byte, length+1)
	} else {
		response = make([]byte, length)
	}

	xfer, err := se.i2cDev.ReadBytes(response)

	if err != nil {
		return
	}

	if xfer != len(response) {
		err = errors.New("response too short")
		return
	}

	if !crcEnabled {

		if se.debug {
			log.Printf("Read name=%s res=[ %02x ]", MotoronCmd2Str[MotoronCmd(cmdByte)], response)
		}

		return
	}

	crc := response[len(response)-1]
	response = response[:len(response)-1]

	if se.debug {
		log.Printf("Read name=%s cmd=[ %02x ] (%02x)", MotoronCmd2Str[MotoronCmd(cmdByte)], response, crc)
	}

	if crc != se.calculateCrc(response) {
		err = errors.New("checksum error")
		return
	}

	return
}

func (se *MXT550) calculateCrc(message []byte) (crc byte) {

	for _, ccc := range message {

		crc ^= ccc

		for j := 0; j < 8; j++ {

			if (crc & 1) == 1 {
				crc ^= 0x91
			}

			crc >>= 1
		}
	}

	return
}
