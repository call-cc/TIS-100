package main

import (
	"testing"
)

func TestSwp(t *testing.T) {
	n := Node{
		Acc:123,
		Bak:456,
	}

	oldAcc := n.Acc
	oldBak := n.Bak

	Swp(&n)

	if oldBak != n.Acc {
		t.Errorf("old BAK %d != new ACC %d", oldBak, n.Acc)
	}

	if oldAcc != n.Bak {
		t.Errorf("old ACC %d != new BAK %d", oldAcc, n.Bak)
	}
}

func TestSav(t *testing.T) {
	n := Node{
		Acc:123,
	}

	Sav(&n)

	if n.Bak != n.Acc {
		t.Errorf("BAK %d != ACC %d", n.Bak, n.Acc)
	}
}

func TestNeg(t *testing.T) {
	vals := []int{-999, -100, 0, 50, 200}

	n := Node{}

	for _, val := range vals{
		n.Acc = val
		Neg(&n)

		if n.Acc != -val {
			t.Errorf("ACC %d != -%d", n.Acc, -val)
		}
	}
}
