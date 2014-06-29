package main

import (
	"fmt"
	"math"
	"strconv"
)

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func main() {
	var p = make(map[int]int)
	for i := 1; i < 10; i++ {
		p[i] = int(math.Pow(float64(i), float64(5)))
	}
	sum := 0
	for n := 1000; n <= 355000; n++ {
		s := strconv.Itoa(n)
		ssum := 0
		for i := 0; i < len(s); i++ {
			ssum += p[atoi(s[i:i+1])]
		}
		if ssum == n {
			sum += n
		}
	}
	fmt.Println(sum)
}
