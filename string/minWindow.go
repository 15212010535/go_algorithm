package main

import "fmt"

/*
给定一个源字符串 s，再给一个字符串 T，要求在源字符串中找到一个窗口，这个窗口包含由字符串各种排列组合组成的，
窗口中可以包含 T 中没有的字符，如果存在多个，在结果中输出最小的窗口，如果找不到这样的窗口，输出空字符串。
*/
func minWindow(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}
	var tFreq, sFreq [256]int
	result, left, right, finalLeft, finalRight, minWin, count := "", 0, -1, -1, -1, len(s)+1, 0
	// 存储目标字符串对应个数
	for i := 0; i < len(t); i++ {
		tFreq[t[i]-'a']++
	}
	for left < len(s) {
		// 需要包含t
		if right+1 < len(s) && count < len(t) {
			sFreq[s[right+1]-'a']++
			if sFreq[s[right+1]-'a'] <= tFreq[s[right+1]-'a'] {
				count++
			}
			right++
		} else {
			// 寻找到
			if right-left+1 < minWin && count == len(t) {
				minWin = right - left + 1
				finalLeft = left
				finalRight = right
			}
			// 相等为什么--????
			if sFreq[s[left]-'a'] == tFreq[s[left]-'a'] {
				count--
			}
			// 去掉当前值,为了遍历
			sFreq[s[left]-'a']--
			left++
		}
		if finalLeft != -1 {
			result = s[finalLeft : finalRight+1]
		}
	}
	return result
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}
