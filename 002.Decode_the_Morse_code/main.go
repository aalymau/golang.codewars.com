package main

import (
	"fmt"
	"strings"
)

var (
	MORSE_CODE_DICT = map[string]string{
		"A": ".-", "B": "-...",
		"C": "-.-.", "D": "-..", "E": ".",
		"F": "..-.", "G": "--.", "H": "....",
		"I": "..", "J": ".---", "K": "-.-",
		"L": ".-..", "M": "--", "N": "-.",
		"O": "---", "P": ".--.", "Q": "--.-",
		"R": ".-.", "S": "...", "T": "-",
		"U": "..-", "V": "...-", "W": ".--",
		"X": "-..-", "Y": "-.--", "Z": "--..",
		"1": ".----", "2": "..---", "3": "...--",
		"4": "....-", "5": ".....", "6": "-....",
		"7": "--...", "8": "---..", "9": "----.",
		"0": "-----", ", ": "--..--", ".": ".-.-.-",
		"?": "..--..", "/": "-..-.", "-": "-....-",
		"(": "-.--.", ")": "-.--.-", "!": "-.-.--",
		"SOS": "...---...",
	}
	MORSE_CODE = reverseMap(MORSE_CODE_DICT)
)

func reverseMap(m map[string]string) map[string]string {
	n := make(map[string]string, len(m))
	for k, v := range m {
		n[v] = k
	}
	return n
}

// https://www.codewars.com/kata/54b724efac3d5402db00065e/train/go
// func DecodeMorse(morseCode string) string {
// 	res := ""
// 	for _, word := range strings.Split(morseCode, "   ") {
// 		for _, symbol := range strings.Split(word, " ") {
// 			res += MORSE_CODE[symbol]
// 		}
// 		res += " "
// 	}
// 	return strings.Trim(res, " ")
// }

// Best solution
func DecodeMorse(morseCode string) (decoded string) {
	for _, word := range strings.Split(strings.TrimSpace(morseCode), "   ") {
		for _, char := range strings.Split(word, " ") {
			decoded += MORSE_CODE[char]
		}
		decoded += " "
	}
	return strings.TrimSpace(decoded)
}

func main() {
	res := DecodeMorse(".... . -.--   .--- ..- -.. .")
	fmt.Printf("Result: <%s>\n", res)
}
