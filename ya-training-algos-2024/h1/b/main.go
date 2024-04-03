package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var s1, s2 string
	var d int
	fmt.Fscan(in, &s1)
	fmt.Fscan(in, &s2)
	fmt.Fscan(in, &d)

	arr1 := strings.Split(s1, ":")
	arr2 := strings.Split(s2, ":")

	temp1, _ := strconv.ParseInt(arr1[0], 10, 64)
	temp2, _ := strconv.ParseInt(arr1[1], 10, 64)
	temp3, _ := strconv.ParseInt(arr2[0], 10, 64)
	temp4, _ := strconv.ParseInt(arr2[1], 10, 64)

	firstTeamPoints := int(temp1) + int(temp3)
	secondTeamPoints := int(temp2) + int(temp4)

	var firstTeamGuestPoints, secondTeamGuestPoints int
	if d == 1 {
		firstTeamGuestPoints = int(temp3)
		secondTeamGuestPoints = int(temp2)
	} else {
		firstTeamGuestPoints = int(temp1)
		secondTeamGuestPoints = int(temp4)
	}

	if firstTeamPoints > secondTeamPoints {
		fmt.Println(0)
		return
	}

	res := 0
	if firstTeamPoints < secondTeamPoints {
		res += secondTeamPoints - firstTeamPoints
		if d == 1 {
			firstTeamGuestPoints += res
		}
	}

	if firstTeamGuestPoints <= secondTeamGuestPoints {
		res += 1
	}
	fmt.Println(res)
}
