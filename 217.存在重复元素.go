/*
 * @lc app=leetcode.cn id=217 lang=golang
 *
 * [217] 存在重复元素
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	Map := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if Map[nums[i]] {
			return true
		} else {
			Map[nums[i]] = true
		}
	}
	return false
}

// @lc code=end
