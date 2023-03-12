package main

import (
	"container/heap"
	"fmt"
)

type LFUCache struct {
	capacity int
	mp       map[int]*page
	minHeap  *hp
	opCnt    int
}

type page struct {
	key       int
	value     int
	c, lastop int
	index     int
}

type hp []*page

func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	a, b := h[i], h[j]
	if a.c == b.c {
		return a.lastop < b.lastop
	}
	return a.c < b.c
}
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i]; h[i].index, h[j].index = h[j].index, h[i].index }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(*page)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func Constructor(capacity int) LFUCache {
	mp := make(map[int]*page)
	minHeap := &hp{}
	return LFUCache{capacity, mp, minHeap, 0}
}

func (this *LFUCache) Get(key int) int {
	this.opCnt++
	if p, ok := this.mp[key]; ok {
		//
		p.c++
		p.lastop = this.opCnt
		heap.Fix(this.minHeap, p.index)
		return p.value
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	this.opCnt++
	if page, ok := this.mp[key]; ok {
		//update and fix
		page.value = value
		page.c++
		page.lastop = this.opCnt
		heap.Fix(this.minHeap, page.index)
		return
	}
	//if not existed
	if this.minHeap.Len() == this.capacity {
		v := heap.Pop(this.minHeap)
		delete(this.mp, v.(*page).key)
	}
	//这里的index是最后一个字段，需要注意
	p := &page{key, value, 1, this.opCnt, this.minHeap.Len()}
	heap.Push(this.minHeap, p)
	this.mp[key] = p
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	lfu := Constructor(2)
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	lfu.Get(1)
	lfu.Put(3, 3) // 2 is the LFU key because cnt(2)=1 is the smallest, invalidate 2.
	lfu.Get(2)    // return -1 (not found)
	lfu.Get(3)    // return 3
	lfu.Put(4, 4) // Both 1 and 3 have the same cnt, but 1 is LRU, invalidate 1.
	// cache=[4,3], cnt(4)=1, cnt(3)=2
	fmt.Printf("%v\n", lfu.Get(1)) // return -1 (not found)
	fmt.Printf("%v\n", lfu.Get(3)) // return 3
	lfu.Get(4)                     // return 4

}
