package unionfindset


func findCircleNum(isConnected [][]int) int {
	parent := make([]int, len(isConnected))
	for i := 0; i < len(parent); i++{
		parent[i] = i
	}

	for i := 0; i < len(isConnected); i++{
		for j := 0; j < len(isConnected[0]); j++{
			if isConnected[i][j] == 1{
				union(parent, i, j)
			}
		}
	}

	mp := make(map[int]bool)
	count := 0
	for i := 0; i < len(parent); i++{
		if _, ok := mp[find(parent, i)]; !ok{
			count++
			mp[find(parent, i)] = true
		}
	}
	return count
}

func union(parent[]int, index1 int, index2 int){
	parent[find(parent, index1)] = find(parent, index2)
}

// key point
func find(parent []int, index int) int{
	for parent[index] != index{
		parent[index] = parent[parent[index]]
		index = parent[index]
	}
	return index
}


