package main

import (
	"fmt"
	"time"
	"tubes2/crawl"
)

func main() {
	var bruh Node = *makeNode("Bandung_Institute_of_Technology", 0, crawl.Scrape("Bandung_Institute_of_Technology"))
	fmt.Println("Starting")
	flag := make(chan bool)
	go clock(flag)
	flag <- false
	var result []Node = IDS(bruh, "Laos", 2, true)
	flag <- true
	//MUTEX BABEY
	//var thing []string = scrape("File:Diploma_icon.png")
	fmt.Println("Winners:")
	for i := 0; i < len(result); i++ {
		printNode(result[i])
	}
	//fmt.Println(bruh.link[0])
}

func printNode(_node Node) {
	for i := 0; i < len(_node.trail); i++ {
		fmt.Println(_node.trail[i])
	}
}

func clock(flag chan bool) {
	var ms int = 0
	var seconds int = 0
	stop := <-flag
	for !stop {
		if !stop {
			time.Sleep(10 * time.Millisecond)
			ms = ms + 1
		}
		select {
		case newstop := <-flag:
			stop = newstop
		default:
			stop = stop
			if ms/100 > seconds {
				seconds = ms / 100
				fmt.Println(seconds, visits)
			}
		}

	}
	fmt.Printf("Done, time taken : %d.%02d seconds \n", seconds, ms-seconds*100)
}
