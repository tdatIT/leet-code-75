package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "the sky is blue"
	fmt.Println(reverseWords(str))
}

// 1768. Merge Strings Alternately
func mergeAlternately(word1 string, word2 string) string {
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

// 345. Reverse Vowels of a String
func reverseVowels(s string) string {

	left := 0
	right := len(s) - 1
	strArr := []byte(s)
	for left < right {
		leftOk := isVowels(strArr[left])
		rightOk := isVowels(strArr[right])

		if !leftOk {
			left++
		}

		if !rightOk {
			right--
		}

		if leftOk && rightOk {
			t := strArr[left]
			strArr[left] = strArr[right]
			strArr[right] = t
			left++
			right--
		}
	}

	return string(strArr)
}

func isVowels(c byte) bool {
	switch c {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}

	return false
}

// 151. Reverse Words in a String
func reverseWords(s string) string {
	var result []string
	i := len(s) - 1

	for i >= 0 {
		if s[i] == ' ' {
			i--
		} else {
			start := i

			for start >= 0 && s[start] != ' ' {
				start--
			}

			result = append(result, s[start+1:i+1])
			i = start
		}
	}

	return strings.Join(result, " ")
}
