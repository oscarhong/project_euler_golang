package main

func main() {
	sum, t := 1, 1
	for width := 3; width <= 1001; width += 2 {
		for i := 0; i < 4; i++ {
			t += width - 1
			sum += t
		}
	}
	print(sum)
}
