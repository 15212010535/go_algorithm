package main

import "fmt"

/*
	给定一个二维网格和一个单词，找出该单词是否存在于网格中。单词必须按照字母顺序，通过相邻的单元格内的字母构成，
	其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
*/

// 定义移动方向
var dir = [][]int{
	[]int{-1, 0},
	[]int{0, 1},
	[]int{1, 0},
	[]int{0, -1},
}

func exit(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	// 遍历通过每个board
	for i, v := range board {
		for j := range v {
			if searchWord(board, visited, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func searchWord(board [][]byte, visited [][]bool, word string, index, x, y int) bool {
	// 最后判断
	if index == len(word)-1 {
		return board[x][y] == word[index]
	}
	// 比较
	if board[x][y] == word[index] {
		visited[x][y] = true // 访问过了
		// 访问邻居
		for i := 0; i < 4; i++ {
			nx := x + dir[i][0]
			ny := y + dir[i][1]
			if isInBoard(board, nx, ny) && !visited[nx][ny] && searchWord(board, visited, word, index+1, nx, ny) {
				return true
			}
		}
		// 清空
		visited[x][y] = false
	}
	return false
}

// 判断是否合法
func isInBoard(board [][]byte, x int, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}

func main() {
	fmt.Println(exit([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCCED"))
}
