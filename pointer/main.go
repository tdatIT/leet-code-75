package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(maxOperations(nums, 5))
}

// 283. Move Zeroes
func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			t := nums[j]
			nums[j] = nums[i]
			nums[i] = t
			j++
		}
	}

	fmt.Println(nums)
}

// 392. Is Subsequence
func isSubsequence(s string, t string) bool {
	i, j := 0, 0

	if s == "" {
		return true
	}

	for i < len(t) {
		if byte(t[i]) == s[j] {
			j++
		}
		i++

		if j == len(s) {
			return true
		}
	}

	return false
}

// 11. Container With Most Water
func maxArea(height []int) int {
	max := 0
	l, r := 0, len(height)-1

	for l < r {
		minHeight, lower := min(height[l], height[r])
		area := minHeight * (r - l)

		if max < area {
			max = area
		}

		if lower == 0 {
			l++
		} else {
			r--
		}
	}

	return max
}

func min(a, b int) (int, int) {
	if a < b {
		return a, 0
	}

	return b, 1
}

// 1679. Max Number of K-Sum Pairs
func maxOperations(nums []int, k int) int {
	slices.Sort(nums)

	count := 0
	l, r := 0, len(nums)-1

	for l < r {
		sum := nums[l] + nums[r]
		if sum == k {
			count++
			l++
			r--
		} else if sum > k {
			r--
		} else {
			l++
		}
	}

	return count
}
