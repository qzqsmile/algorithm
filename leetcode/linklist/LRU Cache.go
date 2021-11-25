//dummyHead和dummyTail是关键

type Node struct{
    key int 
    val int 
    pre *Node
    next *Node
}

type LRUCache struct {
    head *Node
    tail *Node 
    capacity int
    count int
    mp map[int]*Node
}


func Constructor(capacity int) LRUCache {
    dummyHead := &Node{-1, -1, nil, nil}
    dummyTail := &Node{-1, -1, nil, nil}
    dummyHead.next = dummyTail
    dummyTail.pre = dummyHead
    return LRUCache{dummyHead, dummyTail, capacity, 0, make(map[int]*Node)}
}


func (this *LRUCache) Get(key int) int {
    if node, ok := this.mp[key]; ok{
        pre := node.pre
        next := node.next
        pre.next = next
        next.pre = pre 
        
        headNext := this.head.next
        this.head.next = node 
        node.pre = this.head 
        node.next = headNext
        headNext.pre = node
        return node.val
    }else{
        return -1
    }
}


func (this *LRUCache) Put(key int, value int)  {
    if _, ok := this.mp[key]; ok{
        this.Get(key)
        this.head.next.val = value
    }else{
        if this.count < this.capacity{
            this.count++

            node := &Node{key, value, nil, nil}
            nextHead := this.head.next

            node.pre = this.head 
            this.head.next = node 
            node.next = nextHead
            nextHead.pre = node

            this.mp[key] = node
        }else{
            this.tail = this.tail.pre 
            this.tail.next = nil 

            delete(this.mp, this.tail.key)
            this.count--
            this.Put(key, value)
        }
    }
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */