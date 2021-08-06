func evalRPN(tokens []string) int {
	var stack []int
	 for _, v := range tokens {
		 if v == "+" || v == "-" || v == "*" || v == "/" {
			 r1 := stack[len(stack)-1]
			 r2 := stack[len(stack)-2]
			 stack = stack[:len(stack)-2]
			 if v == "+" {
				 stack = append(stack, r1+r2)
			 } else if v == "-" {
				 stack = append(stack, r2-r1)
			 } else if v == "*" {
				 stack = append(stack, r1*r2)
			 } else {
				 stack = append(stack, r2/r1)
			 }
		 } else {
			 v1, _ := strconv.Atoi(v)
			 stack = append(stack, v1)
		 }
	 }
	 return stack[len(stack)-1]
 }

 