import "sort"

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {

	type KV struct {
		k string
		v int
	}

	if len(strs) == 0 {
		return ""
	}

	sort.Strings(strs) // 按字典序排序

	first := strs[0]
	last := strs[len(strs)-1]

	limit := len(first)
	if len(last) < limit {
		limit = len(last)
	}

	i := 0
	for i < limit && first[i] == last[i] {
		i++
	}
	return first[:i]
}

// @lc code=end
