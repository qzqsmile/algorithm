package sort

import "sort"

/*
	这里的小套路就是这种二维数组，需要经过排序。即首先根据第一个以及第二个元素进行排序
	一个正向排，一个反向排
*/

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		p1, p2 := people[i], people[j]
		return p1[0] > p2[0] || (p1[0] == p2[0] && p1[1] < p2[1])
	})
	ans := [][]int{}
	for _, p := range people {
		idx := p[1]
		tmp := append(append([][]int{}, ans[:idx]...), p)
		// tmp = append(tmp, p)
		ans = append(tmp, ans[idx:]...)
	}
	return ans
}

// func reconstructQueue(people [][]int) (ans [][]int) {
//     sort.Slice(people, func(i, j int) bool {
//         a, b := people[i], people[j]
//         return a[0] > b[0] || a[0] == b[0] && a[1] < b[1]
//     })
//     for _, person := range people {
//         idx := person[1]
//         ans = append(ans[:idx], append([][]int{person}, ans[idx:]...)...)
//     }
//     return
// }

// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/queue-reconstruction-by-height/solution/gen-ju-shen-gao-zhong-jian-dui-lie-by-leetcode-sol/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
