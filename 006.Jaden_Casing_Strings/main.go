package main

import (
	"fmt"
	"strings"
)

// https://www.codewars.com/kata/54b724efac3d5402db00065e/train/go
func ToJadenCase(str string) string {
	return strings.Title(str)
}

func main() {
	res := ToJadenCase("most trees are blue")
	fmt.Printf("Result: <%s>\n", res)
}
