package twopointer


func twoSum(numbers []int, target int) []int {
	b, e := 0, len(numbers)-1
	for; b < e;{
		if numbers[b] + numbers[e] == target{
			return []int{b+1, e+1}
		}else if numbers[b] + numbers[e] < target{
			b += 1
		}else{
			e -= 1
		}
	}
	return []int{-1, -1}
}

