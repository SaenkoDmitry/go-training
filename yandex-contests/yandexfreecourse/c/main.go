package main

import (
	"bufio"
	"fmt"
	"os"
)

func findStatusOK(arr []int) int {
	numbers := make([]int, 200)
	for i := 0; i < len(arr); i++ {
		numbers[arr[i]%200]++
	}

	result := 0
	for i := range numbers {
		if numbers[i] != 0 {
			result += numbers[i] * (numbers[i] - 1) / 2
		}
	}
	return result
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, a int
	fmt.Fscan(in, &n)

	arr := make([]int, 0, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a)
		arr = append(arr, a)
	}

	fmt.Println(findStatusOK(arr))
}
