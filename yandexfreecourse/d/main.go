package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func maxIncreasedSequence(arr [][]int) int {
	visited := make([][]bool, 0)
	for i := range arr {
		visited = append(visited, make([]bool, 0, len(arr)))
		for range arr[i] {
			visited[i] = append(visited[i], false)
		}
	}

	accum := make(map[Pair]int)

	max := 0
	for i := range arr {
		for j := range arr[i] {
			temp := maxSequence(arr, i, j, accum, visited) + 1
			if temp > max {
				max = temp
			}
		}
	}
	return max
}

type Pair struct {
	x, y int
}

func maxSequence(arr [][]int, x, y int, accum map[Pair]int, visited [][]bool) int {
	if v, ok := accum[Pair{x, y}]; ok {
		return v
	}

	left, right, up, down := 0, 0, 0, 0
	if x-1 >= 0 && arr[x][y] > arr[x-1][y] {
		up = maxSequence(arr, x-1, y, accum, visited) + 1
	}
	if x+1 <= len(arr)-1 && arr[x][y] > arr[x+1][y] {
		down = maxSequence(arr, x+1, y, accum, visited) + 1
	}
	if y-1 >= 0 && arr[x][y] > arr[x][y-1] {
		left = maxSequence(arr, x, y-1, accum, visited) + 1
	}
	if y+1 <= len(arr[0])-1 && arr[x][y] > arr[x][y+1] {
		right = maxSequence(arr, x, y+1, accum, visited) + 1
	}
	res := max(left, right, up, down)
	accum[Pair{x, y}] = res
	visited[x][y] = true
	return res
}

func max(left, right, up, down int) int {
	return int(math.Max(math.Max(float64(left), float64(right)), math.Max(float64(up), float64(down))))
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, a int
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &m)

	arr := make([][]int, 0)

	for i := 0; i < n; i++ {
		arr = append(arr, make([]int, 0, n))
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &a)
			arr[i] = append(arr[i], a)
		}
	}

	fmt.Println(maxIncreasedSequence(arr))
}
