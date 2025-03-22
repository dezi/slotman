package i2c

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slotman/utils/log"
	"strings"
	"syscall"
	"time"
)

// GetDevicePaths Retrieve all I2C device paths.
func GetDevicePaths() (devicePaths []string, err error) {

	dirEntries, err := os.ReadDir("/dev")
	if err != nil {
		return
	}

	for _, dirEntry := range dirEntries {
		if strings.HasPrefix(dirEntry.Name(), "i2c-") && len(dirEntry.Name()) == 5 {
			devicePaths = append(devicePaths, filepath.Join("/dev", dirEntry.Name()))
			continue
		}
	}

	return
}

// Open I2C-connection.
func (i2c *Device) Open() (err error) {

	i2c.rc, err = os.OpenFile(i2c.device, os.O_RDWR, 0600)
	if err != nil {
		return
	}

	//
	// Attach device address to file descriptor.
	//

	err = ioctl(i2c.rc.Fd(), 0x0703, uintptr(i2c.addr))
	if err != nil {
		return
	}

	return
}

// Close I2C-connection.
func (i2c *Device) Close() (err error) {
	err = i2c.rc.Close()
	return
}

func (i2c *Device) TransLock() (err error) {

	transLockDA := fmt.Sprintf("%s-%02x", i2c.device, i2c.addr)

	//
	// A transaction lock can be acquired for max 1 second.
	// After that time, a new lock is unconditionally granted.
	//

	for try := 0; try < 12; try++ {

		transLock.Lock()

		if transLocks[transLockDA] == 0 {
			transLocks[transLockDA] = time.Now().UnixMilli()
			transLock.Unlock()

			if i2c.addr == 0x59 {
				log.Printf("TransLock clean...")
			}

			return
		}

		if time.Now().UnixMilli()-transLocks[transLockDA] > 1000 {
			transLocks[transLockDA] = time.Now().UnixMilli()
			transLock.Unlock()

			if i2c.addr == 0x59 {
				log.Printf("TransLock dirty...")
			}

			return
		}

		transLock.Unlock()

		if i2c.addr == 0x59 {
			log.Printf("TransLock wait...")
		}

		time.Sleep(time.Millisecond * 100)
	}

	if i2c.addr == 0x59 {
		log.Printf("TransLock fail...")
	}

	err = errors.New("transaction lock not acquired")
	return
}

func (i2c *Device) TransUnlock() (err error) {

	if i2c.addr == 0x59 {
		log.Printf("TransUnlock...")
	}

	transLockDA := fmt.Sprintf("%s-%02x", i2c.device, i2c.addr)

	transLock.Lock()
	defer transLock.Unlock()

	//
	// Unlock unconditionally.
	//

	transLocks[transLockDA] = 0
	return
}

func (i2c *Device) Write(data []byte) (xfer int, err error) {

	locks[i2c.device].Lock()
	defer locks[i2c.device].Unlock()

	for try := 1; try <= 2; try++ {

		xfer, err = i2c.rc.Write(data)
		if err == nil {
			return
		}

		time.Sleep(time.Millisecond)
	}

	txt := strings.Replace(err.Error(), ": ", fmt.Sprintf("-%02x: ", i2c.addr), 1)
	err = errors.New(txt)

	return
}

func (i2c *Device) Read(data []byte) (xfer int, err error) {

	locks[i2c.device].Lock()
	defer locks[i2c.device].Unlock()

	for try := 1; try <= 2; try++ {

		xfer, err = i2c.rc.Read(data)
		if err == nil {
			return
		}

		time.Sleep(time.Millisecond)
	}

	txt := strings.Replace(err.Error(), ": ", fmt.Sprintf("-%02x: ", i2c.addr), 1)
	err = errors.New(txt)

	return
}

// ReadUart
//
// Specialized function to read high-speed
// I2C dual uart devices of type SC16IS752 by proxy.
func (i2c *Device) ReadUart(channel byte, timeOut int, data []byte) (xfer int, err error) {

	//
	// Required SC16IS752 registers.
	//

	var RegRHR byte = 0x00
	var RegRxLvl byte = 0x09

	var size = len(data)
	var startTime = time.Now().UnixMilli()

	//log.Printf("ReadUart channel=%d timeOut=%d size=%d dev=%s addr=%02x",
	//	channel, timeOut, size, i2c.device, i2c.addr)

	var avail byte
	var temp []byte

	for size > len(temp) {

		for {

			avail, err = i2c.ReadRegByte(RegRxLvl<<3 | channel<<1)
			if err != nil {
				log.Cerror(err)
				return
			}

			if avail > 0 {
				break
			}

			if timeOut > 0 && time.Now().UnixMilli()-startTime > int64(timeOut) {
				err = errors.New("read timeout")
				return
			}
		}

		if int(avail) > size-len(temp) {
			avail = byte(size - len(temp))
		}

		var read []byte

		read, _, err = i2c.ReadRegBytes(RegRHR<<3|channel<<1, int(avail))
		if err != nil {
			log.Printf("ReadRegBytes fail should not happen...")
			log.Cerror(err)
			return
		}

		temp = append(temp, read...)
		xfer = len(temp)

		copy(data, temp)

		//
		// If data was available sleep one
		// millisecond to allow new data in
		// flow to come in.
		//

		time.Sleep(time.Millisecond * 1)
	}

	return
}

func ioctl(fd, cmd, arg uintptr) (err error) {

	_, _, err = syscall.Syscall6(syscall.SYS_IOCTL, fd, cmd, arg, 0, 0, 0)

	if err.Error() == "errno 0" {
		err = nil
	}

	return
}
