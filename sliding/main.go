package main

import "fmt"

func main() {
	nums := []int{0, 1, 1, 3, 3}
	fmt.Println(findMaxAverage(nums, 4))
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
