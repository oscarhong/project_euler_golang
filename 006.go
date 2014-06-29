package main

func main() {
	sum, ssum := 0, 0
	for i := 1; i <= 100; i++ {
		sum += i
		ssum += i * i
	}
	print(sum*sum - ssum)
}
