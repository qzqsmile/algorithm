package twopointer


/**
 * @param nums: The integer array you should partition
 * @param k: An integer
 * @return: The index after partition
 */
func partitionArray (nums []int, k int) int {
	// write your code here
	// nums = append(nums, k+1)
	b, e := 0, len(nums)-1

	for;b <= e;{
		//这里是b,也就是左指针先移动，所以会指向第一个比K大的数
		for;b <= e && nums[b] < k; b++{}
		for;b <= e && nums[e] >= k; e--{}
		if b <= e{
			nums[b], nums[e] = nums[e], nums[b]
			b++
			e--
		}
	}
	// if b == e{
	//     return b
	// }
	return b
}

//非常值得学习的一个template，有几点要注意
//1，初始化，以及中间的操作要始终保持nums[b], nume[e]分别指向比k大的，以及比k小的值。避免corner case的出现
//2. 全部使用b < e的比较法，保证出了这个循环必然是 b==e
/***
python version
class Solution:
    """
    @param nums: The integer array you should partition
    @param k: An integer
    @return: The index after partition
    """
    def partitionArray(self, nums, k):
        # write your code here
        b, e = 0, len(nums)-1
        while b < len(nums) and nums[b] < k:
            b += 1
        while e >= 0 and nums[e] >= k:
            e -= 1
        while b < e:
            while b < e and nums[b] < k:
                b += 1
            while b < e and nums[e] >= k:
                e -= 1
            if b < e:
                nums[b], nums[e] = nums[e], nums[b]
        return b
***/