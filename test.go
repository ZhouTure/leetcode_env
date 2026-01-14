package main

import (
	"fmt"
)

func mySqrt(x int) int {
	// 二分查找 l-r <= 1 返回l
	l, r := 0, x
	ans := 0
	for l-r < 1 {
		mid := int((l + r) / 2)
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}

func main() {
	num := mySqrt(2)
	fmt.Println("res: ", num)
}
