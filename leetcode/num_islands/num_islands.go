package main

import "fmt"

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	visited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[i]))
	}

	k := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 || visited[i][j] {
				continue
			}
			markVertex(grid, &visited, i, j)
			k++
		}
	}
	return k
}

func markVertex(grid [][]byte, visited *[][]bool, i, j int) {
	if grid[i][j] == 0 || (*visited)[i][j] {
		return
	}
	(*visited)[i][j] = true
	if i > 0 {
		markVertex(grid, visited, i - 1, j)
	}
	if i < len(grid) - 1 {
		markVertex(grid, visited, i + 1, j)
	}
	if j > 0 {
		markVertex(grid, visited, i, j - 1)
	}
	if j < len(grid[i]) - 1 {
		markVertex(grid, visited, i, j + 1)
	}
}

func main() {
	fmt.Println(numIslands([][]byte{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}))
	fmt.Println(numIslands([][]byte{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}))
}
