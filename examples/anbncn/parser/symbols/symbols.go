
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
	NT_A1 NT = iota
	NT_Ac 
	NT_B1 
	NT_G1 
	NT_Repa0x 
)

const NumNTs = 5

type NTs []NT

// T is the type of terminals symbols
type T int
const( 
	T_0 T = iota // a 
	T_1  // b 
	T_2  // c 
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

func (nt NT) IsOrdered() bool {
	return ordered[nt]
}

var ntToString = []string { 
	"A1", /* NT_A1 */
	"Ac", /* NT_Ac */
	"B1", /* NT_B1 */
	"G1", /* NT_G1 */
	"Repa0x", /* NT_Repa0x */ 
}

var tToString = []string { 
	"a", /* T_0 */
	"b", /* T_1 */
	"c", /* T_2 */ 
}

var stringNT = map[string]NT{ 
	"A1":NT_A1,
	"Ac":NT_Ac,
	"B1":NT_B1,
	"G1":NT_G1,
	"Repa0x":NT_Repa0x,
}

var leftRec = map[NT]NTs { 
	NT_A1: NTs {  },
	NT_Ac: NTs {  NT_A1,  },
	NT_B1: NTs {  },
	NT_G1: NTs {  NT_Ac,  NT_A1,  },
	NT_Repa0x: NTs {  },
}

var ordered = map[NT]bool { 
	NT_A1:true,
	NT_B1:true,
	NT_Repa0x:true,
}
