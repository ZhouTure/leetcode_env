/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	lenth := len(digits)

	digits[lenth-1] += 1

	for i := lenth - 1; i >= 0; i-- {
		if i != 0 && digits[i] >= 10 {
			digits[i] -= 10
			digits[i-1] += 1
		}

		if i == 0 && digits[i] >= 10 {
			digits[i] -= 10
			digits = append([]int{1}, digits...)
		}

	}
	return digits
}

// @lc code=end
