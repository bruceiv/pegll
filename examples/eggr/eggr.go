package main

import (
	"fmt"

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
