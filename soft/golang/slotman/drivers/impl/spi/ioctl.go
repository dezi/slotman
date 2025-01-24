package spi

import (
	"slotman/drivers/impl/ioctl"
	"syscall"
	"unsafe"
)

type SdIoctlTransfer struct {
	txBuf       uint64
	rxBuf       uint64
	length      uint32
	speedHz     uint32
	delayUSecs  uint16
	bitsPerWord uint8
	csChange    uint8
	pad         uint32
}

const SdIoctlMagic = 107

// SdIoctlRdMode
// Read SPI mode (SPI_MODE_0..SPI_MODE_3)
func SdIoctlRdMode() uintptr {
	return ioctl.IOR(SdIoctlMagic, 1, 1)
}

// SdIoctlWrMode
// Write SPI mode (SPI_MODE_0..SPI_MODE_3)
func SdIoctlWrMode() uintptr {
	return ioctl.IOW(SdIoctlMagic, 1, 1)
}

// SdIoctlRdLsbFirst
// Read SPI bit justification
func SdIoctlRdLsbFirst() uintptr {
	return ioctl.IOR(SdIoctlMagic, 2, 1)
}

// SdIoctlWrLsbFirst
// Write SPI bit justification
func SdIoctlWrLsbFirst() uintptr {
	return ioctl.IOW(SdIoctlMagic, 2, 1)
}

// SdIoctlRdBitsPerWord
// Read SPI device word length
func SdIoctlRdBitsPerWord() uintptr {
	return ioctl.IOR(SdIoctlMagic, 3, 1)
}

// SdIoctlWrBitsPerWord
// write SPI device word length
func SdIoctlWrBitsPerWord() uintptr {
	return ioctl.IOW(SdIoctlMagic, 3, 1)
}

// SdIoctlRdMaxSpeedHz
// Read SPI device default max speed hz
func SdIoctlRdMaxSpeedHz() uintptr {
	return ioctl.IOR(SdIoctlMagic, 4, 4)
}

// SdIoctlWrMaxSpeedHz
// Write SPI device default max speed hz
func SdIoctlWrMaxSpeedHz() uintptr {
	return ioctl.IOW(SdIoctlMagic, 4, 4)
}

// SdIoctlWrCustomMessage
// Write custom SPI message
func SdIoctlWrCustomMessage(n uintptr) uintptr {
	return ioctl.IOW(SdIoctlMagic, 0, SdMessageSize(n))
}

func SdMessageSize(n uintptr) (size uintptr) {
	if (n * unsafe.Sizeof(SdIoctlTransfer{})) < (1 << ioctl.IOC_SIZEBITS) {
		size = n * unsafe.Sizeof(SdIoctlTransfer{})
	}
	return
}

func IOCTL(fd, op, arg uintptr) error {
	_, _, ep := syscall.Syscall(syscall.SYS_IOCTL, fd, op, arg)
	if ep != 0 {
		return ep
	}
	return nil
}
