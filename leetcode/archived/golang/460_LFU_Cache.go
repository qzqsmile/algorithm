package main

import "fmt"

type LFUCache struct {
	m map[int] *linknode1
	head *linknode1
	left_size int
}

type linknode1 struct{
	pre, next *linknode1
	key int
	val int
	ref int
}

func Constructor22(capacity int) LFUCache {
	var l LFUCache
	l.left_size = capacity
	l.m = make(map[int] * linknode1)
	return l
}

func (this *LFUCache) Get(key int) int {
	print(this.head)
	if node, ok := this.m[key]; ok{
		//need to refactor
		node.ref += 1
		pre, next := node.pre, node.next
		var cur *linknode1
		//remove node
		if next == nil{
			return node.val
		}else if pre == nil {
			this.head = node.next
			this.head.pre = nil
			cur = node.next
			node.next = nil
		} else{
			cur = node.next
			node.next.pre = node.pre
			node.pre.next = node.next
		}
		//
		var pre_node *linknode1
		pre_node = cur.pre
		for ; (cur != nil) && (cur.ref <= node.ref);{
			pre_node = cur
			cur = cur.next
		}
		if cur == nil{
			pre_node.next = node
			node.pre = pre_node
		} else if pre_node != nil{
			pre_node.next = node
			node.pre = pre_node
			node.next = cur
			cur.pre = node

		}
		return node.val
	}
	return -1
}

func print(head *linknode1){
	cur := head
	for ; cur != nil; {
		fmt.Printf("%v ref: %v ", cur.val, cur.ref);
		cur = cur.next
	}
	fmt.Println("")
}

func (this *LFUCache) Put(key int, value int)  {
	if node, ok := this.m[key]; ok {
		if node.val == value{
			this.Get(key)
		}else{
			pre, next := node.pre, node.next
			if pre == nil{
				this.head.val = value
				this.head.ref = 1
			}
		}
	}else if this.left_size > 0{
		this.left_size -= 1
		new_node := &linknode1{val:value, key:key, ref:1}
		this.m[key] = new_node
		if this.head == nil{
			this.head = new_node
			return
		}
		if this.head.ref > new_node.ref{
			new_node.next = this.head
			this.head.pre = new_node
			this.head = new_node
			return
		}
		cur := this.head
		for ; (cur.next!=nil) && (cur.next.ref==this.head.ref);{
			cur = cur.next
		}
		if cur.next == nil{
			cur.next = new_node
			new_node.pre = cur
		}else{
			next_node := cur.next
			cur.next = new_node
			new_node.next = next_node
			new_node.pre = cur
			next_node.pre = new_node
		}
	}else{
		delete(this.m, this.head.key)
		this.head.val = value
		this.head.ref = 1
		this.head.key = key
		this.m[key] = this.head
	}
}


/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */


func main() {
	cache := Constructor22(2);
	cache.Put(3, 1);
	cache.Put(2, 1);
	cache.Put(2, 2);
	cache.Put(4, 4);
	cache.Get(2)
}
