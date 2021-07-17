package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		n, _ = strconv.Atoi(scanner.Text())
	}

	var curr, prev string

	if n < 1 {
		return
	}

	// print only first element, because it's always unique
	if scanner.Scan() {
		curr = scanner.Text()
	}
	fmt.Println(curr)
	prev = curr

	// print next elements if its unique
	for i := 1; i < n; i++ {
		if scanner.Scan() {
			curr = scanner.Text()
		}
		if prev != curr {
			fmt.Println(curr)
		}
		prev = curr
	}
}
