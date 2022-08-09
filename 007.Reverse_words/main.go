package main

import (
	"fmt"
)

// https://www.codewars.com/kata/5259b20d6021e9e14c0010d4/train/go
// func ReverseWords(str string) (ret string) {
// 	for _, word := range strings.Split(str, " ") {
// 		for i := len(word) - 1; i >= 0; i-- {
// 			ret += string(word[i])
// 		}
// 		ret += " "
// 	}
// 	return strings.TrimSpace(ret)
// }

// Clever solution
func ReverseWords(str string) string {
	var rev string
	var word string

	for _, i := range str {
		if i == ' ' {
			rev = rev + word + " " // Adds word and space to result
			word = ""              // Empties word variable
		} else {
			word = string(i) + word // Adds letter to temporary word variable
		}
	}

	return rev + word // reverse those words
}

func main() {
	res := ReverseWords("The quick brown fox jumps over the lazy dog.")
	fmt.Printf("Result: <%s>\n", res)
}
