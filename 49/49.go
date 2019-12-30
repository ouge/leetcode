package leetcode

func groupAnagrams(strs []string) [][]string {
	ret := make([][]string, 0, len(strs))
	strMap := make(map[[26]byte]int)
	for i := range strs {
		var key [26]byte
		for j := range strs[i] {
			key[strs[i][j]-'a']++
		}
		k, ok := strMap[key]
		if !ok {
			strMap[key] = len(ret)
			ret = append(ret, []string{})
			k = len(ret) - 1
		}
		ret[k] = append(ret[k], strs[i])
	}
	return ret
}
