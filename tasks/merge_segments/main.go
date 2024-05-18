package main

import "fmt"

func solution(arr []int) []string {
	arr = append(arr, arr[len(arr)-1])
	res := make([]string, 0)
	start := 0
	for i := 1; i < len(arr); i++ {
		if arr[i-1]+1 == arr[i] {
			continue
		}
		end := i - 1
		if arr[start] != arr[end] {
			res = append(res, fmt.Sprintf("%d->%d", arr[start], arr[end]))
		} else {
			res = append(res, fmt.Sprintf("%d", arr[start]))
		}
		start = i
	}
	return res
}

type SomeInterface interface {
	Get() string
}

type SomeStruct struct{}

func (s SomeStruct) Get() string {
	return ""
}

func GetInterface() SomeStruct {
	return SomeStruct{}
}

func main() {
	arr := []int{0, 1, 2, 4, 5, 7}
	//arr := []int{0, 1, 2, 4, 5}
	//arr := []int{0, 1, 2, 4, 7}
	GetInterface().Get()
	fmt.Println(solution(arr))
}
