package main

import (
	"fmt"
	"sort"
)

/*
*
给定一个数组，要求在这个数组中找出 3 个数之和为 0 的所有组合。
*/
// 解法一 最优解, 双指针 + 排序
// 整体思想是： 固定中间的数，左右两边往中间遍历，直到找到满足要求的条件
func threeSumOne(nums []int) [][]int {
	// 数组排序
	sort.Ints(nums)
	// 定义遍历
	result, start, end, index, addNum, length := make([][]int, 0), 0, 0, 0, 0, len(nums)
	// 循环, 固定index
	for index = 1; index < length-1; index++ {
		// 定义初始化
		start, end = 0, length-1
		// 跳过相同的数
		if index > 1 && nums[index] == nums[index-1] {
			start = index - 1
		}

		for start < index && end > index {
			// 跳过相同的数
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			// 跳过相同的数
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			// 求和，判断
			addNum = nums[start] + nums[end] + nums[index]
			if addNum == 0 {
				result = append(result, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			} else if addNum > 0 { // 相加之所以大于0，是因为最后一个数太大了
				end--
			} else { //  同理
				start++
			}
		}
	}
	return result
}

func threeSumTwo(nums []int) [][]int {
	var result [][]int
	// 先将数组存入map集合中,key为当前值.value为出现次数
	counter := map[int]int{}
	for _, value := range nums {
		counter[value]++
	}
	var uniqNums []int
	for key := range counter {
		uniqNums = append(uniqNums, key)
	}
	sort.Ints(uniqNums)

	for i := 0; i < len(uniqNums); i++ {
		// 数组中有三个0的情况
		if (uniqNums[i]*3 == 0) && counter[uniqNums[i]] >= 3 {
			result = append(result, []int{uniqNums[i], uniqNums[i], uniqNums[i]})
		}

		for j := i + 1; j < len(uniqNums); j++ {
			// 数组中有至少两个数一样,且相加为0
			if (uniqNums[i]*2+uniqNums[j] == 0) && counter[uniqNums[i]] >= 2 {
				result = append(result, []int{uniqNums[i], uniqNums[i], uniqNums[j]})
			}
			if (uniqNums[j]*2+uniqNums[i] == 0) && counter[uniqNums[j]] >= 2 {
				result = append(result, []int{uniqNums[i], uniqNums[j], uniqNums[j]})
			}
			// 剩余一个数的值
			c := 0 - uniqNums[i] - uniqNums[j]
			// 通过map取值,存在则符合，且c的值应该大于中间的数这样才有可以为0
			if c > uniqNums[j] && counter[c] > 0 {
				result = append(result, []int{uniqNums[i], uniqNums[j], c})
			}
		}
	}
	return result
}

func main() {
	result := threeSumOne([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(result)
	result = threeSumTwo([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(result)
}
