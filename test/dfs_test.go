package test

import (
	"testing"
)

func dfs(grid [][]int, i int, j int, visited [][]bool, left *int, right *int) {
	// 检查当前位置是否已经被访问过，或者是否为0
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 || visited[i][j] {
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
	dfs(grid, i-1, j, visited, left, right) // 上
	dfs(grid, i+1, j, visited, left, right) // 下
	dfs(grid, i, j-1, visited, left, right) // 左
	dfs(grid, i, j+1, visited, left, right) // 右
}

func TestDFS(t *testing.T) {
	tests := []struct {
		grid          [][]int
		expectedLeft  int
		expectedRight int
	}{
		{
			grid: [][]int{
				{1, 0, 1, 1, 0},
				{1, 1, 0, 1, 0},
				{0, 1, 1, 0, 1},
				{1, 1, 1, 0, 1},
				{1, 1, 0, 1, 0},
			},
			expectedLeft:  0,
			expectedRight: 2,
		},
		{
			grid: [][]int{
				{1, 0, 1, 1, 0},
				{1, 1, 0, 1, 0},
				{0, 1, 1, 0, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 0, 1, 0},
			},
			expectedLeft:  0,
			expectedRight: 4,
		},
		{
			grid: [][]int{
				{1, 0, 1, 1, 0},
				{1, 0, 0, 1, 0},
				{0, 0, 1, 0, 1},
				{1, 0, 1, 1, 1},
				{1, 0, 0, 1, 0},
			},
			expectedLeft:  0,
			expectedRight: 0,
		},
		{
			grid: [][]int{
				{0, 0, 1, 1, 0},
				{0, 0, 0, 1, 0},
				{0, 1, 1, 0, 1},
				{0, 0, 1, 1, 0},
				{0, 0, 0, 1, 0},
			},
			expectedLeft:  1,
			expectedRight: 3,
		},
	}

	for _, test := range tests {
		visited := make([][]bool, len(test.grid))
		for i := range visited {
			visited[i] = make([]bool, len(test.grid[i]))
		}

		left, right := len(test.grid[0]), 0
		dfs(test.grid, 0, 0, visited, &left, &right)

		if left != test.expectedLeft || right != test.expectedRight {
			t.Errorf("对于测试数据%v，连通区域左边界应该为%d，右边界应该为%d，但得到左边界为%d，右边界为%d", test.grid, test.expectedLeft, test.expectedRight, left, right)
		}
	}
}
