package main

import (
	"fmt"
	"miniegg/lexer"
	"miniegg/parser"
)

// Should match
const ab = "A=B C=D"

const ri = "R = I O E E = I N"

//Should fail to match
const ff = "F=F F="

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
	//parseAndPrint(ri)
	//parseAndPrint(ff)
}
