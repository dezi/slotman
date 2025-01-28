package ads1115

// i2cdetect -F 1
// i2cdetect -y 1

//goland:noinspection GoUnusedConst
var (
	ThingI2CAddresses = []byte{0x48, 0x49, 0x4a, 0x4b}
)

type Register byte

type Gain uint16
type Rate uint16

//goland:noinspection GoUnusedConst
const (
	RegisterConversion Register = 0x00
	RegisterConfig     Register = 0x01
	RegisterLoThresh   Register = 0x02
	RegisterHiThresh   Register = 0x03

	//
	// Operational status or single-shot conversion start
	// This bit determines the operational status of the device.
	// OS can only be written when in power-down state and
	// has no effect when a conversion is ongoing.
	//
	// When writing:
	// 0 : No effect
	// 1 : Start a single conversion (when in power-down state)
	//
	// When reading:
	// 0 : Device is currently performing a conversion
	// 1 : Device is not currently performing a conversion
	//

	OsShift        = 15
	OsMask  uint16 = 0x01

	OsWriteNoop  uint16 = 0x00
	OsWriteStart uint16 = 0x01
	OsReadBusy   uint16 = 0x00
	OsReadIdle   uint16 = 0x01

	//
	// Input multiplexer configuration (ADS1115 only).
	// These bits configure the input multiplexer.
	// These bits serve no function on the ADS1113 and ADS1114.
	//
	// 000 : AINP = AIN0 and AINN = AIN1 (default)
	// 001 : AINP = AIN0 and AINN = AIN3
	// 010 : AINP = AIN1 and AINN = AIN3
	// 011 : AINP = AIN2 and AINN = AIN3
	// 100 : AINP = AIN0 and AINN = GND
	// 101 : AINP = AIN1 and AINN = GND
	// 110 : AINP = AIN2 and AINN = GND
	// 111 : AINP = AIN3 and AINN = GND
	//

	MuxShift        = 12
	MuxMask  uint16 = 0x07

	Mux0And1   uint16 = 0x00
	Mux0And3   uint16 = 0x01
	Mux1And3   uint16 = 0x02
	Mux2And3   uint16 = 0x03
	Mux0AndGnd uint16 = 0x04
	Mux1AndGnd uint16 = 0x05
	Mux2AndGnd uint16 = 0x06
	Mux3AndGnd uint16 = 0x07

	//
	// Programmable gain amplifier configuration
	// These bits set the FSR of the programmable gain amplifier.
	// These bits serve no function on the ADS1113.
	//
	// 000 : FSR = ±6.144 V(1)
	// 001 : FSR = ±4.096 V(1)
	// 010 : FSR = ±2.048 V (default)
	// 011 : FSR = ±1.024 V
	// 100 : FSR = ±0.512 V
	// 101 : FSR = ±0.256 V
	// 110 : FSR = ±0.256 V
	// 111 : FSR = ±0.256 V
	//

	GainShift        = 9
	GainMask  uint16 = 0x07

	Gain0 Gain = 0x00
	Gain1 Gain = 0x01
	Gain2 Gain = 0x02
	Gain3 Gain = 0x03
	Gain4 Gain = 0x04
	Gain5 Gain = 0x05
	Gain6 Gain = 0x06
	Gain7 Gain = 0x07

	//
	// Device operating mode.
	// This bit controls the operating mode.
	//
	// 0 : Continuous-conversion mode
	// 1 : Single-shot mode or power-down state (default)
	//

	ModeShift        = 8
	ModeMask  uint16 = 0x01

	ModeContinuous uint16 = 0x00
	ModeSingleShot uint16 = 0x01

	//
	// Data rate.
	// These bits control the data rate setting.
	//
	// 000 : 8 SPS
	// 001 : 16 SPS
	// 010 : 32 SPS
	// 011 : 64 SPS
	// 100 : 128 SPS (default)
	// 101 : 250 SPS
	// 110 : 475 SPS
	// 111 : 860 SPS
	//

	RateShift        = 5
	RateMask  uint16 = 0x07

	Rate8Sps   Rate = 0x00
	Rate16Sps  Rate = 0x01
	Rate32Sps  Rate = 0x02
	Rate64Sps  Rate = 0x03
	Rate128ps  Rate = 0x04
	Rate250Sps Rate = 0x05
	Rate475Sps Rate = 0x06
	Rate860Sps Rate = 0x07

	//
	// Comparator mode (ADS1114 and ADS1115 only)
	// This bit configures the comparator operating mode.
	// This bit serves no function on the ADS1113.
	//
	// 0 : Traditional comparator (default)
	// 1 : Window comparator
	//

	CompModeShift        = 4
	CompModeMask  uint16 = 0x01

	CompModeTraditional uint16 = 0x00
	CompModeWindow      uint16 = 0x01

	//
	// Comparator polarity (ADS1114 and ADS1115 only)
	// This bit controls the polarity of the ALERT/RDY pin.
	// This bit serves no function on the ADS1113.
	//
	// 0 : Active low (default)
	// 1 : Active high
	//

	CompPolarityShift        = 3
	CompPolarityMask  uint16 = 0x01

	CompPolarityLow  uint16 = 0x00
	CompPolarityHigh uint16 = 0x01

	//
	// Latching comparator (ADS1114 and ADS1115 only)
	// This bit controls whether the ALERT/RDY pin latches after being
	// asserted or clears after conversions are within the margin
	// of the upper and lower threshold values.
	// This bit serves no function on the ADS1113.
	//
	// 0 : Non latching comparator.
	//	   The ALERT/RDY pin does not latch when asserted (default).
	// 1 : Latching comparator.
	//	   The asserted ALERT/RDY pin remains latched until
	//	   conversion data are read by the master or an
	//	   appropriate SMBus alert response is sent by the master.
	//	   The device responds with its address, and it is the lowest
	//	   address currently asserting the ALERT/RDY bus line.
	//

	CompLatchingShift        = 2
	CompLatchingMask  uint16 = 0x01

	CompLatchingOff uint16 = 0x00
	CompLatchingOn  uint16 = 0x01

	//
	// Comparator queue and disable (ADS1114 and ADS1115 only).
	// These bits perform two functions.
	// When set to 11, the comparator is disabled and the ALERT/RDY
	// pin is set to a high-impedance state.
	// When set to any other value, the ALERT/RDY pin and
	// the comparator function are enabled, and the set
	// value determines the number of successive conversions
	// exceeding the upper or lower threshold required before asserting
	// the ALERT/RDY pin. These bits serve no function on the ADS1113.
	//
	// 00 : Assert after one conversion
	// 01 : Assert after two conversions
	// 10 : Assert after four conversions
	// 11 : Disable comparator and set ALERT/RDY pin to high-impedance (default)
	//

	CompQueueShift        = 0
	CompQueueMask  uint16 = 0x03

	CompQueue1   uint16 = 0x00
	CompQueue2   uint16 = 0x01
	CompQueue4   uint16 = 0x02
	CompQueueOff uint16 = 0x03
)
