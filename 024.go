package main

func main() {
	var perm = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	count, numPerm := 1, 1000000
	for count < numPerm {
		N := len(perm)
		i := N - 1
		for perm[i-1] >= perm[i] {
			i--
		}
		j := N
		for perm[j-1] <= perm[i-1] {
			j--
		}
		perm[i-1], perm[j-1] = perm[j-1], perm[i-1]
		i++
		j = N
		for i < j {
			perm[i-1], perm[j-1] = perm[j-1], perm[i-1]
			i++
			j--
		}
		count++
	}
	for k := 0; k < len(perm); k++ {
		print(perm[k])
	}
}
