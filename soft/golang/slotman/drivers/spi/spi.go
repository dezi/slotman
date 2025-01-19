package spi

import (
	"errors"
	"os"
	"slotman/drivers/ioctl"
	"sync"
	"unsafe"
)

//
// http://piface.github.io/pifacecommon/installation.html#enable-the-spi-module
//

type Device struct {
	Path string
	Bus  int
	Chip int

	file  *os.File
	mode  uint8
	bpw   uint8
	speed uint32
	delay uint16
}

var (
	validDevicePaths = []string{
		"/dev/spidev0.0",
		"/dev/spidev0.1",
		"/dev/spidev1.0",
		"/dev/spidev1.1",
		"/dev/spidev1.2"}

	globalLock sync.Mutex
)

func GetDevicePaths() (devicePaths []string, err error) {

	for _, devicePath := range validDevicePaths {
		_, tryErr := os.Stat(devicePath)
		if tryErr == nil {
			devicePaths = append(devicePaths, devicePath)
		}
	}

	return
}

func NewDevice(devicePath string) (spi *Device) {
	spi = &Device{Path: devicePath}
	return
}

func (spi *Device) Open() (err error) {
	spi.file, err = os.OpenFile(spi.Path, os.O_RDWR, 0)
	return
}

func (spi *Device) Close() (err error) {

	if spi.file != nil {
		err = spi.file.Close()
		spi.file = nil
	}

	return
}

func (spi *Device) Send(request []byte) (result []byte, err error) {

	file := spi.file
	if file == nil {
		err = errors.New("device not open")
		return
	}

	var wBuffer [256 * 1024]byte
	var rBuffer [256 * 1024]byte

	if len(request) > len(wBuffer) {
		return nil, errors.New("request size to large")
	}

	//copy(request, wBuffer[:])

	for index, byt := range request {
		wBuffer[index] = byt
	}

	transfer := SdIoctlTransfer{}
	transfer.txBuf = uint64(uintptr(unsafe.Pointer(&request)))
	transfer.rxBuf = uint64(uintptr(unsafe.Pointer(&rBuffer)))
	transfer.length = uint32(len(request))
	transfer.delayUSecs = spi.delay
	transfer.bitsPerWord = spi.bpw
	transfer.speedHz = spi.speed

	globalLock.Lock()

	err = ioctl.IOCTL(file.Fd(), SdIoctlWrCustomMessage(1), uintptr(unsafe.Pointer(&transfer)))

	globalLock.Unlock()

	if err != nil {
		err = errors.New("ioctl send request failed")
		return
	}

	result = make([]byte, len(request))
	for index := range result {
		result[index] = rBuffer[index]
	}

	return
}

func (spi *Device) SetMode(mode uint8) (err error) {

	spi.mode = mode

	err = ioctl.IOCTL(spi.file.Fd(), SdIoctlWrMode(), uintptr(unsafe.Pointer(&mode)))
	if err != nil {
		err = errors.New("error setting spi mode")
		return err
	}

	return
}

func (spi *Device) SetBitsPerWord(bpw uint8) (err error) {

	spi.bpw = bpw

	err = ioctl.IOCTL(spi.file.Fd(), SdIoctlWrBitsPerWord(), uintptr(unsafe.Pointer(&bpw)))
	if err != nil {
		err = errors.New("error setting bits per word")
		return
	}

	return
}

func (spi *Device) SetSpeed(speed uint32) (err error) {

	spi.speed = speed

	err = ioctl.IOCTL(spi.file.Fd(), SdIoctlWrMaxSpeedHz(), uintptr(unsafe.Pointer(&speed)))
	if err != nil {
		err = errors.New("error setting speed")
		return
	}

	return
}
