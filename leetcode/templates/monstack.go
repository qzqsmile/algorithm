func monistack(){
	stk := []int{}
	for i := len(nums)-1; i >= 0; i--{
		for;len(stk) && nums[i] >= stk[len(stk)-1];{
			stk = stk[0:len(stk)-1]
		}
		// action aera

		
		stk = append(stk, i)
	}
}