package main

import (
	"fmt"
	"strings"
	"tubes2/crawl"
)

func hasJank(text string) bool {
	if strings.Contains(text, ":") {
		return true
	} else {
		return false
	}
}

var IDSvisits int = 0

func IDSWrapper(start string, end string) ([]Node, int) {
	var bruh Node = *makeNode(start, 0, []string{})
	var result []Node
	var i int = 1
	fmt.Println("SEEKING", start, end)
	for len(result) == 0 && i < 6 {
		fmt.Println("Current depth : ", i)
		IDSvisits = 0
		result = IDS(bruh, end, i, true)
		i++
	}
	return result, IDSvisits
}

func IDS(actor Node, target string, maxdepth int, firstonly bool) []Node {
	IDSvisits++
	//define
	var retval []Node //rerturn value
	//check if last or found
	if actor.name == target {
		if actor.name == target {
			actor.trail = append(actor.trail, actor.name)
			retval = append(retval, actor)
			fmt.Println("Found", target, "in", actor.depth, "by", actor.trail)
			if firstonly == true {
				return retval
			}
		}
	} else if actor.depth < maxdepth { //seek
		//for each link in actor.link
		for i := 0; i < len(actor.link); i++ {
			//fmt.Println("Current[", actor.depth+1, "] : ", actor.link[i])
			//fmt.Printf("Current[%d] : %s ", actor.depth+1, actor.link[i])

			//skip if link has been visitted before within self history to prevent useless looping
			for _, oldlink := range actor.trail {
				if actor.link[i] == oldlink {
					//fmt.Printf("%s == %s, skipping\n", actor.link[i], oldlink)
					continue
				}
			}

			if !hasJank(actor.link[i]) {
				var child Node = *newNode(actor.link[i], actor.depth+1, append(actor.trail, actor.name))
				if child.depth < maxdepth {
					child.link = crawl.Scrape(actor.link[i]) //fill with link in child
					//fmt.Printf("|| Scrapped!")
				}
				// Itterate over result of IDS, if empty then stop
				for _, item := range IDS(child, target, maxdepth, firstonly) {
					if firstonly == true {
						return append(retval, item)
					}
					retval = append(retval, item)
				}
			} else {
				//fmt.Printf(" || JANK !, SKIP")
			}
			//fmt.Println()

		}
	}
	//check

	return retval
}
