package main

// https://leetcode.com/problems/same-tree/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	return recursiveIdentity(p, q)
}

func recursiveIdentity(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	if !recursiveIdentity(p.Left, q.Left) {
		return false
	}

	if !recursiveIdentity(p.Right, q.Right) {
		return false
	}

	return true
}

