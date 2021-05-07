package array

import (
	"fmt"
	"sort"
)

func subarraySum(nums []int, k int) int {
	res :=[]int{0}
	mp := make(map[int]int)
	mp[0] = -1
	for i := 0; i < len(nums); i++{
		tmp := res[i] + nums[i]
		res = append(res, tmp)
		mp[tmp] = i
	}
	sort.Ints(res)
	count := 0
	for i := 0; i < len(res); i++{
		for j := i + 1; j < len(res); j++{
			if value, _ := mp[res[j]]; value > mp[res[i]]{
				if res[j] - res[i] == k{
					count++
				}else if res[j] - res[i] > k{
					break
				}
			}
		}
	}
	return count
}

func main()  {
	nums, k := []int{-1,-1,1}, 1
	count := subarraySum(nums, k)
	fmt.Println(count)
}