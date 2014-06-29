package main

import (
	"fmt"
)

func contains(s []int, e int) int {
	for i, a := range s {
		if a == e {
			return i
		}
	}
	return -1
}

func clen(d int) int {
	remain, count := 1, 1
	var remains = []int{1}
	for remain < d && remain%d != 0 {
		remain *= 10
		remain %= d
		i := contains(remains, remain)
		if i >= 0 {
			return count - i
		} else {
			remains = append(remains, remain)
			count++
		}
	}
	return 0

}

func main() {
	g := 0
	d := 0
	for i := 1000; i > 950; i-- {
		l := clen(i)
		if l > g {
			g = l
			d = i
		}
	}
	fmt.Println(d)
}
