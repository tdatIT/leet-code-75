package main

import (
	"fmt"
	"math"
)

func main() {
	flowed := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(flowed))
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

// 605. Can Place Flowers
func canPlaceFlowers(flowerbed []int, n int) bool {
	check := 0

	if n == 0 {
		return true
	}

	for i := 0; i < len(flowerbed); i++ {
		leftEmpty := i == 0 || flowerbed[i-1] == 0
		rightEmpty := i == len(flowerbed)-1 || flowerbed[i+1] == 0

		if flowerbed[i] == 0 && leftEmpty && rightEmpty {
			flowerbed[i] = 1
			check++
		}

		if check == n {
			return true
		}

	}

	return false
}

// 238. Product of Array Except Self
func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	prefix := 1

	for i := 0; i < len(nums); i++ {
		result[i] = prefix
		prefix = prefix * nums[i]
	}

	suff := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] = result[i] * suff
		suff = suff * nums[i]
	}

	return result
}

// 334. Increasing Triplet Subsequence
func increasingTriplet(nums []int) bool {
	first, second := math.MaxInt, math.MaxInt

	for i := 0; i < len(nums); i++ {
		if nums[i] <= first {
			first = nums[i]
		} else if nums[i] <= second {
			second = nums[i]
		}

		if nums[i] > second {
			return true
		}
	}

	return false
}
