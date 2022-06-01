package main

import (
	"fmt"
	"testing"
)

// TestBlockTree_AppendBlock is a unit test.
func TestBlockTree_AppendBlock(t *testing.T) {
	bt := NewBlockTree()
	bt.AppendBlockByNumber(genesisBlock)
	if len(bt[0]) == 0 {
		t.Fatal("missing genesis at index=0")
	}
	bt.AppendBlockByNumber(&Block{
		i: 1,
	})
	if len(bt[1]) == 0 {
		t.Fatal("missing block i=1 at index=1")
	}
}

func TestDistributeHashrates(t *testing.T) {
	res := generateMinerHashrates(HashrateDistLongtail, 12)
	print := "\n"
	for _, r := range res {
		print += fmt.Sprintf("%0.3f\n", r)
	}
	t.Log(print)
}
