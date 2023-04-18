package main

import "fmt"

type Node struct {
	title    string
	children []Node
}

type Tree = []Node

var input = Tree{
	Node{
		title: "Вещи",
		children: Tree{
			Node{
				title: "Одежда",
				children: Tree{
					Node{
						title: "Мужская",
					},
					Node{
						title: "Женская",
					},
				},
			},
		},
	},
	Node{
		title: "Хобби",
		children: Tree{
			Node{
				title: "Велосипеды",
				children: Tree{
					Node{
						title: "Горные",
					},
				},
			},
			Node{
				title: "Мангалы",
			},
		},
	},
	Node{
		title: "Транспорт",
	},
}

func main() {
	printTree(input)
}

func printRecursive(parentTitle string, childrenNodes Tree) {
	if len(childrenNodes) == 0 {
		fmt.Printf("%s\n", parentTitle)
		return
	}
	for _, childrenNode := range childrenNodes {
		printRecursive(parentTitle+" => "+childrenNode.title, childrenNode.children)
	}
}

func printTree(rootNodes Tree) {
	for _, rootNode := range rootNodes {
		printRecursive(rootNode.title, rootNode.children)
	}
}
