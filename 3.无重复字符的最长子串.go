/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	right, left := 0, 0
	maxlen := 0
	map_wind := make(map[byte]int, len(s))
	for right < len(s) {
		if _, ok := map_wind[s[right]]; ok {
			// fmt.Println(map_wind[s[right]], right, left)
			if map_wind[s[right]] <= right && map_wind[s[right]] >= left {
				left = map_wind[s[right]] + 1
			}
		}

		map_wind[s[right]] = right
		right++

		if right-left > maxlen {
			maxlen = right - left
		}

		// fmt.Println(right, left, map_wind)
	}

	return maxlen
}

// @lc code=end
