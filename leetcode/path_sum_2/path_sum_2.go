package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	sequenceQueue := make([][]int, 0)
	sequenceQueue = append(sequenceQueue, []int{root.Val})

	for len(queue) != 0 {
		temp := queue[0]
		queue = queue[1:]

		tempSequence := sequenceQueue[0]
		sequenceQueue = sequenceQueue[1:]

		if temp.Left == nil && temp.Right == nil {
			accum := 0
			for i := range tempSequence {
				accum += tempSequence[i]
			}
			if accum == targetSum {
				result = append(result, tempSequence)
			}
		} else {
			if temp.Left != nil {
				queue = append(queue, temp.Left)
				newSeq := make([]int, len(tempSequence)+1)
				copy(newSeq, tempSequence)
				newSeq[len(newSeq)-1] = temp.Left.Val
				sequenceQueue = append(sequenceQueue, newSeq)
			}
			if temp.Right != nil {
				queue = append(queue, temp.Right)
				newSeq := make([]int, len(tempSequence)+1)
				copy(newSeq, tempSequence)
				newSeq[len(newSeq)-1] = temp.Right.Val
				sequenceQueue = append(sequenceQueue, newSeq)
			}
		}
	}
	return result
}

func ListToTreeNode(arr []interface{}, k int) *TreeNode {
	if k >= len(arr) {
		return nil
	}
	if v, ok := arr[k].(int); !ok {
		return nil
	} else {
		return &TreeNode{Val: v, Left: ListToTreeNode(arr, 2*k+1), Right: ListToTreeNode(arr, 2*k+2)}
	}
}

func (t *TreeNode) String() string {
	queue := make([]*TreeNode, 0)
	queue = append(queue, t)

	var buffer bytes.Buffer

	for len(queue) > 0 {
		temp := queue[0]
		queue = queue[1:]

		if temp == nil {
			buffer.WriteString("null")
			buffer.WriteString(" ")
		} else {
			buffer.WriteString(strconv.FormatInt(int64(temp.Val), 10))
			buffer.WriteString(" ")
			if temp.Left == nil && temp.Right == nil {
				continue
			}
			queue = append(queue, temp.Left)
			queue = append(queue, temp.Right)
		}
	}
	return buffer.String()
}

func main() {
	root := ListToTreeNode([]interface{}{
		5,
		4,8,
		11,nil,13,4,
		7,2,nil,nil,nil,nil,5,1}, 0)
	fmt.Println(root.String())
	fmt.Println(pathSum(root, 22))
}
