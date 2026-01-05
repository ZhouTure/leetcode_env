/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 轮转数组
 */

// @lc code=start
func rotate(nums []int, k int) {

	lenNums := len(nums)

	k %= lenNums

	reverse := func(s []int) {
		n := len(s)
		for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
	}

	if k == 0 || lenNums == 1 {
		nums = nums
	} else {
		// 反转数组
		reverse(nums)

		// 反转前K
		reverse(nums[:k])

		// 反转后N-K
		reverse(nums[k:])
	}

	// fmt.Println(nums)
}

// @lc code=end
