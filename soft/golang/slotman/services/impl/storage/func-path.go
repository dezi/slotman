package storage

import (
	"errors"
	"os"
	"runtime"
	"slotman/services/iface/identity"
	"slotman/utils/log"
	"strings"
)

func (sv *Service) GetPathPlain() (storagePath string, err error) {

	caller := sv.getCallerPackage(5)
	if caller == "unknown" {
		err = errors.New("cannot resolve call stack")
		log.Cerror(err)
		return
	}

	storagePath = identity.GetStoragePath() +
		"/" + string(plain) +
		"/" + caller

	return
}

func (sv *Service) GetPlainStoragePathForPackage() (storagePath string) {

	storagePath = identity.GetStoragePath() +
		"/" + string(plain) +
		"/" + sv.getCallerPackage(5)

	_, err := os.Stat(storagePath)
	if err == nil {
		return
	}

	err = os.MkdirAll(storagePath, 0755)
	if err != nil {
		log.Cerror(err)
		return
	}

	return
}

func (sv *Service) getCallerPackage(frame int) (caller string) {

	buf := make([]byte, 2048)
	siz := runtime.Stack(buf, false)
	str := strings.Split(string(buf[:siz]), "\n")

	if len(str) <= frame {
		return "unknown"
	}

	caller = str[frame]

	bis := strings.Index(caller, "(")
	if bis < 0 {
		log.Printf("%s", str[frame])
		return "unknown"
	}
	caller = caller[:bis]

	bis = strings.LastIndex(caller, ".")
	if bis < 0 {
		log.Printf("%s", str[frame])
		return "unknown"
	}

	caller = caller[:bis]

	parts := strings.Split(caller, "/")
	if len(parts) < 3 {
		log.Printf("%s", str[frame])
		return "unknown"
	}

	caller = strings.Join(parts[2:], "~")
	caller = strings.TrimPrefix(caller, "project.go.liesa.main~")

	return caller
}
