package i2c

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"
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

func (i2c *Device) Write(data []byte) (xfer int, err error) {

	locks[i2c.device].Lock()
	defer locks[i2c.device].Unlock()

	xfer, err = i2c.rc.Write(data)

	if err != nil {
		txt := strings.Replace(err.Error(), ": ", fmt.Sprintf("%02x: ", i2c.addr), 1)
		err = errors.New(txt)
	}

	return
}

func (i2c *Device) Read(data []byte) (xfer int, err error) {

	locks[i2c.device].Lock()
	defer locks[i2c.device].Unlock()

	xfer, err = i2c.rc.Read(data)

	if err != nil {
		txt := strings.Replace(err.Error(), ": ", fmt.Sprintf("%02x: ", i2c.addr), 1)
		err = errors.New(txt)
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
