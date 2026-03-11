package main

func main() {
	str1 := "ABABAB"
	str2 := "ABAB"
	result := gcdOfStrings(str1, str2)
	println(result) // Output: "apbqcr"
}

// O(n)
func mergeAlternately(word1 string, word2 string) string {
	result := ""
	maxLen := len(word1)

	if maxLen < len(word2) {
		maxLen = len(word2)
	}

	i := 0
	for i < maxLen {
		if i < len(word1) {
			result = result + string(word1[i])
		}
		if i < len(word2) {
			result = result + string(word2[i])
		}

		i++
	}

	return result
}

func mergeAlternatelyBetter(word1 string, word2 string) string {
	result := ""
	i, j := 0, 0

	for i < len(word1) && j < len(word2) {
		result += string(word1[i]) + string(word2[j])
		i++
		j++
	}

	for i < len(word1) {
		result += string(word1[i])
		i++
	}

	for j < len(word2) {
		result += string(word2[j])
		j++
	}

	return result
}

// 1071. Greatest Common Divisor of Strings
func gcdOfStrings(str1 string, str2 string) string {
	if (str1 + str2) != (str2 + str1) {
		return ""
	}

	gcdStrLen := gcd(len(str1), len(str2))

	return str1[:gcdStrLen]
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
