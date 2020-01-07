package leetcode

func maxAreaOfIsland(grid [][]int) int {
	var ret int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				newArea := calArea(grid, i, j)
				if ret < newArea {
					ret = newArea
				}
			}
		}
	}
	return ret
}

type pos struct {
	x int
	y int
}

func calArea(grid [][]int, x, y int) int {
	// bfs
	var area int
	var bfs []pos

	bfs = append(bfs, pos{x, y})
	for len(bfs) > 0 {
		p := bfs[0]
		bfs = bfs[1:]

		if p.x < 0 || p.x >= len(grid) ||
			p.y < 0 || p.y >= len(grid[x]) {
			continue
		}
		if grid[p.x][p.y] != 1 {
			continue
		}
		grid[p.x][p.y] = 2
		area++

		bfs = append(bfs,
			pos{p.x - 1, p.y},
			pos{p.x + 1, p.y},
			pos{p.x, p.y - 1},
			pos{p.x, p.y + 1})
	}
	return area

}
