package main

import (
	"fmt"
	"math"
)

func sumSubarrayMins(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++{
		minNum := math.MaxInt32
		for j := i; j < len(arr); j++{
			minNum = min(minNum, arr[j])
			sum += minNum
		}
		sum %= 10^9+7
	}
	return sum
}

func min(a int, b int) int{
	if a < b{
		return a
	}
	return b
}

func main(){
	fmt.Println(100000000000)
}