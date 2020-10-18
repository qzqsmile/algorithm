package main


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	var head, cur *ListNode
	head, cur = nil, nil
	heap := []*ListNode{}
	for _, v := range lists{
		if v != nil{
			heap = append(heap, v)
		}
	}
	makeHeap(heap)
	for len(heap) > 0{
		if head == nil{
			head = heap[0]
			cur = heap[0]
		}else{
			if heap[0].Next == nil{
				heap[0] = heap[len(heap)-1]
				heap = heap[0:len(heap)-1]
			}else{
				heap[0] = heap[0].Next
			}
			heapSort(heap, 0)
			if len(heap) > 0 {
				cur.Next = heap[0]
				cur = cur.Next
			}else{
				cur.Next = nil
			}
		}
	}
	return head
}

func makeHeap(points []*ListNode) {
	k := len(points)
	for i := k / 2; i >= 0; i-- {
		heapSort(points, i)
	}
}

func heapSort(points []*ListNode, i int) {
	l := 2*i + 1
	r := 2*i + 2
	minIndex := i
	if l < len(points) && points[l].Val < points[minIndex].Val {
		minIndex = l
	}
	if r < len(points) && points[r].Val < points[minIndex].Val{
		minIndex = r
	}
	if i != minIndex {
		points[minIndex], points[i] = points[i], points[minIndex]
		heapSort(points, minIndex)
	}

}