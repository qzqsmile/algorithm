package DFS

type NestedIterator struct {
	nestedList []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	t := NestedIterator{}
	for i := len(nestedList)-1; i >= 0; i--{
		t.nestedList = append(t.nestedList, nestedList[i])
	}
	return &t
}

func (this *NestedIterator) Next() int {
	for ;len(this.nestedList) > 0; {
		n := this.nestedList[len(this.nestedList)-1]
		this.nestedList = this.nestedList[0:len(this.nestedList)-1]

		if n.IsInteger(){
			return n.GetInteger()
		}else{
			l := n.GetList()
			for i := len(l)-1; i >= 0; i--{
				this.nestedList = append(this.nestedList, l[i])
			}
		}
	}
	return -1
}

func (this *NestedIterator) HasNext() bool {
	for ;len(this.nestedList) > 0; {
		n := this.nestedList[len(this.nestedList)-1]

		if n.IsInteger(){
			return true
		}else{
			this.nestedList = this.nestedList[0:len(this.nestedList)-1]
			l := n.GetList()
			for i := len(l)-1; i >= 0; i--{
				this.nestedList = append(this.nestedList, l[i])
			}
		}
	}
	return false
}
