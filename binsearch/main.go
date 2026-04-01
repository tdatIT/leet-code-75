package main

import "fmt"

func main() {
	n := 0
	fmt.Println(guessNumber(n))
}

func guessNumber(n int) int {
	left := 1
	right := n

	for left <= right {
		mid := (left + right) / 2

		switch guess(mid) {
		case -1:
			right = mid - 1
		case 1:
			left = mid + 1
		case 0:
			return mid
		}

	}

	return -1
}

func guess(n int) int {
	pick := 6

	if n > pick {
		return -1
	}

	if n < pick {
		return 1
	}

	return 0

}
