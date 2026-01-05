/*
 * @lc app=leetcode.cn id=961 lang=golang
 *
 * [961] 在长度 2N 的数组中找出重复 N 次的元素
 */

// @lc code=start
func repeatedNTimes(nums []int) int {
	var ans int
	nm := make(map[int]bool)

	for _, v := range nums {
		if _, ok := nm[v]; ok {
			ans = v
			break
		}
		nm[v] = true
	}
	return ans
}

// @lc code=end
