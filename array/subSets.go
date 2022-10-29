package main

import (
	"fmt"
	"sort"
)

/*
	给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。说明：解集不能包含重复的子集。
*/

// 解法一：使用递归
// 先产生一个数的子集，再产生两个数的子集，依次
func subsetsOne(nums []int) [][]int {
	var res [][]int
	var c []int
	for k := 0; k <= len(nums); k++ {
		generateSubsets(nums, k, 0, c, &res)
	}
	return res
}

func generateSubsets(nums []int, k int, start int, c []int, res *[][]int) {
	if len(c) == k {
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}
	// 核心逻辑,都可以len(nums)-(k-len(c))+1，加快作用
	//for i := start; i < len(nums); i++ {
	for i := start; i < len(nums)-(k-len(c))+1; i++ {
		c = append(c, nums[i])
		// 一直产生直到满足要求
		generateSubsets(nums, k, i+1, c, res)
		c = c[:len(c)-1] // 重要一部，保留前len(c)-1个数
	}
}

// 追加
func subsetsTwo(nums []int) [][]int {
	res := make([][]int, 1)
	sort.Ints(nums)
	for i := range nums {
		for _, org := range res { // 初始化需要一个空数组
			// 将res中数组拷贝
			clone := make([]int, len(org), len(org)+1)
			copy(clone, org)
			// 添加一个新的元素
			clone = append(clone, nums[i])
			// 在添加进结果中
			res = append(res, clone)
		}
	}
	return res
}

// 位运算的方法
// 三个数对应二进制，即111，当为1的时候选择。
func subsetsThree(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	var res [][]int
	// 计算一共有多少种可能
	sum := 1 << uint(len(nums))
	for i := 0; i < sum; i++ {
		var stack []int
		tmp := i
		for j := len(nums) - 1; j >= 0; j-- {
			// 判断能不能入栈，按位与为1进栈，
			if tmp&1 == 1 {
				stack = append([]int{nums[j]}, stack...)
			}
			tmp >>= 1
		}
		res = append(res, stack)
	}
	return res
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsetsOne(nums))
	fmt.Println(subsetsTwo(nums))
	fmt.Println(subsetsThree(nums))
}
