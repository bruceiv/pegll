package main

import (
	"fmt"

	"ident/lexer"
	"ident/parser"
)

const ifs = `if`
const ifxs = `ifx`
const gos = `go`

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
	parseAndPrint(ifs)
	parseAndPrint(ifxs)
	parseAndPrint(gos)
}
