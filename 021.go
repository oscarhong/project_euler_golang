package main

import (
	"fmt"
)

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

func sumOfDiv(num int, primes [100]int) int {
	n, sum, p, i, j := num, 1, primes[0], 0, 0
	for p*p <= n && n > 1 && i < len(primes) {
		p = primes[i]
		i++
		if n%p == 0 {
			j, n = p*p, n/p
			for n%p == 0 {
				j, n = j*p, n/p
			}
			sum *= (j - 1) / (p - 1)
		}
	}
	if n > 1 {
		sum *= n + 1
	}
	return sum - num
}

func main() {
	var primes [100]int
	ch := make(chan int) // Create a new channel.
	go Generate(ch)      // Launch Generate goroutine.
	for i := 0; i < 100; i++ {
		prime := <-ch
		primes[i] = prime
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
	sum := 0
	for i := 2; i < 10000; i++ {
		j := sumOfDiv(i, primes)
		if j > i && j < 10000 && sumOfDiv(j, primes) == i {
			sum += i + j
		}
	}
	fmt.Println(sum)
}
