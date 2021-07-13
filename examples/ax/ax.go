// go test file
package main

import (
	"fmt"

	"ax/lexer"
	"ax/parser"
)

// test the repeatability of the grammar

// Match
const a = "a"
const aa = "aa"
const aaa = "aaa"

// Loops infinitely
//const ab = "ab"
//const aba = "aba"

//Fails to match
const space = " "
const nothing = ""

// use the GetRoot(s) function from bsr.go
func parse(s []rune) bool {
	// run GLL parser
	bsrSet, _ := parser.Parse(lexer.New(s))
	// quit early if parse fails
	if bsrSet == nil {
		return false
	}
	// print tree in pre-order
	//bsrSet.FlatDump()
	//bsrSet.FilterByOrderedChoice()
	//fmt.Println("=====")
	//bsrSet.FlatDump()
	//fmt.Println("=====")
	// check that root covers whole input
	root := bsrSet.GetRoot()
	return root.RightExtent() == bsrSet.GetRightExtent()
}

func parseAndPrint(s string) {
	if parse([]rune(s)) {
		fmt.Println("'" + s + "' matched")
	} else {
		fmt.Println("'" + s + "' DID NOT match")
	}
}

func main() {
	parseAndPrint(a)
	parseAndPrint(aa)
	parseAndPrint(aaa)
	//parseAndPrint(ab)
	//parseAndPrint(aba)
	parseAndPrint(space)
	parseAndPrint(nothing)
}
