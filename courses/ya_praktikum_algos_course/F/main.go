package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var n int
	if scanner.Scan() {
		n, _ = strconv.Atoi(scanner.Text())
	}

	var arr []int
	if scanner.Scan() {
		tempApp := strings.Split(scanner.Text(), " ")
		for i := 0; i < len(tempApp); i++ {
			temp, _ := strconv.Atoi(tempApp[i])
			arr = append(arr, temp)
		}
	}

	occurrence := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		occurrence[arr[i]]++
	}

	for k := 0; k < 10000; k++ {
		for i := 0; i < len(arr); i++ {
			for j := 0; j < len(arr); j++ {
				if i < j {
					var calculatedValue int
					calculatedValue = n - (arr[i] + arr[j] + k)
					v1, ok1 := occurrence[calculatedValue]
					if (calculatedValue == arr[i] && v1 == 1) ||
						(calculatedValue == arr[j] && v1 == 1) ||
						(arr[i] == arr[j] && calculatedValue == arr[i] && v1 < 3) {
						continue
					}
					if ok1 {
						fmt.Println(n - k)
						return
					}
					calculatedValue = n - (arr[i] + arr[j] - k)
					v2, ok2 := occurrence[calculatedValue]
					if (calculatedValue == arr[i] && v2 == 1) ||
						(calculatedValue == arr[j] && v2 == 1) ||
						(arr[i] == arr[j] && calculatedValue == arr[i] && v2 < 3) {
						continue
					}
					if ok2 {
						fmt.Println(n + k)
						return
					}
				}
			}
		}
	}
}
