package main

import "fmt"

func checkRecursively(board [][]byte, i int, j int, visited [][]bool, word string) bool {
	if len(word) == 0 {
		return true
	}

	if i > 0 && !visited[i-1][j] && board[i-1][j] == word[0] {
		visited[i-1][j] = true
		if checkRecursively(board, i-1, j, visited, word[1:]) {
			return true
		}
		visited[i-1][j] = false
	}
	if i < len(board)-1 && !visited[i+1][j] && board[i+1][j] == word[0] {
		visited[i+1][j] = true
		if checkRecursively(board, i+1, j, visited, word[1:]) {
			return true
		}
		visited[i+1][j] = false
	}
	if j > 0 && !visited[i][j-1] && board[i][j-1] == word[0] {
		visited[i][j-1] = true
		if checkRecursively(board, i, j-1, visited, word[1:]) {
			return true
		}
		visited[i][j-1] = false
	}
	if j < len(board[i])-1 && !visited[i][j+1] && board[i][j+1] == word[0] {
		visited[i][j+1] = true
		if checkRecursively(board, i, j+1, visited, word[1:]) {
			return true
		}
		visited[i][j+1] = false
	}

	return false
}

func exist(board [][]byte, word string) bool {

	visited := make([][]bool, 0)
	for i := 0; i < len(board); i++ {
		visited = append(visited, make([]bool, len(board[i])))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {

			if board[i][j] == word[0] {
				visited[i][j] = true
				if checkRecursively(board, i, j, visited, word[1:]) {
					return true
				}
				visited[i][j] = false
			}
		}
	}
	return false
}

func main() {
	//fmt.Println(exist([][]byte{{'a', 'a'}}, "aa"))
	//fmt.Println(exist([][]byte{
	//	{'a', 'b', 'c', 'e'},
	//	{'s', 'f', 'e', 's'},
	//	{'a', 'd', 'e', 'e'},
	//}, "abceseeefs"))
	fmt.Println(exist([][]byte{
	{'A','B','C','E'},
	{'S','F','C','S'},
	{'A','D','E','E'},
	}, "ABCCED"))
	//fmt.Println(exist([][]byte{
	//	{'c', 'a', 'a'},
	//	{'a', 'a', 'a'},
	//	{'b', 'c', 'd'},
	//}, "aab"))
}
