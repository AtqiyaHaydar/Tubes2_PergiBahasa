package main

import (
	"fmt"
	"math"

	"github.com/gammazero/deque"
)

func BFS(awal string, tujuan string) ([]Node, int) {
	_depth := 0
	ammtOfArticles := 0
	shortestDepth := math.MaxInt32 // some very large number
	var _trail []string
	_trail = append(_trail, awal)
	seed := newNode(awal, _depth, _trail)
	var q deque.Deque[Node]
	var result []Node
	flag := false
	q.PushBack(*seed)
	visited := make(map[string]bool)
	var notFound bool = true
	visited[awal] = true
	for notFound {
		var livingNode Node
		if flag { // kalau udh pernah ketemu a.k.a. depth terkecil ketemu, gausah bikin anak
			livingNode = q.Front()
			if livingNode.depth > shortestDepth {
				notFound = false
			}
			if q.Len() == 0 {
				notFound = false
			}
		} else {
			livingNode = *makeNode(q.Front().name, q.Front().depth, q.Front().trail)
		}
		q.PopFront()
		ammtOfArticles += 1
		fmt.Println(livingNode.name)
		fmt.Println(livingNode.depth)
		if livingNode.name == tujuan {
			if shortestDepth == math.MaxInt32 {
				shortestDepth = livingNode.depth
			}
			flag = true
			result = append(result, livingNode)
		} else {
			for i := 0; i < len(livingNode.link); i++ {
				if visited[livingNode.link[i]] == false {
					if livingNode.link[i] != tujuan {
						visited[livingNode.link[i]] = true
					}
					tempNode := newNode(livingNode.link[i], livingNode.depth+1, append(livingNode.trail, livingNode.link[i]))
					q.PushBack(*tempNode)
				}
			}

		}
	}
	return result, ammtOfArticles
}

func printNode(_node Node) {
	for i := 0; i < len(_node.trail); i++ {
		fmt.Println(_node.trail[i])
	}
}
