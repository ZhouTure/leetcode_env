package main

import (
	"fmt"
)

func maxDotProduct(nums1 []int, nums2 []int) int {

	var num1 []int
	var num2 []int
	maxi, maxj := 0, 0
	res := nums1[0] * nums2[0]
	for i := range len(nums1) {
		for j := 0; j < len(nums2); j++ {
			if res < nums1[i]*nums2[j] && i < maxi && j < maxj {
				res = nums1[i] * nums2[j]
				maxi = i
				maxj = j
				num1 = append(num1, nums1[maxi])
				num2 = append(num2, nums2[maxj])
			}
		}
	}
	fmt.Println(num1, num2)
	//
	result := 0
	for i := range len(num1) {
		result += num1[i] * num2[i]
	}
	return result
}

func main() {
	num := maxDotProduct([]int{2, 1, -2, 5}, []int{3, 0, -6})
	fmt.Println(num)
}
