package simple

import (
	"os/user"
	"path/filepath"
	"slotman/utils/log"
	"strings"
)

func UserHomeDir() (home string, err error) {

	usr, err := user.Current()
	if err != nil {
		log.Cerror(err)
		return
	}

	home = usr.HomeDir
	return
}

func CurrentDir() (dir string) {

	dir, err := filepath.Abs(filepath.Dir("."))
	if err != nil {
		log.Cerror(err)
		return
	}

	return
}

func ResolvePath(path string) (resolved string, err error) {

	home, err := UserHomeDir()
	if err != nil {
		return
	}

	resolved = strings.Replace(path, "~", home, 1)
	return
}
