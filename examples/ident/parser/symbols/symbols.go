
// Package symbols is generated by gogll. Do not edit.
package symbols

type Symbol interface{
	isSymbol()
	IsNonTerminal() bool
	IsLookahead() bool
	String() string
}

func (NT) isSymbol() {}
func (T) isSymbol() {}
func (L) isSymbol() {}

// NT is the type of non-terminals symbols
type NT int
const( 
	NT_IdChar NT = iota
	NT_Ident 
	NT_Keyword 
	NT_RepidChar0x 
)

const NumNTs = 4

type NTs []NT

// T is the type of terminals symbols
type T int
const( 
	T_0 T = iota // f 
	T_1  // i 
	T_2  // idChar 
	T_3  // o 
	T_4  // r 
)

// L is the type of lookahead symbols
type L int
const( 
	LN_NT_Keyword L = iota
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

func (L) IsNonTerminal() bool {
	return false
}

func (NT) IsLookahead() bool {
	return false
}

func (T) IsLookahead() bool {
	return false
}

func (L) IsLookahead() bool {
	return true
}

func (nt NT) String() string {
	return ntToString[nt]
}

func (t T) String() string {
	return tToString[t]
}

func (lk L) String() string {
	if lk.IsNegative() {
		return "!" + lk.ArgSymbol().String()
	} else {
		return "&" + lk.ArgSymbol().String()
	}
}

func (nt NT) LeftRec() NTs {
	return leftRec[nt]
}

func (nt NT) IsOrdered() bool {
	return ordered[nt]
}

const(
	negTerm    = 0
	negNonterm = 1
	posTerm    = 2
	posNonterm = 3
	isNonterm  = 1
	isPos      = 2
)

func (lk L) IsNegative() bool {
	return lkMode[lk] & isPos == 0
}

func (lk L) IsPositive() bool {
	return lkMode[lk] & isPos != 0
}

func (lk L) ArgSymbol() Symbol {
	switch lkMode[lk] & isNonterm {
	case 0: // terminal
		return T(lkSym[lk])
	case 1: // nonterminal
		return NT(lkSym[lk])
	default:
		panic("Invalid lookahead")
	}
}

var ntToString = []string { 
	"IdChar", /* NT_IdChar */
	"Ident", /* NT_Ident */
	"Keyword", /* NT_Keyword */
	"RepidChar0x", /* NT_RepidChar0x */ 
}

var tToString = []string { 
	"f", /* T_0 */
	"i", /* T_1 */
	"idChar", /* T_2 */
	"o", /* T_3 */
	"r", /* T_4 */ 
}

var stringNT = map[string]NT{ 
	"IdChar":NT_IdChar,
	"Ident":NT_Ident,
	"Keyword":NT_Keyword,
	"RepidChar0x":NT_RepidChar0x,
}

var leftRec = map[NT]NTs { 
	NT_IdChar: NTs {  },
	NT_Ident: NTs {  NT_IdChar,  NT_Keyword,  },
	NT_Keyword: NTs {  },
	NT_RepidChar0x: NTs {  NT_IdChar,  },
}

var ordered = map[NT]bool { 
	NT_RepidChar0x:true,
}

var lkMode = []int { 
	negNonterm, 
}

var lkSym = []int { 
	int(NT_Keyword), 
}