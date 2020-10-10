package main

import "fmt"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil{
		return nil
	}
	visited := make(map[int] bool)
	mp := make(map[int] *Node)
	q := []*Node{node}
	mp[q[0].Val] = &Node{q[0].Val, []*Node{}}

	for;len(q) > 0;{
		t := q[0]
		q = q[1:]
		_, ok := visited[t.Val]
		if !ok {
			for i := 0; i < len(t.Neighbors); i++ {
				if _, val := mp[t.Neighbors[i].Val]; !val {
					mp[t.Neighbors[i].Val] = &Node{t.Neighbors[i].Val, []*Node{}}
				}
				if _, val := visited[t.Neighbors[i].Val]; !val {
					q = append(q, t.Neighbors[i])
				}
				mp[t.Val].Neighbors = append(mp[t.Val].Neighbors, mp[t.Neighbors[i].Val])
			}
		}
		visited[t.Val] = true
	}
	return mp[node.Val]
}