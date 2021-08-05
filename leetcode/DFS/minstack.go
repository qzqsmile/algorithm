package DFS

type MinStack struct {
	stk []node
}


/** initialize your data structure here. */
func Constructor() MinStack {
	t := MinStack{}
	return t
}


func (this *MinStack) Push(val int)  {
	if len(this.stk) == 0{
		this.stk = append(this.stk, node{val, val})
	}else{
		t := this.stk[len(this.stk)-1]
		this.stk = append(this.stk, node{val, min(t.mval, val)})
	}
}


func (this *MinStack) Pop()  {
	this.stk = this.stk[0:len(this.stk)-1]
}


func (this *MinStack) Top() int {
	n := this.stk[len(this.stk)-1]
	return n.val
}


func (this *MinStack) GetMin() int {
	n := this.stk[len(this.stk)-1]
	return n.mval
}

type node struct{
	val int
	mval int
}

func min(a int, b int) int{
	if a < b{
		return a
	}
	return b
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
