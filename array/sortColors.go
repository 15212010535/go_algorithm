package main

import "fmt"

/*
抽象题意其实就是排序。这题可以用快排一次通过。
*/
func sortColors(nums []int) {
	zero, one := 0, 0
	for i, n := range nums {
		// 每次赋值为2，确保有一个数
		nums[i] = 2
		// 当小于等于1时，对前面的2进行覆盖
		if n <= 1 {
			nums[one] = 1
			one++ // 找到所属位置
		}
		// 同理
		if n == 0 {
			nums[zero] = 0
			zero++
		}
	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}
