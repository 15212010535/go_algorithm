package main

import (
	"fmt"
	"math"
	"sort"
)

/*
给定一个数组，要求在这个数组中找出 3 个数之和离 target 最近。
*/
func threeSumClosestOne(nums []int, target int) int {
	n, res, diff := len(nums), 0, math.MaxInt32
	if n > 2 {
		sort.Ints(nums)
		// 固定a
		for i := 0; i < n-2; i++ {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			// 双指针, 寻找b, c
			for j, k := i+1, n-1; j < k; {
				sum := nums[i] + nums[j] + nums[k]
				if abs(sum-target) < diff {
					res, diff = sum, abs(sum-target)
				}
				if sum == target {
					return res
				} else if sum > target {
					k--
				} else {
					j++
				}
			}
		}
	}
	return res
}

// 暴力法, 三重循环
func threeSumClosestTwo(nums []int, target int) int {
	res, diff := 0, math.MaxInt32
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if abs(nums[i]+nums[j]+nums[k]-target) < diff {
					res, diff = nums[i]+nums[j]+nums[k], abs(nums[i]+nums[j]+nums[k]-target)
				}
			}
		}
	}
	return res
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func main() {
	res := threeSumClosestOne([]int{-1, 2, 1, -4}, 1)
	fmt.Println(res)
	res = threeSumClosestTwo([]int{-1, 2, 1, -4}, 1)
	fmt.Println(res)
}
