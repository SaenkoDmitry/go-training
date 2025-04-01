package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Item struct {
	x, y  int
	value int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, a int

	fmt.Fscan(in, &n, &m)

	arr := make([][]int, 0, n)
	items := make([]Item, 0, n*m)

	for i := 0; i < n; i++ {
		arr = append(arr, make([]int, m))
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &a)
			arr[i][j] = a
			items = append(items, Item{i, j, a})
		}
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].value > items[j].value
	})

	// try calc for items[0].x
	x := items[0].x
	y := -1
	coordX, coordY := 0, 0
	res := 0
	for i := 0; i < len(items); i++ {
		if items[i].x == x {
			continue
		}
		if y == -1 {
			y = items[i].y
			continue
		}
		if items[i].y == y {
			continue
		}
		coordX, coordY = x+1, y+1
		res = i
		break
	}

	// try calc for items[0].y
	x = -1
	y = items[0].y
	for i := 0; i < len(items); i++ {
		if items[i].y == y {
			continue
		}
		if x == -1 {
			x = items[i].x
			continue
		}
		if items[i].x == x {
			continue
		}
		if i > res {
			coordX, coordY = x+1, y+1
		}
		break
	}

	fmt.Println(coordX, coordY)
}
