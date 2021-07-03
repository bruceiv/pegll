//Incomplete
package main

import (
	"fmt"

	"calc/lexer"
	"calc/parser"
	//"testing"
)

//Should match
const a = "908 + 9999999 * 912341"

//Should fail to match
const f = "12 +"

/*
Need a way to recognize each important NT (function for each one)
	-> SUM, PRODUCT, ELEM,
	-> maybe TIMESorDIVIDE and PLUSorMINUS

Need to go from recognizing the NT to actually performing
the operation (???)
	- Perform mults first

Might be helpful:
	bsr.getTchild
		.getNTchild
	   .isNonTerminal
	   .dump (for testing)
*/

//EXPR : space SUM
func calculate( /*bsr node I think*/ ) int {
	//Traverse down to node containing number
	//Navigate alternates to see the computation performed on it
	//Send into a coordinating function that manages the operation

	//if x is product then product()
	//same for element, Sum

	//Ooooorrrr
	/// get the number and convert to a string

	//interelate the functions for each important NT to call eachother
	//sum calls product calls element (sometimes) calls sum
	//This is a better idea I think
	//sum := b.GetNTChildI(1) //sum

	return 0 //Temp, will return calculation
}

//SUM : PRODUCT RepPLUSorMINUS0x      ;
func sum( /*bsr node I think*/ ) int {
	//identify product, call that function
	//identify if operation is happening (RepPLUSorMINUS0x)
	//if yes, find the alternate (+/-) and call product again
	//afterwards, perform computation

	// return product(node) ( alternate product(node) )?
	//how do we check if repPlusorMinus is called?
	return 0 //temp, will eventually return computational value
}

func parse(s []rune) bool {
	// run GLL parser
	bsrSet, _ := parser.Parse(lexer.New(s))
	// quit early if parse fails
	if bsrSet == nil {
		return false
	}

	//calculate(bsrSet)
	bsrSet.Dump()
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
	parseAndPrint(f)
}

/*func Test1(t *testing.T) { //Match test
	bsrSet, errs := parser.Parse(lexer.New([]rune(a)))
	if len(errs) > 0 {
		fail(errs)
	}

	if bsrSet == nil {
		panic("BSRSet == nil")
	}

	for i, r := range bsrSet.GetRoots() {
		fmt.Printf("%d: %s\n", i, calculate(r))
	}
}

func Test2(t *testing.T) { //Fail to match test
	bsrSet, errs := parser.Parse(lexer.New([]rune(f)))
	if len(errs) > 0 {
		fail(errs)
	}

	if bsrSet == nil {
		panic("BSRSet == nil")
	}

	for i, r := range bsrSet.GetRoots() {
		fmt.Printf("%d: %s\n", i, calculate(r))
	}
}
*/
/*
Need a way to recognize each important NT (function for each one)
	-> SUM, PRODUCT, ELEM,
	-> maybe TIMESorDIVIDE and PLUSorMINUS

Need to go from recognizing the NT to actually performing
the operation (???)
	- Perform mults first

Might be helpful:
	bsr.getTchild
		.getNTchild
	   .isNonTerminal
	   .dump (for testing)
*/
/*

func fail(errs []*parser.Error) {
	ln := errs[0].Line
	fmt.Println("Parse Error:")
	for _, e := range errs {
		if e.Line == ln {
			fmt.Println(e)
		}
	}
	os.Exit(1)
}*/
