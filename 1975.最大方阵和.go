/*
 * @lc app=leetcode.cn id=1975 lang=golang
 *
 * [1975] 最大方阵和
 */

// @lc code=start
func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func maxMatrixSum(matrix [][]int) int64 {

	// 奇数个负数 后矩阵一定存在负数 偶数个负数 后矩阵不存在负数
	// 找到绝对值最大的负数 对其相邻的绝对值最小的数取负值

	// 奇数个找到最小 然后取负 其他全正求和
	// 复数个直接取正求和

	count, sum, min := 0, 0, absInt(matrix[0][0])
	for i := range len(matrix) {
		for j := range len(matrix) {
			if matrix[i][j] < 0 {
				count++
			}
			if min > absInt(matrix[i][j]) {
				min = absInt(matrix[i][j])
			}
			sum += absInt(matrix[i][j])
		}
	}

	if count%2 == 0 {
		return int64(sum)
	}
	return int64(sum - 2*min)
}

// @lc code=end
