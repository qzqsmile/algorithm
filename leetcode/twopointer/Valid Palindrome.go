package twopointer

func isPalindrome(s string) bool {
	b, e := 0, len(s)-1
	for ; b < e; {
		for ; b < e && !isletter(s[b]); b++ {
		}
		for ; b < e && !isletter(s[e]); e-- {
		}
		if b < e {
			if !equal(s[b], s[e]){
				return false
			}
			b++
			e--
		}

	}
	return true
}

func isletter(a uint8) bool {
	if ('a' <= a && a <= 'z') || ('A' <= a && a <= 'Z') || ('0' <= a && a <= '9'){
		return true
	}
	return false
}

func equal(a uint8,  b uint8) bool {
	if 'A' <= a && a <= 'Z'{
		a = 'a' + a-'A'
	}
	if 'A' <= b && b <= 'Z'{
		b = 'a' + b-'A'
	}
	if a == b {
		return true
	}
	return false
}