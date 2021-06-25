package exp

//These tests all fail under gogll, possibly because of
//the middle recursion of the grammar

import (
	"exp/lexer"
	"exp/parser"
	"testing"
)

// declare the constants to test repeatability
const (
	//Should match
	ab   = "ab"
	aabc = "aabc"
	//Should fail to match
	acc = "acc"
)

// test ab
func Test1(t *testing.T) {
	bs, errs := parser.Parse(lexer.New([]rune(ab)))
	if len(errs) != 0 {
		t.Fail()
	}

	// get the root
	root := bs.GetRoot()

	// get the child to test
	a := root.GetTChildI(0)
	if ab != a.LiteralString() {
		t.Fail()
	}
}

// test aabc
func Test3(t *testing.T) {
	bs, errs := parser.Parse(lexer.New([]rune(aabc)))
	if len(errs) != 0 {
		t.Fail()
	}

	// get the root
	root := bs.GetRoot()

	// get the child to test
	a := root.GetTChildI(0)
	if aabc != a.LiteralString() {
		t.Fail()
	}
}

// test aac
func Test2(t *testing.T) {
	bs, errs := parser.Parse(lexer.New([]rune(aac)))
	if len(errs) != 0 {
		t.Fail()
	}

	// get the root
	root := bs.GetRoot()

	// get the child to test
	a := root.GetTChildI(0)
	if aac != a.LiteralString() {
		t.Fail()
	}
}
