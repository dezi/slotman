package sc16is752

import (
	"fmt"
	"slotman/utils/log"
	"time"
)

func (se *SC15IS752) NewSC15IS752Uart(channel byte) (seUart *SC15IS752Uart) {

	seUart = &SC15IS752Uart{
		devicePath: fmt.Sprintf("%s-%d", se.DevicePath, channel),
		channel:    channel,
		sc15is752:  se,
	}

	return
}

func (seUart *SC15IS752Uart) Open() (err error) {
	log.Printf("UART open device=%s", seUart.devicePath)
	return
}

func (seUart *SC15IS752Uart) Close() (err error) {
	log.Printf("UART close device=%s", seUart.devicePath)
	return
}

func (seUart *SC15IS752Uart) GetDevice() (device string) {
	device = seUart.devicePath
	return
}

func (seUart *SC15IS752Uart) GetBaudrate() (baudrate int) {
	baudrate = seUart.sc15is752.baudrate[seUart.channel]
	return
}

func (seUart *SC15IS752Uart) SetBaudrate(baudrate int) (err error) {
	log.Printf("UART set baudrate device=%s baudrate=%d", seUart.devicePath, baudrate)
	err = seUart.sc15is752.SetBaudrate(seUart.channel, baudrate)
	return
}

func (seUart *SC15IS752Uart) SetReadTimeout(timeout time.Duration) (err error) {

	millis := int(timeout / time.Millisecond)

	log.Printf("UART set read timeout device=%s millis=%d", seUart.devicePath, millis)

	err = seUart.sc15is752.SetReadTimeout(seUart.channel, millis)
	return
}

func (seUart *SC15IS752Uart) Write(data []byte) (xfer int, err error) {

	//log.Printf("UART write device=%s channel=%d data=[ %02x ]",
	//	seUart.devicePath, seUart.channel, data)

	xfer, err = seUart.sc15is752.WriteUartBytes(seUart.channel, data)
	return
}

func (seUart *SC15IS752Uart) Read(data []byte) (xfer int, err error) {

	var read []byte
	xfer, read, err = seUart.sc15is752.ReadUartBytes(seUart.channel, len(data))

	//log.Printf("UART read device=%s channel=%d xfer=%d read=[ %02x ]",
	//	seUart.devicePath, seUart.channel, xfer, read)

	copy(data, read)

	return
}
