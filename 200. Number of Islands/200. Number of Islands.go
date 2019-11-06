package leetcode

func numIslands(grid [][]byte) int {
	var ret int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				ret++
				dye(grid, i, j)
			}
		}
	}
	return ret
}

func dye(grid [][]byte, x, y int) {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] != '1' {
		return
	}
	grid[x][y] = '2'
	dye(grid, x-1, y)
	dye(grid, x+1, y)
	dye(grid, x, y-1)
	dye(grid, x, y+1)
}
