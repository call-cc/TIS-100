/*

===[ INSTRUCTION SET ]===


---[ ENCODING ]---

5 bits for op codes
11 bits for numbers (-999 .. 999)
3 bits for port specification


Port1      Number        Unused
| |        |           | ||
000 00000  00000000  000 00 000
    |   |                   | |
    Opcode                  Port2

---[ PORTS ]---

000 NIL
001 Left
010 Right
011 Up
100 Down
101 Any
110 Last


---[ INSTRUCTIONS ]---

Code  Name [Length in bytes] (Cycles)
-----------------------------
00000 NOP [1] (1)
00001 SWP [1] (1)
00010 SAV [1] (1)
00011 NEG [1] (1)

00100 MOV number to ACC [2] (1)
00101 MOV port to port [2] (*)
00110 MOV ACC to port [1] (1)
00111 MOV port to ACC [1] (1)
01000 MOV number to port [3] (*)
01001 MOV port to number [3] (*)

(*) 1 if writing NIL, 2 if writing Up, Down, Left, or Right

01010 ADD number [2] (1)
01011 ADD port [1] (1)

01100 SUB number [2] (1)
01101 SUB port [1] (1)

01110 JMP label [2] (1)
01111 JEZ label [2] (1)
10000 JNZ label [2] (1)
10001 JGZ label [2] (1)
10010 JLZ label [2] (1)
10011 JRO offset [2] (1)
10100 JRO ACC [1] (1)

11111 HCF [1] (1)


---[ EXAMPLES ]---

MOV 33, ACC
00000100  00000100  00100000
   |   |  |           |
   MOV    33

*/

package main

import (
	"fmt"
	"os"
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

type OpCode struct {
	Code     uint8
	Name     string
	Function func()
}

var OpCodes = map[int]func(){
	0b00000: Nop,
	0b00001: Swp,
	0b00010: Sav,
	0b00011: Neg,
	0b11111: HCF,
}

var Ports []Port

var Nodes []Node

var CurrentNode int

var Cycle uint

type Number int

type Reg string

type PortName string

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

func Swp() {
	Nodes[CurrentNode].Acc, Nodes[CurrentNode].Bak = Nodes[CurrentNode].Bak, Nodes[CurrentNode].Acc
}

func Sav() {
	Nodes[CurrentNode].Bak = Nodes[CurrentNode].Acc
}

func AddNum(a int) {
	Nodes[CurrentNode].Acc += a
}

func SubNum(a int) {
	Nodes[CurrentNode].Acc -= a
}

func Nop() {
}

func Neg() {
	Nodes[CurrentNode].Acc = -Nodes[CurrentNode].Acc
}

func MovNumReg(src Number, dst Reg) {
	// TODO types
	Nodes[CurrentNode].Acc = int(src)
}

func MovNumPort(src Number, dst PortName) error {
	// TODO types
	idx, err := findPort(CurrentNode, string(dst))
	if err != nil {
		return err
	}

	Ports[idx].Value = int(src)

	return nil
}

func MovPortPort(src PortName, dst PortName) {

}

func MovPortNil(src PortName) {
	// read from port and discard
}

func MovNilPort(dst PortName) {
	_ = MovNumPort(0, dst)
}

func HCF() {
	fmt.Println("Halt & Catch Fire")
	os.Exit(0)
}

func GetOp(op uint8) (func(), error) {
	op = op & 0b11111
	fn, ok := OpCodes[int(op)]
	if ok {
		return fn, nil
	}

	return func() {}, fmt.Errorf("no instruction found: %0b", op)
}

func FetchNext() uint8 {
	pc := Nodes[CurrentNode].PC
	op := Nodes[CurrentNode].Code[pc]

	if pc == len(Nodes[CurrentNode].Code)-1 {
		Nodes[CurrentNode].PC = 0
	} else {
		Nodes[CurrentNode].PC += 1
	}

	return op
}

func Run() {
	count := 20
	for count > 0 {
		for CurrentNode = range Nodes {
			code := FetchNext()
			op, err := GetOp(code)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Printf("Executing: %08b\n", code)
			op()
		}
		count -= 1
		// Cycle += 1
	}
}

func main() {
	Nodes = []Node{{Acc: 42, Code: []uint8{2, 0, 0}}}
	Ports = []Port{
		{0, 1, "Right", "Left", 0},
		{0, 2, "Down", "Up", 0},
		{1, 3, "Down", "Up", 0},
		{2, 3, "Right", "Left", 0},
	}

	Run()
}
