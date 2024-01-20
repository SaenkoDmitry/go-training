package tree

import "strings"

func BuildTree(path string) *Tree {
	root := &Tree{Path: "/"}
	curr := root
	nodes := strings.Split(path, "/")
	for _, node := range nodes[1:] {
		curr.Children = append(curr.Children, &Tree{Path: node})
		curr = curr.Children[0]
	}
	return root
}
