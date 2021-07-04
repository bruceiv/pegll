package main

import (
	"fmt"

	"string/lexer"
	"string/parser"
)

// test strings for the JSON grammar
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
		fmt.Println( s + " matched")
	} else {
		fmt.Println( s + " DID NOT match")
	}
}

func main() {
	parseAndPrint(`"aa"`)
	parseAndPrint(`/* the */`)
	//parseAndPrint(no_quotes)
}

