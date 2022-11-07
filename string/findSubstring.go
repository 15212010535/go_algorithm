package main

import "fmt"

/*
给定一个源字符串 s，再给一个字符串数组，要求在源字符串中找到由字符串数组各种组合组成的连续串的起始下标，
如果存在多个，在结果中都需要输出。
*/
func findSubstring(s string, words []string) (ans []int) {
	ls, m, n := len(s), len(words), len(words[0])
	// 遍历
	for i := 0; i < n && i+m+n <= ls; i++ {
		differ := map[string]int{}
		// 遍历，切分并存储单词
		for j := 0; j < m; j++ {
			differ[s[i+j*n:i+(j+1)*n]]++
		}
		// 匹配
		for _, word := range words {
			differ[word]--
			if differ[word] == 0 {
				// 不存在，删除
				delete(differ, word)
			}
		}
		for start := i; start < ls-m*n+1; start += n {
			if start != i {
				word := s[start+(m-1)*n : start+m*n]
				differ[word]++
				if differ[word] == 0 {
					delete(differ, word)
				}
				word = s[start-n : start]
				differ[word]--
				if differ[word] == 0 {
					delete(differ, word)
				}
			}
			if len(differ) == 0 {
				ans = append(ans, start)
			}
		}
	}
	return
}

func main() {
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
}
