package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)


func ReverseNumber(input string) string {
	var buffer bytes.Buffer
	if input[0] == '-' {
		buffer.WriteString("-")
		input = input[1:]
	}

	var nonZeroDigitFound bool
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] == '0' && !nonZeroDigitFound {
			continue
		} else if input[i] != 0 {
			nonZeroDigitFound = true
		}
		buffer.WriteString(string(input[i]))
	}
	if !nonZeroDigitFound {
		return "0"
	}
	return buffer.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var input string
	if scanner.Scan() {
		input = scanner.Text()
	}

	res := ReverseNumber(input)

	fmt.Println(res)
}
