
// Package symbols is generated by gogll. Do not edit.
package symbols

type Symbol interface{
	isSymbol()
	IsNonTerminal() bool
	String() string
}

func (NT) isSymbol() {}
func (T) isSymbol() {}

// NT is the type of non-terminals symbols
type NT int
const( 
	NT_Expr NT = iota
	NT_ExprRep 
	NT_Grammar 
	NT_Rule 
	NT_RuleRep 
)

// T is the type of terminals symbols
type T int
const( 
	T_0 T = iota //   
	T_1  // eq 
	T_2  // id 
	T_3  // neq 
)

type Symbols []Symbol

func (ss Symbols) Strings() []string {
	strs := make([]string, len(ss))
	for i, s := range ss {
		strs[i] = s.String()
	}
	return strs
}

func (NT) IsNonTerminal() bool {
	return true
}

func (T) IsNonTerminal() bool {
	return false
}

func (nt NT) String() string {
	return ntToString[nt]
}

func (t T) String() string {
	return tToString[t]
}

var ntToString = []string { 
	"Expr", /* NT_Expr */
	"ExprRep", /* NT_ExprRep */
	"Grammar", /* NT_Grammar */
	"Rule", /* NT_Rule */
	"RuleRep", /* NT_RuleRep */ 
}

var tToString = []string { 
	" ", /* T_0 */
	"eq", /* T_1 */
	"id", /* T_2 */
	"neq", /* T_3 */ 
}

var stringNT = map[string]NT{ 
	"Expr":NT_Expr,
	"ExprRep":NT_ExprRep,
	"Grammar":NT_Grammar,
	"Rule":NT_Rule,
	"RuleRep":NT_RuleRep,
}