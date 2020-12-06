package BFS
func findOrder(numCourses int, prerequisites [][]int) []int {
	indegree := make(map[int]int)
	mp := make(map[int] []int)

	for i := 0; i  < numCourses; i++{
		indegree[i] = 0
	}
	for _, v := range prerequisites{
		indegree[v[0]]++
		mp[v[1]] = append(mp[v[1]], v[0])
	}

	q := []int{}
	res := []int{}
	for i := 0; i < numCourses; i++{
		if indegree[i] == 0{
			q = append(q, i)
			res = append(res, i)
		}
	}

	for;len(q) > 0;{
		n := q[0]
		q = q[1:]
		for _, v := range mp[n]{
			indegree[v]--
			if indegree[v] == 0{
				res = append(res, v)
				q = append(q, v)
			}
		}
	}
	if len(res) != numCourses{
		return []int{}
	}
	return res
}

