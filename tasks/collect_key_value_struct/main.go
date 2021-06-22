package main

import (
	"strconv"
	"strings"
)

type Tree struct {
	value int
	inner map[string]*Tree
}


func recursiveCollect() {

}

func setValueByPath(path []string, value int, root *Tree) {
	// TODO : implement logic for setting value by path

}

func build(input string, accum *Tree) {
	arr := strings.Split(input, "=")
	value, _ := strconv.Atoi(arr[1])
	setValueByPath(strings.Split(arr[0], "."), value, accum)
}

func buildAll(inputs []string) *Tree {
	root := &Tree{value: nil, inner: make(map[string]*Tree)}
	for i := range inputs {
		build(inputs[i], root)
	}
	return root
}

func main() {
	inputs := []string{
		"key1.key2=value1",
		"key1.key2.key5=value2",
		"key3.key4=value3",
		"key3=value3",
	}
	buildAll(inputs)
}
