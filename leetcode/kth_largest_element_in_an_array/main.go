package main

import (
	"errors"

	"golang.org/x/exp/constraints"
)

func findKthLargest(nums []int, k int) int {
	heap := NewBinaryHeap[int]()
	for _, num := range nums {
		heap.Push(num)
	}
	var result int
	for k > 0 {
		result, _ = heap.Pop()
		k--
	}
	return result
}

// this is max binary heap
type BinaryHeap[T constraints.Ordered] struct {
	arr []T
}

func NewBinaryHeap[T constraints.Ordered]() *BinaryHeap[T] {
	return &BinaryHeap[T]{arr: make([]T, 0)}
}

// добавляем в конец, делаем всплытие этого элемента
func (b *BinaryHeap[T]) Push(elem T) {
	b.arr = append(b.arr, elem)
	i := len(b.arr) - 1
	parent := (i - 1) / 2
	for i > 0 && b.arr[i] > b.arr[parent] {
		b.arr[i], b.arr[parent] = b.arr[parent], b.arr[i]
		i = parent
		parent = (i - 1) / 2
	}
}

// берем верхний элемент, ставим последний на его место, затем вызываем heapify(0)
func (b *BinaryHeap[T]) Pop() (T, error) {
	if len(b.arr) == 0 {
		var zero T
		return zero, errors.New("empty heap")
	}
	top := b.arr[0]
	b.arr[0] = b.arr[len(b.arr)-1]
	b.arr = b.arr[:len(b.arr)-1]
	b.heapify(0)
	return top, nil
}

func (b *BinaryHeap[T]) heapify(i int) {
	largest := i
	for {
		left := 2*i + 1
		right := 2*i + 2
		if left < len(b.arr) && b.arr[left] > b.arr[largest] {
			largest = left
		}
		if right < len(b.arr) && b.arr[right] > b.arr[largest] {
			largest = right
		}
		if i == largest {
			break
		}
		b.arr[largest], b.arr[i] = b.arr[i], b.arr[largest]
		i = largest
	}
}

func (b *BinaryHeap[T]) Len() int {
	return len(b.arr)
}

func main() {
	// Пример использования
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	result := findKthLargest(nums, k)
	println(result)
}
