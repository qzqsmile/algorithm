package array

func winSum (nums []int, k int) []int {
	// write your code here
	mp := make(map[int]bool)
	sum := []int{0}
	r := 0
	for i := 0; i < len(nums); i++{
		r += nums[i]
		sum = append(sum, r)
	}
	res := []int{}
	for i := 3; i < len(nums); i++{
		tmp := sum[i] - sum[i-3]
		if _, ok := mp[tmp]; !ok{
			res = append(res, tmp)
			mp[tmp] = true
		}
	}
	return res
}

func main(){
	nums := []int{1,2,7,7,2}
	winSum(nums, 3)
}
