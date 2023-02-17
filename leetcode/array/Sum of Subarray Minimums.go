package array

func sumSubarrayMins(arr []int) int {
	left, right := make([]int, len(arr)), make([]int, len(arr))
	monoStk := []int{}
	for i := len(arr) - 1; i >= 0; i-- {
		for len(monoStk) > 0 && arr[i] < arr[monoStk[len(monoStk)-1]] {
			monoStk = monoStk[0 : len(monoStk)-1]
		}
		if len(monoStk) == 0 {
			right[i] = len(arr)
		} else {
			right[i] = monoStk[len(monoStk)-1]
		}

		monoStk = append(monoStk, i)
	}

	monoStk = []int{}
	//这里等于要注意，由于right没有等于，所以右边是不重复的，可以保证整个都是不重复的
	for i := 0; i < len(arr); i++ {
		for len(monoStk) > 0 && arr[i] <= arr[monoStk[len(monoStk)-1]] {
			monoStk = monoStk[0 : len(monoStk)-1]
		}
		if len(monoStk) == 0 {
			left[i] = -1
		} else {
			left[i] = monoStk[len(monoStk)-1]
		}

		monoStk = append(monoStk, i)
	}

	var mod int = 1e9 + 7
	ans := 0
	for i := 0; i < len(arr); i++ {
		ans = (ans + (i-left[i])*(right[i]-i)*arr[i]) % mod
		// ans %= mod
	}
	return ans
}
