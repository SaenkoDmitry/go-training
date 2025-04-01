package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, a int
	fmt.Fscan(in, &n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a)
		arr[i] = a
	}

	sort.Ints(arr)

	maxItem := arr[len(arr)-1]
	sum := 0
	for i := 0; i < n; i++ {
		sum += arr[i]
	}

	res := 0
	for i := n - 1; i >= 0; i-- {
		curr := arr[i]
		if curr-(sum-curr) > 0 {
			res = maxItem - (sum - maxItem)
			break
		}
	}
	if res == 0 {
		res = sum
	}
	fmt.Println(res)
}
