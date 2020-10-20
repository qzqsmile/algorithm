package main

//使用了单调栈的技术，单调栈即在栈里维护了一个单调递增或递减的序列
//一般用于计算数组中比给定的坐标大/小的坐标
//要理解单调栈的实际图。


func largestRectangleArea(heights []int) int {
	r, l := []int{}, []int{}
	stk := []int{}
	maxAera := 0
	for i := len(heights)-1; i >= 0; i--{
		for len(stk) > 0 && heights[stk[len(stk)-1]] >= heights[i]{
			stk = stk[0: len(stk)-1]
		}
		if len(stk) == 0{
			r = append(r, -1)
		}else{
			r = append(r, stk[len(stk)-1])
		}
		stk = append(stk, i)
	}

	reverse(r)

	stk = []int{}
	for i := 0; i <= len(heights)-1; i++{
		for len(stk) > 0 && heights[stk[len(stk)-1]] >= heights[i]{
			stk = stk[0: len(stk)-1]
		}
		if len(stk) == 0{
			l = append(l, -1)
		}else{
			l = append(l, stk[len(stk)-1])
		}
		stk = append(stk, i)
	}

	for i := 0; i < len(heights); i++{
		lb, rb := l[i], r[i]
		if l[i] == -1{
			lb = -1
		}
		if r[i] == -1{
			rb = len(heights)
		}
		if heights[i] * (rb-lb-1) > maxAera{
			maxAera = heights[i] * (rb-lb-1)
		}
	}
	return maxAera
}

func reverse(nums []int){
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1{
		nums[i], nums[j] = nums[j], nums[i]
	}
}