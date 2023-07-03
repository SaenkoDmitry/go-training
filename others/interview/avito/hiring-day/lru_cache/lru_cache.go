package main

import (
	"bytes"
	"fmt"
)

type lruCache struct {
	arr []string
	m   map[string]struct{}
	cap int
}

func NewLruCache(size int) *lruCache {
	return &lruCache{
		arr: make([]string, 0, size),
		m:   make(map[string]struct{}),
		cap: size,
	}
}

func (l *lruCache) Store(elem string) {
	if _, ok := l.m[elem]; ok {
		index := 0
		for i := range l.arr {
			if l.arr[i] == elem {
				index = i
				break
			}
		}
		l.arr = append(l.arr, elem)
		l.arr = append(l.arr[:index], l.arr[index+1:]...)
	} else {
		if len(l.arr) == l.cap {
			l.arr = l.arr[1:]
		}
		l.arr = append(l.arr, elem)
		l.m[elem] = struct{}{}
	}
}

func (l *lruCache) Display() string {
	var buffer bytes.Buffer
	if len(l.arr) > 0 {
		buffer.WriteString((l.arr)[0])
	}
	for i := 1; i < len(l.arr); i++ {
		buffer.WriteString("-")
		buffer.WriteString((l.arr)[i])
	}
	return buffer.String()
}

func LRUCache(strArr []string) string {
	cache := NewLruCache(5)
	for i := 0; i < len(strArr); i++ {
		cache.Store(strArr[i])
	}
	return cache.Display()
}

func main() {
	fmt.Println(LRUCache([]string{"A", "B", "A", "C", "A", "B"}))
	fmt.Println(LRUCache([]string{"A", "B", "C", "D", "E", "D", "Q", "Z", "C"}))
	fmt.Println(LRUCache([]string{"A"}))
	fmt.Println(LRUCache([]string{}))
	fmt.Println(LRUCache([]string{"A", "B","C","D","A","A","A","A","A","T"}))
}
