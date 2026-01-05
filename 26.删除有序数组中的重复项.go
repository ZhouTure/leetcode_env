/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {

	n := len(nums)
	if n == 0 {
		return 0
	}

	// 初始化慢指针
	slow := 0

	//快指针从索引1开始遍历
	for fast := range n {
		// 如果发现快指针的值和慢指针不同
		if nums[fast] != nums[slow] {
			// 慢指针前进一步
			slow++
			// 将快指针发现的新值覆盖到慢指针位置
			nums[slow] = nums[fast]
		}
	}

	// 返回新数组的长度 慢指针索引 + 1
	return slow + 1

}

// @lc code=end
