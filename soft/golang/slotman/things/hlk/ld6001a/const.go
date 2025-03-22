package ld6001a

//
// Motion target detection and tracking
//

var (
	baudRates = []int{115200}
)

const (
	//
	// AT+START
	// Start working
	//
	commandAtStart = "AT+START"

	//
	// AT+STOP
	// Stop working
	//
	commandAtStop = "AT+STOP"

	//
	// AT+RESET
	// Reset module
	//
	commandAtReset = "AT+RESET"

	//
	// AT+READ
	// Reading parameters
	//
	commandAtRead = "AT+READ"

	//
	// AT+RESTORE
	// Restore Default Settings
	//
	commandAtRestore = "AT+RESTORE"

	//
	// AT+DPKTH
	// Long-distance detection sensitivity
	// (default is 4 , range 1~9 , the larger the value, the lower the sensitivity)
	//
	commandAtDpkht = "AT+DPKTH=%d"

	//
	// AT+BAUD=XX
	// Configure the serial port baud rate (default value is 115200)
	//
	commandAtBaudrate = "AT+BAUD=%d"

	//
	// AT+RANGE=XX\n
	// Configure the radius of the radar detection circle projected onto the ground
	// (unit: cm , range: 100-500 , default value: 450)
	//
	commandAtRange = "AT+RANGE=%d"

	//
	// AT+HEIGHTD=XXX\n
	// Set vertical distance (unit: cm , setting range: 50~500 , default value: 300)
	//
	commandAtHeight = "AT+HEIGHTD=%d"

	//
	// AT+DEBUG=X\n
	// 0 : Protocol mode (default, simple protocol)
	// 1 : Output string
	// 2 : Debug mode (used by host computer)
	// 3 : Protocol mode (detailed protocol)
	//
	commandAtDebug = "AT+DEBUG=%d"

	//
	// AT+XPosi=XXX\n
	// Configure X Forward range (unit: cm , range: 20 ~ 500 , default value 450)
	//
	commandAtXPosi = "AT+XPosi=%d"

	//
	// AT+XNega=-XXX\n
	// Configure X Negative range (unit: cm , range: -500 ~ - 20 , default - 450)
	//
	commandAtXNega = "AT+XNega=%d"

	//
	// AT+YPosi=XXX\n
	// Configuration Y Forward range (unit: cm , range: 20 ~ 500 , default value 450)
	//
	commandAtYPosi = "AT+YPosi=%d"

	//
	// AT+YNega=-XXX\n
	// Configuration Y Negative range (unit: cm , range: -500 ~ - 20 , default - 450)
	//
	commandAtYNega = "AT+YNega=%d"

	//
	// AT+Moving=XXX\n
	// Configure the moving target disappearance time
	// (unit 100ms, range 5~1000, default value is 110)
	//
	commandAtMoving = "AT+Moving=%d"

	//
	// AT+Static=XXX\n
	// Configure static target disappearance time
	// (unit 100ms, range 5~1000, default value is 100)
	//
	commandAtStatic = "AT+Static=%d"

	//
	// AT+Exit=XXX\n
	// Configure target exit boundary time
	// (unit 100ms, range 2~1000, default value is 5)
	//
	commandAtExit = "AT+Exit=%d"
)
