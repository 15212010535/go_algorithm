package main

import "fmt"

/*
	划分字母区间
*/

func partitionLabels(s string) (partition []int) {
	lastPos := [26]int{}
	for i, c := range s {
		// 记录每个元素最后一次出现的位置
		lastPos[c-'a'] = i
	}
	start, end := 0, 0
	for i, c := range s {
		// 每个字母最后一次出现的位置
		if lastPos[c-'a'] > end {
			end = lastPos[c-'a']
		}
		if i == end {
			partition = append(partition, end-start+1)
			start = end + 1
		}
	}
	return
}

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
}
