package main

import (
	"fmt"

	"JSON/lexer"
	"JSON/parser"
)

//Should match
const t1 = `{
    "fruit": "Apple",
    "size": "Large",
    "color": "Red"
}`

//Should fail to match
const t2 = `not JSON ~~`

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
	parseAndPrint(t1)
	parseAndPrint(t2)
}
