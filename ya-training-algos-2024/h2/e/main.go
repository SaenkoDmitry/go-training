package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Item struct {
	diff  int
	index int
	a, b  int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, a, b int

	fmt.Fscan(in, &n)

	res := make([]Item, 0, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a, &b)
		res = append(res, Item{a - b, i + 1, a, b})
	}

	sort.Slice(res, func(i, j int) bool {
		if res[i].diff < 0 && res[j].diff < 0 {
			return res[i].a > res[j].a
		}
		if res[i].a-res[i].b+res[j].a == res[j].a-res[j].b+res[i].a {
			return res[i].diff > res[j].diff
		}
		return res[i].a-res[i].b+res[j].a > res[j].a-res[j].b+res[i].a
	})

	h, curr := int64(0), int64(0)

	for i := 0; i < n; i++ {
		curr += int64(res[i].a)
		if curr > h {
			h = curr
		}
		curr -= int64(res[i].b)
	}

	fmt.Println(h)
	fmt.Print(res[0].index)
	for i := 1; i < n; i++ {
		fmt.Print(" ", res[i].index)
	}
}
