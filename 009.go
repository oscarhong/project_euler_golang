package main

func main() {
	for c := 500; c > 3; c-- {
		for a := 1; a < (1000-c)/2; a++ {
			b := 1000 - c - a
			if a*a+b*b == c*c {
				print(a * b * c)
				break
			}
		}
	}
}
