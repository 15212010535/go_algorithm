package main

import "fmt"

/*
找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。
*/
func combinationSum(k int, n int) [][]int {
	if k == 0 {
		return [][]int{}
	}
	var c []int
	var res [][]int
	findCombinationSum(k, n, 1, c, &res)
	return res
}

func findCombinationSum(k int, target int, index int, c []int, res *[][]int) {
	// 递归的推出条件，当目标值为0
	if target == 0 {
		if len(c) == k {
			b := make([]int, len(c))
			copy(b, c)
			*res = append(*res, b)
		}
		return
	}
	for i := index; i < 10; i++ {
		if target >= i {
			c = append(c, i)
			findCombinationSum(k, target-i, i+1, c, res)
			c = c[:len(c)-1]
		}
	}
}

func main() {
	fmt.Println(combinationSum(3, 5))
}
