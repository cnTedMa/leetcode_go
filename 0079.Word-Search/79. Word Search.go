package leetcode

var dir = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func exist(board [][]byte, word string) bool {
	hasVisited := make([][]bool, len(board))
	for i := 0; i < len(hasVisited); i++ {
		hasVisited[i] = make([]bool, len(board[0]))
	}
	for i, v := range board {
		for j := range v {
			if searchWord(board, hasVisited, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func isInBoard(board [][]byte, x, y int) bool {
	return x < len(board) && y < len(board[0]) && x >= 0 && y >= 0
}

func searchWord(board [][]byte, hasVisited [][]bool, word string, index, x, y int) bool {
	if index == len(word)-1 {
		return word[index] == board[x][y]
	}
	if board[x][y] == word[index] {
		hasVisited[x][y] = true
		for i := 0; i < 4; i++ {
			nx := x + dir[i][0]
			ny := y + dir[i][1]
			if isInBoard(board, nx, ny) && !hasVisited[nx][ny] && searchWord(board, hasVisited, word, index+1, nx, ny) {
				return true
			}
		}
		hasVisited[x][y] = false
	}
	return false
}
