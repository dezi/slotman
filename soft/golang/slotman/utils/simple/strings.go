package simple

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
