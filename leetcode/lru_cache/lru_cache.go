package main

import (
	"container/list"
	"fmt"
)

type Item struct {
	Key   int
	Value int
}

type LRUCache struct {
	capacity int
	items    map[int]*list.Element
	queue    *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		items:    make(map[int]*list.Element),
		queue:    list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if value, ok := this.items[key]; ok {
		this.queue.MoveToFront(value)
		return value.Value.(*Item).Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// exists
	if elem, ok := this.items[key]; ok {
		this.queue.MoveToFront(elem)
		elem.Value.(*Item).Value = value
		return
	}

	// overflow
	if this.queue.Len() == this.capacity {
		this.purge()
	}

	// not exists
	elem := this.queue.PushFront(&Item{
		Key:   key,
		Value: value,
	})
	this.items[key] = elem
}

func (this *LRUCache) purge() {
	if lastElem := this.queue.Back(); lastElem != nil {
		delete(this.items, lastElem.Value.(*Item).Key)
		this.queue.Remove(lastElem)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	obj := Constructor(2)
	obj.Put(1, 0)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	obj.Put(3, 3)
	fmt.Println(obj.Get(2))
	obj.Put(4, 4)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))
}
