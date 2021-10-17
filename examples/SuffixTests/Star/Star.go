//Everything matches and fails as expected
package main

import (
	"Star/lexer"
	"Star/parser"
	"fmt"
)

const (
	t0 = `Required`
	t1 = `Required Base`
	t2 = `Required Base Base`
	t5 = `Required Base Base Base Base Base`
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
	parseAndPrint(t0)
	parseAndPrint(t1)
	parseAndPrint(t2)
	parseAndPrint(t5)
}
