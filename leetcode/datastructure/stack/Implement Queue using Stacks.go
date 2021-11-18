type MyQueue struct {
    stk1 []int
    stk2 []int
}


func Constructor() MyQueue {
    return MyQueue{}
}


func (this *MyQueue) Push(x int)  {
    this.stk1 = append(this.stk1, x)
}


func (this *MyQueue) Pop() int {
    for;len(this.stk1) > 0;{
        t := this.stk1[len(this.stk1)-1]
        this.stk2 = append(this.stk2, t)
        this.stk1 = this.stk1[0:len(this.stk1)-1]
    }
    pVal := this.stk2[len(this.stk2)-1]
    this.stk2 = this.stk2[0:len(this.stk2)-1]

    for;len(this.stk2) > 0;{
        t := this.stk2[len(this.stk2)-1]
        this.stk1 = append(this.stk1, t)
        this.stk2 = this.stk2[0:len(this.stk2)-1]
    }
    return pVal
}


func (this *MyQueue) Peek() int {
    for;len(this.stk1) > 0;{
        t := this.stk1[len(this.stk1)-1]
        this.stk2 = append(this.stk2, t)
        this.stk1 = this.stk1[0:len(this.stk1)-1]
    }
    pVal := this.stk2[len(this.stk2)-1]
    for;len(this.stk2) > 0;{
        t := this.stk2[len(this.stk2)-1]
        this.stk1 = append(this.stk1, t)
        this.stk2 = this.stk2[0:len(this.stk2)-1]
    }
    return pVal
}


func (this *MyQueue) Empty() bool {
    if len(this.stk1) > 0{
        return false
    }
    return true
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */