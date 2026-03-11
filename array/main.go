package main

import "fmt"

func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	fmt.Println(kidsWithCandies(candies, extraCandies))
}

// 1431. Kids With the Greatest Number of Candies
func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxTerm := candies[0]

	for i := 1; i < len(candies); i++ {
		if candies[i] > maxTerm {
			maxTerm = candies[i]
		}
	}

	result := make([]bool, len(candies))

	for i := 0; i < len(candies); i++ {
		result[i] = candies[i]+extraCandies >= maxTerm
	}

	return result
}
