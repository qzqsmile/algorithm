package main

/**
 * @param nums: an integer array
 * @return: nothing
 */
func moveZeroes (nums *[]int)  {
	// write your code here
	b1, b2 := 0, 0
	for;b2 < len(*nums);{
		for ;b2 < len(*nums) && (*nums)[b2] == 0;{
			b2++
		}
		if b2 < len(*nums){
			(*nums)[b1], (*nums)[b2] = (*nums)[b2], (*nums)[b1]
			b1++
			b2++
		}
	}
}

