package main

import (
	"math/big"
)

func main() {
	var set = make(map[string]bool)
	for a := 2; a <= 100; a++ {
		for b := 2; b <= 100; b++ {
			result := big.NewInt(0).Exp(big.NewInt(int64(a)), big.NewInt(int64(b)), nil)
			set[result.String()] = true
		}
	}
	print(len(set))
}
