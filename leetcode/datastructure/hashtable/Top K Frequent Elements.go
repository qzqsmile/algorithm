func topKFrequent(nums []int, k int) []int {
    mp := make(map[int]int)
    for _, v := range nums{
        mp[v]++
    }
    count := []int{}
    for _, v := range mp{
        count = append(count, v)
    }
    sort.Ints(count)
    slot := count[len(count)-k]
    res := []int{}
    for k, v := range mp{
        if v >=slot{
            res = append(res, k)
        }
    }
    return res
}
