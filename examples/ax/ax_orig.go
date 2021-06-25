package main

import (
	"fmt"

	"ax/lexer"
	"ax/parser"
)

// test the repeatability of the grammar

// should match
const a = `a`
const aa = `aa`
const aaa = `aaa`

// should fail to match
const ab = `ab`

// use the GetRoot(s) function from bsr.go
func parse(s []rune) bool {
	// run GLL parser
	bsrSet, _ := parser.Parse(lexer.New(s))
	// quit early if parse fails
	if bsrSet == nil {
		return false
	}
	// check that root covers whole input
	root := bsrSet.GetRoots()
	return root.LeftExtent == bsrSet.GetLeftExtent()
}

func parseAndPrint(s string) {
	if parse([]rune(s)) {
		fmt.Println("`" + s + "` matched")
	} else {
		fmt.Println("`" + s + "` DID NOT match")
	}
}

func main() {
	parseAndPrint(a)
	parseAndPrint(aa)
	parseAndPrint(aaa)
	parseAndPrint(ab)
}
