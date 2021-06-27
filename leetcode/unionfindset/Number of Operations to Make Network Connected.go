package unionfindset

func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1{
		return -1
	}
	parent := make([]int, n)
	for i := 0; i < len(parent); i++{
		parent[i] = i
	}

	for i := 0; i < len(connections); i++{
		union(parent,connections[i][0], connections[i][1])
	}

	count := 0
	mp := make(map[int] bool)
	for i := 0; i < n; i++{
		if _, ok := mp[find(parent, i)]; !ok{
			count++
			mp[find(parent, i)] = true
		}
	}

	return count-1
}


//func union(parent[]int, index1 int, index2 int){
//	parent[find(parent, index1)] = find(parent, index2)
//}
//
//// key point
//func find(parent []int, index int) int{
//	for parent[index] != index{
//		parent[index] = parent[parent[index]]
//		index = parent[index]
//	}
//	return index
//}

