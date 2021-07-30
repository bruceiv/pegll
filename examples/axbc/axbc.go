package main

import (
	"fmt"

	"axbc/lexer"
	"axbc/parser"
)

const aac = `aac`
const abc = `abc`
const aabc = `aabc`

func parse(s []rune) bool {
	// run GLL parser
	bsrSet, _ := parser.Parse(lexer.New(s))
	// quit early if parse fails
	if bsrSet == nil {
		return false
	}
	// Filter out results that violate ordered choice
	// bsrSet.FlatDump()
	//bsrSet.FilterByOrderedChoice()
	// fmt.Println("=====")
	// bsrSet.FlatDump()

	// check that single root covers whole input
	roots := bsrSet.GetRoots()
	switch len(roots) {
	case 0:
		fmt.Println("No solutions")
		return false
	case 1:
		return true
	default:
		fmt.Println("Ambiguous")
		return false
	}
}

func parseAndPrint(s string) {
	if parse([]rune(s)) {
		fmt.Println("`" + s + "` matched")
	} else {
		fmt.Println("`" + s + "` DID NOT match")
	}
}

func main() {
	parseAndPrint(aac)
	parseAndPrint(abc)
	// TODO investigate lexer (?) infinite-loop
 	parseAndPrint(aabc)
}
