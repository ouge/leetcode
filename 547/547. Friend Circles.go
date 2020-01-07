package leetcode

func findCircleNum(M [][]int) (ret int) {
	ok := make([]bool, len(M)) // 是否统计过该用户
	for i := range M {
		if ok[i] {
			continue
		}
		ok[i] = true
		for j := range M[i] {
			if M[i][j] == 1 {
				ret++
				var bfs []int
				bfs = append(bfs, i)
				for len(bfs) > 0 {
					top := bfs[0]
					bfs = bfs[1:]
					M[top][top] = 0

					for k := range M[top] {
						if M[top][k] == 1 && !ok[k] {
							bfs = append(bfs, k)
							M[top][k] = 0
							M[k][top] = 0
							ok[k] = true
						}
					}
				}
			}
		}
	}
	return ret
}
