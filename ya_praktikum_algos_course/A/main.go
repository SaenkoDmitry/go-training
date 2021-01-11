package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var n int
	if scanner.Scan() {
		if temp, err := strconv.Atoi(scanner.Text()); err == nil {
			n = temp
		} else {
			panic("wrong input")
		}
	}

	var arr, rus, eng []string

	// fill rus playlist
	if scanner.Scan() {
		arr = strings.Split(scanner.Text(), " ")
		for _, item := range arr {
			rus = append(rus, item)
		}
	}

	// fill eng playlist
	if scanner.Scan() {
		arr = strings.Split(scanner.Text(), " ")
		for _, item := range arr {
			eng = append(eng, item)
		}
	}

	if len(rus) < n || len(eng) < n {
		panic("wrong input")
	} else {
		if len(rus) > 0 && len(eng) > 0 {
			fmt.Print(rus[0], " ", eng[0])
		}
		for i := 1; i < n && len(rus) > i && len(eng) > i; i++ {
			fmt.Print(" ", rus[i], " ", eng[i])
		}
	}
}
