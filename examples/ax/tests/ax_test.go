package main

import (
	"ax/lexer"
	"ax/parser"
	"testing"
)

// declare the constants to test repeatability
const (
	rep_a0 = " "
	rep_a1 = "a"
	rep_a2 = "aa"
	ab     = "ab"
)

/* INPUT STRING TESTS */
// rep_a0 test
// FAILS
/* func Test0(t *testing.T) {
	bs, errs := parser.Parse(lexer.New([]rune(rep_a0)))
	if len(errs) != 0 {
		t.Fail()
	}

	// get the root
	// Repa0x : repa0x ;
	root := bs.GetRoot()

	// get the child to test
	a := root.GetTChildI(0)
	if rep_a0 != a.LiteralString() {
		t.Fail()
	}
} */

// rep_a1 test
// Pass
func Test1(t *testing.T) {
	bs, errs := parser.Parse(lexer.New([]rune(rep_a1)))
	if len(errs) != 0 {
		t.Fail()
	}

	// get the root
	// Repa0x : repa0x ;
	root := bs.GetRoot()

	// get the child to test
	a := root.GetTChildI(0)
	if rep_a1 != a.LiteralString() {
		t.Fail()
	}
}

// test rep_a2
// Pass
func Test2(t *testing.T) {
	bs, errs := parser.Parse(lexer.New([]rune(rep_a2)))
	if len(errs) != 0 {
		t.Fail()
	}

	// get the root
	// Repa0x : repa0x ;
	root := bs.GetRoot()

	// get the child to test
	a := root.GetTChildI(0)
	if rep_a2 != a.LiteralString() {
		t.Fail()
	}
}

// test ab
// FAILS - PROGRAM NEVER ENDS
func Test_ab(t *testing.T) {
	bs, errs := parser.Parse(lexer.New([]rune(ab)))
	if len(errs) != 0 {
		t.Fail()
	}

	// get the root
	// Repa0x : repa0x ;
	root := bs.GetRoot()

	// get the child to test
	a := root.GetTChildI(0)
	if ab != a.LiteralString() {
		t.Fail()
	}
}
