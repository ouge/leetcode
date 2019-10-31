package leetcode

func minInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minDistance(word1 string, word2 string) int {
	res := make([][]int, len(word1)+1)
	for i := range res {
		res[i] = make([]int, len(word2)+1)
		res[i][0] = i
	}
	for i := range res[0] {
		res[0][i] = i
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			res[i][j] = minInt(res[i][j-1]+1, res[i-1][j]+1)
			if word1[i-1] == word2[j-1] {
				res[i][j] = minInt(res[i][j], res[i-1][j-1])
			} else {
				res[i][j] = minInt(res[i][j], res[i-1][j-1]+1)
			}
		}
	}
	return res[len(word1)][len(word2)]
}
