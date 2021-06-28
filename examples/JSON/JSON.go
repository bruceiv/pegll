package main

import (
	"fmt"

	"JSON/lexer"
	"JSON/parser"
)

//Matches
const t00 = `{}`
const t0 = `{ "name" : "John" }`

const t1 = `{ "name" : "John", "name" : "John" }` // doesn't accept repeated letters

//Doesn't Match
const t2 = `123`
const t3 = `not JSON ~~`
const t4 = `{ "name" : "Johhn" }`

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
	parseAndPrint(t00)
	parseAndPrint(t0)
	parseAndPrint(t1)
	parseAndPrint(t2)
	parseAndPrint(t3)
	parseAndPrint(t4)
}
