package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, a int
	var s string

	fmt.Fscan(in, &n)
	fmt.Fscan(in, &m)

	acc := make(map[string][]int)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &a)
		fmt.Fscan(in, &s)

		clean(acc, a)

		if len(acc[s]) < n {
			fmt.Println(a, s)
			acc[s] = append(acc[s], a)
		}
	}
}

func clean(acc map[string][]int, a int) {
	for k := range acc {

		for i := 0; i < len(acc[k]); i++ {
			if a-acc[k][i] >= 1000 {
				acc[k] = acc[k][1:]
				i--
			}
		}
		if len(acc[k]) == 0 {
			delete(acc, k)
		}
	}
}
