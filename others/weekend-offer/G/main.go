package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


/**
1
1
 */

/**
2
4
2
*/

/**
2
1
2
*/

/**
2
2
2
*/

/**
3
2
4
5
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var n int
	arr := make([]int, 0, n)
	discount := make([]int, 0, n)

	if scanner.Scan() {
		n, _ = strconv.Atoi(scanner.Text())
	}

	for i := 0; i < n; i++ {
		if scanner.Scan() {
			temp, _ := strconv.Atoi(scanner.Text())
			arr = append(arr, temp)
		}
	}

	if len(arr) == 2 {
		if arr[0] != arr[1] {
			fmt.Println(1500)
			return
		}
	}

	for i := 0; i < len(arr); i++ {
		discount = append(discount, 500)
	}

	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] && discount[i] <= discount[i-1] {
			discount[i] = discount[i-1] + 500
		} else if arr[i] < arr[i-1] && discount[i] >= discount[i-1] {
			discount[i-1] = discount[i] + 500
			for j := i-1; j >= 1; j-- {
				if arr[j] < arr[j-1] && discount[j] >= discount[j-1] {
					discount[j-1] = discount[j] + 500
				}
			}
		}
	}

	sum := 0
	for i := 0; i < len(discount); i++ {
		sum += discount[i]
	}

	fmt.Println(sum)
}
