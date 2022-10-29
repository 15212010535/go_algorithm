package main

import (
	"fmt"
	"sort"
)

/*
	给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
	说明：解集不能包含重复的子集。
*/

func subsetsWithDup(nums []int) [][]int {
	var c []int
	var res [][]int
	sort.Ints(nums)
	for k := 0; k <= len(nums); k++ {
		generateSubsetsWithDup(nums, k, 0, c, &res)
	}
	return res
}

func generateSubsetsWithDup(nums []int, k int, start int, c []int, res *[][]int) {
	if len(c) == k {
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}
	for i := start; i < len(nums)-(k-len(c))+1; i++ {
		fmt.Printf("i= %v start = %v c = %v\n", i, start, c)
		if i > start && nums[i] == nums[i-1] { // 去重复
			continue
		}
		c = append(c, nums[i])
		generateSubsetsWithDup(nums, k, i+1, c, res)
		c = c[:len(c)-1]
	}
	return
}

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(subsetsWithDup(nums))
}
