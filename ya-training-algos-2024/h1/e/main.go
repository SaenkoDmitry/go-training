package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k, d int
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &k)
	fmt.Fscan(in, &d)

	var i int
	for i = 1; i <= d; i++ {
		n = n * 10
		if n%k == 0 {
			break
		} else if k-n%k < 10 {
			n += k - n%k
		} else {
			fmt.Println(-1)
			return
		}
	}
	fmt.Print(n)
	for i < d {
		fmt.Print(0)
		i++
	}
	fmt.Println()
}
