package _200_Number_of_Islands

func numIslands(grid [][]byte) int {
	var (
		count int
	)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				count++
				dfs(i, j, &grid)
			}
		}
	}
	return count
}

func dfs(i, j int, grid *([][]byte)) {
	if i < 0 || i > len(*grid)-1 || j < 0 || j > len((*grid)[i])-1 || (*grid)[i][j] == '0' {
		return
	}
	(*grid)[i][j] = '0'
	dfs(i-1, j, grid)
	dfs(i+1, j, grid)
	dfs(i, j-1, grid)
	dfs(i, j+1, grid)
}
