package mcp23017

// i2cdetect -F 1
// i2cdetect -y 1

var (
	ThingAddresses = []byte{0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27}
)

type Register byte

//goland:noinspection GoUnusedConst
const (

	//
	// 3.5.1 I/O DIRECTION REGISTER
	//
	// Controls the direction of the data I/O.
	// When a bit is set, the corresponding pin becomes an input.
	// When a bit is clear, the corresponding pin becomes an output.
	//

	RegisterIoDirA Register = 0x00
	RegisterIoDirB Register = 0x01

	IoDirOutput = 0x0 // 0 = Pin is configured as an output.
	IoDirInput  = 0x1 // 1 = Pin is configured as an input.

	//
	// 3.5.2 INPUT POLARITY REGISTER
	//
	// This register allows the user to configure
	// the polarity on the corresponding GPIO port bits.
	// If a bit is set, the corresponding GPIO register bit
	// will reflect the inverted value on the pin.
	//

	RegisterIPolA Register = 0x02
	RegisterIPolB Register = 0x03

	IPolSame   = 0x0 // 0 = GPIO register bit reflects the same logic state of the input pin.
	IPolInvert = 0x1 // 1 = GPIO register bit reflects the opposite logic state of the input pin.

	//
	// 3.5.3 INTERRUPT-ON-CHANGE CONTROL REGISTER
	//
	// The GPINTEN register controls the interrupt-on-change
	// feature for each pin.
	// If a bit is set, the corresponding pin is enabled for interrupt-on-change.
	// The DEFVAL and INTCON registers must also be configured if any
	// pins are enabled for interrupt-on-change.
	//

	RegisterGpIntEnA Register = 0x04
	RegisterGpIntEnB Register = 0x05

	GpIntDis = 0x0 // 0 = Disables GPIO input pin for interrupt-on-change event.
	GpIntEna = 0x1 // 1 = Enables GPIO input pin for interrupt-on-change event.

	//
	// 3.5.4 DEFAULT COMPARE REGISTER FOR INTERRUPT-ON-CHANGE
	// The default comparison value is configured in the DEFVAL register.
	// If enabled (via GPINTEN and INTCON) to compare against the DEFVAL
	// register, an opposite value on the associated pin will cause an interrupt to occur.
	//

	RegisterDefValA Register = 0x06
	RegisterDefValB Register = 0x07

	//
	// 3.5.5 INTERRUPT CONTROL REGISTER
	//
	// The INTCON register controls how the associated pin
	// value is compared for the interrupt-on-change feature.
	// If a bit is set, the corresponding I/O pin is compared
	// against the associated bit in the DEFVAL register.
	// If a bit value is clear, the corresponding I/O pin
	// is compared against the previous value.
	//

	RegisterIntConA Register = 0x08
	RegisterIntConB Register = 0x09

	IntConPrev   = 0x0 // 0 = Pin value is compared against the previous pin value.
	IntConDefVal = 0x1 // 1 = Pin value is compared against the associated bit in the DEFVAL register.

	//
	// 3.5.6 CONFIGURATION REGISTER
	//
	// Bit 7 BANK: Controls how the registers are addressed
	// 1 = The registers associated with each port are separated into different banks.
	// 0 = The registers are in the same bank (addresses are sequential).
	//
	// Bit 6 MIRROR: INT Pins Mirror bit
	// 1 = The INT pins are internally connected
	// 0 = The INT pins are not connected. INTA is associated with PORTA and INTB is associated with PORTB
	//
	// Bit 5 SEQOP: Sequential Operation mode bit
	// 1 = Sequential operation disabled, address pointer does not increment.
	// 0 = Sequential operation enabled, address pointer increments.
	//
	// Bit 4 DISSLW: Slew Rate control-bit for SDA output
	// 1 = Slew rate disabled.
	// 0 = Slew rate enabled.
	//
	// Bit 3 HAEN: Hardware Address Enable bit (MCP23S17 only) (Note 1)
	// 1 = Enables the MCP23S17 address pins.
	// 0 = Disables the MCP23S17 address pins.
	//
	// Bit 2 ODR: Configures the INT pin as an open-drain output
	// 1 = Open-drain output (overrides the INTPOL bit.)
	// 0 = Active driver output (INTPOL bit sets the polarity.)
	//
	// Bit 1 INTPOL: This bit sets the polarity of the INT output pin
	// 1 = Active-high.
	// 0 = Active-low.
	//
	// Bit 0 Unimplemented: Read as ‘0’
	//

	RegisterIoConA Register = 0x0a
	RegisterIoConB Register = 0x0b

	IoConBankBit  = 7
	IoConBankDiff = 0x1 // 1 = The registers associated with each port are separated into different banks.
	IoConBankSame = 0x0 // 0 = The registers are in the same bank (addresses are sequential).

	IoConMirrorBit = 6
	IoConMirrorInt = 0x1 // 1 = The INT pins are internally connected
	IoConMirrorUn  = 0x0 // 0 = The INT pins are not connected. INTA is associated with PORTA and INTB is associated with PORTB

	IoConSeqOpBit = 5
	IoConSeqOpDis = 0x1 // 1 = Sequential operation disabled, address pointer does not increment.
	IoConSeqOpEna = 0x0 // 0 = Sequential operation enabled, address pointer increments.

	IoConDisSlwBit = 4
	IoConDisSlwDis = 0x1 // 1 = Slew rate disabled.
	IoConDisSlwEna = 0x0 // 0 = Slew rate enabled.

	IoConHaenBit = 3
	IoConHaenEna = 0x1 // 1 = Enables the MCP23S17 address pins.
	IoConHaenDis = 0x0 // 0 = Disables the MCP23S17 address pins.

	IoConOdrBit = 2
	IoConOdrOdr = 0x1 // 1 = Open-drain output (overrides the INTPOL bit.)
	IoConOdrAct = 0x0 // 0 = Active driver output (INTPOL bit sets the polarity.)

	IoConIntPolBit = 1
	IoConIntPolHi  = 0x1 // 1 = Active-high.
	IoConIntPolLo  = 0x0 // 0 = Active-low.

	//
	// 3.5.7 PULL-UP RESISTOR CONFIGURATION REGISTER
	//
	// The GPPU register controls the pull-up resistors for the port pins.
	// If a bit is set and the corresponding pin is configured as an input,
	// the corresponding port pin is internally pulled up with a 100 kOhm resistor.
	//

	RegisterGpPuA Register = 0x0c
	RegisterGpPuB Register = 0x0d

	GpPuEna = 0x1 // 1 = Pull-up enabled.
	GpPuDis = 0x0 // 0 = Pull-up disabled.

	//
	// 3.5.8 INTERRUPT FLAG REGISTER
	//
	// The INTF register reflects the interrupt condition on the
	// port pins of any pin that is enabled for interrupts via the GPINTEN register.
	// A set bit indicates that the associated pin caused the interrupt.
	//

	RegisterIntfA Register = 0x0e
	RegisterIntfB Register = 0x0f

	IntfInter = 0x1 // 1 = Pin caused interrupt.
	IntfIdle  = 0x0 // 0 = Interrupt not pending.

	//
	// 3.5.9 INTERRUPT CAPTURED REGISTER
	//
	// The INTCAP register captures the GPIO port value at
	// the time the interrupt occurred.
	// The register is read-only and is updated only when an interrupt occurs.
	// The register remains unchanged until the interrupt is cleared via a read of INTCAP or GPIO.
	//

	RegisterIntCapA Register = 0x10
	RegisterIntCapB Register = 0x11

	IntCapHi = 0x1 // 1 = Logic-high.
	IntCapLo = 0x0 // 0 = Logic-low.

	//
	// 3.5.10 PORT REGISTER
	//
	// The GPIO register reflects the value on the port.
	// Reading from this register reads the port.
	// Writing to this register modifies the Output Latch (OLAT) register.
	//

	RegisterGpioA Register = 0x12
	RegisterGpioB Register = 0x13

	GpioHi = 0x1 // 1 = Logic-high.
	GpioLo = 0x0 // 0 = Logic-low.

	//
	// 3.5.11 OUTPUT LATCH REGISTER (OLAT)
	//
	// The OLAT register provides access to the output latches.
	// A read from this register results in a read of the OLAT and not the port itself.
	// A write to this register modifies the output latches that modifies the pins configured as outputs.
	//

	RegisterOlatA Register = 0x14
	RegisterOlatB Register = 0x15

	OlatHi = 0x1 // 1 = Logic-high.
	OlatLo = 0x0 // 0 = Logic-low.
)
