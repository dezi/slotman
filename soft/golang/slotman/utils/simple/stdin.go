package simple

import (
	"golang.org/x/sys/unix"
	"os"
)

func SetKeyboardMode(raw bool) (err error) {

	tio, err := unix.IoctlGetTermios(int(os.Stdin.Fd()), ioctlReadTermios)
	if err != nil {
		return
	}

	if raw {
		tio.Lflag &^= unix.ECHO | unix.ECHONL | unix.ICANON
	} else {
		tio.Lflag |= unix.ECHO | unix.ECHONL | unix.ICANON
	}

	err = unix.IoctlSetTermios(int(os.Stdin.Fd()), ioctlWriteTermios, tio)
	return
}
