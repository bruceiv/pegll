//Incomplete
//look into else if - line 63
//and several other times including line 169-173
//return errors?
/// need to figure out if need errors for sum / product
package main

import (
	"fmt"
	"strconv"

	"calc/lexer"
	"calc/parser"
	"calc/parser/bsr"
	//"testing"
)

//Should match
const plus_test = `8 + 6` // passed
const minus_test = `8 - 6`
const mult_test = `8 * 6` // passed
const mult_ops_test1 = "8 + 16 * 24" // passed

//Should fail to match
const f = "12 +"

//Calculates the value of the input
//EXPR : WS SUM
func calculate(c bsr.BSR) int {
	// get the sum from the root (EXPR - expression )
	val := sum(c.GetNTChildI(1)) //sum
	return val
}

//SUM : PRODUCT RepPLUSorMINUS0x ;
func sum(s bsr.BSR) int {
	//PRODUCT child of SUM
	prod_node := s.GetNTChildI(0)
	//get the product
	val := product(prod_node)

	//RepPLUSorMINUS0x token
	pORmrep_node := s.GetNTChildI(1)

	// plus or minus child of sum
	return repPLUSorMINUS(val, pORmrep_node)

}

//RepPLUSorMINUS0x : PLUSorMINUS RepPLUSorMINUS0x
//                 / empty ;
//PLUSorMINUS      : PLUS PRODUCT
//                 | MINUS PRODUCT ;
func repPLUSorMINUS(val int, pORmrep bsr.BSR) int {

	//empty alternate - base case
	if pORmrep.Alternate() == 1 {
		return val //If rep is done, return value of product
	}
	//PLUSorMINUS NT
	pORm := pORmrep.GetNTChildI(0)

	//PRODUCT aspect of operator
	prod := pORm.GetNTChildI(1)

	//self-assignment aspect of RepPLUSorMINUS0x NT
	repChild := pORmrep.GetNTChildI(1)

	//alt0 -> addition
	if pORm.Alternate() == 0 {
		return repPLUSorMINUS((val + product(prod)), repChild)
	}

	//alt1 -> subtraction
	if pORm.Alternate() == 1 {
		return repPLUSorMINUS((val - product(prod)), repChild)
	}

	// panic if error occurs
	panic(fmt.Sprintf("Error in RepPLUSorMINUS0x: %d \t %T", val, pORmrep ))

}

//PRODUCT         : ELEMENT RepTIMESorDIV0x ;
func product(p bsr.BSR) int {
	//ELEMENT child of PRODUCT
	elem_node := p.GetNTChildI(0)

	//get element
	val := element(elem_node)

	//RepTIMEorDIVIDE0x token
	tORdrep_node := p.GetNTChildI(1)

	// times or divide child of sum
	return repTIMESorDIV(val, tORdrep_node)
}

//RepTIMESorDIV0x : TIMESorDIVIDE RepTIMESorDIV0x
//      		  / empty ;
//TIMESorDIVIDE   : TIMES ELEMENT
//  		      | DIVIDE ELEMENT ;
func repTIMESorDIV(val int, tORdrep bsr.BSR) int {

	//empty alternate - base case
	if tORdrep.Alternate() == 1 {
		return val //If rep is done, return value of element
	}

	//TIMESorDIVIDE NT
	tORd := tORdrep.GetNTChildI(0)

	//ELEMENT aspect of operator
	elem := tORd.GetNTChildI(1)

	//self-assignment aspect of RepTIMESorDIVIDE0x NT
	repChild := tORdrep.GetNTChildI(1)

	//alt0 -> multiplication
	if tORd.Alternate() == 0 {
		return repTIMESorDIV((val * element(elem)), repChild)
	}

	//alt1 -> division
	if tORd.Alternate() == 1 {
		return repTIMESorDIV((val / element(elem)), repChild)
	}

	// panic if error occurs
	panic(fmt.Sprintf("Error in RepTIMESorDIV0x: %d \t %T", val, tORdrep ))
}

//ELEMENT : OPEN SUM CLOSE
//        | Number ;
func element(e bsr.BSR) int {
	//Alt1 - OPEN SUM CLOSE
	if e.Alternate() == 0 {
		//Get SUM NT
		su := e.GetNTChildI(1)
		//Calculate SUM
		val := sum(su)
		//return value of SUM
		return val
	}

	//Alt2 - Number
	if e.Alternate() == 1 {
		//Get Number NT
		num := e.GetNTChildI(0)
		//Get value of Number
		val := number(num)
		// return value of Number
		return val
	}
	// panic if error occurs 
	panic(fmt.Sprintf("Error in element %T", e))
}

//Number : repNumber1x WS ;
//repNumber1x : < number > ;
func number(n bsr.BSR) int {
	// get the terminal child of the number node
	// repNumber1x : < number > 
	num_node := n.GetTChildI(0)

	// convert that terminal to a string
	num_string := num_node.LiteralString()

	//testing
	//fmt.Println(num_string)

	// convert the string version of the number to numberic
	num_digits, err := strconv.Atoi(num_string)

	// return the numeric version if no error
	if err == nil {
		return num_digits
	}

	// otherwise, panic with error
	panic("Error occurred: " + err.Error() )
}

func parse(s []rune) bool {
	// run GLL parser
	bsrSet, _ := parser.Parse(lexer.New(s))
	// quit early if parse fails
	if bsrSet == nil {
		return false
	}

	//Run calculation (functions above)
	value := calculate(bsrSet.GetRoot())
	fmt.Println("The calculation result is:", value)

	//bsrSet.Dump()
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
	//parseAndPrint(plus_test)
	//parseAndPrint(mult_test)
	parseAndPrint(mult_ops_test1)
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
