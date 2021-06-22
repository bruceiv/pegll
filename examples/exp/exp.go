// NEED TO IMPLEMENT
package main

import (
	"fmt"

	"exp/lexer"
	"exp/parser"
)

//Should match
const ab = `ab`
const aabc = `aabc`

//Should fail to match
const acc = `acc`

func parse(s []rune) bool {
	// run GLL parser
	bsrSet, _ := parser.Parse(lexer.New(s))
	// quit early if parse fails
	if bsrSet == nil {
		return false
	}
	// check that root covers whole input
	root := bsrSet.GetRoots()
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
