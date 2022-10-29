package main

import "fmt"

/*
	给出每个直方图的高度，要求在这些直方图之中找到面积最大的矩形，输出矩形的面积
*/

func largestRectangleArea(heights []int) int {
	maxArea := 0
	// 在开始和介绍添加一个标志,防止越界
	n := len(heights) + 2
	getHeight := func(i int) int {
		if i == 0 || n-1 == i {
			return 0
		}
		return heights[i-1]
	}
	// 构造栈，存放位置下标
	st := make([]int, 0, n/2)
	for i := 0; i < n; i++ {
		// 栈顶元素前一个的高度高于当前下标前一个高度
		for len(st) > 0 && getHeight(st[len(st)-1]) > getHeight(i) {
			idx := st[len(st)-1]
			st = st[:len(st)-1]
			maxArea = max(maxArea, getHeight(idx)*(i-st[len(st)-1]-1))
		}
		st = append(st, i)
	}
	return maxArea
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 5, 3}))
}
