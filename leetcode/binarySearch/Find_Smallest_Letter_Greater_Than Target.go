package binarySearch

func nextGreatestLetter(letters []byte, target byte) byte {
	if letters[len(letters)-1] <= target || letters[0] > target {
		return letters[0]
	}
	b, e := 0, len(letters)-1
	for ; b+1 < e; {
		m := b + (e-b)/2
		if letters[m] <= target {
			b = m
		} else {
			e = m
		}
	}
	if letters[b] > target {
		return letters[b]
	}
	return letters[e]
}
