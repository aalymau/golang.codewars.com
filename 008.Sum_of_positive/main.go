package main

import (
	"fmt"
)

// https://www.codewars.com/kata/5715eaedb436cf5606000381/train/go
func PositiveSum(numbers []int) (sum int) {
	for _, num := range numbers {
		if num > 0 {
			sum += num
		}
	}
	return
}

func main() {
	res := PositiveSum([]int{1, 2, 3, 4, 5})
	fmt.Printf("Result: <%v>\n", res)
}
