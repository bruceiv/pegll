/* ax.go
* Test the repeatability of the grammar 
*/
// import the ax package 
package ax

// import the parser files generated
import (
	"ax/lexer"
	"ax/parser"
	"ax/token"
	"testing"
	"strings"
)


// testing repeatability of the character a 
// last test should fail to match
const (
	rep_a1 	= `a`
	//rep_a2 	= `aa`
	//rep_a3 	= `aaa`
	//non_rep	= `b`
)


// test function for repeatability 
//	Repa0x : repa0x  ;
// 	repa0x : { 'a' } ;

func Test1( t *testing.T )  {
	bs, errs := parser.Parse(lexer.New([]rune(rep_a1)))

	// determine if any errors exist 
	if len( errs ) != 0 {
		t.Fail()
	}
	
	// Repa0x : repa0x
	root := bs.GetRoot()

	// repa0x : { 'a' }
	a_test := root.GetNTChildI(0)
	if rep_a1 != a_test.LiteralString() {
		t.Fail()
	}
}

