package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func score(name string) int {
	sum := 0
	for i := 0; i < len(name); i++ {
		if int(name[i]) >= 'A' && int(name[i]) <= 'Z' {
			sum += int(name[i]) - 'A' + 1
		}
	}
	return sum
}

func main() {
	buf, err := ioutil.ReadFile("022-names.txt")
	if err != nil {
		panic(err)
	}
	s := string(buf)
	names := strings.Split(s, ",")
	sort.Strings(names)
	sum := 0
	for i, s := range names {
		sum += (i + 1) * score(s)
	}
	fmt.Println(sum)
}
