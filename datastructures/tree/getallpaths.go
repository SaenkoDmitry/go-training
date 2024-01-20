package tree

func (root *Tree) GetAllPaths() []string {
	if root == nil {
		return make([]string, 0)
	}
	return getPaths(root, "")
}

func getPaths(root *Tree, acc string) []string {
	if len(root.Children) == 0 {
		return []string{acc}
	}
	res := make([]string, 0)
	for i := range root.Children {
		res = append(res, getPaths(root.Children[i], acc+"/"+root.Children[i].Path)...)
	}
	return res
}
