package main

type Port struct {
	Name string
	Enabled bool
	Ready bool
}

type Node struct {
	Acc int
	Bak int
	Last string
	Ports []Port
}



func Swp(n *Node) {
	n.Acc, n.Bak = n.Bak, n.Acc
}

func Sav(n *Node) {
	n.Bak = n.Acc
}

func Neg(n *Node) {
	n.Acc = -n.Acc
}

func main() {
}

