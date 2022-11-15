package main

import "fmt"

/*
	字符串的排列
*/

// 方法一：滑动窗口
func checkInclusionOne(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	var cnt1, cnt2 [26]int
	// 遍历s1长度
	for i, ch := range s1 {
		cnt1[ch-'a']++
		cnt2[s2[i]-'a']++
	}
	if cnt1 == cnt2 {
		return true
	}
	// s1的窗口内滑动
	for i := n; i < m; i++ {
		// 窗口滑动，大小s1的长度
		cnt2[s2[i]-'a']++
		cnt2[s2[i-n]-'a']--
		if cnt1 == cnt2 {
			return true
		}
	}
	return false
}

// 方法二：滑动窗口改进版
func checkInclusionTwo(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	cnt := [26]int{}
	for i, ch := range s1 {
		cnt[ch-'a']--
		cnt[s2[i]-'a']++
	}
	diff := 0
	// 剩余字符
	for _, c := range cnt[:] {
		if c != 0 {
			diff++
		}
	}
	if diff == 0 {
		return true
	}
	for i := n; i < m; i++ {
		// 一个窗口
		x, y := s2[i]-'a', s2[i-n]-'a'
		if x == y {
			continue
		}
		if cnt[x] == 0 {
			diff++
		}
		cnt[x]++
		if cnt[x] == 0 {
			diff--
		}
		if cnt[y] == 0 {
			diff++
		}
		cnt[y]--
		if cnt[y] == 0 {
			diff--
		}
		if diff == 0 {
			return true
		}
	}
	return false
}

// 双指针
func checkInclusionThree(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	cnt := [26]int{}
	for _, ch := range s1 {
		cnt[ch-'a']--
	}
	left := 0
	for right, ch := range s2 {
		x := ch - 'a'
		cnt[x]++
		for cnt[x] > 0 {
			cnt[s2[left]-'a']--
			left++
		}
		if right-left+1 == n {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(checkInclusionOne("ab", "eidbaooo"))
	fmt.Println(checkInclusionTwo("ab", "eidbaooo"))
	fmt.Println(checkInclusionThree("ab", "eidbaooo"))
}
