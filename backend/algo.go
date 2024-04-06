package main

type Node struct {
	name  string
	trail []string
	link  []string
	depth int
}

func makeNode(start string, _depth int, _trail []string) *Node {
	return &Node{
		link:  scrape(start),
		depth: _depth,
		name:  start,
		trail: _trail,
	}
}

// INITIAL NODE, parent = makeNode("Test",0,{})
//
// MAKE CHILD, child = makeNode(parent.link[1],parent.depth + 1, append(parent.trail, parent.name))
// alternatively. to make it simple, makechild(actor Node, index int)

// fungsi ini tidak membuat semua Node children sekaligus, sehingga imo lebih ringan
func makeChildAlt(parent Node, index int) *Node {
	if index < len(parent.link) {
		return nil
	} else {
	return	makeNode(parent.link[index], parent.depth + 1, append(parent.trail,parent.name))
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
