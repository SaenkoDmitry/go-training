package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var s string
	arr := make([][]rune, 0, 8)
	for i := 0; i < 8; i++ {
		fmt.Fscan(in, &s)
		arr = append(arr, []rune(s))
	}

	res := 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if arr[i][j] == 'R' || arr[i][j] == 'B' {
				continue
			}
			if !isDangerousR(arr, i, j) && !isDangerousB(arr, i, j) {
				res++
			}
		}
	}
	fmt.Println(res)
}

func isDangerousR(arr [][]rune, i, j int) bool {
	for x := i + 1; x < len(arr); x++ {
		if arr[x][j] == 'R' {
			return true
		} else if arr[x][j] == 'B' {
			break
		}
	}

	for x := i - 1; x >= 0; x-- {
		if arr[x][j] == 'R' {
			return true
		} else if arr[x][j] == 'B' {
			break
		}
	}

	for y := j + 1; y < len(arr); y++ {
		if arr[i][y] == 'R' {
			return true
		} else if arr[i][y] == 'B' {
			break
		}
	}

	for y := j - 1; y >= 0; y-- {
		if arr[i][y] == 'R' {
			return true
		} else if arr[i][y] == 'B' {
			break
		}
	}
	return false
}

func isDangerousB(arr [][]rune, i, j int) bool {
	for x, y := i+1, j+1; x < len(arr) && y < len(arr); x, y = x+1, y+1 {
		if arr[x][y] == 'B' {
			return true
		} else if arr[x][y] == 'R' {
			break
		}
	}

	for x, y := i-1, j-1; x >= 0 && y >= 0; x, y = x-1, y-1 {
		if arr[x][y] == 'B' {
			return true
		} else if arr[x][y] == 'R' {
			break
		}
	}

	for x, y := i+1, j-1; x < len(arr) && y >= 0; x, y = x+1, y-1 {
		if arr[x][y] == 'B' {
			return true
		} else if arr[x][y] == 'R' {
			break
		}
	}

	for x, y := i-1, j+1; x >= 0 && y < len(arr); x, y = x-1, y+1 {
		if arr[x][y] == 'B' {
			return true
		} else if arr[x][y] == 'R' {
			break
		}
	}
	return false
}
