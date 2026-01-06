package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type iv struct {
	cell int
	val  int
}

func maxLevelSum(nums [][]any) int {
	//左节点2n + 1 右节点2n + 2
	//i != 0  n+1是层数
	// (x - 1)% 2 == 0 左节点  x % 2 == 0 右节点
	// 则 (x - 1) /2 层数  (x - 2)/2层数
	// map or slice
	// return x or min x
	// 先序遍历得到数组
	// var nums [][]any

	// var bfs func(node *TreeNode)
	// bfs = func(node *TreeNode) {
	// 	if node == nil {
	// 		return
	// 	}

	// 	q := []*TreeNode{node}

	// 	for len(q) > 0 {
	// 		sz := len(q) // 本层节点数
	// 		level := make([]any, 0, sz)

	// 		for i := 0; i < sz; i++ {
	// 			node := q[0]
	// 			q = q[1:]

	// 			if node == nil {
	// 				level = append(level, nil)
	// 				// nil 不再扩展孩子，否则会无限增大
	// 				continue
	// 			}

	// 			level = append(level, node.Val)
	// 			q = append(q, node.Left)
	// 			q = append(q, node.Right)
	// 		}

	// 		nums = append(nums, level) // 一层结束
	// 	}

	// 	// 去掉末尾多余的“全 nil 层”（LeetCode 会省略）
	// 	for len(nums) > 0 {
	// 		last := nums[len(nums)-1]
	// 		allNil := true
	// 		for _, v := range last {
	// 			if v != nil {
	// 				allNil = false
	// 				break
	// 			}
	// 		}
	// 		if !allNil {
	// 			break
	// 		}
	// 		nums = nums[:len(nums)-1]
	// 	}
	// }

	// bfs(root)

	fmt.Println(nums)

	queue := make([]iv, 0, len(nums))

	sum := 0
	for i, v := range nums {
		sum = 0
		for _, val := range v {
			if val == nil {
				continue
			} else {
				sum += val.(int)
			}
		}
		queue = append(queue, iv{cell: i, val: sum})
	}
	fmt.Println(queue)

	sort.Slice(queue, func(i, j int) bool {
		if queue[i].val != queue[j].val {
			return queue[i].val > queue[j].val
		}
		return queue[i].cell < queue[j].cell
	})

	fmt.Println(queue)

	return queue[0].cell
}

func main() {
	num := maxLevelSum([][]any{{1}, {7, 0}, {7, -8, nil, nil}})
	fmt.Println(num)
}
