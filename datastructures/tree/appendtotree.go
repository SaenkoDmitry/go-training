package tree

import "strings"

func (root *Tree) AppendByPath(path string) *Tree {
	curr := root
	nodes := strings.Split(path, "/")
	nodes = nodes[1:] // skip root path
	for len(nodes) > 0 {
		top := nodes[0]
		nodes = nodes[1:]
		samePath := false
		for _, item := range curr.Children {
			if item.Path == top {
				curr = item
				samePath = true
				break
			}
		}
		if samePath {
			continue
		}
		curr.Children = append(curr.Children, &Tree{Path: top})
		curr = curr.Children[len(curr.Children)-1]
	}
	return root
}
