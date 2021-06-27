package unionfindset


func findRedundantConnection(edges [][]int) []int {
	parent := make([]int, len(edges)+1)
	for i := 0; i < len(parent); i++{
		parent[i] = i
	}

	for i := 0; i < len(edges); i++{
		if find(parent, edges[i][0]) == find(parent, edges[i][1]){
			return edges[i]
		}else{
			union(parent, edges[i][0], edges[i][1])
		}
	}
	return []int{}
}
//
//func union(parent []int, index1 int, index2 int){
//	parent[find(parent, index1)] = find(parent, index2)
//}
//
//func find(parent []int, index int) int{
//	for parent[index] != index{
//		parent[index] = parent[parent[index]]
//		index = parent[index]
//	}
//	return index
//}