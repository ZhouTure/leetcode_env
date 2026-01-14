package main

import (
	"fmt"
)

func BinarySearch(target int, nums []int) int {
	right, left := len(nums)-1, 0
	for left <= right {
		mid := int((left + right) / 2)
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		}
		if nums[mid] > target {
			right = mid - 1
		}
	}

	return -1
}

func main1() {
	num := BinarySearch(321443, []int{1, 4, 5, 21, 99, 882, 2213, 321443, 1122333})
	fmt.Println(num)
}
