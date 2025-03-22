package ld6001a

import (
	"errors"
	"fmt"
	"slotman/utils/log"
)

func (se *LD6001a) SetHandler(handler Handler) {
	se.handler = handler
}

// SetBaudRate
//
// Configure the serial port baud rate (default value is 115200)
func (se *LD6001a) SetBaudRate(baudRate int) (err error) {

	//command := fmt.Sprintf(commandAtBaudrate, baudRate)
	//
	//_ = se.writeWithOk(command)
	//
	//time.Sleep(time.Second)
	//
	//_ = se.uart.SetBaudrate(baudRate)
	//return

	_ = baudRate
	err = errors.New("wtf: command does not work")
	return
}

func (se *LD6001a) StartWorking() (err error) {
	err = se.writeWithOk(commandAtStart)
	return
}

func (se *LD6001a) StopWorking() (err error) {

	err = se.writeWithOk(commandAtStop)

	if errors.Is(err, ErrSerialTimeout) {

		//
		// Fuck this shit.
		//
		// If sensor is already stopped,
		// this command returns nothing...
		//

		err = nil
	}

	return
}

func (se *LD6001a) Reset() (err error) {
	err = se.writeWithOk(commandAtReset)
	return
}

func (se *LD6001a) Restore() (err error) {
	err = se.writeWithOk(commandAtRestore)
	return
}

func (se *LD6001a) ReadParams() (err error) {
	err = se.writeWithOk(commandAtRead)
	return
}

// SetDetectionSensitivity
//
// Long-distance detection sensitivity
// (default is 4 , range 1~9 , the larger the value, the lower the sensitivity)
func (se *LD6001a) SetSensitivity(sensitivity int) (err error) {

	command := fmt.Sprintf(commandAtDpkht, sensitivity)
	err = se.writeWithOk(command)

	return
}

// SetRange
//
// Configure the radius of the radar detection circle projected onto the ground
// (unit: cm , range: 100-500 , default value: 450)
func (se *LD6001a) SetRange(xrange int) (err error) {

	command := fmt.Sprintf(commandAtRange, xrange)
	err = se.writeWithOk(command)

	return
}

// SetHeight
//
// Set vertical distance (unit: cm , setting range: 50~500 , default value: 300)
func (se *LD6001a) SetHeight(height int) (err error) {

	command := fmt.Sprintf(commandAtHeight, height)
	err = se.writeWithOk(command)

	return
}

// SetProtocol
//
// 0 : Protocol mode (default, simple protocol)
// 1 : Output string
// 2 : Debug mode (used by host computer)
// 3 : Protocol mode (detailed protocol)
func (se *LD6001a) SetProtocol(mode int) (err error) {

	command := fmt.Sprintf(commandAtDebug, mode)
	err = se.writeWithOk(command)

	return
}

// SetRanges
//
// Configure X Forward range (unit: cm , range: 20 ~ 500 , default value 450)
// Configure X Negative range (unit: cm , range: -500 ~ - 20 , default - 450)
// Configuration Y Forward range (unit: cm , range: 20 ~ 500 , default value 450)
// Configuration Y Negative range (unit: cm , range: -500 ~ - 20 , default - 450)
func (se *LD6001a) SetRanges(xPosi, xNega, yPosi, yNega int) (err error) {

	command := fmt.Sprintf(commandAtXPosi, xPosi)

	err = se.writeWithOk(command)
	if err != nil {
		log.Cerror(err)
		return
	}

	command = fmt.Sprintf(commandAtXNega, xNega)

	err = se.writeWithOk(command)
	if err != nil {
		log.Cerror(err)
		return
	}

	command = fmt.Sprintf(commandAtYPosi, yPosi)

	err = se.writeWithOk(command)
	if err != nil {
		log.Cerror(err)
		return
	}

	command = fmt.Sprintf(commandAtYNega, yNega)

	err = se.writeWithOk(command)
	if err != nil {
		log.Cerror(err)
		return
	}

	return
}

// SetMoving
//
// Configure the moving target disappearance time
// (unit 100ms, range 5~1000, default value is 110)
func (se *LD6001a) SetMoving(time int) (err error) {

	command := fmt.Sprintf(commandAtMoving, time)
	err = se.writeWithOk(command)

	return
}

// SetStatic
//
// Configure static target disappearance time
// (unit 100ms, range 5~1000, default value is 100)
func (se *LD6001a) SetStatic(time int) (err error) {

	command := fmt.Sprintf(commandAtStatic, time)
	err = se.writeWithOk(command)

	return
}

// SetExit
//
// Configure target exit boundary time
// (unit 100ms, range 2~1000, default value is 5)
func (se *LD6001a) SetExit(time int) (err error) {

	command := fmt.Sprintf(commandAtExit, time)
	err = se.writeWithOk(command)

	return
}
