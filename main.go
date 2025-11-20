/*

 - check numbers: -999..999


===[ INSTRUCTION SET ]===

5 bits for op code
10 bits for numbers


---[ PORTS ]---

000 NIL
001 Left
010 Right
011 Up
100 Down
101 Any
110 Last


---[ INSTRUCTIONS ]---

00000000 NOP
00001000 SWP
00010000 SAV
00011000 NEG

00100000 MOV number to ACC
00101000 MOV port to port
00110000 MOV ACC to port
00111000 MOV port to ACC

01000000 ADD number
01001000 ADD port

01010000 SUB number
01011000 SUB port

01100000 JMP
01101000 JEZ
01110000 JNZ
01111000 JGZ
10000000 JLZ
10001000 JRO offset
10010000 JRO ACC


---[ EXAMPLES ]---

MOV 33, ACC
00100000 00100001

*/

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

