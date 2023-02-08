package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hash_map := make(map[int]int)
	for index, element := range nums {
		if value, existed := hash_map[target-element]; existed == true {
			return []int{value, index}
		}
		hash_map[element] = index
	}
	return nil
}

func main() {
	num := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum(num, target))
}
