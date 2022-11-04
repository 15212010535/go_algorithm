package main

import "fmt"

/*
	在一个字符串中寻找没有重复字母的最长子串。
*/

// 解法一 位图
func lengthOfLongestSubstringOne(s string) int {
	if len(s) == 0 {
		return 0
	}
	var bitSet [256]bool
	result, left, right := 0, 0, 0
	for len(s) > left {
		// 右侧字符对应的bitSet被标记为true，说明此字符在x位置重复，
		// 需要左侧向前移动，指导x标记为false
		if bitSet[s[right]] {
			bitSet[s[left]] = false
			left++
		} else {
			bitSet[s[right]] = true
			right++
		}
		if result < right-left {
			result = right - left
		}
		if left+result >= len(s) || right >= len(s) {
			break
		}
	}
	return result
}

// 解法二：滑动窗口
func lengthOfLongestSubstringTwo(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [127]int
	result, left, right := 0, 0, -1
	for len(s) > left {
		if right+1 < len(s) && freq[s[right+1]] == 0 {
			freq[s[right+1]]++
			right++
		} else { // 遇到重复字母，把重复字母去掉，再次循环
			freq[s[left]]--
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// 方法三：滑动窗口-哈希桶(map)
func lengthOfLongestSubstringThree(s string) int {
	right, left, res := 0, 0, 0
	indexes := make(map[byte]int, len(s))
	for left < len(s) {
		// 当桶里有数据时候，取出对应坐标
		if idx, ok := indexes[s[left]]; ok && idx >= right {
			right = idx + 1
		}
		// 向桶里存放数据，key: 字母， value: 下标
		indexes[s[left]] = left
		left++
		res = max(res, left-right)
	}
	return res
}

func main() {
	fmt.Println(lengthOfLongestSubstringOne("abcabcbb"))
	fmt.Println(lengthOfLongestSubstringTwo("abcabcbb"))
	fmt.Println(lengthOfLongestSubstringThree("abcabcbb"))
}
