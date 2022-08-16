package main

import (
	"fmt"
)

func CalculateAverage(numbers []int) (sum float64) {
	for _, num := range numbers {
		sum += float64(num)
	}
	if len(numbers) == 0 {
		return 0.0
	} else {
		return sum / float64(len(numbers))
	}
}

func main() {
	res := CalculateAverage([]int{1, 2, 3, 4, 5})
	fmt.Printf("Result: <%v>\n", res)
}
