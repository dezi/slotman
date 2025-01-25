package uart

import "go.bug.st/serial"

type Device struct {
	Path     string
	BaudRate int

	port serial.Port
}
