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

import (
	"fmt"
)

type Node struct {
	Acc  int
	Bak  int
	Last string
	Code []uint8
	PC   int
}

type Port struct {
	Node1     int
	Node2     int
	PortName1 string
	PortName2 string
	Value     int
}

type Number int

type Reg string

type PortName string

var Ports []Port

var Nodes []Node

func findPort(n int, port string) (int, error) {
	for idx := range Ports {
		if Ports[idx].Node1 == n && Ports[idx].PortName1 == port {
			return idx, nil
		}

		if Ports[idx].Node2 == n && Ports[idx].PortName2 == port {
			return idx, nil
		}
	}

	return 0, fmt.Errorf("Port '%s' not found", port)
}

func Swp(n int) {
	Nodes[n].Acc, Nodes[n].Bak = Nodes[n].Bak, Nodes[n].Acc
}

func Sav(n int) {
	Nodes[n].Bak = Nodes[n].Acc
}

func AddNum(n, a int) {
	Nodes[n].Acc += a
}

func SubNum(n, a int) {
	Nodes[n].Acc -= a
}

func Nop(n int) {
}

func Neg(n int) {
	Nodes[n].Acc = -Nodes[n].Acc
}

func MovNumReg(n int, src Number, dst Reg) {
	// TODO types
	Nodes[n].Acc = int(src)
}

func MovNumPort(n int, src Number, dst PortName) error {
	// TODO types
	idx, err := findPort(n, string(dst))
	if err != nil {
		return err
	}

	Ports[idx].Value = int(src)

	return nil
}

func MovPortPort(n int, src PortName, dst PortName) {

}

func MovPortNil(n int, src PortName) {
	// read from port and discard
}

func MovNilPort(n int, dst PortName) {
	MovNumPort(n, 0, dst)
}

func main() {
	Nodes = []Node{{}, {}, {}, {}}
	Ports = []Port{
		{0, 1, "Right", "Left", 0},
		{0, 2, "Down", "Up", 0},
		{1, 3, "Down", "Up", 0},
		{2, 3, "Right", "Left", 0},
	}
}
