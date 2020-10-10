package main

/**
 * @param org: a permutation of the integers from 1 to n
 * @param seqs: a list of sequences
 * @return: true if it can be reconstructed only one or false
 */
func sequenceReconstruction(org []int, seqs [][]int) bool {
	// write your code here
	if len(org) == 0 {
		return true
	}
	indegree := make(map[int]int)
	mp := make(map[int][]int)
	for _, v := range seqs {
		if _, ok := indegree[v[1]]; ok {
			indegree[v[1]]++
		} else {
			indegree[v[1]] = 1
		}
		mp[v[0]] = append(mp[v[0]], v[1])
	}
	q := []int{}

	for _, v := range seqs {
		if _, ok := indegree[v[0]]; !ok {
			q = append(q, v[0])
		}
	}
	if len(q) == 0 {
		return false
	}
	l := 0
	for len(q) > 0 {
		if len(q) == 1 && l < len(org) && q[0] == org[l] {
			l++
			t := q[0]
			q = q[1:]
			for _, v := range mp[t] {
				indegree[v]--
				if indegree[v] == 0 {
					q = append(q, v)
				}
			}
		} else {
			break
		}
	}
	if l == len(org) {
		return true
	}
	return false
}
