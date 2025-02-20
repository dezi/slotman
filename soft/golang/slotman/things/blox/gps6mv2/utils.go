package gps6mv2

func bytesAreEqual(b1, b2 []byte) (equal bool) {

	if b1 == nil || b2 == nil || len(b1) != len(b2) {
		return
	}

	for inx := 0; inx < len(b1); inx++ {
		if b1[inx] != b2[inx] {
			return
		}
	}

	equal = true
	return
}
