package BFS

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
	q := []*Node{}
	mp := make(map[int]*Node)
	q = append(q, node)
	visited := make(map[int] bool)
	for;len(q) > 0;{
		n := q[0]
		q = q[1:]
		visited[n.Val] = true
		if _, ok := mp[n.Val]; !ok{
			mp[n.Val] = &Node{n.Val, []*Node{}}
		}
		for _, each := range n.Neighbors{
			if _, ok := visited[each.Val]; !ok{
				q = append(q, each)
			}
		}
	}

	q = append(q, node)
	visited = make(map[int] bool)
	for;len(q) > 0;{
		n := q[0]
		q = q[1:]
		if len(mp[n.Val].Neighbors) > 0{
			continue
		}else{
			for _, each := range n.Neighbors{
				mp[n.Val].Neighbors = append(mp[n.Val].Neighbors, mp[each.Val])
			}
		}
		for _, each := range n.Neighbors{
			if _, ok := visited[each.Val]; !ok{
				q = append(q, each)
			}
		}
	}
	return mp[node.Val]
}
