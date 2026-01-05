package main

import (
	"fmt"
)

func quicksort(nums []int) {

	// 长度小于一返回
	if len(nums) <= 1 {
		return
	}

	right, left := len(nums)-1, 0
	base := nums[left]
	for left < right {
		for left < right && nums[right] > base {
			right--
		}

		if left < right {
			nums[left] = nums[right]
			left++
		}

		for left < right && nums[left] < base {
			left++
		}

		if left < right {
			nums[right] = nums[left]
			right--
		}
	}
	nums[left] = base
	// fmt.Println(nums)
	quicksort(nums[:left])
	quicksort(nums[left+1:])
}

func main1() {
	nums := []int{4, 5, 7, 6, 3, 1, 2}
	// 3 4 2 5 1 6
	// 1 4 2 5 1 6
	quicksort(nums)
	fmt.Println(nums)
}
