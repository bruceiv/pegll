
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
	NT_Base NT = iota
	NT_Rep 
	NT_Required 
	NT_S1 
	NT_SuffBase 
)

const NumNTs = 5

type NTs []NT

// T is the type of terminals symbols
type T int
const( 
	T_0 T = iota // Base 
	T_1  // Required 
)

// L is the type of lookahead symbols
type L int
const( 
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
	"Base", /* NT_Base */
	"Rep", /* NT_Rep */
	"Required", /* NT_Required */
	"S1", /* NT_S1 */
	"SuffBase", /* NT_SuffBase */ 
}

var tToString = []string { 
	"Base", /* T_0 */
	"Required", /* T_1 */ 
}

var stringNT = map[string]NT{ 
	"Base":NT_Base,
	"Rep":NT_Rep,
	"Required":NT_Required,
	"S1":NT_S1,
	"SuffBase":NT_SuffBase,
}

var leftRec = map[NT]NTs { 
	NT_Base: NTs {  },
	NT_Rep: NTs {  NT_SuffBase,  NT_Base,  },
	NT_Required: NTs {  },
	NT_S1: NTs {  NT_Required,  },
	NT_SuffBase: NTs {  NT_Base,  },
}

var ordered = map[NT]bool { 
	NT_SuffBase:true,
}

var lkMode = []int { 
}

var lkSym = []int { 
}