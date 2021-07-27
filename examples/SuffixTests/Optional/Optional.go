package main

import (
	"Optional/lexer"
	"Optional/parser"
	"fmt"
)

const (
	//Both of these should match because base is optional
	with    = `Required Base Required` //Matches
	without = `Required Required`      //Currently doesn't match
)

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
		fmt.Println("`" + s + "` matched")
	} else {
		fmt.Println("`" + s + "` DID NOT match")
	}
}

func main() {
	parseAndPrint(with)
	parseAndPrint(without)
}
