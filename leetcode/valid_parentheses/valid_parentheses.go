package main

import "fmt"

type stack []rune

func (s *stack) Push(a rune) {
	*s = append([]rune{a}, *s...)
}

func (s *stack) Pop() (rune, bool) {
	if len(*s) == 0 {
		return 0, false
	}
	lastElem := []rune(*s)[0]
	*s = []rune(*s)[1:]
	return lastElem, true
}

func isValid(s string) bool {
	st := stack(make([]rune, 0))
	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			st.Push(ch)
		} else {
			lastElem, ok := st.Pop()
			if !ok {
				return false
			}
			if ch == ')' && lastElem == '(' ||
				ch == ']' && lastElem == '[' ||
				ch == '}' && lastElem == '{' {
				continue
			} else {
				return false
			}
		}
	}
	return len(st) == 0
}

func main() {
	s := "()"
	fmt.Println(isValid(s))
	s = "()[]{}"
	fmt.Println(isValid(s))
	s = "(]"
	fmt.Println(isValid(s))
	s = "([)]"
	fmt.Println(isValid(s))
	s = "{[]}"
	fmt.Println(isValid(s))
	s = "("
	fmt.Println(isValid(s))
}
