package linklist


/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

//对于此题，或者说对于多达数linklist的题，1.就是检查是否head为nil
//2,时刻警惕着当前指针指向nil
func copyRandomList(head *Node) *Node {
	if head == nil{
		return head
	}
	cur := head

	for;cur != nil;{
		tmp := cur.Next
		cur.Next = &Node{cur.Val, nil , nil}
		cur.Next.Next = tmp
		cur = tmp
	}

	cur = head
	for;cur != nil;{
		if cur.Random != nil{
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	cur = head
	copyhead := head.Next
	copycur := copyhead

	for;cur != nil;{
		cur.Next = copycur.Next
		cur = copycur.Next
		if cur != nil{
			copycur.Next = cur.Next
			copycur = cur.Next
		}
	}

	return copyhead
}
