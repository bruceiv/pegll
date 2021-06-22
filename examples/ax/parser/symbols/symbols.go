
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
	NT_S1 NT = iota
)

// T is the type of terminals symbols
type T int
const( 
	T_0 T = iota // ax 
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
	"S1", /* NT_S1 */ 
}

var tToString = []string { 
	"ax", /* T_0 */ 
}

var stringNT = map[string]NT{ 
	"S1":NT_S1,
}