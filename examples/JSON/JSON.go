package main

import (
	"fmt"

	"JSON/lexer"
	"JSON/parser"
)

//Matches
const bracket_test 	= `{ }`
const simple_test 	= `{ "name" : "John" }`
const num_test		= `{ "num" : "N123" }` 
const empty_test = `{ "num" : "" }` 
	// issue with reading numbers
	// issue with empty string 

//Doesn't Match
//const t2 = `123`
//const t3 = `not JSON ~~`
 // doesn't accept repeated letters
//const t4 = `{ "name" : "Johhn" }` 

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
	parseAndPrint(bracket_test)
	parseAndPrint(simple_test)
	parseAndPrint(num_test)
	parseAndPrint(empty_test)
	//parseAndPrint(t2)
	//parseAndPrint(t3)
	//parseAndPrint(t4) 
}
