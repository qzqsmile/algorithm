package main

type linknode struct{
	key, value int
	pre, next *linknode
}

type LRUCache struct {
	capacity int
	m map[int] *linknode
	b, e *linknode
}


func Constructor(capacity int) LRUCache {
	lru := LRUCache{capacity: capacity}
	lru.b = &linknode{key: 0, value:0}
	lru.m = make(map[int] *linknode)
	lru.e = lru.b
	return lru
}


func (this *LRUCache) Get(key int) int {
	if _, ok:= this.m[key]; !ok{
		return -1
	}
	if this.m[key].next != nil{
		this.m[key].next.pre = this.m[key].pre
		this.m[key].pre.next = this.m[key].next
		this.e.next = this.m[key]
		this.m[key].pre = this.e
		this.e = this.e.next
		this.e.next = nil
	}
	return this.m[key].value
}


func (this *LRUCache) Put(key int, value int)  {
	if _, ok := this.m[key]; !ok{
		if this.capacity > 0{
			this.e.next = &linknode{key: key, value:value, pre:this.e, next:nil}
			this.e = this.e.next
			this.m[key] = this.e
			this.capacity--
		}else{
			this.b = this.b.next
			delete(this.m, this.b.key)
			this.e.next = &linknode{key: key, value:value, pre:this.e, next:nil}
			this.e = this.e.next
			this.m[key] = this.e
		}
	}else{
		if this.m[key].next != nil{
			this.m[key].next.pre = this.m[key].pre
			this.m[key].pre.next = this.m[key].next
			this.e.next = this.m[key]
			this.e.next.pre = this.e
			this.e = this.e.next
			this.e.next = nil
		}
		this.m[key].value = value
	}
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */


