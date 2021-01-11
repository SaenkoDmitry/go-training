package main

import (
	"fmt"
	"math"
	"time"
)

type Pair struct {
	first float64
	second float64
}

type Stack []Pair

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(newElem Pair) {
	*s = append(*s, newElem)
}

func (s *Stack) Top() Pair {
	return (*s)[len(*s) - 1]
}

func (s *Stack) Pop() Pair {
	index := len(*s) - 1
	elem := (*s)[index]
	*s = (*s)[:len(*s) - 1]
	return elem
}

type QueueOn2Stacks struct {
	s1 *Stack
	s2 *Stack
}

func (q *QueueOn2Stacks) Add(newElem float64) {
	var minimum float64
	if q.s1.IsEmpty() {
		minimum = newElem
	} else {
		minimum = math.Min(newElem, q.s1.Top().second)
	}
	q.s1.Push(Pair{newElem, minimum})
}

func (q *QueueOn2Stacks) Min() float64 {
	if q.s1.IsEmpty() || q.s2.IsEmpty() {
		if q.s1.IsEmpty() {
			return q.s2.Top().second
		}
		return q.s1.Top().second
	}
	return math.Min(q.s1.Top().second, q.s2.Top().second)
}

func (q *QueueOn2Stacks) Pop() float64 {
	var result float64
	if q.s2.IsEmpty() {
		for !q.s1.IsEmpty() {
			elem := q.s1.Top().first
			q.s1.Pop()
			var minimum float64
			if q.s2.IsEmpty() {
				minimum = elem
			} else {
				minimum = math.Min(elem, q.s2.Top().second)
			}
			q.s2.Push(Pair{elem, minimum})
		}
	}
	result = q.s2.Top().second
	q.s2.Pop()
	return result
}

func NewQueueOn2Stacks() *QueueOn2Stacks {
	return &QueueOn2Stacks{
		new(Stack),
		new(Stack),
	}
}

func main() {
	queue := NewQueueOn2Stacks()

	queue.Add(5)
	queue.Add(4)
	queue.Add(15)
	queue.Add(3)
	queue.Add(2)

	fmt.Println(queue.Min())

	queue.Pop()
	fmt.Println(queue.Min())

	queue.Pop()
	fmt.Println(queue.Min())

	ch := make(chan int, 3)
	go func() {
		ch <- 5
		ch <- 4
		fmt.Println(1)
	}()

	<-time.After(3 * time.Second)

	go func() {
		for e := range ch {
			fmt.Println("loop:", e)
		}
	}()
	<-time.After(2 * time.Second)
}
