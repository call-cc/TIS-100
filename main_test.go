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
