package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func factorial(n int64) *big.Int {
	if n < 0 {
		return big.NewInt(1)
	}
	if n == 0 {
		return big.NewInt(1)
	}
	bigN := big.NewInt(n)
	return bigN.Mul(bigN, factorial(n-1))
}

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func main() {
	s, sum := factorial(100).String(), 0
	for i := 0; i < len(s); i++ {
		sum += atoi(s[i : i+1])
	}
	fmt.Println(sum)
}
