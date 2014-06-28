package main

import (
	"fmt"
)

func main() {
	sum := 2
	f0, f1, f2 := 1, 2, 3
	for f2 < 4e6 {
		if f2%2 == 0 {
			sum += f2
		}
		f2 = f0 + f1
		f0, f1 = f1, f2
	}
	fmt.Println(sum)
}
