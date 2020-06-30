
func canFinish(numCourses int, prerequisites [][]int) bool {
	degree := make([]int, numCourses)
	mp := make([][]int, numCourses)

	for i, _ := range degree {
		degree[i] = 0
		mp[i] = make([]int, 0)
	}

	for _, v := range prerequisites {
		mp[v[1]] = append(mp[v[1]], v[0])
		degree[v[0]]++
	}

	for i := 0; i < numCourses; i++ {
		for j := 0; j < numCourses; j++ {
			if degree[j] == 0 {
				for _, k := range mp[j] {
					degree[k]--
				}
				degree[j] = -1
			}
		}
	}

	for _, v := range degree {
		if (v != 0) && (v != -1) {
			return false
		}
	}
	return true
}

