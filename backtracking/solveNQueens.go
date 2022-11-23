package main

import (
	"fmt"
	"math/bits"
)

var solutions [][]string

// 方法一： 基于集合的回溯
func solveNQueensOne(n int) [][]string {
	solutions = [][]string{}
	// 每一行第几个位置为Q
	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	columns := map[int]bool{}
	diagonals1, diagonals2 := map[int]bool{}, map[int]bool{}
	backtrackOne(queens, n, 0, columns, diagonals1, diagonals2)
	return solutions
}

func backtrackOne(queens []int, n int, row int, columns map[int]bool, diagonals1 map[int]bool, diagonals2 map[int]bool) {
	// 构造
	if row == n {
		board := generateBoard(queens, n)
		solutions = append(solutions, board)
		return
	}
	for i := 0; i < n; i++ {
		// 行不可以重复
		if columns[i] {
			continue
		}
		// 左对角线不可以重复
		diagonal1 := row - i
		if diagonals1[diagonal1] {
			continue
		}
		// 右对角线不可以放置
		diagonal2 := row + i
		if diagonals2[diagonal2] {
			continue
		}
		queens[row] = i
		columns[i] = true
		diagonals1[diagonal1], diagonals2[diagonal2] = true, true
		backtrackOne(queens, n, row+1, columns, diagonals1, diagonals2)
		queens[row] = -1
		delete(columns, i)
		delete(diagonals1, diagonal1)
		delete(diagonals2, diagonal2)
	}
}

func generateBoard(queens []int, n int) []string {
	var board []string
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		// 某一行中那个位置防止Q
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return board
}

// 方法二：基于位运算的回溯
func solveNQueensTwo(n int) [][]string {
	solutions = [][]string{}
	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	solve(queens, n, 0, 0, 0, 0)
	return solutions
}

func solve(queens []int, n int, row int, columns int, diagonals1 int, diagonals2 int) {
	if row == n {
		board := generateBoard(queens, n)
		solutions = append(solutions, board)
		return
	}
	availablePositions := ((1 << n) - 1) & (^(columns | diagonals1 | diagonals2))
	for availablePositions != 0 {
		position := availablePositions & (-availablePositions)
		availablePositions = availablePositions & (availablePositions - 1)
		columns := bits.OnesCount(uint(position - 1))
		queens[row] = columns
		solve(queens, n, row+1, columns|position, (diagonals1|position)>>1, (diagonals2|position)<<1)
		queens[row] = -1
	}
}

func main() {
	fmt.Println(solveNQueensOne(4))
	fmt.Println(solveNQueensTwo(4))
}
