package main

import (
	"fmt"
	"sort"
)

/*
	给出 n + 1 个数，这些数是在 1-n 中取值的，同一个数字可以出现多次。
	要求找出这些数中重复的数字。时间复杂度最好低于 O(n^2)，空间复杂度为 O(1)。
*/

// 方法一： 快慢指针
func findDuplicateOne(nums []int) int {
	slow := nums[0]
	fast := nums[nums[0]]
	// 找到不同下标取到相同的值
	for fast != slow {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	// 找到重复的值
	walker := 0
	for walker != slow {
		walker = nums[walker]
		slow = nums[slow]
	}
	return walker
}

// 方法二：二分查找
// 二分不是二分数组，而是二分范围0 ~ n-1
func findDuplicateTwo(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid, count := low+(high-low)>>1, 0
		for _, num := range nums {
			if num <= mid {
				count++
			}
		}
		if count > mid {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

// 解法三，排序之后当前值与下标+1之间有差距，则该元素有问题
func findDuplicateThree(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	diff := -1
	for i := 0; i < len(nums); i++ {
		if nums[i]-i-1 >= diff {
			diff = nums[i] - i - 1
		} else {
			return nums[i]
		}
	}
	return 0
}
func main() {
	fmt.Println(findDuplicateOne([]int{1, 3, 4, 2, 4}))
	fmt.Println(findDuplicateTwo([]int{1, 3, 4, 2, 4}))
	fmt.Println(findDuplicateThree([]int{1, 3, 4, 2, 4}))
}
