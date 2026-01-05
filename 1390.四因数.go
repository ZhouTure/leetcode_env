/*
 * @lc app=leetcode.cn id=1390 lang=golang
 *
 * [1390] 四因数
 */

// @lc code=start
func sumFourDivisors(nums []int) int {
	count := 0
	for _, i := range nums {
		temp_map := make(map[int]bool)
		for j := 1; j <= int(i/2); j++ {
			if i%j == 0 {
				temp_map[j] = true
			}
			temp_map[i] = true
		}
		if len(temp_map) == 4 {
			for k, _ := range temp_map {
				count += k
			}
		}
		// fmt.Println(temp_map)
	}
	return count
}

// @lc code=end
