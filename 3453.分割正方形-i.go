/*
 * @lc app=leetcode.cn id=3453 lang=golang
 *
 * [3453] 分割正方形 I
 */

// @lc code=start
func getAreaBelow(m float64, squares [][]int) float64 {
	res := 0.0
	for _, s := range squares {
		y, side := float64(s[1]), float64(s[2])
		if y >= m {
			continue
		} else if y+side <= m {

			res += side * side
		} else {
			res += (m - y) * side
		}
	}
	return res
}

func separateSquares(squares [][]int) float64 {
	var totalArea float64
	minY, maxY := 2e9, 0.0

	for _, s := range squares {
		side := float64(s[2])
		totalArea += side * side
		y := float64(s[1])
		if y < minY {
			minY = y
		}
		if y+side > maxY {
			maxY = y + side
		}
	}

	target := totalArea / 2.0
	l, r := minY, maxY

	for i := 0; i < 100; i++ {
		mid := l + (r-l)/2.0
		if getAreaBelow(mid, squares) < target {
			l = mid
		} else {
			r = mid
		}
	}

	return l
}

// @lc code=end
