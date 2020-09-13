package LinkedList

import "net/http"

func main(){
	dummy := ListNode(0)
	dummy.next = head

	listNode curt = dummy
	while curt != nil {
		curt = reverseNextK(curt, k)
	}

	return dummy.next
}