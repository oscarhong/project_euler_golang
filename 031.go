package main

func main() {
	var sizes = []int{1, 2, 5, 10, 20, 50, 100, 200}
	var ways = [201]int{}
	ways[0] = 1

	for i := 0; i < len(sizes); i++ {
		for j := sizes[i]; j <= 200; j++ {
			ways[j] += ways[j-sizes[i]]
		}
	}
	print(ways[200])
}
