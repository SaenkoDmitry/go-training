package main

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"strconv"
)

// training 1: https://ya.cc/t/ZvQwZrmuiUYc3

// BinaryHeap куча с возвращением максимального элемента
type BinaryHeap struct {
	arr []int
}

// NewBinaryHeap создание бинарной кучи
func NewBinaryHeap() *BinaryHeap {
	return &BinaryHeap{arr: make([]int, 0)}
}

func HeapSort(arr []int) []int {
	result := make([]int, len(arr))
	binaryHeap := BuildHeapFromArray(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		result[i], _ = binaryHeap.Pop()
	}
	return result
}

// печать двоичной кучи в виде дерева
func (h *BinaryHeap) String() string {
	var buffer bytes.Buffer
	level := 0
	height := math.Log2(float64(len(h.arr)))
	currentCount := math.Pow(float64(2), float64(level))
	if len(h.arr) > 0 {
		for i := 0; i < int(height); i++ {
			buffer.WriteRune(' ')
		}
	}
	for i := range h.arr {
		buffer.WriteString(strconv.FormatInt(int64(h.arr[i]), 10))
		buffer.WriteRune(' ')
		currentCount--
		if currentCount == 0 {
			buffer.WriteRune('\n')
			level++
			currentCount = math.Pow(float64(2), float64(level))
			height--
			for i := 0; i < int(height); i++ {
				buffer.WriteRune(' ')
			}
		}
	}
	return buffer.String()
}

// Size размер бинарной кучи
func (h *BinaryHeap) Size() int {
	return len(h.arr)
}

// Push добавление нового элемента
// Идея: инсертим в конец списка и делаем всплытие элемента вверх до корня до тех пор, пока это необходимо
func (h *BinaryHeap) Push(elem int) {
	h.arr = append(h.arr, elem)
	i := len(h.arr) - 1
	parent := (i - 1) / 2
	for i > 0 && h.arr[i] > h.arr[parent] {
		h.arr[i], h.arr[parent] = h.arr[parent], h.arr[i]

		i = parent
		parent = (i - 1) / 2
	}
}

// реализация выравнивания бинарной кучи
func (h *BinaryHeap) heapify(i int) {
	for {
		left := 2 * i + 1
		right := 2 * i + 2
		largest := i

		if left < len(h.arr) && h.arr[left] > h.arr[largest] {
			largest = left
		}

		if right < len(h.arr) && h.arr[right] > h.arr[largest] {
			largest = right
		}

		if largest == i {
			break
		}

		h.arr[i], h.arr[largest] = h.arr[largest], h.arr[i]

		i = largest
	}
}

// Top извлечение максимального элемента
func (h *BinaryHeap) Top() (int, error) {
	if len(h.arr) == 0 {
		return 0, errors.New("binary heap is empty")
	}
	return h.arr[0], nil
}

// Pop извлечение максимального элемента + его удаление;
// Идея: ставим на место корня последний элемент и вызываем для него процедуру heap
func (h *BinaryHeap) Pop() (int, error) {
	if len(h.arr) == 0 {
		return 0, errors.New("nothing to extract")
	}
	top := h.arr[0]
	h.arr[0] = h.arr[len(h.arr) - 1]
	h.arr = h.arr[:len(h.arr) - 1]
	h.heapify(0)
	return top, nil
}

// BuildHeapFromArray O(N)-реализация построения бинарной кучи из массива
func BuildHeapFromArray(arr []int) *BinaryHeap {
	heap := NewBinaryHeap()
	heap.buildFromListWithoutRules(arr)
	for i := len(heap.arr) / 2; i >= 0; i-- {
		heap.heapify(i)
	}
	return heap
}

func (h *BinaryHeap) buildFromListWithoutRules(arr []int) {
	for i := range arr {
		h.arr = append(h.arr, arr[i])
	}
}

func main() {
	binaryHeap := NewBinaryHeap()

	binaryHeap.Push(1)
	binaryHeap.Push(4)
	binaryHeap.Push(3)
	binaryHeap.Push(5)
	binaryHeap.Push(15)
	binaryHeap.Push(11)
	binaryHeap.Push(17)
	binaryHeap.Push(13)
	binaryHeap.Push(2)
	binaryHeap.Push(3)
	binaryHeap.Push(2)
	binaryHeap.Push(4)
	binaryHeap.Push(7)
	binaryHeap.Push(8)
	binaryHeap.Push(9)

	if topElement, err := binaryHeap.Top(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(topElement)
	}

	fmt.Println(binaryHeap)

	fmt.Println(BuildHeapFromArray([]int{1,4,3,5,15,11,17,13,2,3,2,4,7,8,9}))
}
