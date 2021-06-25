package eggr

import (
	"fmt"
	"testing"

	"eggr/lexer"
	"eggr/parser"
)

//Should match
const test1 = `C5 = & A 123 | B 4 B +`

//Should fail to match
const test2 = `A &`

func parse(s []rune) bool {
	// run GLL parser
	bsrSet, _ := parser.Parse(lexer.New(s))
	// quit early if parse fails
	if bsrSet == nil {
		return false
	}
	// check that root covers whole input
	root := bsrSet.GetOrderedRoot()
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
	parseAndPrint(test1)
	parseAndPrint(test2)
}

// test1
// untested
func Test1(t *testing.T) {
	bs, errs := parser.Parse(lexer.New([]rune(test1)))
	if len(errs) != 0 {
		t.Fail()
	}

	// get the root
	root := bs.GetRoot()
	// get the child to test
	a := root.GetTChildI(0)
	if test1 != a.LiteralString() {
		t.Fail()
	}
}

// test2
// untested
func Test2(t *testing.T) {
	bs, errs := parser.Parse(lexer.New([]rune(test2)))
	if len(errs) != 0 {
		t.Fail()
	}

	// get the root
	root := bs.GetRoot()
	// get the child to test
	a := root.GetTChildI(0)
	if test2 != a.LiteralString() {
		t.Fail()
	}
}
