package utils

func Dfs(grid [][]int, i int, j int, visited [][]bool, left *int, right *int) {
	// 检查当前位置是否已经被访问过，或者是否为0
	if i < 0 || i >= len(grid) ||
		j < 0 || j >= len(grid[0]) ||
		grid[i][j] == 0 || visited[i][j] {
		return
	}
	// 将当前位置标记为已访问
	visited[i][j] = true
	// 更新左右边界
	if *left > j {
		*left = j
	}
	if *right < j {
		*right = j
	}
	// 递归遍历相邻的位置
	Dfs(grid, i-1, j, visited, left, right) // 上
	Dfs(grid, i+1, j, visited, left, right) // 下
	Dfs(grid, i, j-1, visited, left, right) // 左
	Dfs(grid, i, j+1, visited, left, right) // 右
}
