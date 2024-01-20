package tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTree(t *testing.T) {
	file1 := "/home/bob/info.txt"
	file2 := "/home/bob/a/info.txt"
	file3 := "/home/bob/b/info.txt"
	file4 := "/home/john/info.txt"

	root := BuildTree(file1) //OR root.AppendByPath(&Tree{Path: "/"}, file1)
	root = root.AppendByPath(file2)
	root = root.AppendByPath(file3)
	root = root.AppendByPath(file4)
	_ = root

	fmt.Println(root.GetAllPaths())
	got := root.GetAllPaths()
	assert.Equal(t, 4, len(got))
	assert.Equal(t, file1, got[0])
	assert.Equal(t, file2, got[1])
	assert.Equal(t, file3, got[2])
	assert.Equal(t, file4, got[3])
}
