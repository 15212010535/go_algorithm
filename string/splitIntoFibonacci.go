package main

import (
	"fmt"
	"math"
)

/*
	将数组拆分成斐波那契序列
*/

func splitIntoFibonacci(num string) (F []int) {
	n := len(num)
	var backtrack func(index, sum, prev int) bool
	backtrack = func(index, sum, prev int) bool {
		if index == n {
			return len(F) >= 3
		}
		cur := 0
		// 循环遍历
		for i := index; i < n; i++ {
			if i > index && num[index] == '0' {
				break
			}
			cur = cur*10 + int(num[i]-'0')
			if cur > math.MaxInt32 {
				break
			}

			if len(F) >= 2 {
				if cur < sum {
					continue
				}
				if cur > sum {
					break
				}
			}

			F = append(F, cur)
			if backtrack(i+1, prev+cur, cur) {
				return true
			}
			F = F[:len(F)-1]
		}
		return false
	}
	backtrack(0, 0, 0)
	return
}

func main() {
	fmt.Println(splitIntoFibonacci("1101111"))
}
