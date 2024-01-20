package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

// An IntHeap is addtotree min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minEnergy(arr []int) int {
	h := IntHeap(arr)
	heap.Init(&h)

	acc := 0
	for h.Len() != 1 {
		left := heap.Pop(&h).(int)
		right := heap.Pop(&h).(int)
		temp := left + right
		acc += temp
		heap.Push(&h, temp)
	}
	return acc
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

	fmt.Println(minEnergy(arr))
}
