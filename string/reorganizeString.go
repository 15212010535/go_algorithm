package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/*
	给定一个字符串，要求重新排列字符串，让字符串两两字符不相同，如果可以实现，
	即输出最终的字符串，如果不能让两两不相同，则输出空字符串。
*/

var cnt [26]int

type hp struct {
	sort.IntSlice
}

func (h hp) Less(i, j int) bool {
	return cnt[h.IntSlice[i]] > cnt[h.IntSlice[j]]
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func (h *hp) push(v int) {
	heap.Push(h, v)
}

func (h *hp) pop() int {
	return heap.Pop(h).(int)
}

func reorganizeStringOne(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	cnt = [26]int{}
	maxCnt := 0

	for _, ch := range s {
		ch -= 'a'
		cnt[ch]++
		if cnt[ch] > maxCnt {
			maxCnt = cnt[ch]
		}
	}

	if maxCnt > (n+1)/2 {
		return ""
	}

	h := &hp{}
	for i, c := range cnt[:] {
		if c > 0 {
			h.IntSlice = append(h.IntSlice, i)
		}
	}
	heap.Init(h)
	ans := make([]byte, 0, n)
	for len(h.IntSlice) > 1 {
		i, j := h.pop(), h.pop()
		ans = append(ans, byte('a'+i), byte('a'+j))
		if cnt[i]--; cnt[i] > 0 {
			h.push(i)
		}
		if cnt[j]--; cnt[j] > 0 {
			h.push(j)
		}
	}
	if len(h.IntSlice) > 0 {
		ans = append(ans, byte('a'+h.IntSlice[0]))
	}
	return string(ans)
}

func reorganizeStringTwo(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	cnt := [26]int{}
	maxCnt := 0
	for _, ch := range s {
		ch -= 'a'
		cnt[ch]++
		if cnt[ch] > maxCnt {
			maxCnt = cnt[ch]
		}
	}
	if maxCnt > (n+1)/2 {
		return ""
	}

	ans := make([]byte, n)
	evenIdx, oddIdx, halfLen := 0, 1, n/2
	for i, c := range cnt[:] {
		ch := byte('a' + i)
		for c > 0 && c <= halfLen && oddIdx < n {
			ans[oddIdx] = ch
			c--
			oddIdx += 2
		}
		for c > 0 {
			ans[evenIdx] = ch
			c--
			evenIdx += 2
		}
	}
	return string(ans)
}

func main() {
	fmt.Println(reorganizeStringOne("aab"))
	fmt.Println(reorganizeStringTwo("aab"))
}
