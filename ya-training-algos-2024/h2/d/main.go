package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, a, b int
	arr := make([][]int, n)

	fmt.Fscan(in, &n)

	size := 8

	for i := 0; i < size; i++ {
		arr = append(arr, make([]int, 8))
	}

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a, &b)
		arr[a-1][b-1] = 1
	}

	count := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if arr[i][j] == 1 {
				count += 4
				if i-1 >= 0 && arr[i-1][j] == 1 {
					count--
				}
				if j-1 >= 0 && arr[i][j-1] == 1 {
					count--
				}
				if i+1 < size && arr[i+1][j] == 1 {
					count--
				}
				if j+1 < size && arr[i][j+1] == 1 {
					count--
				}
			}
		}
	}
	fmt.Println(count)
}
