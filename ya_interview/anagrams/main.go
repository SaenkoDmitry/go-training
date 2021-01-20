package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

func main() {
	var first, second string
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		first = scanner.Text()
	}

	if scanner.Scan() {
		second = scanner.Text()
	}

	firstOccurrence := make(map[rune]int)
	secondOccurrence := make(map[rune]int)

	for _, elem := range first {
		firstOccurrence[elem]++
	}

	for _, elem := range second {
		secondOccurrence[elem]++
	}

	if len(first) == len(second) && reflect.DeepEqual(firstOccurrence, secondOccurrence) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
