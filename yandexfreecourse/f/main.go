package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func makeGoodString(input string) string {
	buf := bytes.Buffer{}
	runes := []rune(input)

	skipped := make([]bool, len(runes))

	for i := 0; i < len(runes); i++ {
		if skipped[i] {
			continue
		}
		j := i + 1
		for j < len(runes) && skipped[j] {
			j++
		}
		if j < len(runes) && runes[i] != runes[j] && strings.EqualFold(string(runes[i]), string(runes[j])) {
			skipped[i] = true
			skipped[j] = true
			for i > 0 && skipped[i] {
				i--
			}
			i--
		}
	}

	for i := 0; i < len(runes); i++ {
		if skipped[i] {
			continue
		}
		buf.WriteRune(runes[i])
	}
	return buf.String()
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s string
	fmt.Fscan(in, &s)
	fmt.Println(makeGoodString(s))
}
