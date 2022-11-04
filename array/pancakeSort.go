package main

import "fmt"

/*
	给定一个数组，要求输出“煎饼排序”的步骤，使得最终数组是从小到大有序的。
	“煎饼排序”，每次排序都反转前 n 个数，n 小于数组的长度。
*/

func pancakeSort(arr []int) (ans []int) {
	for n := len(arr); n > 1; n-- {
		index := 0
		// 寻找最大值下标
		for i, v := range arr[:n] {
			if v > arr[index] {
				index = i
			}
		}
		// 当前下标为最大
		if index == n-1 {
			continue
		}
		// 反转
		for i, m := 0, index; i < (m+1)/2; i++ {
			arr[i], arr[m-1] = arr[m-i], arr[i]
		}
		for i := 0; i < n/2; i++ {
			arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
		}
		// 最大值放在最后
		ans = append(ans, index+1, n)
	}
	return
}

func main() {
	fmt.Println(pancakeSort([]int{3, 2, 4, 1}))
}
