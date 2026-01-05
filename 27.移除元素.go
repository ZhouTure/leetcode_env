/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {

	n := len(nums)
	if n == 0 {
		return 0
	}
	slow := 0
	for i := range n {
		if nums[i] != val {
			nums[slow] = nums[i]
			slow++
		}
	}

	return slow
}

// @lc code=end
