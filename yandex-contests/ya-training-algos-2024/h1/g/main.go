package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var x, y, p int

	fmt.Fscan(in, &x, &y, &p)

	defenders := 0

	rounds := 0
	for y > 0 || defenders > 0 {
		rounds++
		attackers := x
		if x > p { // attack defenders firstly
			if defenders > 0 {
				temp := min(x, defenders)
				defenders -= temp
				attackers -= temp
			}
			if y > 0 {
				temp := min(attackers, y)
				y -= temp
			}
		} else {
			if y > 0 {
				temp := min(x, y)
				y -= temp
				attackers -= temp
			}
			if defenders > 0 {
				temp := min(attackers, defenders)
				defenders -= temp
			}
		}
		if defenders > 0 {
			x -= defenders
		}
		if y > 0 {
			defenders += p
		}
		if x <= 0 {
			fmt.Println(-1)
			return
		}
	}
	fmt.Println(rounds)
}
