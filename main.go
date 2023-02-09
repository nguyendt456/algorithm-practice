package main

import "fmt"

func expandLeftRight(target string, left int, right int, max string) string {
	var substr string
	for left >= 0 && right < len(target) && target[left] == target[right] {
		fmt.Println(left, right)
		substr = target[left : right+1]
		left--
		right++
	}
	if len(max) < len(substr) {
		return substr
	}
	return max
}

func longestPalindrome(str string) string {
	result := ""
	for i := 0; i < len(str)-1; i++ {
		result = expandLeftRight(str, i, i, result)
		result = expandLeftRight(str, i, i+1, result)
	}
	return result
}

func main() {
	str := "alkbb"
	fmt.Println(longestPalindrome(str))
}
