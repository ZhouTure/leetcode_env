import "math"

/*
 * @lc app=leetcode.cn id=1266 lang=golang
 *
 * [1266] 访问所有点的最小时间
 */

// @lc code=start
func distance(point1 []int, point2 []int) (res float64) {
	return math.Sqrt(math.Pow(float64(point1[0]-point2[0]), 2) + math.Pow(float64(point1[1]-point2[1]), 2))
}

func minTimeToVisitAllPoints(points [][]int) int {
	// step, count := 0, 0
	// short := math.Sqrt(2)
	step := 0

	for i := 0; i < len(points)-1; i++ {
		dis := distance(points[i], points[i+1])
		// 如果是同y轴或者x轴 则直接x或y移动 不需要走 根号二 格
		if points[i][0]-points[i+1][0] == 0 || points[i][1]-points[i+1][1] == 0 {
			step += int(dis)
		} else {
			// xy全不相同 则 Dx Dy的差值是要走的直线步数 例如 3,3 4,5 那要走1补斜线 1步y
			// fmt.Println(dis)
			Dx, Dy := math.Abs(float64(points[i][0]-points[i+1][0])), math.Abs(float64(points[i][1]-points[i+1][1]))
			step += int(math.Abs(Dy - Dx))
			if Dy > Dx {
				step += int(math.Abs(Dx))
			} else {
				step += int(math.Abs(Dy))
			}
		}
	}

	return step
}

// @lc code=end
