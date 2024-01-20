package main

import "fmt"

func getLayeredRepresentation(root *Tree) [][]int {
	result := make([][]int, 0)

	var deepFirstSearch func(root *Tree, depth int)

	deepFirstSearch = func(root *Tree, depth int) {
		if root == nil {
			return
		}
		if depth >= len(result) {
			result = append(result, make([]int, 0))
		}
		result[depth] = append(result[depth], root.Val)
		deepFirstSearch(root.Left, depth+1)
		deepFirstSearch(root.Right, depth+1)
	}

	deepFirstSearch(root, 0)
	return result
}

type Tree struct {
	Val         int
	Left, Right *Tree
}

func main() {
	root := &Tree{
		Left: &Tree{
			Left: nil,
			Right: &Tree{
				Val: 15,
			},
			Val: 20,
		},
		Right: &Tree{
			Left: &Tree{
				Left: &Tree{
					Val: 42,
				},
				Val: 25,
			},
			Right: &Tree{
				Val: 18,
			},
			Val: 70,
		},
		Val: 1,
	}
	result := getLayeredRepresentation(root)
	for i := range result {
		fmt.Println(result[i])
	}
}
