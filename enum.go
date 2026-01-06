package main

import "fmt"

const mx = 3

var divisorNum, divisorSum [mx]int

func initenum() {
	for i := 1; i < mx; i++ {
		for j := i; j < mx; j += i { // 枚举 i 的倍数 j
			divisorNum[j]++    // i 是 j 的因子
			divisorSum[j] += i // 因子和
			fmt.Println(i, j, divisorNum, divisorSum)
		}
	}
}

func main2() {
	initenum()
}
