package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total_pos := 0
	mid := (len(nums1) + len(nums2)) / 2
	var value, prev int
	var result float64
	for pos1, pos2 := 0, 0; pos1 <= len(nums1)-1 || pos2 <= len(nums2)-1; {
		if pos1 > len(nums1)-1 && pos2 <= len(nums2)-1 {
			value = nums2[pos2]
			pos2++
		}
		if pos2 > len(nums2)-1 && pos1 <= len(nums1)-1 {
			value = nums1[pos1]
			pos1++
		}

		if pos1 <= len(nums1)-1 && pos2 <= len(nums2)-1 {
			if nums1[pos1] < nums2[pos2] {
				value = nums1[pos1]
				pos1++
			} else {
				value = nums2[pos2]
				pos2++
			}
		}
		if total_pos >= mid {
			if (len(nums1)+len(nums2))%2 == 0 {
				result = float64(value+prev) / 2
				break
			} else {
				result = float64(value)
			}
		}
		prev = value
		total_pos++
	}
	return result
}

func main() {
	arr1 := []int{1, 2, 6, 7}
	arr2 := []int{3, 8}

	fmt.Println(findMedianSortedArrays(arr1, arr2))
}
