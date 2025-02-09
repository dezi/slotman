package simple

import "strings"

func StringInArray(haystack []string, needle string) bool {

	if haystack == nil {
		return false
	}

	for _, value := range haystack {
		if value == needle {
			return true
		}
	}

	return false
}

func FirstUpper(str string) (newStr string) {

	newStr = str

	if len(str) > 0 {
		runes := []rune(str)
		newStr = strings.ToUpper(string(runes[:1]))
		newStr += string(runes[1:])
	}

	return
}
