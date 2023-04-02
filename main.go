package main

import "fmt"

func SlurredTalk(words *string) {
	result := ""
	for _, c := range *words {

		if c == 's' || c == 'r' || c == 'z' {
			result += "l"
		} else if c == 'S' || c == 'R' || c == 'Z' {
			result += "L"
		} else {
			result += string(c)
		}
	}
	*words = result
}

func main() {
	// bisa dicoba untuk pengujian test case
	var words string = "Rangga Surya Prayoga"
	fmt.Println("Before",words)
	SlurredTalk(&words)
	fmt.Println("After",words)
}
