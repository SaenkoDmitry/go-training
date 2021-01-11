package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func BypassSpiral(arr [][]int16) string {
	n := len(arr)

	var buffer bytes.Buffer

	counter := 1

	startI, startJ := n / 2, n / 2
	buffer.WriteString(strconv.Itoa(int(arr[startI][startJ])))

	cycleNumber := 1

	for startI >= 0 && startI < n && startJ >= 0 && startJ < n && counter < n * n {
		// вверх = уменьшаем i x N, где N = номер цикла
		for i := 0; i < cycleNumber && counter < n * n; i++ {
			startI--
			counter++
			buffer.WriteRune('\n')
			buffer.WriteString(strconv.Itoa(int(arr[startI][startJ])))
		}

		// вправо = увеличиваем j x N
		for i := 0; i < cycleNumber && counter < n * n; i++ {
			startJ++
			counter++
			buffer.WriteRune('\n')
			buffer.WriteString(strconv.Itoa(int(arr[startI][startJ])))
		}

		// увеличиваем номер цикла
		cycleNumber++

		// вниз = увеличиваем i x N
		for i := 0; i < cycleNumber && counter < n * n; i++ {
			startI++
			counter++
			buffer.WriteRune('\n')
			buffer.WriteString(strconv.Itoa(int(arr[startI][startJ])))
		}

		// влево = уменьшаем j x N
		for i := 0; i < cycleNumber && counter < n * n; i++ {
			startJ--
			counter++
			buffer.WriteRune('\n')
			buffer.WriteString(strconv.Itoa(int(arr[startI][startJ])))
		}

		// увеличиваем номер цикла
		cycleNumber++
	}
	return buffer.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	var n int
	if scanner.Scan() {
		n, _ = strconv.Atoi(scanner.Text())
	}

	start := time.Now()
	// read matrix
	arr := make([][]int16, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int16, n)
		for j := 0; j < n; j++ {
			if scanner.Scan() {
				temp, _ := strconv.Atoi(scanner.Text())
				arr[i][j] = int16(temp)
			}
		}
	}
	end := time.Now()

	log.Println("reading ended in ", end.Sub(start).Seconds())

	result := BypassSpiral(arr)
	fmt.Println(result)
}
