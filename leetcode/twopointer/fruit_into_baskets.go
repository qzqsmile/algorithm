package main

func totalFruit(tree []int) int {
	left, right, sum := 0, 0, 0
	win := make(map[int]int)
	for right < len(tree){
		c := tree[right]
		right++
		win[c]++
		if len(win) <= 2{
			sum = max(sum, right-left)
		}
		for len(win) > 2{
			c := tree[left]
			left++
			win[c]--
			if win[c] == 0{
				delete(win, c)
			}
		}
	}
	return sum
}

func max(a int, b int) int{
	if a > b{
		return a
	}
	return b
}

