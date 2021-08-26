package main

import (
	"bytes"
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

type Tree struct {
	value int
	inner map[string]*Tree
}

// recursive
func (t *Tree) String() string {
	if len(t.inner) == 0 {
		return buildVertex("root", t.value)
	}

	var resultBuffer bytes.Buffer
	buf := make([]bytes.Buffer, 0)
	buf = append(buf, bytes.Buffer{})

	buf[0].WriteString(buildVertex("root", t.value))
	for key := range t.inner {
		recursive(t.inner[key], key, 1, &buf)
	}

	for i := range buf {
		resultBuffer.WriteString(buf[i].String())
		resultBuffer.WriteString("\n")
	}
	return resultBuffer.String()
}

func recursive(vertex *Tree, key string, level int, buf *[]bytes.Buffer) {
	if len(*buf) <= level {
		*buf = append(*buf, bytes.Buffer{})
	}
	(*buf)[level].WriteString(buildVertex(key, vertex.value) + " ")
	for key := range vertex.inner {
		recursive(vertex.inner[key], key, level + 1, buf)
	}
}

func (t *Tree) String1() string {
	var buf bytes.Buffer

	currQueue := list.New()
	currQueue.PushBack(t)
	if len(t.inner) == 0 {
		return buildVertex("root", t.value)
	}
	currLevel := len(t.inner) - 1
	currIndex := 0
	for currQueue.Len() != 0 {
		elem := currQueue.Front()
		currValue := elem.Value.(*Tree).value
		buf.WriteString(buildVertex("root", currValue))
		if currIndex == currLevel {
			buf.WriteString("\n")
		}
		items := elem.Value.(*Tree).inner
		if currIndex == currLevel {

		}
		for key := range items {
			buf.WriteString(key + " ")
			currQueue.PushBack(items[key])
		}
		currQueue.Remove(elem)
		currIndex++
	}
	return buf.String()
}

func buildVertex(key string, value int) string {
	if value == 0 {
		return key
	}
	return key + ":" + strconv.Itoa(value)
}

func setValueByPath(path []string, value int, root *Tree) {
	if len(path) == 0 {
		return
	} else if len(path) == 1 {
		if _, ok := root.inner[path[0]]; !ok {
			root.inner[path[0]] = &Tree{
				value: value,
				inner: make(map[string]*Tree),
			}
		}
	}

	if _, ok := root.inner[path[0]]; !ok {
		root.inner[path[0]] = &Tree{
			inner: make(map[string]*Tree),
		}
	}
	newRoot := root.inner[path[0]]
	setValueByPath(path[1:], value, newRoot)
}

func build(input string, accum *Tree) {
	arr := strings.Split(input, "=")
	value, _ := strconv.Atoi(arr[1])
	setValueByPath(strings.Split(arr[0], "."), value, accum)
}

func buildAll(inputs []string) *Tree {
	root := &Tree{inner: make(map[string]*Tree)}
	for i := range inputs {
		build(inputs[i], root)
	}
	return root
}

// [key1 [ key2 [] key3 [] ] key3]

func main() {
	inputs := []string{
		"key1.key2=4",
		"key1.key2.key5=2",
		"key1.key3.key7=8",
		"key10.key4=9",
		"key11=10",
	}
	fmt.Println()
	fmt.Println(buildAll(inputs))
}
