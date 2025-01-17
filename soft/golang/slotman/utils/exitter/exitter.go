package exitter

import (
	"slotman/utils/log"
	"time"
)

type Handler struct {
}

func StartService() (err error) {
	handler := &Handler{}
	CaptureSignals(handler)
	return
}

func (ex *Handler) SigHup() (exit bool) {
	log.Printf("SigHup.")
	ex.stopAll()
	return true
}

func (ex *Handler) SigInt() (exit bool) {
	log.Printf("SigInt.")
	ex.stopAll()
	return true
}

func (ex *Handler) SigTerm() (exit bool) {
	log.Printf("SigTerm.")
	ex.stopAll()
	return true
}

func (ex *Handler) SigSegv() (exit bool) {
	log.Printf("SigSegv.")
	ex.stopAll()
	return true
}

func (ex *Handler) stopAll() {
	log.Printf("Exiting now.")
	time.Sleep(time.Millisecond * 100)
}
