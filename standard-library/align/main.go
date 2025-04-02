package main

import (
	"fmt"
	"unsafe"
)

type TestStr struct {
	aa bool
	bb string
	cc []int64
}

func main() {
	str := TestStr{}
	fmt.Println("size:", unsafe.Sizeof(str))
	fmt.Println("offset:", unsafe.Offsetof(str.cc))
	fmt.Println("align:", unsafe.Alignof(str))
}
