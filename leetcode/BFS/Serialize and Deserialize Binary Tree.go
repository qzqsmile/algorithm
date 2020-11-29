package BFS

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {

}

func Constructor() Codec {
	c := Codec{}
	return c
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	res := ""
	if root == nil{
		return res
	}
	q := [] *TreeNode{root}
	for;len(q) > 0;{
		n := q[0]
		q = q[1:]
		if n == nil{
			res += " #"
		}else{
			res += " " + strconv.Itoa(n.Val)
			q = append(q, n.Left)
			q = append(q, n.Right)
		}
	}
	return res
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0{
		return nil
	}
	q := []*TreeNode{}
	sp := mysplit(data)
	var r *TreeNode
	if v, e := strconv.Atoi(sp[0]); e == nil{
		r = &TreeNode{v, nil, nil}
	}
	q = append(q, r)
	i := 1
	for;len(q) > 0;{
		n := q[0]
		q = q[1:]
		if i < len(sp){
			if sp[i] != "#"{
				if v, e := strconv.Atoi(sp[i]); e == nil{
					node := &TreeNode{v, nil, nil}
					n.Left = node
					q = append(q, n.Left)
				}
			}
			i++
			if sp[i] != "#"{
				if v, e := strconv.Atoi(sp[i]); e == nil{
					node := &TreeNode{v, nil, nil}
					n.Right = node
					q = append(q, n.Right)
				}
			}
			i++
		}
	}

	return r
}

func mysplit(res string) []string{
	r := []string{}
	for b := 0; b < len(res); {
		if res[b] == ' '{
			b++
			continue
		}
		e := b
		for ; e < len(res) && res[e] != ' '; e++{}
		r = append(r, res[b:e])
		b = e
	}
	return r
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */