package main

import (
	"fmt"
	"time"
)

func main() {
	var bruh Node = *makeNode("Basketball", scrape("Basketball"))
	fmt.Println("Starting")
	flag := make(chan bool)
	go clock(flag)
	flag <- false
	var result []Node = IDS(bruh, "Joko_Widodo", 2, true)
	flag <- true
	//MUTEX BABEY
	//var thing []string = scrape("File:Diploma_icon.png")
	fmt.Println("Winners:")
	for _, item := range result {
		fmt.Println(item)
	}
	//fmt.Println(bruh.link[0])
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