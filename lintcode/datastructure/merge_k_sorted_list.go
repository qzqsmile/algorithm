package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


 type ListNode struct {
	     Val int
	     Next *ListNode
 }

func mergeKLists(lists []*ListNode) *ListNode {
	var head, cur *ListNode
	head, cur = nil, nil
	heap := []*ListNode{}
	for _, v := range lists{
		if v != nil{
			heap = append(heap, v)
		}
	}
	makeHeap1(heap)
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
			heapSort1(heap, 0)
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

func makeHeap1(points []*ListNode) {
	k := len(points)
	for i := k / 2; i >= 0; i-- {
		heapSort1(points, i)
	}
}

func heapSort1(points []*ListNode, i int) {
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
		heapSort1(points, minIndex)
	}

}

func main(){
	//[[1,4,5],[1,3,4],[2,6]]
	t1 := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}
	t2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{
		Val:  4,
		Next: nil,
	}}}
	t3 := &ListNode{
		Val:  2,
		Next: &ListNode{
			Val:  6,
			Next: nil,
		},
	}
	mergeKLists([]*ListNode{t1,t2,t3})
	fmt.Print("---")
}
