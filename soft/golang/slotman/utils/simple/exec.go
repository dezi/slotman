package simple

import (
	"os"
	"path"
)

func GetExecName() (execName string) {

	execName, err := os.Executable()
	if err != nil {
		execName = "unknown"
		return
	}

	execName = path.Base(execName)

	for execName[0] == '_' {
		execName = execName[1:]
	}

	for execName[0] == '1' || execName[0] == '2' {
		execName = execName[1:]
	}

	return
}
