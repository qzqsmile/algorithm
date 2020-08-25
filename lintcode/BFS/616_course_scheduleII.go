package main


func findOrder(numCourses int, prerequisites [][]int) []int {
	degree := make([]int, numCourses)
	mp := make([][]int, numCourses)

	res := []int{}

	for _, v := range prerequisites {
		mp[v[1]] = append(mp[v[1]], v[0])
		degree[v[0]]++
	}

	q := []int{}

	for i, _ := range degree{
		if degree[i] == 0{
			q = append(q, i)
		}
	}

	for; len(q) > 0;{
		n := q[0]
		res = append(res, n)
		q = q[1:]
		degree[n]--
		for _, v := range mp[n]{
			degree[v]--
			if degree[v] == 0{
				q = append(q, v)
			}
		}
	}

	if numCourses == len(res){
		return res
	}
	return []int{}
}
