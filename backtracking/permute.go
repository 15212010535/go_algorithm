package main

import "fmt"

/*
	全排列
*/

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	used, p, res := make([]bool, len(nums)), []int{}, [][]int{}
	generatePermutation(nums, 0, p, &res, &used)
	return res
}

func generatePermutation(nums []int, index int, p []int, res *[][]int, used *[]bool) {
	// 出口
	if index == len(nums) {
		temp := make([]int, len(p))
		copy(temp, p)
		*res = append(*res, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		// 是否出现
		if !(*used)[i] {
			// 添加
			(*used)[i] = true
			p = append(p, nums[i])
			generatePermutation(nums, index+1, p, res, used)
			// 清除
			p = p[:len(p)-1]
			(*used)[i] = false
		}
	}
	return
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
