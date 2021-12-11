package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, x, k int
	fmt.Scanf("%d %d %d", &n, &x, &k)
	arr := make([]int, 0, n)
	maxDay := -1
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		temp := scanner.Text()
		tempArr := strings.Split(temp, " ")
		for i := range tempArr {
			n, _ = strconv.Atoi(tempArr[i])
			arr = append(arr, n)
			if arr[i] > maxDay {
				maxDay = arr[i]
			}
		}
	}

	remainderForDays := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		temp := arr[i] % x
		if v, ok := remainderForDays[temp]; !ok {
			remainderForDays[temp] = arr[i]
		} else {
			if v > arr[i] {
				remainderForDays[temp] = arr[i]
			}
		}
	}
	i := 0
	for k > 0 {
		i++
		temp := i % x
		if v, ok := remainderForDays[temp]; ok && v <= i {
			k--
		}
		if maxDay == i {
			break
		}
	}
	for x <= k {
		i = i + x
		k -= len(remainderForDays)
	}
	for k > 0 {
		i++
		temp := i % x
		if v, ok := remainderForDays[temp]; ok && v <= i {
			k--
		}
	}
	fmt.Println(i)
}
