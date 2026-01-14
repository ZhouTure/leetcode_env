/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
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

// @lc code=end
