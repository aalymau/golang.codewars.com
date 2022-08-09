package main

import (
	"fmt"
	"strings"
)

// https://www.codewars.com/kata/52bc74d4ac05d0945d00054e/train/go
// func FirstNonRepeating(str string) string {
// 	for _, char := range str {
// 		if len(strings.Split(strings.ToLower(str), strings.ToLower(string(char)))) == 2 {
// 			return string(char)
// 		}
// 	}
// 	return ""
// }

// Clever solution
func FirstNonRepeating(str string) string {
	for _, c := range str {
		if strings.Count(strings.ToLower(str), strings.ToLower(string(c))) < 2 {
			return string(c)
		}
	}
	return ""
}

func main() {
	res := FirstNonRepeating("sTreSS")
	fmt.Printf("Result: %s\n", res)
}
