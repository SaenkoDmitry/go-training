package main

import (
	"fmt"
	"testing"
)

func TestIsSameTreeWhenIsEqual(t *testing.T) {
	p := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	q := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	equal := IsSameTree(p, q)
	if !equal {
		t.Fatal(fmt.Printf("trees not equal : %v and %v", p, q))
	}
}

func TestIsSameTreeWhenIsNotEqual(t *testing.T) {
	p := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	q := &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}
	equal := IsSameTree(p, q)
	if equal {
		t.Fatal(fmt.Printf("trees not equal : %v and %v", p, q))
	}
}
