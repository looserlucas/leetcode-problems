package main

// word search 1
func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] {
				var mark [20][20]bool
				if dfs1(board, word, mark, 0, i, j) {
					return true
				}
			}
		}
	}
	return false
}

func dfs1(board [][]byte, word string, mark [20][20]bool, idx, i, j int) bool {
	mark[i][j] = true
	if idx == len(word)-1 {
		if board[i][j] == word[idx] {
			return true
		} else {
			return false
		}
	}
	var ok1, ok2, ok3, ok4 bool
	if i-1 >= 0 && !mark[i-1][j] && board[i-1][j] == word[idx+1] {
		ok1 = dfs1(board, word, mark, idx+1, i-1, j)
	}
	if j-1 >= 0 && !mark[i][j-1] && board[i][j-1] == word[idx+1] {
		ok2 = dfs1(board, word, mark, idx+1, i, j-1)
	}
	if i+1 < len(board) && !mark[i+1][j] && board[i+1][j] == word[idx+1] {
		ok3 = dfs1(board, word, mark, idx+1, i+1, j)
	}
	if j+1 < len(board[i]) && !mark[i][j+1] && board[i][j+1] == word[idx+1] {
		ok4 = dfs1(board, word, mark, idx+1, i, j+1)
	}
	return ok1 || ok2 || ok3 || ok4
}
