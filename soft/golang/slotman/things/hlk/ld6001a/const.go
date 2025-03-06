package ld6001a

//
// Motion target detection and tracking
//

var (
	baudRates = []int{115200}
)

const (
// AT+RANGE=XX\n
// Configure the radius of the radar detection circle projected onto the ground ( unit: cm , range: 100-500 , default value: 450 )
// AT+HEIGHTD=XXX\n
// Set vertical distance ( unit: cm , setting range: 50~500 , default value: 300)
// AT+DEBUG=X\n
// 0 : Protocol mode (default, simple protocol ) 1 : Output string
// 2 : Debug mode (used by host computer )
// 3 : Protocol mode (detailed protocol )
// AT+XPosi=XXX\n
// Configure X Forward range ( unit: cm , range: 20 ~ 500 , default value 450 )
// AT+XNega=-XXX \n
// Configure X Negative range ( unit: cm , range: -500 ~ - 20 , default - 450 )
// AT+YPosi=XXX\n
// Configuration Y Forward range ( unit: cm , range: 20 ~ 500 , default value 450 )
// AT+YNega=-XXX \n
// Configuration Y Negative range ( unit: cm , range: -500 ~ - 20 , default - 450 )
// AT+Moving=XXX\n
// Configure the moving target disappearance time ( unit 100ms, range 5~1000, default value is 110 )
// AT+Static=XXX\n
// Configure static target disappearance time ( unit 100ms, range 5~1000, default value is 100 )
// AT+Exit=XXX\n
// Configure target exit boundary time ( unit 100ms, range 2~1000, default value is 5 )
)

const (
	//
	// AT+BAUD=XX
	// Configure the serial port baud rate (default value is 115200)
	//
	commandBaudrate = "AT+BAUD=%d"
)
