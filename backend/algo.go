package main

type Node struct {
	name  string
	trail []string
	link  []string
	depth int
}

func makeNode(start string, links []string, _depth int, _trail []string) *Node {
	return &Node{
		link:  links,
		depth: _depth,
		name:  start,
		trail: _trail,
	}
}

// nanti diganti links dengan implementasi scrapper
func makeChildren(start Node) []Node {
	var tempStr []string = scrape(start.name)
	var result []Node
	for i := 0; i < len(tempStr); i++ {
		var tempNode Node
		var container []string
		tempNode = *makeNode(tempStr[i], container, start.depth+1, append(start.trail, start.name))
		result = append(result, tempNode)
	}
	return result
}

// func IDS(Node actor,name target, maxdepth int) []Node{
// 	//seek
// 	if (actor.name = name || )

// 	//check

// }
