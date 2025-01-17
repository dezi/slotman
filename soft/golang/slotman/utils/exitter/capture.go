package exitter

import (
	"fmt"
	"os"
	"os/signal"
	"slotman/utils/log"
	"syscall"
)

func CaptureSignals(handler *Handler) {

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGSEGV)

loop:
	for {

		switch <-sigs {

		case syscall.SIGHUP:
			fmt.Println()
			log.Printf("SigHup...")
			if handler.SigHup() {
				break loop
			}

		case syscall.SIGINT:
			fmt.Println()
			log.Printf("SigInt...")
			if handler.SigInt() {
				break loop
			}

		case syscall.SIGTERM:
			fmt.Println()
			log.Printf("SigTerm...")
			if handler.SigTerm() {
				break loop
			}

		case syscall.SIGSEGV:
			fmt.Println()
			log.Printf("SigSegv...")
			if handler.SigSegv() {
				break loop
			}
		}
	}
}
