package main

import (
	"math"
)

func numOfDiv(n int) int {
	num, ter := 0, int(math.Sqrt(float64(n)))
	for i := 1; i <= ter; i++ {
		if n%i == 0 {
			num += 2
			if n == i {
				num -= 1
			}
		}
	}
	return num
}

func main() {
	t := 0
	d := 0
	for i := 1; d <= 500; i++ {
		t += i
		d = numOfDiv(t)
	}
	print(t)
}
