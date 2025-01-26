package simple

func IntInArray(haystack []int, needle int) bool {

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
