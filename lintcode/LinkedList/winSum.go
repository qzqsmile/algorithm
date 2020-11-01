package main
/**
 * @param nums: a list of integers.
 * @param k: length of window.
 * @return: the sum of the element inside the window at each moving.
 */
func winSum (nums []int, k int) []int {
	// write your code here
	if len(nums) == 0 || k == 0 || k == 1{
		return nums
	}
	cur := 0
	mp := make(map[int]int)
	mp[-1] = 0
	res := []int{}
	for i, v := range nums{
		cur += v
		mp[i] = cur
		if i >= k-1{
			res = append(res, mp[i]-mp[i-k])
		}
	}
	return res

}

