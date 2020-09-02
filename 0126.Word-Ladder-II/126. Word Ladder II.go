package leetcode

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
		for i := 0; i < len(lastWord); i++ {
			for c := 'a'; c <= 'z'; c++ {

			}
		}
	}
}
