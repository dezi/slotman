package uart

import (
	"errors"
	"go.bug.st/serial"
	"slotman/services/iface/proxy"
	"strings"
	"time"
)

const wantLocal = true

func GetDevicePaths() (devicePaths []string, err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	//goland:noinspection GoBoolExpressions
	if prx.CheckTarget() && !wantLocal {
		devicePaths, err = prx.UartGetDevicePaths()
		return
	}

	checkPaths, err := serial.GetPortsList()

	for _, checkPath := range checkPaths {

		if checkPath == "/dev/cu.wlan-debug" ||
			checkPath == "/dev/cu.debug-console" ||
			checkPath == "/dev/cu.Bluetooth-Incoming-Port" ||
			checkPath == "/dev/tty.wlan-debug" ||
			checkPath == "/dev/tty.debug-console" ||
			checkPath == "/dev/tty.Bluetooth-Incoming-Port" {
			continue
		}

		if strings.HasPrefix(checkPath, "/dev/cu.") {
			continue
		}

		devicePaths = append(devicePaths, checkPath)
	}

	for _, i2cUart := range i2cUarts {
		devicePaths = append(devicePaths, i2cUart.GetDevice())
	}

	return
}

func (uart *Device) Open() (err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	//goland:noinspection GoBoolExpressions
	if prx.CheckTarget() && !wantLocal {
		err = prx.UartOpen(uart)
		return
	}

	port, err := serial.Open(uart.Path, &serial.Mode{
		BaudRate: uart.BaudRate,
		DataBits: 8,
		Parity:   serial.NoParity,
		StopBits: serial.OneStopBit,
	})

	if err == nil {
		uart.port = port
	}

	return
}

func (uart *Device) Close() (err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	//goland:noinspection GoBoolExpressions
	if prx.CheckTarget() && !wantLocal {
		err = prx.UartClose(uart)
		return
	}

	err = uart.port.Close()
	return
}

func (uart *Device) SetReadTimeout(timeout time.Duration) (err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	//goland:noinspection GoBoolExpressions
	if prx.CheckTarget() && !wantLocal {
		err = prx.UartSetReadTimeout(uart, timeout)
		return
	}

	err = uart.port.SetReadTimeout(timeout)
	return
}

func (uart *Device) Write(data []byte) (xfer int, err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	//goland:noinspection GoBoolExpressions
	if prx.CheckTarget() && !wantLocal {
		xfer, err = prx.UartWrite(uart, data)
		return
	}

	port := uart.port
	if port == nil {
		err = errors.New("uart port is nil")
		return
	}

	xfer, err = port.Write(data)
	return
}

func (uart *Device) Read(data []byte) (xfer int, err error) {

	prx, err := proxy.GetInstance()
	if err != nil {
		return
	}

	//goland:noinspection GoBoolExpressions
	if prx.CheckTarget() && !wantLocal {
		xfer, err = prx.UartRead(uart, data)
		return
	}

	port := uart.port
	if port == nil {
		err = errors.New("uart port is nil")
		return
	}

	xfer, err = port.Read(data)
	return
}
