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

	leftArr := make([]int, n)
	rightArr := make([]int, n)
	ageArr := make([]int, n)
	for i := 0; i < n; i++ {
		var index, age, left, right int
		var temp []string
		if scanner.Scan() {
			temp = strings.Split(scanner.Text(), " ")
		}
		if len(temp) == 4 {
			index, _ = strconv.Atoi(temp[0])
			age, _ = strconv.Atoi(temp[1])
			left, _ = strconv.Atoi(temp[2])
			right, _ = strconv.Atoi(temp[3])
		}
		index--
		if left != -1 {
			left--
		}
		if right != -1 {
			right--
		}
		leftArr[index] = left
		rightArr[index] = right
		ageArr[index] = age
	}

	var currentLevelOfTree []int
	currentLevelOfTree = append(currentLevelOfTree, 0)

	var result []float64
	for len(currentLevelOfTree) > 0 {
		var nextLevelOfTree []int

		var accumSum float64
		var num int
		for len(currentLevelOfTree) > 0 {
			index := currentLevelOfTree[0]
			accumSum += float64(ageArr[index])
			num++
			if leftArr[index] != -1 {
				nextLevelOfTree = append(nextLevelOfTree, leftArr[index])
			}
			if rightArr[index] != -1 {
				nextLevelOfTree = append(nextLevelOfTree, rightArr[index])
			}
			currentLevelOfTree = currentLevelOfTree[1:]
		}
		result = append(result, accumSum / float64(num))
		currentLevelOfTree = nextLevelOfTree
	}

	if len(result) > 0 {
		fmt.Printf("%.2f", result[0])
	}
	for i := 1; i < len(result); i++ {
		fmt.Printf(" %.2f", result[i])
	}
}
