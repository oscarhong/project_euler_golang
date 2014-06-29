package main

func main() {
	count, day := 0, 1
	for year := 1900; year <= 2000; year++ {
		for month := 1; month <= 12; month++ {
			switch month {
			case 4, 6, 9, 11:
				day += 30
			case 1, 3, 5, 7, 8, 10, 12:
				day += 31
			case 2:
				{
					if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
						day += 29
					} else {
						day += 28
					}
				}
			}
			day %= 7
			if day == 0 && year >= 1901 {
				count++
			}
		}
	}
	print(count)
}
