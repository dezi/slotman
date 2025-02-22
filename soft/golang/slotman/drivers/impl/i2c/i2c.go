package i2c

import (
	"fmt"
	"sync"
)

var (
	locks = make(map[string]*sync.Mutex)

	transLock  sync.Mutex
	transLocks = make(map[string]int64)
)

// NewDevice opens a connection for I2C-device.
// SMBus (System Management Bus) protocol over I2C
// supported as well: you should preliminarily specify
// register address to read from, either write register
// together with the data in case of write operations.
func NewDevice(device string, addr uint8) (i2c *Device) {

	i2c = &Device{
		device: device,
		addr:   addr,
		lock:   sync.Mutex{},
	}

	//
	// Create global mutexes for device.
	//

	if _, ok := locks[i2c.device]; !ok {
		locks[i2c.device] = &sync.Mutex{}
	}

	//
	// Create transaction mutexes for device plus addr.
	//

	transLockDA := fmt.Sprintf("%s-%02x", i2c.device, i2c.addr)

	transLock.Lock()

	if _, ok := locks[transLockDA]; !ok {
		transLocks[transLockDA] = 0
	}

	transLock.Unlock()

	return
}

// GetDevice return bus line where I2C-device is allocated.
func (i2c *Device) GetDevice() (device string) {
	device = i2c.device
	return
}

// GetAddr return device occupied address in the bus.
func (i2c *Device) GetAddr() (addr uint8) {
	addr = i2c.addr
	return
}

// WriteBytes send bytes to the remote I2C-device. The interpretation of
// the message is implementation-dependent.
func (i2c *Device) WriteBytes(data []byte) (xfer int, err error) {
	xfer, err = i2c.Write(data)
	return
}

// ReadBytes read bytes from I2C-device.
// Number of bytes read correspond to buf parameter length.
func (i2c *Device) ReadBytes(data []byte) (xfer int, err error) {
	xfer, err = i2c.Read(data)
	return
}

// ReadRegByte reads byte from I2C-device register specified in reg.
// SMBus (System Management Bus) protocol over I2C.
func (i2c *Device) ReadRegByte(reg byte) (value byte, err error) {

	i2c.lock.Lock()
	defer i2c.lock.Unlock()

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

// ReadRegBytes read count of n byte's sequence from I2C-device
// starting from reg address.
// SMBus (System Management Bus) protocol over I2C.
func (i2c *Device) ReadRegBytes(reg byte, n int) (data []byte, xfer int, err error) {

	i2c.lock.Lock()
	defer i2c.lock.Unlock()

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

// WriteReg writes byte to I2C-device register specified in reg.
// SMBus (System Management Bus) protocol over I2C.
func (i2c *Device) WriteReg(reg byte) (err error) {
	_, err = i2c.WriteBytes([]byte{reg})
	return
}

// WriteRegByte writes byte to I2C-device register specified in reg.
// SMBus (System Management Bus) protocol over I2C.
func (i2c *Device) WriteRegByte(reg byte, value byte) (err error) {
	_, err = i2c.WriteBytes([]byte{reg, value})
	return
}

// WriteRegBytes writes bytes to I2C-device register specified in reg.
// SMBus (System Management Bus) protocol over I2C.
func (i2c *Device) WriteRegBytes(reg byte, data []byte) (err error) {

	buf := make([]byte, 1)
	buf[0] = reg
	buf = append(buf, data...)

	_, err = i2c.WriteBytes(buf)
	return
}

// ReadRegU16BE reads unsigned big endian word (16 bits)
// from I2C-device starting from address specified in reg.
// SMBus (System Management Bus) protocol over I2C.
func (i2c *Device) ReadRegU16BE(reg byte) (value uint16, err error) {

	i2c.lock.Lock()
	defer i2c.lock.Unlock()

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
func (i2c *Device) ReadRegU16LE(reg byte) (value uint16, err error) {

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
func (i2c *Device) ReadRegS16BE(reg byte) (value int16, err error) {

	i2c.lock.Lock()
	defer i2c.lock.Unlock()

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
func (i2c *Device) ReadRegS16LE(reg byte) (value int16, err error) {

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
func (i2c *Device) WriteRegU16BE(reg byte, value uint16) (err error) {
	buf := []byte{reg, byte((value & 0xff00) >> 8), byte(value & 0xff)}
	_, err = i2c.WriteBytes(buf)
	return
}

// WriteRegU16LE writes unsigned little endian word (16 bits)
// value to I2C-device starting from address specified in reg.
// SMBus (System Management Bus) protocol over I2C.
func (i2c *Device) WriteRegU16LE(reg byte, value uint16) (err error) {
	w := (value*0xff00)>>8 + value<<8
	err = i2c.WriteRegU16BE(reg, w)
	return
}

// WriteRegS16BE writes signed big endian word (16 bits)
// value to I2C-device starting from address specified in reg.
// SMBus (System Management Bus) protocol over I2C.
func (i2c *Device) WriteRegS16BE(reg byte, value int16) (err error) {
	buf := []byte{reg, byte((uint16(value) & 0xff00) >> 8), byte(value & 0xff)}
	_, err = i2c.WriteBytes(buf)
	return
}

// WriteRegS16LE writes signed little endian word (16 bits)
// value to I2C-device starting from address specified in reg.
// SMBus (System Management Bus) protocol over I2C.
func (i2c *Device) WriteRegS16LE(reg byte, value int16) (err error) {
	w := int16((uint16(value)*0xff00)>>8) + value<<8
	err = i2c.WriteRegS16BE(reg, w)
	return
}
