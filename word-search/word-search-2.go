package main

// Trie Implementation
type TrieNode struct {
	Children    []*TrieNode
	IsEndOfWord bool
	Word        string
}

const mASize int = 27

func NewNode() (n *TrieNode) {
	return &TrieNode{
		Children:    make([]*TrieNode, mASize),
		IsEndOfWord: false,
		Word:        "",
	}
}

func Insert(root *TrieNode, key string) {
	currentNode := root
	for i := 0; i < len(key); i++ {
		index := int(key[i]) - 97
		if currentNode.Children[index] == nil {
			currentNode.Children[index] = NewNode()
		}
		currentNode = currentNode.Children[index]
	}

	currentNode.IsEndOfWord = true
	currentNode.Word = key
}

func Search(root *TrieNode, key string) (ok bool) {
	currentNode := root
	i := 0
	for i < len(key) {
		index := int(key[i]) - 97
		if currentNode.Children[index] == nil {
			return false
		} else {
			currentNode = currentNode.Children[index]
		}
	}

	return true
}

var res []string

func findWords(board [][]byte, words []string) []string {
	root := NewNode()
	res = nil
	for i := 0; i < len(words); i++ {
		Insert(root, words[i])
	}

	var marker [20][20]bool
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			dfs(board, marker, root, i, j)
		}
	}
	return res
}

func dfs(board [][]byte, marker [20][20]bool, node *TrieNode, i, j int) {
	if marker[i][j] || node.Children[int(board[i][j])-97] == nil {
		return
	}

	node = node.Children[int(board[i][j])-97]
	// found the word
	if node.Word != "" {
		res = append(res, node.Word)
		node.Word = ""
	}

	marker[i][j] = true
	if i > 0 {
		dfs(board, marker, node, i-1, j)
	}
	if j > 0 {
		dfs(board, marker, node, i, j-1)
	}
	if i+1 < len(board) {
		dfs(board, marker, node, i+1, j)
	}
	if j+1 < len(board[i]) {
		dfs(board, marker, node, i, j+1)
	}
	marker[i][j] = false
}
