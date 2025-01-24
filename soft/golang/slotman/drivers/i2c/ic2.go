package i2c

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
)

type I2C struct {
	device string
	addr   uint8
	rc     *os.File
}

var (
	locks = make(map[string]*sync.Mutex)
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

// NewI2C opens a connection for I2C-device.
// SMBus (System Management Bus) protocol over I2C
// supported as well: you should preliminarily specify
// register address to read from, either write register
// together with the data in case of write operations.
func NewI2C(device string, addr uint8) (i2c *I2C) {

	i2c = &I2C{
		device: device,
		addr:   addr,
	}

	//
	// Create global mutexes for device and device address.
	//

	if _, ok := locks[i2c.device]; !ok {
		locks[i2c.device] = &sync.Mutex{}
	}

	return
}

// Open I2C-connection.
func (i2c *I2C) Open() (err error) {

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
func (i2c *I2C) Close() (err error) {
	err = i2c.rc.Close()
	return
}

// GetDevice return bus line where I2C-device is allocated.
func (i2c *I2C) GetDevice() (device string) {
	device = i2c.device
	return
}

// GetAddr return device occupied address in the bus.
func (i2c *I2C) GetAddr() uint8 {
	return i2c.addr
}

// WriteBytes send bytes to the remote I2C-device. The interpretation of
// the message is implementation-dependent.
func (i2c *I2C) WriteBytes(buf []byte) (int, error) {
	return i2c.write(buf)
}

// BeginTransaction Begin a write and read transaction for device.
func (i2c *I2C) BeginTransaction() {

	locks[i2c.device].Lock()

	return
}

// EndTransaction End a write and read transaction for device.
func (i2c *I2C) EndTransaction() {

	locks[i2c.device].Unlock()

	return
}

// ReadBytes read bytes from I2C-device.
// Number of bytes read correspond to buf parameter length.
func (i2c *I2C) ReadBytes(buf []byte) (xfer int, err error) {
	xfer, err = i2c.read(buf)
	return
}

// ReadRegBytes read count of n byte's sequence from I2C-device
// starting from reg address.
// SMBus (System Management Bus) protocol over I2C.
func (i2c *I2C) ReadRegBytes(reg byte, n int) (data []byte, xfer int, err error) {

	_, err = i2c.WriteBytes([]byte{reg})
	if err != nil {
		return nil, 0, err
	}

	data = make([]byte, n)
	xfer, err = i2c.ReadBytes(data)
	if err != nil {
		return
	}

	return
}

// ReadRegByte reads byte from I2C-device register specified in reg.
// SMBus (System Management Bus) protocol over I2C.
func (i2c *I2C) ReadRegByte(reg byte) (value byte, err error) {

	_, err = i2c.WriteBytes([]byte{reg})
	if err != nil {
		return 0, err
	}

	buf := make([]byte, 1)
	_, err = i2c.ReadBytes(buf)
	if err != nil {
		return
	}

	value = buf[0]
	return
}

// WriteReg writes byte to I2C-device register specified in reg.
// SMBus (System Management Bus) protocol over I2C.
func (i2c *I2C) WriteReg(reg byte) (err error) {
	_, err = i2c.WriteBytes([]byte{reg})
	return
}

// WriteRegByte writes byte to I2C-device register specified in reg.
// SMBus (System Management Bus) protocol over I2C.
func (i2c *I2C) WriteRegByte(reg byte, value byte) (err error) {
	_, err = i2c.WriteBytes([]byte{reg, value})
	return
}

// WriteRegBytes writes bytes to I2C-device register specified in reg.
// SMBus (System Management Bus) protocol over I2C.
func (i2c *I2C) WriteRegBytes(reg byte, data []byte) (err error) {

	buf := make([]byte, 1)
	buf[0] = reg
	buf = append(buf, data...)

	_, err = i2c.WriteBytes(buf)
	return
}

// ReadRegU16BE reads unsigned big endian word (16 bits)
// from I2C-device starting from address specified in reg.
// SMBus (System Management Bus) protocol over I2C.
func (i2c *I2C) ReadRegU16BE(reg byte) (value uint16, err error) {

	_, err = i2c.WriteBytes([]byte{reg})
	if err != nil {
		return
	}

	buf := make([]byte, 2)
	_, err = i2c.ReadBytes(buf)
	if err != nil {
		return
	}

	value = uint16(buf[0])<<8 + uint16(buf[1])
	return
}

// ReadRegU16LE reads unsigned little endian word (16 bits)
// from I2C-device starting from address specified in reg.
// SMBus (System Management Bus) protocol over I2C.
func (i2c *I2C) ReadRegU16LE(reg byte) (value uint16, err error) {

	value, err = i2c.ReadRegU16BE(reg)
	if err != nil {
		return
	}

	value = (value&0xff)<<8 + value>>8
	return
}

// ReadRegS16BE reads signed big endian word (16 bits)
// from I2C-device starting from address specified in reg.
// SMBus (System Management Bus) protocol over I2C.
func (i2c *I2C) ReadRegS16BE(reg byte) (value int16, err error) {

	_, err = i2c.WriteBytes([]byte{reg})
	if err != nil {
		return
	}

	buf := make([]byte, 2)
	_, err = i2c.ReadBytes(buf)
	if err != nil {
		return
	}

	value = int16(buf[0])<<8 + int16(buf[1])
	return
}

// ReadRegS16LE reads signed little endian word (16 bits)
// from I2C-device starting from address specified in reg.
// SMBus (System Management Bus) protocol over I2C.
func (i2c *I2C) ReadRegS16LE(reg byte) (value int16, err error) {

	value, err = i2c.ReadRegS16BE(reg)
	if err != nil {
		return
	}

	value = (value&0xff)<<8 + value>>8
	return

}

// WriteRegU16BE writes unsigned big endian word (16 bits)
// value to I2C-device starting from address specified in reg.
// SMBus (System Management Bus) protocol over I2C.
func (i2c *I2C) WriteRegU16BE(reg byte, value uint16) (err error) {
	buf := []byte{reg, byte((value & 0xff00) >> 8), byte(value & 0xff)}
	_, err = i2c.WriteBytes(buf)
	return
}

// WriteRegU16LE writes unsigned little endian word (16 bits)
// value to I2C-device starting from address specified in reg.
// SMBus (System Management Bus) protocol over I2C.
func (i2c *I2C) WriteRegU16LE(reg byte, value uint16) (err error) {
	w := (value*0xff00)>>8 + value<<8
	err = i2c.WriteRegU16BE(reg, w)
	return
}

// WriteRegS16BE writes signed big endian word (16 bits)
// value to I2C-device starting from address specified in reg.
// SMBus (System Management Bus) protocol over I2C.
func (i2c *I2C) WriteRegS16BE(reg byte, value int16) (err error) {
	buf := []byte{reg, byte((uint16(value) & 0xff00) >> 8), byte(value & 0xff)}
	_, err = i2c.WriteBytes(buf)
	return
}

// WriteRegS16LE writes signed little endian word (16 bits)
// value to I2C-device starting from address specified in reg.
// SMBus (System Management Bus) protocol over I2C.
func (i2c *I2C) WriteRegS16LE(reg byte, value int16) (err error) {
	w := int16((uint16(value)*0xff00)>>8) + value<<8
	err = i2c.WriteRegS16BE(reg, w)
	return
}

func (i2c *I2C) write(buf []byte) (xfer int, err error) {
	xfer, err = i2c.rc.Write(buf)
	return
}

func (i2c *I2C) read(buf []byte) (xfer int, err error) {
	xfer, err = i2c.rc.Read(buf)
	return
}

func ioctl(fd, cmd, arg uintptr) (err error) {

	_, _, err = syscall.Syscall6(syscall.SYS_IOCTL, fd, cmd, arg, 0, 0, 0)

	if err.Error() == "errno 0" {
		err = nil
	}

	return
}
