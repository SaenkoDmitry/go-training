package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	var n int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	if scanner.Scan() {
		n, _ = strconv.Atoi(scanner.Text())
	}

	arr := make([]bool, n)
	start := time.Now()
	for i := 0; i < n-2; i++ {
		if scanner.Scan() {
			number, _ := strconv.Atoi(scanner.Text())
			arr[number - 1] = true
		}
	}
	end := time.Now()

	log.Println("first cycle", end.Sub(start).Seconds())

	var founded int
	start = time.Now()
	for i := 0; i < len(arr); i++ {
		if !arr[i] {
			if founded == 1 {
				fmt.Print(i + 1)
				end = time.Now()
				log.Println("second cycle", end.Sub(start).Seconds())
				return
			} else {
				fmt.Print(i + 1, " ")
			}
			founded++
		}
	}
}
