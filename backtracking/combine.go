package main

/*
组合
*/
func combineOne(n int, k int) (ans [][]int) {
	var temp []int
	var dfs func(int)
	dfs = func(cur int) {
		if len(temp)+(n-cur+1) < k {
			return
		}
		if len(temp) == k {
			comb := make([]int, k)
			copy(comb, temp)
			ans = append(ans, comb)
			return
		}
		// 选择当前位置
		temp = append(temp, cur)
		dfs(cur + 1)
		temp = temp[:len(temp)-1]
		// 不选当前位置
		dfs(cur + 1)
	}
	dfs(1)
	return
}

func combineTwo(n int, k int) (ans [][]int) {
	var temp []int
	for i := 1; i <= k; i++ {
		temp = append(temp, i)
	}
	temp = append(temp, n+1)
	for j := 0; j < k; {
		comb := make([]int, k)
		copy(comb, temp[:k])
		ans = append(ans, comb)
		for j = 0; j < k && temp[j]+1 == temp[j+1]; j++ {
			temp[j] = j + 1
		}
		temp[j]++
	}
	return
}
