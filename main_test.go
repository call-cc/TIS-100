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

	oldAcc := Nodes[0].Acc
	oldBak := Nodes[0].Bak

	Swp(0)

	if oldBak != Nodes[0].Acc {
		t.Errorf("old BAK %d != new ACC %d", oldBak, Nodes[0].Acc)
	}

	if oldAcc != Nodes[0].Bak {
		t.Errorf("old ACC %d != new BAK %d", oldAcc, Nodes[0].Bak)
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

	for _, test := range tests {
		Nodes[0].Acc = test[0]
		AddNum(0, test[1])
		if Nodes[0].Acc != test[2] {
			t.Errorf("ACC %d + %d != %d", Nodes[0].Acc, test[1], test[2])
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

	for _, test := range tests {
		Nodes[0].Acc = test[0]
		SubNum(0, test[1])
		if Nodes[0].Acc != test[2] {
			t.Errorf("ACC %d - %d != %d", Nodes[0].Acc, test[1], test[2])
		}
	}
}

func TestBak(t *testing.T) {
	Nodes = []Node{
		{
			Acc: 123,
		},
	}

	Sav(0)

	if Nodes[0].Bak != Nodes[0].Acc {
		t.Errorf("BAK %d != ACC %d", Nodes[0].Bak, Nodes[0].Acc)
	}
}

func TestNeg(t *testing.T) {
	vals := []int{-999, -100, 0, 50, 200}

	Nodes = []Node{
		{},
	}

	for _, val := range vals {
		Nodes[0].Acc = val
		Neg(0)

		if Nodes[0].Acc != -val {
			t.Errorf("ACC %d != -%d", Nodes[0].Acc, -val)
		}
	}
}

func TestFetchNextOp(t *testing.T) {
	Nodes = []Node{{}}

	Nodes[0].Code = []uint8{0, 0b11000, 0b1000}
	Nodes[0].PC = 1

	op := FetchNext(0)

	if Nodes[0].PC != 2 {
		t.Errorf("PC should be 2, not %d", Nodes[0].PC)
	}

	opReq := uint8(0b11000)
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
	Nodes[0].Code = []uint8{0, 0, 0}

	for _, test := range tests {
		Nodes[0].PC = test[0]
		FetchNext(0)

		if Nodes[0].PC != test[1] {
			t.Errorf("PC %d should be %d", Nodes[0].PC, test[1])
		}
	}
}
