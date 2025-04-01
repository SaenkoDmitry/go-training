package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var p, v int
	var q, m int

	fmt.Fscan(in, &p)
	fmt.Fscan(in, &v)
	fmt.Fscan(in, &q)
	fmt.Fscan(in, &m)

	res := v*2 + m*2 + 2

	if p > q {
		p, v, q, m = q, m, p, v
	}

	if p+v < q-m {
		fmt.Println(res)
		return
	}

	left, right := max(p-v, q-m), min(p+v, q+m)
	fmt.Println(res - right + left - 1)
}
