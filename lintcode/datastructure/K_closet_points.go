package main

func kClosest(points [][]int, K int) [][]int {
	res := [][]int{}
	res = append(res, points[0:K]...)
	makeHeap(res)

	for i := K; i < len(points); i++ {
		if calPoint(res[0]) > calPoint(points[i]) {
			res[0] = points[i]
			heapSort(res, 0)
		}
	}
	return res
}

func makeHeap(points [][]int) {
	k := len(points)
	for i := k / 2; i >= 0; i-- {
		heapSort(points, i)
	}
}

func heapSort(points [][]int, i int) {
	l := 2*i + 1
	r := 2*i + 2
	maxIndex := i
	if l < len(points) && calPoint(points[l]) > calPoint(points[maxIndex]) {
		maxIndex = l
	}
	if r < len(points) && calPoint(points[r]) > calPoint(points[maxIndex]) {
		maxIndex = r
	}
	if i != maxIndex {
		points[maxIndex], points[i] = points[i], points[maxIndex]
		heapSort(points, maxIndex)
	}

}

func calPoint(point []int) int {
	return point[0]*point[0] + point[1]*point[1]
}
