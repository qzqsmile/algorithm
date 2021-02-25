package binarySearch

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	b, e := 1, n
	for; b+1 < e;{
		m := b + (e-b)/2
		if isBadVersion(m){
			e = m
		}else{
			b = m
		}
	}
	if isBadVersion(b){
		return b
	}
	return e
}

