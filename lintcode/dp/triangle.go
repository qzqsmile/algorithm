package main

/**
 * @param triangle: a list of lists of integers
 * @return: An integer, minimum path sum
 */
func minimumTotal(triangle [][]int) int {
	// write your code here
	if len(triangle) == 0 {
		return 0
	}
	l1, l2 := []int{triangle[0][0]}, []int{triangle[0][0]}

	for i := 1; i <= len(triangle)-1; i++ {
		l2 = []int{}
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				l2 = append(l2, triangle[i][j]+l1[j])
			} else if j == len(triangle[i])-1 {
				l2 = append(l2, triangle[i][j]+l1[j-1])
			} else {
				l2 = append(l2, min(triangle[i][j]+l1[j-1], triangle[i][j]+l1[j]))
			}
		}
		l1 = l2
	}
	r := l2[0]
	for i := 0; i < len(l2); i++ {
		if l2[i] < r {
			r = l2[i]
		}
	}
	return r
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {

}
