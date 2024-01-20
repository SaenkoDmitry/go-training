package main

import (
	"bufio"
	"fmt"
	"os"
)

type Event struct {
	Name string
	T    int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	var events []Event

	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		var name string
		var t int
		fmt.Fscan(in, &name)
		fmt.Fscan(in, &t)
		events = append(events, Event{Name: name, T: t})
	}

	team := make(map[string]Event)
	prevExperience := make(map[string]int)
	for _, event := range events {
		if prev, ok := team[event.Name]; !ok {
			team[event.Name] = event
		} else {
			delete(team, event.Name)
			prevExperience[event.Name] += event.T - prev.T
		}
		fmt.Println(calcDiff(team, prevExperience, event.T))
	}
}

func calcDiff(team map[string]Event, prevExperience map[string]int, t int) string {
	max := -1
	maxName := ""

	res := 0
	for _, teammate := range team {
		prevExp, _ := prevExperience[teammate.Name]
		if t-teammate.T+prevExp > max || t-teammate.T+prevExp == max && teammate.Name[0] < maxName[0] {
			maxName = teammate.Name
			max = t - teammate.T + prevExp
		}

		res = res + (t - teammate.T + prevExp)
	}
	res -= max + max

	return fmt.Sprintf("%s %d", maxName, res)
}
