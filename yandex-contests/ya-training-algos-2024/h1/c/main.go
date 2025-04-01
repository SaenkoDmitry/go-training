package main

import (
	"bufio"
	"fmt"
	"os"
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

	m := make(map[int]int)
	m[1] = 1
	m[2] = 2
	m[3] = 2
	m[4] = 1

	res := 0
	for i := 0; i < n; i++ {
		if arr[i] >= 4 {
			res += arr[i] / 4
		}
		res += m[arr[i]%4]
	}
	fmt.Println(res)
}
