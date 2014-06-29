package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	n, sum := big.NewInt(2), 0
	n.Exp(n, big.NewInt(1000), nil)
	s := n.String()
	for i := 0; i < len(s); i++ {
		num, _ := strconv.Atoi(s[i : i+1])
		sum += num
	}
	fmt.Println(sum)
}
