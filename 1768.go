package main

import (
	"fmt"
	"math"
	"strings"
)

func mergeAlternately(word1 string, word2 string) string {
	result := ""
	arr1 := strings.Split(word1, "")
	arr2 := strings.Split(word2, "")
	for i := 0; i < int(math.Min(float64(len(arr1)), float64(len(arr2)))); i++ {
		result += arr1[i] + arr2[i]
	}

	if len(arr1) > len(arr2) {
		for i := len(arr2); i < len(arr1); i++ {
			result += arr1[i]
		}
	} else {
		for i := len(arr1); i < len(arr2); i++ {
			result += arr2[i]
		}
	}
	return result
}

func main() {
	fmt.Println(mergeAlternately("sadsadasdas", "dsa"))
}
