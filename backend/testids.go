package main

import (
	"fmt"
	"time"
)

//muneki's main
// func main() {
// 	var bruh Node = *makeNode("Institut_Teknologi_Bandung", scrape("Institut_Teknologi_Bandung"))
// 	flag := make(chan bool)
// 	go clock(flag)
// 	flag <- false
// 	var result []Node = IDS(bruh, "Laos", 4)
// 	flag <- true
// 	//MUTEX BABEY
// 	//var thing []string = scrape("File:Diploma_icon.png")
// 	fmt.Println("Winners:")
// 	for _, item := range result {
// 		fmt.Println(item)
// 	}
// 	//fmt.Println(bruh.link[0])
// }

// Sean's main
func main() {
	flag := make(chan bool)
	go clock(flag)
	flag <- false
	result, someInt := BFS("Nintendo_Entertainment_System", "Slot_machine")
	flag <- true
	//MUTEX BABEY
	//var thing []string = scrape("File:Diploma_icon.png")
	fmt.Println("Winners:")
	for i := 0; i < len(result); i++ {
		printNode(result[i])
	}
	fmt.Println("Ammount of links checked : ", someInt)
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
			}
		}

	}
	fmt.Printf("Done, time taken : %d.%02d seconds \n", seconds, ms-seconds*100)
}
