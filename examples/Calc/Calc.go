// NEED TO IMPLEMENT
package main

import (
	"fmt"
	"os"

	"calc/lexer"
	"calc/parser"
	"testing"
)

//Should match
const a = "1 + 2 * 3"

//Should fail to match
const f = "12 +"

func Test1(t *testing.T) { //Match test
	bsrSet, errs := parser.Parse(lexer.New([]rune(a)))
	if len(errs) > 0 {
		fail(errs)
	}

	if bsrSet == nil {
		panic("BSRSet == nil")
	}

	for i, r := range bsrSet.GetRoots() {
		fmt.Printf("%d: %s\n", i, calculate(r))
	}
}

func Test2(t *testing.T) { //Fail to match test
	bsrSet, errs := parser.Parse(lexer.New([]rune(f)))
	if len(errs) > 0 {
		fail(errs)
	}

	if bsrSet == nil {
		panic("BSRSet == nil")
	}

	for i, r := range bsrSet.GetRoots() {
		fmt.Printf("%d: %s\n", i, calculate(r))
	}
}

/*
Need a way to recognize each important NT (function for each one)
	-> SUM, PRODUCT, ELEM,
	-> maybe TIMESorDIVIDE and PLUSorMINUS

Need to go from recognizing the NT to actually performing
the operation (???)
	- Perform mults first

Might be helpful:
	bsr.getTchild
		.getNTchild
	   .isNonTerminal
	   .dump (for testing)
*/
func calculate(b bsr.BSR) int {

}

func fail(errs []*parser.Error) {
	ln := errs[0].Line
	fmt.Println("Parse Error:")
	for _, e := range errs {
		if e.Line == ln {
			fmt.Println(e)
		}
	}
	os.Exit(1)
}
