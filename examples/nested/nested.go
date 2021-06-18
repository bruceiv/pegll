// NEED TO IMPLEMENT
package main

import (
	"fmt"

	"nested/lexer"
	"nested/parser"
)

//Should match
const abcd = `abcd`
const parens = `()`
//Should fail to match
const unclosed = `((`
const nums = `123`

func parse(s []rune) bool {
	// run GLL parser
	bsrSet, _ := parser.Parse(lexer.New(s))
	// quit early if parse fails
	if bsrSet == nil {
		return false
	}
	// check that root covers whole input
	//root := bsrSet.GetOrderedRoot()
	//return root.RightExtent() == bsrSet.GetRightExtent();
}

func parseAndPrint(s string) {
	if parse([]rune(s)) {
		fmt.Println("`" + s + "` matched")
	} else {
		fmt.Println("`" + s + "` DID NOT match")
	}
}

func main() {
	parseAndPrint(abcd)
	parseAndPrint(parens)
	parseAndPrint(unclosed)
	parseAndPrint(nums)
}
