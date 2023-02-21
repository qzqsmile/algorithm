import "sort"

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	mp := make(map[string]int)
	for _, v := range strs{
		if i, ok :=  mp[convertStr(v)]; ok{
			res[i] = append(res[i], v)
		}else{
			res = append(res, []string{v})
			mp[convertStr(v)] = len(res)-1
		}
	}
	return res
}

func convertStr(str string) string{
	by := []byte(str)
	sort.Slice(by, func(i, j int) bool{
		return by[i] < by[j]
	})
	return  string(by[:])
}