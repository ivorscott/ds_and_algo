package main

import "fmt"

func main() {
	fmt.Println(isPalindrome("lol"))
	fmt.Println(isPalindrome("scott"))
}

func isPalindrome(str string) bool {
	return str == reverseStr(str)
}

func reverseStr(str string) string {
	// when iterating over characters of a string, first convert it to runes
	runes := []rune(str)

	// start at both ends and swap the positions
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[j], runes[i] = runes[i], runes[j]
	}
	// convert back to a string
	return string(runes)
}
