package main

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for ;i < j;{
		for;i < j && numbers[i]+numbers[j] < target; i++{
		}
		for;i < j && numbers[i]+numbers[j] > target; j--{
		}
		if i < j{
			if numbers[i] + numbers[j] == target{
				return []int{i+1, j+1}
			}
		}
	}
	return []int{}
}
