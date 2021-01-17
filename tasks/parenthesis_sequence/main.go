package main

import (
	"fmt"
)

func parenthesisGenerator(n int) {
	recursive("", []rune{}, n, n)
}

func recursive(accum string, stack []rune, l int, r int) {
	if l == 0 && r == 0 {
		fmt.Println(accum)
		return
	}

	if l > 0 {
		recursive(accum+"(", append([]rune{'('}, stack...), l-1, r)
		//recursive(accum+"{", append([]rune{'{'}, stack...), l-1, r)
		//recursive(accum+"[", append([]rune{'['}, stack...), l-1, r)
	}

	if l < r {
		if stack[0] == '(' {
			recursive(accum+")", stack[1:], l, r-1)
		}
		//if stack[0] == '{' {
		//	recursive(accum+"}", stack[1:], l, r-1)
		//}
		//if stack[0] == '[' {
		//	recursive(accum+"]", stack[1:], l, r-1)
		//}
	}
}

func main() {
	//scanner := bufio.NewScanner(os.Stdin)
	//var n int
	//if scanner.Scan() {
	//	s := scanner.Text()
	//	var err error
	//	n, err = strconv.Atoi(s)
	//	if err != nil {
	//		panic(err)
	//	}
	//}
	parenthesisGenerator(2)
}
