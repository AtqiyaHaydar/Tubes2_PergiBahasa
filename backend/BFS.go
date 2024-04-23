package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gammazero/deque"
)

const maxGoRoutine = 5

func BFS(awal string, tujuan string) ([]Node, int) {
	_depth := 0
	ammtOfArticles := 0
	var _trail []string
	_trail = append(_trail, awal)
	seed := newNode(awal, _depth, _trail)
	//var runningGoRoutines int
	var buffer deque.Deque[Node]
	var buffer2 deque.Deque[Node]
	var result []Node
	flag := false
	buffer.PushBack(*seed)
	visited := make(map[string]bool)
	//sem := make(chan struct{}, maxGoRoutine)
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
		ammtOfArticles += 1
		fmt.Println(livingNode.name)
		fmt.Println(livingNode.depth)
		if strings.EqualFold(livingNode.name, tujuan) {
			flag = true
			result = append(result, livingNode)
		} else {
			visited[livingNode.name] = true
			tempNode2 := *newNode(livingNode.name, livingNode.depth, livingNode.trail)
			buffer2.PushBack(tempNode2)
		}

		if buffer.Len() == 0 {
			finishedGoroutine := 0
			for i := 0; i < maxGoRoutine; i++ {
				go func() {
					for buffer.Len() == 0 {
						parentNode := makeNode(buffer2.Front().name, buffer2.Front().depth, buffer2.Front().trail)
						buffer2.PopFront()
						for i := 0; i < len(parentNode.link); i++ {
							if visited[parentNode.link[i]] == false {
								tempNode := newNode(parentNode.link[i], parentNode.depth+1, append(parentNode.trail, parentNode.link[i]))
								buffer.PushBack(*tempNode)
							}
						}
					}
					finishedGoroutine += 1
				}()
				time.Sleep(10 * time.Millisecond)
			}
			for finishedGoroutine < maxGoRoutine {
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
