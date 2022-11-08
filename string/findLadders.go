package main

import "fmt"

/*
	给定两个单词（beginWord 和 endWord）和一个字典 wordList，找出所有从 beginWord 到 endWord 的最短转换序列。转换需遵循如下规则：
		每次转换只能改变一个字母。
		转换过程中的中间单词必须是字典中的单词
*/
// 广度遍历，需要一个队列
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	result, wordMap := make([][]string, 0), make(map[string]bool)
	for _, w := range wordList {
		wordMap[w] = true
	}
	if !wordMap[endWord] {
		return result
	}
	queue := make([][]string, 0)
	queue = append(queue, []string{beginWord})
	queueLen := 1

	levelMap := make(map[string]bool)
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		lastWord := path[len(path)-1]
		// 广度遍历
		for i := 0; i < len(lastWord); i++ {
			// 遍历替换单词
			for c := 'a'; c <= 'z'; c++ {
				nextWord := lastWord[:i] + string(c) + lastWord[i+1:]
				// 找到结束
				if nextWord == endWord {
					path = append(path, endWord)
					result = append(result, path)
					continue
				}
				// 字符数组中存在，接着遍历
				if wordMap[nextWord] {
					levelMap[nextWord] = true
					newPath := make([]string, len(path))
					copy(newPath, path)
					newPath = append(newPath, nextWord)
					queue = append(queue, newPath)
				}
			}
		}
		queueLen--
		if queueLen == 0 {
			if len(result) > 0 {
				return result
			}
			// 删除结点
			for k := range levelMap {
				delete(wordMap, k)
			}
			// 清除levelMap
			levelMap = make(map[string]bool)
			queueLen = len(queue)
		}
	}
	return result
}

func main() {
	fmt.Println(findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}
