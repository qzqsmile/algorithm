package main

import "fmt"


func solveNQueens(n int) [][]string {
	res := [][]string{}
	each := []string{}
	var tmp string
	for i := 0;i < n; i++{
		tmp += "."
	}
	for i := 0; i < n; i++{
		each = append(each, tmp)
	}
	helper(0, n, each, &res)
	return res
}

func helper(row int, n int, each []string,res *[][]string){
	if row >= n{
		copyeach := append([]string{}, each...)
		*res = append(*res, copyeach)
		return
	}

	for i := 0; i < n; i++{
		if canpalce(each, row, i){
			each[row] = replaceAtIndex(each[row], 'Q', i)
			helper(row+1, n, each, res)
			each[row] = replaceAtIndex(each[row], '.', i)
		}
	}
}

func canpalce(each []string, row int, col int) bool{
	n := len(each)
	for i := 0; i < n; i++{
		if each[i][col] == 'Q'{
			return false
		}
		if row + i < n && col + i < n && each[row+i][col+i] == 'Q'{
			return false
		}
		if row-i >= 0 && col-i >= 0 && each[row-i][col-i] == 'Q'{
			return false
		}
		if row+i < n && col-i >= 0 && each[row+i][col-i] == 'Q'{
			return false
		}
		if col+i < n && row-i >= 0 && each[row-i][col+i] == 'Q'{
			return false
		}
	}
	return true
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
