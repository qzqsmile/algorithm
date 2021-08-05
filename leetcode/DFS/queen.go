package DFS

func solveNQueens(n int) [][]string {
	var res [][]string
	var tmp []string
	queendfs(n, 0, tmp, &res)
	return res
}

func queendfs(n int, row int, tmp []string, res *[][]string){
	if row == n{
		*res = append(*res, append([]string{}, tmp...))
	}
	for i := 0; i < n; i++{
		if meet(row, i, n, tmp){
			tmp = append(tmp, construct(i, n))
			queendfs(n, row+1, tmp, res)
			tmp = tmp[0:len(tmp)-1]
		}
	}
}

func construct(index int, n int) string{
	res := ""
	for i := 0; i < n; i++{
		if i == index{
			res += "Q"
		}else{
			res += "."
		}
	}
	return res
}

func meet(r int, c int, n int, tmp []string) bool{
	for i := r-1; i >= 0; i--{
		if tmp[i][c] == 'Q'{
			return false
		}
		if c+r-i < n && tmp[i][c+r-i] == 'Q'{
			return false
		}
		if c-r+i >=0 && tmp[i][c-r+i] == 'Q'{
			return false
		}
	}
	return true
}
