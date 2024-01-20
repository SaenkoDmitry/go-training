package main

import (
	"bufio"
	"fmt"
	"os"
)

func parseRome(number string) int {

	iCount, xCount, cCount, mCount := 0, 0, 0, 0 // x3 раз подряд
	vCount, lCount, dCount := 0, 0, 0            // x1 раз макс во всей записи

	runes := []rune(number)
	for i := range runes {
		if runes[i] != 'I' && runes[i] != 'V' && runes[i] != 'X' && runes[i] != 'L' &&
			runes[i] != 'C' && runes[i] != 'D' && runes[i] != 'M' {
			return -1
		}

		if runes[i] == 'I' {
			iCount++
			xCount = 0
			cCount = 0
			mCount = 0
		} else if runes[i] == 'X' {
			xCount++
			iCount = 0
			cCount = 0
			mCount = 0
		} else if runes[i] == 'C' {
			cCount++
			iCount = 0
			xCount = 0
			mCount = 0
		} else if runes[i] == 'M' {
			mCount++
			iCount = 0
			xCount = 0
			cCount = 0
		} else if runes[i] == 'V' {
			vCount++
			iCount = 0
			xCount = 0
			cCount = 0
			mCount = 0
		} else if runes[i] == 'L' {
			lCount++
			iCount = 0
			xCount = 0
			cCount = 0
			mCount = 0
		} else if runes[i] == 'D' {
			dCount++
			iCount = 0
			xCount = 0
			cCount = 0
			mCount = 0
		}

		if iCount > 3 || xCount > 3 || cCount > 3 || mCount > 3 || vCount > 1 || lCount > 1 || dCount > 1 {
			return -1
		}
	}
	if iCount > 3 || xCount > 3 || cCount > 3 || mCount > 3 || vCount > 1 || lCount > 1 || dCount > 1 {
		return -1
	}

	result := 0
	prev := 0
	curr := 0
	for i := 0; i < len(runes); i++ {
		if runes[i] == 'I' {
			if i+1 < len(runes) && runes[i+1] == 'V' {
				curr = 4
				i++
			} else if i+1 < len(runes) && runes[i+1] == 'X' {
				curr = 9
				i++
			} else {
				curr = 1
			}
		} else if runes[i] == 'X' {
			if i+1 < len(runes) && runes[i+1] == 'L' {
				curr = 40
				i++
			} else if i+1 < len(runes) && runes[i+1] == 'C' {
				curr = 90
				i++
			} else {
				curr = 10
			}
		} else if runes[i] == 'C' {
			if i+1 < len(runes) && runes[i+1] == 'D' {
				curr = 400
				i++
			} else if i+1 < len(runes) && runes[i+1] == 'M' {
				curr = 900
				i++
			} else {
				curr = 100
			}
		} else if runes[i] == 'V' {
			curr = 5
		} else if runes[i] == 'L' {
			curr = 50
		} else if runes[i] == 'D' {
			curr = 500
		} else if runes[i] == 'M' {
			curr = 1000
		}

		if prev != 0 && prev < curr {
			return -1
		}
		result += curr
		prev = curr
	}

	return result
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s string
	fmt.Fscan(in, &s)
	fmt.Println(parseRome(s))
}
