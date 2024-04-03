package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, a int
	arr := make([]int, 0)

	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a)
		arr = append(arr, a)
	}

	start := -1
	end := n
	for i := 0; i < n; i++ {
		if arr[i]%2 == 1 || arr[i]%2 == -1 {
			start = i
			for j := i + 1; j < n; j++ {
				if arr[j]%2 == 0 {
					end = j
					break
				}
			}
			break
		}
	}

	for i := 1; i < n; i++ {
		if i < start || i > end {
			fmt.Printf("x")
		} else if i == start || i == end {
			fmt.Printf("+")
		} else {
			fmt.Printf("x")
		}
	}
}
