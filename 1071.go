package main

import (
	"fmt"
)

func gcdOfStrings(str1 string, str2 string) string {
	if len(str1) < len(str2) {
		str1, str2 = str2, str1
	}
	
	if str1 + str2 != str2 + str1 {
		return ""
	}
	return str1[:gcd(len(str1),len(str2))]
}

func gcd(a, b int) int {
	for a != 0 && b != 0 {
		if a > b {
			a = a % b
		} else {
			b = b % a
		}
	}
	return a + b
}

func main() {
	fmt.Println(gcdOfStrings("ABCABC","ABC"))
}