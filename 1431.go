package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandies := 0
    for i := 0; i < len(candies); i++ {
		if candies[i] > maxCandies {
			maxCandies = candies[i]
		}
	}
	
	result := make([]bool, len(candies))
	for i, candy := range candies {
		result[i] = candy + extraCandies >= maxCandies
	}

	return result
}

func main() {
	candies := []int{2,3,5,1,3}
	extraCandies := 1
	fmt.Println(kidsWithCandies(candies,extraCandies))
}