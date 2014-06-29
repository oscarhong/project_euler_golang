package main

import (
	"fmt"
	"math"
)

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ESieve(n int) []int {

	var primes []int
	ch := make(chan int) // Create a new channel.
	go Generate(ch)      // Launch Generate goroutine.
	for i := 0; ; i++ {
		prime := <-ch
		if prime > n {
			break
		}
		primes = append(primes, prime)
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
	return primes
}

//http://golang.org/doc/play/sieve.go
// Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

// The prime sieve: Daisy-chain Filter processes.

func main() {
	aMax, bMax, nMax := 0, 0, 0
	primes := ESieve(87400)
	bPos := ESieve(1000)

	for a := -999; a < 1001; a += 2 {
		for i := 1; i < len(bPos); i++ {
			for j := 0; j < 2; j++ {
				n := 0
				sign := -1
				if j == 0 {
					sign = 1
				}
				aodd := 0
				if i%2 == 0 {
					aodd = -1
				}
				for contains(primes, int(math.Abs(float64(n*n+(a+aodd)*n+sign*bPos[i])))) {
					n++
				}
				if n > nMax {
					aMax = a
					bMax = bPos[i]
					nMax = n
				}
			}
		}
	}
	fmt.Printf("a: %d, b: %d, p: %d", aMax, bMax, aMax*bMax)
}
