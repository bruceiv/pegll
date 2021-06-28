package main

import (
	"fmt"

	"Java/lexer"
	"Java/parser"
)

//Should match
const t1 = `class HelloWorld {
    public static void main(String[] args) {
        System.out.println("Hello, World!"); 
    }
}`

//Should fail to match
const t2 = `println("Hello")`

func parse(s []rune) bool {
	// run GLL parser
	bsrSet, _ := parser.Parse(lexer.New(s))
	// quit early if parse fails
	if bsrSet == nil {
		return false
	}
	// check that root covers whole input
	root := bsrSet.GetOrderedRoot()
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
