package main

import "fmt"

type Node struct {
	Next *Node
	Elem int
}

// Reverse Идея: нужно пройтись по списку и на каждом шаге текущий элемент делать рутом, а следующим к нему делать старый рут
func (n *Node) Reverse() *Node {
	var newRoot, prevRoot *Node

	for curr := n; curr != nil; curr = curr.Next {
		newRoot = &Node{
			Next: prevRoot,
			Elem: curr.Elem,
		}
		prevRoot = newRoot
	}

	return newRoot
}

func (n *Node) Print() string {
	if n.Next != nil {
		return fmt.Sprintf("%d -> ", n.Elem) + n.Next.Print()
	}
	return fmt.Sprintf("%d", n.Elem)
}

func main() {
	root := &Node{
		Elem: 5,
		Next: &Node{
			Elem: 4,
			Next: &Node{
				Elem: 3,
				Next: nil,
			},
		},
	}
	fmt.Println(root.Print())
	fmt.Println(root.Reverse().Print())
}
