package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int

	fmt.Fscan(in, &n)
	fmt.Fscan(in, &m)

	digitsCount := maxDigitsCount(n, m)

	res := 0
	for i := 0; i < n; i++ {
		reversed := reverseNumber(i, digitsCount)
		if reversed < m {
			res++
		}
	}

	fmt.Println(res)
}

func reverseNumber(input int, digitsCount int) int {
	temp := strconv.FormatInt(int64(input), 10)
	temp = Reverse(temp)
	for len(temp) < digitsCount {
		temp = temp + "0"
	}
	res, _ := strconv.ParseInt(temp, 10, 64)
	return int(res)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func maxDigitsCount(a, b int) int {
	digitsCount := 0
	var temp int
	if a > b {
		temp = a
	} else {
		temp = b
	}
	for temp > 0 {
		temp = temp / 10
		digitsCount++
	}
	return digitsCount
}
