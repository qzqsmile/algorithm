package main

type MinStack struct {
	stk []int
	min int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	m := MinStack{stk: []int{}}
	return m
}


func (this *MinStack) Push(x int)  {
	if len(this.stk) == 0{
		this.min = x
	}
	if x <= this.min{
		this.stk = append(this.stk, this.min)
		this.min = x
	}
	this.stk = append(this.stk, x)
}


func (this *MinStack) Pop()  {
	v := this.stk[len(this.stk)-1]
	this.stk = this.stk[0:len(this.stk)-1]
	if v == this.min{
		this.min = this.stk[len(this.stk)-1]
		this.stk = this.stk[0:len(this.stk)-1]
	}
}


func (this *MinStack) Top() int {
	return this.stk[len(this.stk)-1]
}


func (this *MinStack) GetMin() int {
	return this.min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
