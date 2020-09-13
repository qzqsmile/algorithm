package main


func sortColors (nums *[]int)  {
	// write your code here
	i, j := 0, len(*nums)-1
	c := 0

	for ;c <= j;{
		if (*nums)[c] == 0{
			(*nums)[c], (*nums)[i] = (*nums)[i], (*nums)[c]
			i++
			c++
		}else if (*nums)[c] == 2{
			(*nums)[c], (*nums)[j] = (*nums)[j], (*nums)[c]
			j--
		}else{
			c++
		}
	}
}