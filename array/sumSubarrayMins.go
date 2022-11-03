package main

import "fmt"

/*
	给定一个整数数组 A，找到 min(B) 的总和，其中 B 的范围为 A 的每个（连续）子数组。
	每个子区间内最小值的和
*/

// 方法一： DP + 单调栈
func sumSubarrayMinsOne(A []int) int {
	stack, dp, res, mod := []int{}, make([]int, len(A)+1), 0, 1000000007
	stack = append(stack, -1)

	for i := 0; i < len(A); i++ {
		for stack[len(stack)-1] != -1 && A[i] <= A[stack[len(stack)-1]] { // 栈顶元素不为空，且当前元素小于栈顶元素值
			stack = stack[:len(stack)-1]
		}
		dp[i+1] = (dp[stack[len(stack)-1]+1] + (i-stack[len(stack)-1])*A[i]) % mod
		// 存储区间内最小值的下标
		stack = append(stack, i)
		res += dp[i+1]
		res %= mod
	}
	return res
}

type pair struct {
	val   int
	count int
}

// 方法二：两个单调栈
// 从左往右寻找左子集、从右往左寻找右子集
func sumSubarrayMinsTwo(A []int) int {
	res, n, mod := 0, len(A), 1000000007
	lefts, rights, leftStack, rightStack := make([]int, n), make([]int, n), []*pair{}, []*pair{}
	for i := 0; i < n; i++ {
		count := 1
		for len(leftStack) != 0 && leftStack[len(leftStack)-1].val > A[i] { // 栈顶元素大于当前元素，目的是寻找小的元素
			count += leftStack[len(leftStack)-1].count
			leftStack = leftStack[:len(leftStack)-1]
		}
		leftStack = append(leftStack, &pair{val: A[i], count: count})
		// 存储与当前值左区间范围
		lefts[i] = count
	}

	for i := n - 1; i >= 0; i-- {
		count := 1
		for len(rightStack) != 0 && rightStack[len(rightStack)-1].val >= A[i] {
			count += rightStack[len(rightStack)-1].count
			rightStack = rightStack[:len(rightStack)-1]
		}
		rightStack = append(rightStack, &pair{val: A[i], count: count})
		rights[i] = count
	}

	for i := 0; i < n; i++ {
		res = (res + A[i]*lefts[i]*rights[i]) % mod
	}
	return res
}

// 方法三：暴力法
func sumSubarrayMinsThree(A []int) int {
	res, mod := 0, 1000000007
	for i := 0; i < len(A); i++ {
		var stack []int
		stack = append(stack, A[i])
		for j := i; j < len(A); j++ {
			if stack[len(stack)-1] >= A[j] {
				stack = stack[:len(stack)-1]
				stack = append(stack, A[j])
			}
			res += stack[len(stack)-1]
		}
	}
	return res % mod
}

func main() {
	fmt.Println(sumSubarrayMinsOne([]int{3, 1, 2, 4}))
	fmt.Println(sumSubarrayMinsTwo([]int{3, 1, 2, 4}))
	fmt.Println(sumSubarrayMinsThree([]int{3, 1, 2, 4}))
}
