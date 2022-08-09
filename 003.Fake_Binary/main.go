package main

import "fmt"

// https://www.codewars.com/kata/57eae65a4321032ce000002d/train/go
func FakeBin(x string) (res string) {
	for _, symbol := range x {
		if symbol < '5' {
			res += "0"
		} else {
			res += "1"
		}
	}
	return
}

func main() {
	res := FakeBin("45385593107843568")
	fmt.Printf("Result: <%s>\n", res)
}
