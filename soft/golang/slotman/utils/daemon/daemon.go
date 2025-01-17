package daemon

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
	"syscall"
)

func Daemonize(startup func()) {

	for _, arg := range os.Args[1:] {
		if arg == "-f" {
			daemonize()
			return
		}
	}

	startup()
}

func daemonize() {

	daemonUser := "slotman"
	daemonGroup := "slotman"

	var childArgs []string

	args := os.Args[1:]
	for index := 0; index < len(args); index++ {

		arg := args[index]

		if arg == "-f" {
			continue
		}

		if arg == "-p" {
			if index+1 < len(args) {
				index++
			}
			continue
		}

		if arg == "-u" {
			if index+1 < len(args) {
				index++
				daemonUser = args[index]
			}
			continue
		}

		if arg == "-g" {
			if index+1 < len(args) {
				index++
				daemonGroup = args[index]
			}
			continue
		}

		childArgs = append(childArgs, arg)
	}

	uxUser, err := user.Lookup(daemonUser)
	if err != nil {
		panic(err)
	}

	uxGroup, err := user.LookupGroup(daemonGroup)
	if err != nil {
		panic(err)
	}

	uid, err := strconv.ParseInt(uxUser.Uid, 10, 32)
	if err != nil {
		panic(err)
	}

	gid, err := strconv.ParseInt(uxGroup.Gid, 10, 32)
	if err != nil {
		panic(err)
	}

	groupIdsUint32, err := getGroups(daemonUser)
	if err != nil {
		panic(err)
	}

	cmd := exec.Command(os.Args[0], childArgs...)
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.SysProcAttr.Credential = &syscall.Credential{
		Uid:    uint32(uid),
		Gid:    uint32(gid),
		Groups: groupIdsUint32,
	}

	err = cmd.Start()
	if err != nil {
		panic(err)
	}

	writePid(cmd.Process.Pid)

	os.Exit(0)
}

func getGroups(daemonUser string) (groups []uint32, err error) {

	cOut, err := exec.Command("groups", daemonUser).CombinedOutput()
	if err != nil {
		panic(err)
	}

	groupNames := strings.Split(string(cOut), " ")

	for _, groupName := range groupNames {

		var uxGroup *user.Group
		uxGroup, err = user.LookupGroup(groupName)
		if err != nil {
			return
		}

		var uxGid int64
		uxGid, err = strconv.ParseInt(uxGroup.Gid, 10, 32)
		if err != nil {
			return
		}

		groups = append(groups, uint32(uxGid))
	}

	return
}

func writePid(pid int) {

	var pidFile string

	args := os.Args[1:]
	for index := 0; index < len(args); index++ {

		arg := args[index]

		if arg == "-p" && index+1 < len(args) {
			index++
			pidFile = args[index]
			continue
		}
	}

	if pidFile != "" {
		pidString := fmt.Sprintf("%d", pid)
		_ = os.WriteFile(pidFile, []byte(pidString), 0644)
	}
}
