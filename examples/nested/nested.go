package main

import (
	"fmt"

	"nested/lexer"
	"nested/parser"
)

//Matches
const abcd = `abcd`
const parens = `()`
const nested_parens = `( a ( b ( c ) ) )`

//Doesn't Match
const unclosed = `((`
const nums = `123`
const nested_parens2 = `( a ( 1 ( 2 ) ) )` //  doesn't match numbers based on grammar

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
	parseAndPrint(abcd)
	parseAndPrint(parens)
	parseAndPrint(nested_parens)
	parseAndPrint(unclosed)
	parseAndPrint(nums)
	parseAndPrint(nested_parens2)
}
