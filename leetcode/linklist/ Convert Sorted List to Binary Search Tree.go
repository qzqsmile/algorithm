/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
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
 func sortedListToBST(head *ListNode) *TreeNode {
    if head == nil{
        return nil
    }
    nums := []int{}
    cur := head 
    
    for;cur != nil;{
        nums = append(nums, cur.Val)
        cur = cur.Next
    }
    return convert(nums, 0, len(nums)-1)
}   


func convert(nums []int, b int, e int) *TreeNode{
    if b > e{
        return nil
    }
    m := (b+e)/2
    node := &TreeNode{nums[m], nil, nil}
    node.Left = convert(nums, b, m-1)
    node.Right = convert(nums, m+1, e)
    return node
}