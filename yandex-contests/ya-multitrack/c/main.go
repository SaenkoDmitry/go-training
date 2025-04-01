package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Tree struct {
	Path     string
	Read     bool
	Write    bool
	Children []*Tree
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var a, b, c, d string

	trees := make(map[string]*Tree)

	for {
		fmt.Fscan(in, &a)
		fmt.Fscan(in, &b)
		fmt.Fscan(in, &c)
		fmt.Fscan(in, &d)

		if a == "grant" {
			addToTree(trees, b, c, d)
			fmt.Println("done")
			continue
		}

		if a == "block" {
			fmt.Println("done")
			continue
		}

		if hasAccess(trees, b, c, d) {
			fmt.Println("allowed")
			continue
		}

		fmt.Println("forbidden")
	}
}

func addToTree(trees map[string]*Tree, name, command, path string) {
	root, ok := trees[name]
	if !ok {
		root = &Tree{Path: "/"}
		trees[name] = root
	}

	nodes := strings.Split(path, "/")
	head := root
	for len(nodes) > 1 {
		top := nodes[0]
		nodes = nodes[1:]
		head.Children = append(head.Children, &Tree{Path: top})
		head = head.Children[0]
	}
	head.Children = append(head.Children, &Tree{Path: nodes[0]})
	if command == "read" {
		head.Children[0].Read = true
	} else {
		head.Children[0].Write = true
	}
}

func hasAccess(rules map[string]*Tree, name, command, path string) bool {
	return false
}

/* input
grant bob read /home/bob
block bob read /
check bob read /home/bob/Pictures/gary.jpg
check bob read /etc/shadow
check bob write /home/bob/Pictures/gary.jpg
grant bob write /home/bob
check bob write /home/bob/Pictures/gary.jpg
block bob write /home/bob/Pictures
check bob write /home/bob/Pictures/gary.jpg
check bob read /home/bob/Pictures/gary.jpg
grant root read /
check root read /home/bob/Pictures/gary.jpg
check root read /etc/shadow
*/

//done
//done
//allowed
//forbidden
//forbidden
//done
//allowed
//done
//forbidden
//allowed
//done
//allowed
//allowed
