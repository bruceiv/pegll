// NEED TO IMPLEMENT
package main

import (
	"fmt"

	"miniegg/lexer"
	"miniegg/parser"
)

// Should match
const ab = "AAA = BBB"
const eg = "EE = EE GG GGG"
//Should fail to match
const ff = "FF = FFF ="

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
	parseAndPrint(ab)
	parseAndPrint(eg)
	parseAndPrint(ff)
}