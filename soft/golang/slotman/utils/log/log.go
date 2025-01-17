package log

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	pidLen        = 12
	callerLen     = 28
	logMutex      sync.Mutex
	debugPacks    []string
	remoteHandler func(err error)
)

func SetRemoteErrorHandle(handler func(err error)) {
	remoteHandler = handler
}

//goland:noinspection GoUnusedExportedFunction
func AddDebugPack(debugPack string) {
	logMutex.Lock()
	debugPacks = append(debugPacks, debugPack)
	logMutex.Unlock()
}

func SetCallerLength(length int) {
	callerLen = length
}

//
// Debug logging.
//

//goland:noinspection GoUnusedExportedFunction
func Derror(err error) error {

	if logDat() {
		normalLn(false, "Fail! err=%s", err)
	}

	return err
}

func Debugf(format string, v ...interface{}) {

	if logDat() {
		normalLn(true, format, v...)
	}
}

//
// Normal logging.
//

func Cerror(err error) {

	if err != nil {

		if remoteHandler != nil {
			remoteHandler(err)
		}

		normalLn(false, "Fail! err=%s", err)
	}
}

func Perror(err error) error {

	if err != nil {

		if remoteHandler != nil {
			remoteHandler(err)
		}

		normalLn(false, "Fail! err=%s", err)
	}

	return err
}

func Ferror(err error) {

	if err != nil {

		if remoteHandler != nil {
			remoteHandler(err)
		}

		fatalLn("Fail! err=%s", err)
	}
}

func Werror(err error) {

	if err != nil {
		normalLn(false, fmt.Sprintf("Warn: %s.", err.Error()))
	}
}

func Printf(format string, v ...interface{}) {
	normalLn(true, format, v...)
}

func PrintfFrame(frame int, format string, v ...interface{}) {
	normalLnFrame(frame, true, format, v...)
}

//
// Fatal logging WITH exit.
//

func Fatalf(format string, v ...interface{}) {
	fatalLn(format, v...)
}

func fatalLn(format string, v ...interface{}) {
	timeLine := time.Now().Format("2006/01/02 15:04:05.000")
	textLine := pFormat(9, false, format, v...)
	logMutex.Lock()
	fmt.Println(timeLine + " " + textLine)
	logMutex.Unlock()
	os.Exit(1)
}

func normalLn(noNegs bool, format string, v ...interface{}) {
	timeLine := time.Now().Format("2006/01/02 15:04:05.000")
	textLine := pFormat(9, noNegs, format, v...)
	if textLine != "" {
		logMutex.Lock()
		fmt.Println(timeLine + " " + textLine)
		logMutex.Unlock()
	}
}

func normalLnFrame(frame int, noNegs bool, format string, v ...interface{}) {
	timeLine := time.Now().Format("2006/01/02 15:04:05.000")
	textLine := pFormat(frame, noNegs, format, v...)
	if textLine != "" {
		logMutex.Lock()
		fmt.Println(timeLine + " " + textLine)
		logMutex.Unlock()
	}
}

func getFrame(frame int) (goroutine int, caller string) {

	buf := make([]byte, 2048)
	siz := runtime.Stack(buf, false)
	str := strings.Split(string(buf[:siz]), "\n")

	if len(str) <= frame {
		return 0, "unknown"
	}

	temp := strings.Split(str[0], " ")

	if len(temp) >= 3 {
		goroutine, _ = strconv.Atoi(temp[1])
	}

	stack := str[frame]

	von := strings.LastIndex(stack, "/")
	bis := strings.LastIndex(stack, "(")

	//
	// If von == -1 did not find the string.
	// -1 is just the value we want in this case.
	// LOL
	//

	if bis >= 0 {
		caller = stack[von+1 : bis]
	} else {
		Printf("%s", str[7])
	}

	return goroutine, caller
}

func pFormat(frame int, noNegs bool, format string, v ...interface{}) string {

	goroutine, caller := getFrame(frame)

	if noNegs {

		hostName, _ := os.Hostname()

		if len(positives[hostName]) != 0 {

			isPositive := false

			for _, pos := range positives[hostName] {
				if strings.HasPrefix(caller, pos) {
					isPositive = true
					break
				}
			}

			if !isPositive {
				return ""
			}

		} else {

			for _, neg := range negatives[hostName] {
				if strings.HasPrefix(caller, neg) {
					return ""
				}
			}
		}
	}

	pidFmt := fmt.Sprintf("%0d/%d", os.Getpid(), goroutine)

	if len(pidFmt) > pidLen {
		pidLen = len(pidFmt)
	}

	for len(pidFmt) < pidLen {
		pidFmt += " "
	}

	//
	// Dezi hates this shit in his logs...
	//

	caller = strings.ReplaceAll(caller, "paho%2emqtt%2egolang", "pahoMqtt")

	caller += ":"

	for len(caller) < callerLen {
		caller += " "
	}

	extFmt := fmt.Sprintf("%s %s", pidFmt, caller)
	orgFmt := fmt.Sprintf(format, v...)

	return extFmt + " " + orgFmt
}

func logDat() (logIt bool) {

	if debugPacks != nil {

		_, caller := getFrame(7)

		for _, pack := range debugPacks {

			if pack == "*" {
				logIt = true
				continue
			}

			if strings.HasPrefix(caller, pack+".") {
				logIt = true
				continue
			}

			if strings.HasPrefix(pack, "!") &&
				strings.HasPrefix(caller, pack[1:]+".") {
				logIt = false
				continue
			}
		}
	}

	return
}
