package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
)

func PlusMinus(num int) string {
	arr := make([]int, 0)
	temp := num
	for temp > 0 {
		arr = append([]int{temp % 10}, arr...)
		temp = temp / 10
	}

	resultArr := make([]int, int(math.Pow(float64(2), float64(len(arr) - 1))))
	zeroMax := -1
	var result string
	for n := 0; n < len(resultArr); n++ {
		binaryString := strconv.FormatInt(int64(n), 2)
		for len(binaryString) < len(arr) {
			binaryString = "0" + binaryString
		}
		resultArr[n] = arr[0]
		for i := len(arr) - 1; i > 0; i-- {
			if binaryString[i] == '0' {
				resultArr[n] -= arr[i]
			} else {
				resultArr[n] += arr[i]
			}
		}
		if resultArr[n] == 0 {
			tempResult := binaryString[1:]
			temp := 0
			for j := range tempResult {
				if tempResult[j] == '0' {
					temp++
				}
			}
			if temp > zeroMax {
				zeroMax = temp
				result = tempResult
			}
		}
	}

	var buffer bytes.Buffer
	for i := range result {
		if result[i] == '0' {
			buffer.WriteRune('-')
		} else {
			buffer.WriteRune('+')
		}
	}

	if zeroMax == -1 {
		return "not possible"
	}
	return buffer.String()
}

func main() {
	fmt.Println(PlusMinus(35132))
	fmt.Println(PlusMinus(26712))
	fmt.Println(PlusMinus(1))
	fmt.Println(PlusMinus(1000))
	fmt.Println(PlusMinus(10))
	fmt.Println(PlusMinus(0))
}
