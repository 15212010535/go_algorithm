package main

// 方法一：基于集合的回溯
func totalNQueensOne(n int) (ans int) {
	columns := make([]bool, n)
	diagonals1 := make([]bool, 2*n-1)
	diagonals2 := make([]bool, 2*n-1)
	var backtrack func(int)
	backtrack = func(row int) {
		if row == n {
			ans++
			return
		}
		for col, hasQueen := range columns {
			d1, d2 := row+n-1-col, row+col // 两个对角线位置
			if hasQueen || diagonals1[d1] || diagonals2[d2] {
				continue
			}
			columns[col] = true
			diagonals1[d1] = true
			diagonals2[d2] = true
			backtrack(row + 1)
			columns[col] = false
			diagonals1[d1] = false
			diagonals2[d2] = false
		}
	}
	backtrack(0)
	return
}

// 方法二位运算
func totalNQueens(n int) (ans int) {
	var solve func(row, columns, diagonals1, diagonals2 int)
	solve = func(row, columns, diagonals1, diagonals2 int) {
		if row == n {
			ans++
			return
		}
		availablePositions := (1<<n - 1) &^ (columns | diagonals1 | diagonals2)
		for availablePositions > 0 {
			position := availablePositions & -availablePositions
			solve(row+1, columns|position, (diagonals1|position)<<1, (diagonals2|position)>>1)
			availablePositions &^= position // 移除该比特位
		}
	}
	solve(0, 0, 0, 0)
	return
}
