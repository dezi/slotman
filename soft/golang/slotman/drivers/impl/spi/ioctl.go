package spi

import (
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
	return IOR(SdIoctlMagic, 1, 1)
}

// SdIoctlWrMode
// Write SPI mode (SPI_MODE_0..SPI_MODE_3)
func SdIoctlWrMode() uintptr {
	return IOW(SdIoctlMagic, 1, 1)
}

// SdIoctlRdLsbFirst
// Read SPI bit justification
func SdIoctlRdLsbFirst() uintptr {
	return IOR(SdIoctlMagic, 2, 1)
}

// SdIoctlWrLsbFirst
// Write SPI bit justification
func SdIoctlWrLsbFirst() uintptr {
	return IOW(SdIoctlMagic, 2, 1)
}

// SdIoctlRdBitsPerWord
// Read SPI device word length
func SdIoctlRdBitsPerWord() uintptr {
	return IOR(SdIoctlMagic, 3, 1)
}

// SdIoctlWrBitsPerWord
// write SPI device word length
func SdIoctlWrBitsPerWord() uintptr {
	return IOW(SdIoctlMagic, 3, 1)
}

// SdIoctlRdMaxSpeedHz
// Read SPI device default max speed hz
func SdIoctlRdMaxSpeedHz() uintptr {
	return IOR(SdIoctlMagic, 4, 4)
}

// SdIoctlWrMaxSpeedHz
// Write SPI device default max speed hz
func SdIoctlWrMaxSpeedHz() uintptr {
	return IOW(SdIoctlMagic, 4, 4)
}

// SdIoctlWrCustomMessage
// Write custom SPI message
func SdIoctlWrCustomMessage(n uintptr) uintptr {
	return IOW(SdIoctlMagic, 0, SdMessageSize(n))
}

func SdMessageSize(n uintptr) (size uintptr) {
	if (n * unsafe.Sizeof(SdIoctlTransfer{})) < (1 << IOC_SIZEBITS) {
		size = n * unsafe.Sizeof(SdIoctlTransfer{})
	}
	return
}

const (
	IOC_NONE  = 0
	IOC_WRITE = 1
	IOC_READ  = 2
)

func IOCTL(fd, op, arg uintptr) error {
	_, _, ep := syscall.Syscall(syscall.SYS_IOCTL, fd, op, arg)
	if ep != 0 {
		return ep
	}
	return nil
}

const (
	IOC_NRBITS   = 8
	IOC_TYPEBITS = 8

	IOC_SIZEBITS = 14
	IOC_DIRBITS  = 2

	IOC_NRSHIFT   = 0
	IOC_TYPESHIFT = IOC_NRSHIFT + IOC_NRBITS
	IOC_SIZESHIFT = IOC_TYPESHIFT + IOC_TYPEBITS
	IOC_DIRSHIFT  = IOC_SIZESHIFT + IOC_SIZEBITS
)

func IOC(dir, t, nr, size uintptr) uintptr {
	return (dir << IOC_DIRSHIFT) |
		(t << IOC_TYPESHIFT) |
		(nr << IOC_NRSHIFT) |
		(size << IOC_SIZESHIFT)
}

func IOR(t, nr, size uintptr) uintptr {
	return IOC(IOC_READ, t, nr, size)
}

func IOW(t, nr, size uintptr) uintptr {
	return IOC(IOC_WRITE, t, nr, size)
}
