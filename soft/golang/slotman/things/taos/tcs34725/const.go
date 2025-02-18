package tcs34725

// sudo modprobe i2c-dev
// i2cdetect -F 1
// i2cdetect -y 1

type Register byte

//goland:noinspection GoUnusedConst
const (
	ThingAddress = 0x29

	RegisterEnable  Register = 0x80 // Enables states and interrupts
	RegisterATime   Register = 0x81 // RGBC time
	RegisterWTime   Register = 0x83 // Wait time
	RegisterAiltL   Register = 0x84 // Clear interrupt low threshold low byte
	RegisterAiltH   Register = 0x85 // Clear interrupt low threshold high byte
	RegisterAihtL   Register = 0x86 // Clear interrupt high threshold low byte
	RegisterAihtH   Register = 0x87 // Clear interrupt high threshold high byte
	RegisterPers    Register = 0x8c // Interrupt persistence filter
	RegisterConfig  Register = 0x8d // Configuration
	RegisterControl Register = 0x8f // Control
	RegisterId      Register = 0x92 // Device ID
	RegisterStatus  Register = 0x93 // Device status
	RegisterCDataL  Register = 0x94 // Clear data low byte
	RegisterCDataH  Register = 0x95 // Clear data high byte
	RegisterRDataL  Register = 0x96 // Red data low byte
	RegisterRDataH  Register = 0x97 // Red data high byte
	RegisterGDataL  Register = 0x98 // Green data low byte
	RegisterGDataH  Register = 0x99 // Green data high byte
	RegisterBDataL  Register = 0x9a // Blue data low byte
	RegisterBDataH  Register = 0x9b // Blue data high byte
)

const (
	EnableWEN = byte(0x08) // Wait Enable - Writing 1 activates the wait timer
	EnableAEN = byte(0x02) // RGBC Enable - Writing 1 actives the ADC, 0 disables it
	EnablePON = byte(0x01) // Power on - Writing 1 activates the internal oscillator, 0 disables it
)

//
// Integration time settings for TCS34725
//
// 60-Hz period: 16.67ms, 50-Hz period: 20ms
// 100ms is evenly divisible by 50Hz periods and by 60Hz periods
//

type IntegrationTime byte

//goland:noinspection GoUnusedConst
const (
	IntegrationTime2dot4ms = 0xFF // 2.4ms - 1 cycle - Max Count: 1024
	IntegrationTime24ms    = 0xF6 // 24.0ms - 10 cycles - Max Count: 10240
	IntegrationTime50ms    = 0xEB // 50.4ms - 21 cycles - Max Count: 21504
	IntegrationTime60ms    = 0xE7 // 60.0ms - 25 cycles - Max Count: 25700
	IntegrationTime101ms   = 0xD6 // 100.8ms - 42 cycles - Max Count: 43008
	IntegrationTime120ms   = 0xCE // 120.0ms - 50 cycles - Max Count: 51200
	IntegrationTime154ms   = 0xC0 // 153.6ms - 64 cycles - Max Count: 65535
	IntegrationTime180ms   = 0xB5 // 180.0ms - 75 cycles - Max Count: 65535
	IntegrationTime199ms   = 0xAD // 199.2ms - 83 cycles - Max Count: 65535
	IntegrationTime240ms   = 0x9C // 240.0ms - 100 cycles - Max Count: 65535
	IntegrationTime300ms   = 0x83 // 300.0ms - 125 cycles - Max Count: 65535
	IntegrationTime360ms   = 0x6A // 360.0ms - 150 cycles - Max Count: 65535
	IntegrationTime401ms   = 0x59 // 400.8ms - 167 cycles - Max Count: 65535
	IntegrationTime420ms   = 0x51 // 420.0ms - 175 cycles - Max Count: 65535
	IntegrationTime480ms   = 0x38 // 480.0ms - 200 cycles - Max Count: 65535
	IntegrationTime499ms   = 0x30 // 499.2ms - 208 cycles - Max Count: 65535
	IntegrationTime540ms   = 0x1F // 540.0ms - 225 cycles - Max Count: 65535
	IntegrationTime600ms   = 0x06 // 600.0ms - 250 cycles - Max Count: 65535
	IntegrationTime614ms   = 0x00 // 614.4ms - 256 cycles - Max Count: 65535
)

type Gain byte

//goland:noinspection GoUnusedConst
const (
	Gain1x  Gain = 0x00 // No gain
	Gain4x  Gain = 0x01 // 4x gain
	Gain16x Gain = 0x02 // 16x gain
	Gain60x Gain = 0x03 // 60x gain
)
