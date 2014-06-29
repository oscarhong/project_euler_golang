package main

var ones = [...]int{0, 3, 3, 5, 4, 4, 3, 5, 5, 4, 3, 6, 6, 8, 8, 7, 7, 9, 8, 8, 6}
var tens = [...]int{0, 3, 6, 6, 5, 5, 5, 7, 6, 6, 7}

func nLetters(n int) int {
	if n == 1000 {
		return 11
	} else {
		count, flag := 0, false
		if n >= 100 {
			count += ones[n/100] + tens[10]
			flag = true
			n %= 100
		}
		if n >= 20 {
			count += tens[n/10]
			n %= 10
			if flag {
				count += 3
				flag = false
			}
		}
		if flag && n != 0 {
			count += 3
		}
		count += ones[n]
		return count
	}
}

func main() {
	sum := 0
	for i := 1; i <= 1000; i++ {
		sum += nLetters(i)
	}
	print(sum)
}
