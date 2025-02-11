package simple

import (
	"encoding/json"
	"strings"
)

func MarshalJsonClean(data interface{}) (clean []byte, err error) {

	clean, err = json.MarshalIndent(data, "", "  ")
	if err != nil {
		return
	}

	//
	// Get rid of ampersand UTF-code.
	//

	clean = []byte(MarshalDefuck(string(clean)))
	return
}

func MarshalDefuck(fucked string) (unFucked string) {
	unFucked = strings.Replace(fucked, "\\u0026", "&", -1)
	unFucked = strings.Replace(unFucked, "\\u003c", "<", -1)
	unFucked = strings.Replace(unFucked, "\\u003e", ">", -1)
	return
}
