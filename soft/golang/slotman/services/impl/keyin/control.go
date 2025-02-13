package keyin

import (
	"os"
	"slotman/utils/log"
	"slotman/utils/simple"
)

func (sv *Service) DoControlTask() {
	sv.checkReader()
}

func (sv *Service) checkReader() {

	if sv.consoleReader != nil {
		return
	}

	sv.consoleReader = os.Stdin

	err := simple.SetKeyboardMode(true)
	if err != nil {
		log.Cerror(err)
		return
	}

	go sv.looper()
}
