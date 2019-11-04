package main

func getRow(rowIndex int) []int {
	var res[][]int = [][]int{[]int{1}, []int{1,1}}
	if rowIndex <= 1 {
		return res[rowIndex]
	}
	lastest_row := res[len(res)-1]
	for i := 2; i <= rowIndex; i++{
		l1 := append([]int{0},lastest_row...)
		l2 := append(lastest_row, []int{0}...)
		lastest_row = lastest_row[:0]
		for j := 0; j < len(l1); j++{
			lastest_row = append(lastest_row, l1[j]+l2[j])
		}

	}
	return lastest_row
}

func main()  {

}
