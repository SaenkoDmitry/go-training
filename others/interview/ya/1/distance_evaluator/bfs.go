package main

import (
	"fmt"
)

type Node struct {
	value int
	children *[]Node
}

type DistanceEvaluator struct {
	m map[int]int
}

func (d *DistanceEvaluator) init(root *Node) {
	if root == nil {
		return
	}
	d.m[0] = root.value

	for i := range *root.children {
		d.m[1] += (*root.children)[i].value
		d.recursive((*root.children)[i], 2)
	}

}

func (d *DistanceEvaluator) recursive(vertex Node, level int) {
	if vertex.children == nil {
		return
	}
	for i := range *vertex.children {
		d.m[level] += (*vertex.children)[i].value
		d.recursive((*vertex.children)[i], level + 1)
	}
}


func (d *DistanceEvaluator) calcSumAtDistance(distance int) int {
	if val, ok := d.m[distance]; ok {
		return val
	}
	return 0
}

//    1
//   / \
//  2   3
// | \  |
// 4  5 7

func main() {
	d := DistanceEvaluator{
		m: make(map[int]int),
	}

	d.init(&Node{
		value: 1,
		children: &[]Node{
			Node{
				value: 2,
				children: &[]Node{
					Node{
						value: 4,
						children: nil,
					},
					Node{
						value: 5,
						children: nil,
					},
				},
			},
			Node{
				value: 3,
				children: &[]Node{
					Node{
						value: 7,
						children: nil,
					},
				},
			},
		},
	})

	fmt.Println(d.calcSumAtDistance(0))
	fmt.Println(d.calcSumAtDistance(1))
	fmt.Println(d.calcSumAtDistance(2))
	fmt.Println(d.calcSumAtDistance(3))
}
