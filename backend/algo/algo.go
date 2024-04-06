package main

type Node struct {
	name  string
	trail []string
	link  []string
	depth int
}

// nanti diganti links dengan implementasi scrapper
func makeNode(start string, links []string) *Node {
	return &Node{
		link:  links,
		depth: 0,
		name:  start,
	}
}

// func IDS(Node actor,name target, maxdepth int) []Node{
// 	//seek
// 	if (actor.name = name || )

// 	//check

// }
