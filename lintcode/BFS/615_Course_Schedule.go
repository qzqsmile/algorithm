package main


func canFinish(numCourses int, prerequisites [][]int) bool {
	degree := make([]int, numCourses)
	mp := make([][]int, numCourses)

	for _, v := range prerequisites{
		degree[v[0]]++
		mp[v[1]] = append(mp[v[1]], v[0])
	}

	q := make([]int, 0)
	for i, _ := range degree{
		if degree[i] == 0{
			q = append(q, i)
		}
	}

	for ; len(q) > 0;{
		n := q[0]
		q = q[1:]
		degree[n]--
		for _, v := range mp[n]{
			degree[v]--
			if degree[v] == 0{
				q = append(q, v)
			}
		}
	}

	for i, _ := range degree{
		if degree[i] > 0{
			return false
		}
	}
	return true
}