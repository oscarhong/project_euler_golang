package main

import (
	"fmt"
)

const upperLimit int = 28123
const sqrt int = 168

var primes [28123]int

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
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

func sumOfDiv(num int) int {
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

func hasSum(n int, nums []int) bool {
	if n%2 == 0 && contains(nums, n/2) {
		fmt.Println(n, " = ", n/2, " + ", n/2)
		return true
	}
	i, j := 0, len(nums)-1
	for i < j {
		s := nums[i] + nums[j]
		if s == n {
			fmt.Println(s, " = ", nums[i], " + ", nums[j])
			return true
		} else if s > n {
			i++
		} else {
			j--
		}
	}
	return false
}

func main() {
	ch := make(chan int) // Create a new channel.
	go Generate(ch)      // Launch Generate goroutine.
	for i := 0; i < sqrt; i++ {
		prime := <-ch
		primes[i] = prime
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
	var abundant []int
	for i := 12; i < 28123; i++ {
		if sumOfDiv(i) > i {
			abundant = append(abundant, i)
		}
	}
	var can [28123]bool
	for i := 0; i < len(abundant); i++ {
		for j := 0; j < len(abundant); j++ {
			s := abundant[i] + abundant[j]
			if s < 28123 {
				can[s] = true
			} else {
				break
			}
		}
	}
	sum := 0
	for i := 1; i < 28123; i++ {
		if !can[i] {
			sum += i
		}
	}
	fmt.Println(sum)
}
