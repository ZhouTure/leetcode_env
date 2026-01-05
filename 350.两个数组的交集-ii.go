/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {

	map1 := make(map[int]int)
	for _, n := range nums1 {
		map1[n]++
	}

	var res []int

	for _, m := range nums2 {
		if map1[m] > 0 {
			res = append(res, m)
			map1[m]--
		}
	}

	return res

}

// @lc code=end
