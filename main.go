package main

import "fmt"

func lengthOfLongestSubstring(str string) int {
	indexMap, leftBound, maxResult := make(map[rune]int), 0, 0
	for index, char := range str {
		if _, existed := indexMap[char]; existed == true && indexMap[char] >= leftBound {
			if index-leftBound > maxResult {
				maxResult = index - leftBound
			}
			leftBound = indexMap[char] + 1
		}
		indexMap[char] = index
	}

	if len(str)-leftBound > maxResult {
		maxResult = len(str) - leftBound
	}
	return maxResult
}

func main() {
	str := "asdabdssuj"
	fmt.Println(lengthOfLongestSubstring(str))
}
