package main

import (
	"container/list"
	"fmt"
)

type Item struct {
	Key   string
	Value interface{}
}

type LRU struct {
	capacity int
	items    map[string]*list.Element
	queue    *list.List
}

func NewLRU(capacity int) *LRU {
	return &LRU{
		capacity: capacity,
		items:    make(map[string]*list.Element),
		queue:    list.New(),
	}
}

func (l *LRU) Set(key string, value interface{}) bool {
	if element, exists := l.items[key]; exists {
		l.queue.MoveToFront(element)
		element.Value.(*Item).Value = value
		return true
	}

	if l.queue.Len() == l.capacity {
		l.purge()
	}

	item := &Item{
		Key:   key,
		Value: value,
	}

	element := l.queue.PushFront(item)
	l.items[key] = element

	return true
}

func (l *LRU) Get(key string) interface{} {
	if element, exists := l.items[key]; exists {
		l.queue.MoveToFront(element)
		return element.Value.(*Item).Value
	}
	return nil
}

func (l *LRU) purge() {
	if element := l.queue.Back(); element != nil {
		item := l.queue.Remove(element).(*Item)
		delete(l.items, item.Key)
	}
}

func (l *LRU) Len() int {
	return l.queue.Len()
}

func main() {
	lru := NewLRU(5)
	lru.Set("addtotree", "1")
	lru.Set("b", "2")
	lru.Set("c", "3")
	lru.Set("d", "4")
	lru.Set("addtotree", "100")
	lru.Set("e", "5")
	lru.Set("f", "6")
	lru.Set("g", "7")
	lru.Set("h", "8")
	fmt.Println(lru.Len())
	fmt.Println(lru.Get("g"))
	fmt.Println(lru.Get("addtotree"))
	fmt.Println(lru.Get("d"))
}
