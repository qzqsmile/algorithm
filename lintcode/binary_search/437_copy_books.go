package main

/**
 * @param pages: an array of integers
 * @param k: An integer
 * @return: an integer
 */
func copyBooks (pages []int, k int) int {
	// write your code here
	if len(pages) == 0{
		return 0
	}
	b, e := 0, 0
	for i := 0; i < len(pages); i++{
		b = max(b, pages[i])
		e += pages[i]
	}

	for;b + 1 < e;{
		m := b + (e - b)/2
		if check(pages, m) > k{
			b = m
		}else{
			e = m
		}
	}
	if check(pages, b) <= k{
		return b
	}
	return e
}

func check(pages []int, m int) int{
	t, c := 0, 0
	for i := 0; i < len(pages); i++{
		if t + pages[i] > m{
			c++
			t = 0
		}
		t += pages[i]
	}
	c++
	return c
}


func max(a int, b int) int{
	if a > b{
		return a
	}
	return b
}
