package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k, a int

	fmt.Fscan(in, &n, &k)
	arr := make([]int, 0, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a)
		arr = append(arr, a)
	}

	res := 0
	for i := k; i < n; i++ {
		minElem := arr[i-k]
		diff := 0
		for j := i - k; j < i+1; j++ {
			if arr[j] < minElem {
				minElem = arr[j]
			}
			if arr[j]-minElem > diff {
				diff = arr[j] - minElem
			}
		}
		if diff > res {
			res = diff
		}
	}

	if k == n {
		minElem := arr[0]
		diff := 0
		for j := 0; j < n; j++ {
			if arr[j] < minElem {
				minElem = arr[j]
			}
			if arr[j]-minElem > diff {
				diff = arr[j] - minElem
			}
		}
		if diff > res {
			res = diff
		}
	}
	fmt.Println(res)
}
