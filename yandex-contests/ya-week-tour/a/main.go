package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, a, k int

	fmt.Fscan(in, &n)
	arr := make([]int, 0, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a)
		arr = append(arr, a)
	}

	if n%2 == 1 {
		fmt.Println(-1)
		return
	}

	left := 0
	right := n - 1

	k = arr[left] + arr[right]

	for left < right && arr[left]+arr[right] == k {
		left++
		right--
	}
	if left < right {
		fmt.Println(-1)
		return
	}

	fmt.Println(k)
}
