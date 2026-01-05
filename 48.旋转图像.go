/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */

// @lc code=start
func rotate(matrix [][]int) {
	lenx, leny := len(matrix), len(matrix[0])

	// 转置
	for i := range lenx {
		for j := i; j < leny; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 行反转
	for i := range lenx {
		for j := range (leny) / 2 {
			matrix[i][j], matrix[i][leny-1-j] = matrix[i][leny-1-j], matrix[i][j]
		}
	}

}

// @lc code=end

// 1 2 3    1 4 7   7 4 1
// 4 5 6 -> 2 5 8 ->8 5 2
// 7 8 9    3 6 9   9 6 3