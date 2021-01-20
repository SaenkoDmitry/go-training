package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var j, s string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		j = scanner.Text()
	}
	if scanner.Scan() {
		s = scanner.Text()
	}

	occurrence := make(map[rune]bool)
	for _, elem := range j {
		occurrence[elem] = true
	}

	var numOfJ int
	for _, elem := range s {
		if occurrence[elem] {
			numOfJ++
		}
	}

	fmt.Println(numOfJ)
}
