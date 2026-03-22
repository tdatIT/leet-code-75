package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	fmt.Println(longestOnes(nums, 2))
}

// 643. Maximum Average Subarray I
func findMaxAverage(nums []int, k int) float64 {
	sum := 0

	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	max := sum
	fmt.Println(len(nums) - k)

	for j := 1; j <= len(nums)-k; j++ {
		sum = sum - nums[j-1] + nums[j+k-1]

		if max < sum {
			max = sum
		}
	}

	return float64(max) / float64(k)
}

//1456. Maximum Number of Vowels in a Substring of Given Length
func maxVowels(s string, k int) int {

	count, max := 0, 0

	for i := 0; i < k; i++ {
		if checkVowels(s[i]) {
			count++
		}
	}

	max = count

	for i := k; i < len(s); i++ {
		if checkVowels(s[i]) {
			count++
		}

		if checkVowels(s[i-k]) {
			count--
		}

		if max < count {
			max = count
		}

	}

	return max
}

func checkVowels(c byte) bool {
	switch c {
	case 'a', 'u', 'i', 'o', 'e':
		return true
	}
	return false
}

// 1004. Max Consecutive Ones III
func longestOnes(nums []int, k int) int {
	l, r, zeroCount, result := 0, 0, 0, 0

	for r < len(nums) {
		if nums[r] == 0 {
			zeroCount++
		}

		for zeroCount > k {
			if nums[l] == 0 {
				zeroCount--
			}

			l++
		}

		result = max(result, r-l+1)
		r++
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
