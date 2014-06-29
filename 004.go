package main

import (
	"strconv"
)

func isPalindrome(in string) bool {
	for i := 0; i < len(in); i++ {
		if in[i] != in[len(in)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	bp := 0
	for i := 999; i > 900; i-- {
		for j := 999; j > 900; j-- {
			if isPalindrome(strconv.Itoa(i*j)) && i*j > bp {
				bp = i * j
			}
		}
	}
	print(bp)
}
