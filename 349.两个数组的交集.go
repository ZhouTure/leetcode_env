/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	// struct{}占位 内存损耗小
	set := make(map[int]struct{}, len(nums1))
	for _, x := range nums1 {
		set[x] = struct{}{}
	}

	anset := make(map[int]struct{})
	for _, y := range nums2 {
		if _, ok := set[y]; ok {
			anset[y] = struct{}{}
		}
	}

	ans := make([]int, 0, len(anset))
	for k := range anset {
		ans = append(ans, k)
	}

	return ans

}

// @lc code=end
