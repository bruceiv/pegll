package main

import (
	"fmt"

	"XML/lexer"
	"XML/parser"
)

//Should match
const t0 = `<?xml version="1.0" encoding="UTF-8"?>`
const t1 = `<?xml version="1.0" encoding="UTF-8"?> 
<note>
<to> & Tove ; </to>
<from> & Jani ; </from>
<heading> & Reminder ; </heading>
<body> & Don't forget me this weekend! ; </body>
</note>`
const comment = `<!--Students grades are uploaded by months-->`
const t3 = `<?xml version="1.0" encoding="UTF-8"?> <note> & name ; </note>`

//Should fail to match
const t2 = `not XML ~~` //Infinite loop

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
	//parseAndPrint(t1)
	//parseAndPrint(t2)
	parseAndPrint(t3)
	parseAndPrint(comment)
}
