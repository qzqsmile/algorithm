package BFS

// topological sort 重要的一点在于要看每个节点的入度是否为0，这是可以进队列的条件
// 而不能简单的采用BFS，访问到了就可以进队列的方法

func canFinish(numCourses int, prerequisites [][]int) bool {
	indegree := make(map[int]int)
	mp := make(map[int] []int)

	for i := 0; i  < numCourses; i++{
		indegree[i] = 0
	}
	for _, v := range prerequisites{
		indegree[v[1]]++
		mp[v[0]] = append(mp[v[0]], v[1])
	}

	q := []int{}
	for i := 0; i < numCourses; i++{
		if indegree[i] == 0{
			q = append(q, i)
		}
	}
	count := len(q)
	for;len(q) > 0;{
		n := q[0]
		q = q[1:]
		for _, v := range mp[n]{
			indegree[v]--
			if indegree[v] == 0{
				count++
				q = append(q, v)
			}
		}
	}
	return count == numCourses
}
