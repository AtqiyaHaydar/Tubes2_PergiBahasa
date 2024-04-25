package main

import (
	"fmt"
	"strings"
	"sync"
)

const (
	producerCount = 50
	consumerCount = 5
)

func CPBFS(awal string, tujuan string) ([]Node, int) {
	_depth := 0
	ammtOfArticles := 0
	var _trail []string
	_trail = append(_trail, awal)
	seed := makeNode(awal, _depth, _trail)
	queue1 := make(chan *Node)
	queue2 := make(chan *Node)
	var result []Node
	flag := false
	queue1 <- seed
	notFound := true
	var wg sync.WaitGroup
	for notFound {
		wg.Add(producerCount)
		for i := 0; i < producerCount; i++ {
			go producer(queue1, queue2, &wg)
		}
		wg.Wait()

		wg.Add(consumerCount)
		for i := 0; i < consumerCount; i++ {
			go consumer(queue1, queue2, tujuan, &result, &flag, &wg)
		}
		wg.Wait()

		// Check for termination condition here
		// Update notFound accordingly
	}
	return result, ammtOfArticles
}

func producer(ch1 chan<- *Node, ch2 chan *Node, wg *sync.WaitGroup) {
	defer wg.Done()
	tempBuffer2 := <-ch2
	fmt.Println(tempBuffer2.name)
	fmt.Println(tempBuffer2.depth)
	parentNode := makeNode(tempBuffer2.name, tempBuffer2.depth, tempBuffer2.trail)
	for i := 0; i < len(parentNode.link); i++ {
		tempNode := newNode(parentNode.link[i], parentNode.depth+1, append(parentNode.trail, parentNode.link[i]))
		ch1 <- tempNode
	}
}

func consumer(ch1 <-chan *Node, ch2 chan<- *Node, tujuan string, result *[]Node, flag *bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range ch1 {
		if strings.EqualFold(data.name, tujuan) {
			*flag = true
			*result = append(*result, *data)
		} else {
			ch2 <- data
		}
	}
}
