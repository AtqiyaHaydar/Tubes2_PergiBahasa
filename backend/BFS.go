package main

import (
	"fmt"

	"github.com/gammazero/deque"
)

func BFS(awal string, tujuan string) ([]Node, int) {
	_depth := 0
	ammtOfArticles := 0
	var _trail []string
	_trail = append(_trail, awal)
	seed := newNode(awal, _depth, _trail)
	var buffer deque.Deque[Node]
	var buffer2 deque.Deque[Node]
	var result []Node
	flag := false
	buffer.PushBack(*seed)
	visited := make(map[string]bool)
	visited2 := make(map[string]bool)
	var notFound bool = true
	visited[awal] = true
	for notFound {
		livingNode := buffer.Front()
		if flag {
			if livingNode.depth > result[0].depth {
				notFound = false
			}
		}
		buffer.PopFront()
		if livingNode.depth == 0 {
			parentNode := makeNode(livingNode.name, livingNode.depth, livingNode.trail)
			for i := 0; i < len(parentNode.link); i++ {
				if visited[parentNode.link[i]] == false {
					tempNode := newNode(parentNode.link[i], parentNode.depth+1, append(parentNode.trail, parentNode.link[i]))
					buffer.PushBack(*tempNode)
				}
			}
		}
		if buffer.Len() == 0 {
			for buffer.Len() == 0 {
				for visited2[buffer2.Front().name] {
					buffer2.PopFront()
				}
				parentNode := makeNode(buffer2.Front().name, buffer2.Front().depth, buffer2.Front().trail)
				visited2[buffer2.Front().name] = true
				buffer2.PopFront()
				for i := 0; i < len(parentNode.link); i++ {
					if visited[parentNode.link[i]] == false {
						tempNode := newNode(parentNode.link[i], parentNode.depth+1, append(parentNode.trail, parentNode.link[i]))
						buffer.PushBack(*tempNode)
					}
				}
			}
		}
		ammtOfArticles += 1
		fmt.Println(livingNode.name)
		fmt.Println(livingNode.depth)
		if livingNode.name == tujuan {
			flag = true
			result = append(result, livingNode)
		} else {
			tempNode2 := *newNode(livingNode.name, livingNode.depth, livingNode.trail)
			buffer2.PushBack(tempNode2)
		}
		visited[livingNode.name] = true
	}
	return result, ammtOfArticles
}

func printNode(_node Node) {
	for i := 0; i < len(_node.trail); i++ {
		fmt.Println(_node.trail[i])
	}
}

func printSlice(slc []Node) {
	fmt.Printf("[")
	for i := 0; i < len(slc); i++ {
		if i == len(slc)-1 {
			fmt.Print(slc[i].name)
		} else {
			fmt.Print(slc[i].name, ",")
		}
	}
	fmt.Println("]")
}
