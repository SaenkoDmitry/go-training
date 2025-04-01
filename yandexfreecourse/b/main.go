package main

import (
	"bufio"
	"fmt"
	"os"
)

func findMaxSum(n, k int, arr []int) int {
	max := 0
	left, right := k-1, len(arr)-1

	leftSum, rightSum := 0, 0

	for i := 0; i < k; i++ {
		leftSum += arr[i]
	}
	max = leftSum

	for left >= 0 {
		leftSum -= arr[left]
		rightSum += arr[right]
		left--
		right--
		temp := leftSum + rightSum
		if temp > max {
			max = temp
		}
	}

	return max
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k int
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &k)

	arr := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(in, &a)
		arr = append(arr, a)
	}
	fmt.Println(findMaxSum(n, k, arr))
}

// output:
// 0
//
