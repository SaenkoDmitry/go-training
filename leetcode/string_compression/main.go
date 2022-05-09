package main

import (
	"strconv"
)

// Compress0 https://leetcode.com/problems/string-compression/
func Compress0(chars []byte) int {
	write, slow := 0, 0
	for fast := 0; fast <= len(chars); fast++ {
		if fast == len(chars) || chars[slow] != chars[fast] {
			chars[write] = chars[slow]
			write++
			length := fast - slow
			if length > 1 {
				for _, digit := range strconv.Itoa(length) {
					chars[write] = byte(digit)
					write++
				}
			}
			slow = fast
		}
	}
	return write
}

// Compress1 идея в том, чтобы завести два индекса: read index and write index
func Compress1(chars []byte) int {
	writeIndex := 0
	prev, temp := chars[0], chars[0]
	for i := 0; i < len(chars); i++ {
		prev, temp = temp, chars[i]
		if i > 0 && chars[i] == prev {
			count := 1
			for ; i < len(chars) && temp == chars[i]; i++ {
				count++
			}
			number := []byte(strconv.FormatInt(int64(count), 10))
			for k := 0; k < len(number); k++ {
				chars[writeIndex] = number[k]
				writeIndex++
			}
			i--
		} else {
			prev = chars[i]
			chars[writeIndex] = chars[i]
			writeIndex++
		}
	}
	return writeIndex
}
