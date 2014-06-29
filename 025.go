package main

import (
	"fmt"
	"math/big"
)

func main() {
	f1, f2, fn, count := big.NewInt(1), big.NewInt(1), big.NewInt(2), 3
	for len(fn.String()) < 1000 {
		f1, f2 = f2, fn
		fn = big.NewInt(0).Add(f1, f2)
		count++
	}
	fmt.Println("F", count, " = ", fn.String())
}
