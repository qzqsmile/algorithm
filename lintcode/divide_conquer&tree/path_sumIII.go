package main

//not conquer&divide it's travese

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) int {
	mp := make(map[int]int)
	mp[0] = 1
	return helper(root, sum, 0, mp)

}

func helper(root *TreeNode, sum int, cur int, mp map[int]int) int{
	if root == nil{
		return 0
	}
	res := 0
	cur += root.Val
	if _, val := mp[cur-sum]; val{
		res = mp[cur-sum]
	}
	if _, val := mp[cur]; val{
		mp[cur] += 1
	}else{
		mp[cur] = 1
	}
	res += helper(root.Left, sum, cur, mp) + helper(root.Right, sum, cur, mp)
	mp[cur] -= 1
	return res
}
