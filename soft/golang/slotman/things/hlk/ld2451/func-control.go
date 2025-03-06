package ld2451

import (
	"errors"
	"slotman/drivers/impl/uart"
	"slotman/utils/log"
	"time"
)

func (se *LD2451) SetHandler(handler Handler) {
	se.handler = handler
}

func (se *LD2451) EnableConfigurations() (protocol int, err error) {

	result, err := se.writeAndRead(
		commandEnableConfigurations,
		[]byte{0x01, 0x00}, 8)

	if err != nil {
		log.Cerror(err)
		return
	}

	protocol = int(result[4]) | int(result[5])<<8

	if protocol != 0x0001 {
		err = errors.New("invalid protocol version")
		return
	}

	return
}

func (se *LD2451) EndConfiguration() (err error) {

	_, err = se.writeAndRead(commandEndConfiguration, nil, 4)
	log.Cerror(err)

	return
}

func (se *LD2451) FactoryReset() (err error) {

	_, err = se.writeAndRead(commandRestoreFactorySetting, nil, 4)
	log.Cerror(err)

	return
}

func (se *LD2451) RestartModule() (err error) {

	_, err = se.writeAndRead(commandRestartModule, nil, 4)
	if err != nil {
		log.Cerror(err)
		return
	}

	time.Sleep(time.Millisecond * 1000)

	return
}

func (se *LD2451) SetBaudRate(baudRate int) (err error) {

	value := validBaudRates[baudRate]
	if value == 0 {
		err = errors.New("invalid baud rate")
		return
	}

	_, err = se.writeAndRead(
		commandSetBaudrate,
		[]byte{byte(value), byte(value >> 8)}, 4)

	if err != nil {
		log.Cerror(err)
		return
	}

	err = se.RestartModule()
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

	_, err = se.EnableConfigurations()
	log.Cerror(err)

	return
}

func (se *LD2451) GetVersion() (fType, major, minor int, err error) {

	result, err := se.writeAndRead(commandReadFirmwareVersion, nil, 12)
	if err != nil {
		log.Cerror(err)
		return
	}

	fType = int(result[4]) + int(result[5])<<8
	major = int(result[6]) + int(result[7])<<8
	minor = int(result[8]) + int(result[9])<<8 +
		int(result[10])<<16 + int(result[11])<<24

	return
}

func (se *LD2451) GetDetectionParams() (
	maxDist, minSpeed, delay byte,
	mode DetectionMode, err error) {

	result, err := se.writeAndRead(commandGetDetection, nil, 8)
	if err != nil {
		log.Cerror(err)
		return
	}

	maxDist = result[4]
	minSpeed = result[6]
	delay = result[7]
	mode = DetectionMode(result[5])

	return
}

func (se *LD2451) SetDetectionParams(
	maxDist, minSpeed, delay byte,
	mode DetectionMode) (err error) {

	data := make([]byte, 4)
	data[0] = maxDist
	data[1] = byte(mode)
	data[2] = minSpeed
	data[3] = delay

	_, err = se.writeAndRead(commandSetDetection, data, 4)
	log.Cerror(err)

	time.Sleep(time.Millisecond * 250)
	return
}

func (se *LD2451) GetSensitivityParams() (trigger, noise byte, err error) {

	result, err := se.writeAndRead(commandGetSensitivity, nil, 8)
	if err != nil {
		log.Cerror(err)
		return
	}

	trigger = result[4]
	noise = result[5]

	return
}

func (se *LD2451) SetSensitivityParams(trigger, noise byte) (err error) {

	data := make([]byte, 4)
	data[0] = trigger
	data[1] = noise
	data[2] = 0
	data[3] = 0

	_, err = se.writeAndRead(commandSetSensitivity, data, 4)
	log.Cerror(err)

	time.Sleep(time.Millisecond * 250)
	return
}
