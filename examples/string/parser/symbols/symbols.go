
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
	NT_EscOrComment NT = iota
	NT_String 
	NT_WS 
)

const NumNTs = 3

type NTs []NT

// T is the type of terminals symbols
type T int
const( 
	T_0 T = iota // block_comment 
	T_1  // escCharSpace 
	T_2  // line_comment 
	T_3  // string_ns 
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

func (nt NT) LeftRec() NTs {
	return leftRec[nt]
}

var ntToString = []string { 
	"EscOrComment", /* NT_EscOrComment */
	"String", /* NT_String */
	"WS", /* NT_WS */ 
}

var tToString = []string { 
	"block_comment", /* T_0 */
	"escCharSpace", /* T_1 */
	"line_comment", /* T_2 */
	"string_ns", /* T_3 */ 
}

var stringNT = map[string]NT{ 
	"EscOrComment":NT_EscOrComment,
	"String":NT_String,
	"WS":NT_WS,
}

var leftRec = map[NT]NTs { 
	NT_EscOrComment: NTs {  },
	NT_String: NTs {  },
	NT_WS: NTs {  NT_EscOrComment,  NT_WS,  },
}
