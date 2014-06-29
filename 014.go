package main

import (
	"fmt"
	"runtime"
)

const NCPU = 4

type result struct {
	chain int
	from  int
}

func collatz(n int) int {
	count := 1
	n64 := int64(n)
	for n64 != 1 {
		if n64%2 == 0 {
			n64 /= 2
		} else {
			n64 = 3*n64 + 1
		}
		count++
	}
	return count
}

func worker(n, total int, c chan result) {
	g := result{}
	for i := total - 1; i > total/2; i -= 2 * NCPU {
		r := result{collatz(i), i}
		if r.chain > g.chain {
			g = r
		}
	}
	c <- g
}

func main() {
	g := result{}
	runtime.GOMAXPROCS(NCPU)
	c := make(chan result, NCPU)
	for i := 0; i < NCPU; i++ {
		go worker(i, 1e6, c)
	}
	for i := 0; i < NCPU; i++ {
		r := <-c
		if r.chain > g.chain {
			g = r
		}
	}
	fmt.Println(g)
}
