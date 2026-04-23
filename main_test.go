package main

import (
	"testing"
)

func TestSwp(t *testing.T) {
	Nodes = []Node{
		{
			Acc: 123,
			Bak: 456,
		},
	}

	CurrentNode = 0
	oldAcc := Nodes[CurrentNode].Acc
	oldBak := Nodes[CurrentNode].Bak

	Swp()

	if oldBak != Nodes[CurrentNode].Acc {
		t.Errorf("old BAK %d != new ACC %d", oldBak, Nodes[CurrentNode].Acc)
	}

	if oldAcc != Nodes[CurrentNode].Bak {
		t.Errorf("old ACC %d != new BAK %d", oldAcc, Nodes[CurrentNode].Bak)
	}
}

func TestAddNum(t *testing.T) {
	tests := [][3]int{
		{0, 0, 0},
		{0, 1, 1},
		{0, -1, -1},
		{-5, -10, -15},
		{12, 35, 47},
	}

	Nodes = []Node{{}}
	CurrentNode = 0

	for _, test := range tests {
		Nodes[CurrentNode].Acc = test[0]
		AddNum(test[1])
		if Nodes[CurrentNode].Acc != test[2] {
			t.Errorf("ACC %d + %d != %d", Nodes[CurrentNode].Acc, test[1], test[2])
		}
	}
}

func TestSubNum(t *testing.T) {
	tests := [][3]int{
		{0, 0, 0},
		{0, 1, -1},
		{-1, -1, 0},
		{-5, -3, -2},
		{10, 1, 9},
	}

	Nodes = []Node{{}}
	CurrentNode = 0

	for _, test := range tests {
		Nodes[CurrentNode].Acc = test[0]
		SubNum(test[1])
		if Nodes[CurrentNode].Acc != test[2] {
			t.Errorf("ACC %d - %d != %d", Nodes[CurrentNode].Acc, test[1], test[2])
		}
	}
}

func TestBak(t *testing.T) {
	Nodes = []Node{
		{
			Acc: 123,
		},
	}

	CurrentNode = 0
	Sav()

	if Nodes[CurrentNode].Bak != Nodes[CurrentNode].Acc {
		t.Errorf("BAK %d != ACC %d", Nodes[CurrentNode].Bak, Nodes[CurrentNode].Acc)
	}
}

func TestNeg(t *testing.T) {
	vals := []int{-999, -100, 0, 50, 200}

	Nodes = []Node{{}}
	CurrentNode = 0

	for _, val := range vals {
		Nodes[CurrentNode].Acc = val
		Neg()

		if Nodes[CurrentNode].Acc != -val {
			t.Errorf("ACC %d != -%d", Nodes[CurrentNode].Acc, -val)
		}
	}
}

func TestFetchNextOp(t *testing.T) {
	Nodes = []Node{{}}
	CurrentNode = 0

	Nodes[CurrentNode].Code = []uint8{0, 0b11, 0b1}
	Nodes[CurrentNode].PC = 1

	op := FetchNext()

	opReq := uint8(0b11)
	if op != opReq {
		t.Errorf("op should be %b, not %b", opReq, op)
	}
}

func TestFetchNextPC(t *testing.T) {
	tests := [][2]int{
		{0, 1},
		{2, 0},
	}

	Nodes = []Node{{}}
	CurrentNode = 0
	Nodes[CurrentNode].Code = []uint8{0, 0, 0}

	for _, test := range tests {
		Nodes[CurrentNode].PC = test[0]
		FetchNext()

		if Nodes[CurrentNode].PC != test[1] {
			t.Errorf("PC %d should be %d", Nodes[CurrentNode].PC, test[1])
		}
	}
}
