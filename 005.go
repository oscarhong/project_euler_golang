package main

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func main() {
	p := 1
	for i := 1; i <= 20; i++ {
		p = lcm(p, i)
	}
	print(p)
}
