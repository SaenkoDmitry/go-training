package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	var year int
	var day int
	var month string
	var dayOfWeek string

	fmt.Fscan(in, &n)
	fmt.Fscan(in, &year)

	holidays := make(map[string]int)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &day, &month)
		t, _ := time.Parse("2 January 2006", fmt.Sprintf("%d %s %d", day, month, year))
		holidays[t.Weekday().String()]++
	}
	fmt.Fscan(in, &dayOfWeek)

	weekDays := make(map[string]int)
	start, _ := time.Parse("2 January 2006", fmt.Sprintf("%d %s %d", 1, "January", year))
	for {
		weekDays[start.Weekday().String()]++
		start = start.Add(time.Hour * 24)
		if start.Month() == 1 && start.Day() == 1 {
			break
		}
	}

	holidaysCount := 0
	for _, v := range holidays {
		holidaysCount += v
	}

	maxVal, maxDay := 0, ""
	minVal, minDay := 367, ""
	for k, v := range weekDays {
		temp := v + holidaysCount - holidays[k]
		if temp > maxVal {
			maxVal, maxDay = temp, k
		}
		if temp < minVal {
			minVal, minDay = temp, k
		}
	}
	fmt.Println(maxDay, minDay)
}
