// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	var i int

// 	fmt.Print("Type a number: ")
// 	fmt.Scan(&i)
// 	if i == 0 {
// 		idsTest()
// 	} else {
// 		bfsTest()
// 	}
// }

// var start string = "Basketball"
// var tujuan string = "Joko_Widodo"

// func idsTest() {
// 	var bruh Node = *makeNode(start, 0, []string{})
// 	fmt.Println("Starting")
// 	flag := make(chan bool)
// 	go clock(flag)
// 	flag <- false
// 	var result []Node = IDS(bruh, tujuan, 4, true)
// 	flag <- true
// 	//MUTEX BABEY
// 	//var thing []string = scrape("File:Diploma_icon.png")
// 	fmt.Println("Winners:")
// 	for i := 0; i < len(result); i++ {
// 		printNode(result[i])
// 	}
// 	//fmt.Println(bruh.link[0])
// }

// // Sean's main
// func bfsTest() {
// 	counter = &visits
// 	flag := make(chan bool)
// 	go clock(flag)
// 	flag <- false
// 	result, someInt := BFS(start, tujuan)
// 	flag <- true
// 	//MUTEX BABEY
// 	//var thing []string = scrape("File:Diploma_icon.png")
// 	fmt.Println("Winners:")
// 	for i := 0; i < len(result); i++ {
// 		printNode(result[i])
// 	}
// 	fmt.Println("Ammount of links checked : ", someInt)
// }

// func printNode(_node Node) {
// 	for i := 0; i < len(_node.trail); i++ {
// 		fmt.Println(_node.trail[i])
// 	}
// }

// func clock(flag chan bool) {
// 	var ms int = 0
// 	var seconds int = 0
// 	stop := <-flag
// 	for !stop {
// 		if !stop {
// 			time.Sleep(10 * time.Millisecond)
// 			ms = ms + 1
// 		}
// 		select {
// 		case newstop := <-flag:
// 			stop = newstop
// 		default:
// 			stop = stop
// 			if ms/100 > seconds {
// 				seconds = ms / 100
// 				fmt.Println(seconds, visits)
// 			}
// 		}

// 	}
// 	fmt.Printf("Done, time taken : %d.%02d seconds \n", seconds, ms-seconds*100)
// }
