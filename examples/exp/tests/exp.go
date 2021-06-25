package exp

import (
	"exp/lexer"
	"exp/parser"
	"testing"
)

const (
	//Should match
	ab = "ab"
	//aabc = "aabc"
	//Should fail to match
 	//acc = "acc"
)

/*
func parse(s []rune) bool {
	// run GLL parser
	bsrSet, _ := parser.Parse(lexer.New(s))
	// quit early if parse fails
	if bsrSet == nil {
		return false
	}
	// check that root covers whole input
	root := bsrSet.GetRoot()
	return root.RightExtent() == bsrSet.GetRightExtent()
}

func parseAndPrint(s string) {
	if parse([]rune(s)) {
		fmt.Println("`" + s + "` matched")
	} else {
		fmt.Println("`" + s + "` DID NOT match")
	}
}

func main() {
	parseAndPrint(ab)
	parseAndPrint(aabc)
	parseAndPrint(acc)
}
 */

// testing ab
// untested
func Test1(t *testing.T) {
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


/* // testing aabc
// untested
func Test2(t *testing.T) {
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

// testing acc
// untested
func Test3(t *testing.T) {
	bs, errs := parser.Parse(lexer.New([]rune(acc)))
	if len(errs) != 0 {
		t.Fail()
	}

	// get the root
	root := bs.GetRoot()
	// get the child to test
	a := root.GetTChildI(0)
	if acc != a.LiteralString() {
		t.Fail()
	}
}
 */
