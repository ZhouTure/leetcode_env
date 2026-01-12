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
			// res = append(res, "")
			left = map_wind[s[right]]
			// fmt.Println(map_wind)
			if len(map_wind) > maxlen {
				maxlen = len(map_wind)
				// fmt.Println(maxlen)
			}

			map_wind = make(map[byte]int, len(s))
		}
		map_wind[s[right]] = right
		right++
		left = left
	}

	return maxlen
}

// @lc code=end
