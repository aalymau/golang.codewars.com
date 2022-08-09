package main

import "fmt"

// https://www.codewars.com/kata/57f781872e3d8ca2a000007e/train/go
// func Maps(x []int) []int {
// 	res := make([]int, len(x))
// 	for i := range x {
// 		res[i] = x[i] * 2
// 	}
// 	return res
// }

// Clever solution
func Maps(x []int) (result []int) {
	for _, j := range x {
		result = append(result, j*2)
	}
	return
}

func main() {
	initial := []int{1, 2, 3}
	res := Maps(initial)
	fmt.Printf("Initial: ")
	for _, val := range initial {
		fmt.Printf("%v ", val)
	}
	fmt.Println()
	fmt.Printf("Result : ")
	for _, val := range res {
		fmt.Printf("%v ", val)
	}
	fmt.Println()
}
