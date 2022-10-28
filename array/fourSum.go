package main

import (
	"fmt"
	"sort"
)

/*
给定一个数组，要求在这个数组中找出 4 个数之和为 0 的所有组合。
*/
func fourSumOne(nums []int, target int) (quadruplets [][]int) {
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			for left, right := j+1, n-1; left < right; {
				if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
					quadruplets = append(quadruplets, []int{nums[i], nums[j], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {
					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return
}

// 递归求解
func fourSumTwo(nums []int, target int) [][]int {
	res, cur := make([][]int, 0), make([]int, 0)
	sort.Ints(nums)
	kSum(nums, 0, len(nums)-1, target, 4, cur, &res)
	return res
}

func kSum(nums []int, left int, right int, target int, k int, cur []int, res *[][]int) {
	// 左右之间的区域小于k个,或者k小于2，或者目标数小于当前左边 * k
	if right-left+1 < k || k < 2 || target < nums[left]*k || target > nums[right]*k {
		return
	}
	if k == 2 {
		twoSum(nums, left, right, target, cur, res)
	} else {
		for i := left; i < len(nums); i++ {
			if i == left || (i > left && nums[i-1] != nums[i]) {
				next := make([]int, len(cur))
				copy(next, cur)
				next = append(next, nums[i])
				kSum(nums, i+1, len(nums)-1, target-nums[i], k-1, next, res)
			}
		}
	}
}

// 双指针找到两个数之和
func twoSum(nums []int, left int, right int, target int, cur []int, res *[][]int) {
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			cur = append(cur, nums[left], nums[right])
			temp := make([]int, len(cur))
			copy(temp, cur)
			*res = append(*res, temp)
			// 重置
			cur = cur[:len(cur)-2]
			left++
			right--
			for left < right && nums[left] == nums[left-1] {
				left++
			}
			for left < right && nums[right] == nums[right+1] {
				right--
			}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
}

// 使用map
func fourSumThree(nums []int, target int) [][]int {
	var res [][]int
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
		if (uniqNums[i]*4 == target) && counter[uniqNums[i]] >= 4 {
			res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i], uniqNums[i]})
		}
		for j := i + 1; j < len(uniqNums); j++ {
			if (uniqNums[i]+uniqNums[j]*3 == target) && counter[uniqNums[j]] >= 3 {
				res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j], uniqNums[j]})
			}
			if (uniqNums[j]+uniqNums[i]*3 == target) && counter[uniqNums[i]] >= 3 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i], uniqNums[j]})
			}
			if (uniqNums[i]*2+uniqNums[j]*2 == target) && counter[uniqNums[i]] >= 2 && counter[uniqNums[j]] >= 2 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j], uniqNums[j]})
			}
			for k := j + 1; k < len(uniqNums); k++ {
				if (uniqNums[i]*2+uniqNums[j]+uniqNums[k] == target) && counter[uniqNums[i]] >= 2 {
					res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j], uniqNums[k]})
				}
				if (uniqNums[i]+uniqNums[j]*2+uniqNums[k] == target) && counter[uniqNums[j]] >= 2 {
					res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j], uniqNums[k]})
				}
				if (uniqNums[i]+uniqNums[j]+uniqNums[k]*2 == target) && counter[uniqNums[k]] >= 2 {
					res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[k], uniqNums[k]})
				}
				c := target - uniqNums[i] - uniqNums[j] - uniqNums[k]
				if c > uniqNums[k] && counter[k] >= 1 {
					res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[k], c})
				}
			}
		}
	}
	return res
}

func main() {
	res1 := fourSumOne([]int{1, 0, -1, 0, -2, 2}, 0)
	fmt.Println(res1)
	res2 := fourSumTwo([]int{1, 0, -1, 0, -2, 2}, 0)
	fmt.Println(res2)
	res3 := fourSumThree([]int{1, 0, -1, 0, -2, 2}, 0)
	fmt.Println(res3)

}
