package mcp23017

func (se *MCP23017) SetHandler(handler Handler) {
	se.handler = handler
}

func (se *MCP23017) SetPinDirections(directions uint16) (err error) {

	err = se.i2cDev.WriteRegByte(byte(RegisterIoDirB), byte(directions>>8))
	if err != nil {
		return
	}

	err = se.i2cDev.WriteRegByte(byte(RegisterIoDirA), byte(directions))
	return
}

func (se *MCP23017) GetPinDirections() (directions uint16, err error) {

	dirsB, err := se.i2cDev.ReadRegByte(byte(RegisterIoDirB))
	if err != nil {
		return
	}

	dirsA, err := se.i2cDev.ReadRegByte(byte(RegisterIoDirA))
	if err != nil {
		return
	}

	directions = uint16(dirsB)<<8 | uint16(dirsA)
	return
}

func (se *MCP23017) SetPinDirection(pin Pin, dir PinDirection) (err error) {

	if pin >= 16 {
		err = ErrInvalidPin
		return
	}

	if dir != PinDirectionInput && dir != PinDirectionOutput {
		err = ErrInvalidDir
		return
	}

	reg := RegisterIoDirA
	if pin >= 8 {
		pin -= 8
		reg = RegisterIoDirB
	}

	dirs, err := se.i2cDev.ReadRegByte(byte(reg))
	if err != nil {
		return
	}

	dirs ^= 1 << pin
	dirs |= byte(dir) << pin

	err = se.i2cDev.WriteRegByte(byte(reg), dirs)
	if err != nil {
		return
	}

	return
}

func (se *MCP23017) GetPinDirection(pin Pin) (dir PinDirection, err error) {

	if pin >= 16 {
		err = ErrInvalidPin
		return
	}

	reg := RegisterIoDirA
	if pin >= 8 {
		pin -= 8
		reg = RegisterIoDirB
	}

	dirs, err := se.i2cDev.ReadRegByte(byte(reg))
	if err != nil {
		return
	}

	dir = PinDirection(0x1 & (dirs >> pin))
	return
}

func (se *MCP23017) WritePin(pin Pin, val PinLogic) (err error) {

	if pin >= 16 {
		err = ErrInvalidPin
		return
	}

	if val != PinLogicLo && val != PinLogicHi {
		err = ErrInvalidLogic
		return
	}

	reg := RegisterOlatA
	if pin >= 8 {
		pin -= 8
		reg = RegisterOlatB
	}

	bits, err := se.i2cDev.ReadRegByte(byte(reg))
	if err != nil {
		return
	}

	bits ^= 1 << pin
	bits |= byte(val) << pin

	err = se.i2cDev.WriteRegByte(byte(reg), bits)
	return
}

func (se *MCP23017) ReadPin(pin Pin) (val PinLogic, err error) {

	if pin >= 16 {
		err = ErrInvalidPin
		return
	}

	reg := RegisterGpioA
	if pin >= 8 {
		pin -= 8
		reg = RegisterGpioB
	}

	bits, err := se.i2cDev.ReadRegByte(byte(reg))
	if err != nil {
		return
	}

	val = PinLogic(0x1 & (bits >> pin))
	return
}

func (se *MCP23017) WritePins(values uint16) (err error) {

	err = se.i2cDev.WriteRegByte(byte(RegisterOlatB), byte(values>>8))
	if err != nil {
		return
	}

	err = se.i2cDev.WriteRegByte(byte(RegisterOlatA), byte(values))
	return
}

func (se *MCP23017) ReadPins() (values uint16, err error) {

	valuesB, err := se.i2cDev.ReadRegByte(byte(RegisterGpioB))
	if err != nil {
		return
	}

	valuesA, err := se.i2cDev.ReadRegByte(byte(RegisterGpioA))
	if err != nil {
		return
	}

	values = uint16(valuesB)<<8 | uint16(valuesA)
	return
}
