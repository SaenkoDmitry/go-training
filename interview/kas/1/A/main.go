package main

import (
	"bytes"
	"fmt"
)

func Solution(N int) string {
	var buffer bytes.Buffer
	for i := 0; i < N-1; i++ {
		buffer.WriteString("x")
	}
	if N % 2 == 0 {
		buffer.WriteString("y")
	} else {
		buffer.WriteString("x")
	}
	return buffer.String()
}

func main() {
	fmt.Println(Solution(1))
	fmt.Println(Solution(4))
	fmt.Println(Solution(7))
}
