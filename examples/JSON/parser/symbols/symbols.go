
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
	NT_Array NT = iota
	NT_COLON 
	NT_COMMA 
	NT_ComPair 
	NT_ComVal 
	NT_EXP 
	NT_Elements 
	NT_EscOrComment 
	NT_FALSE 
	NT_FRAC 
	NT_INT 
	NT_Integers 
	NT_JSON 
	NT_LBRACE 
	NT_LBRACKET 
	NT_Members 
	NT_NUL 
	NT_Number 
	NT_Object 
	NT_OptElem 
	NT_OptExp 
	NT_OptFrac 
	NT_OptMems 
	NT_OptNeg 
	NT_OptPM 
	NT_Pair 
	NT_PlusORMinus 
	NT_RBRACE 
	NT_RBRACKET 
	NT_RepComPair0x 
	NT_RepComVal0x 
	NT_String 
	NT_TRUE 
	NT_Value 
	NT_WS 
)

const NumNTs = 35

type NTs []NT

// T is the type of terminals symbols
type T int
const( 
	T_0 T = iota // + 
	T_1  // , 
	T_2  // - 
	T_3  // . 
	T_4  // 0 
	T_5  // : 
	T_6  // [ 
	T_7  // ] 
	T_8  // block_comment 
	T_9  // eE 
	T_10  // escCharSpace 
	T_11  // false 
	T_12  // hex 
	T_13  // line_comment 
	T_14  // nonZero 
	T_15  // null 
	T_16  // repNum1x 
	T_17  // string_ns 
	T_18  // true 
	T_19  // { 
	T_20  // } 
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
	"Array", /* NT_Array */
	"COLON", /* NT_COLON */
	"COMMA", /* NT_COMMA */
	"ComPair", /* NT_ComPair */
	"ComVal", /* NT_ComVal */
	"EXP", /* NT_EXP */
	"Elements", /* NT_Elements */
	"EscOrComment", /* NT_EscOrComment */
	"FALSE", /* NT_FALSE */
	"FRAC", /* NT_FRAC */
	"INT", /* NT_INT */
	"Integers", /* NT_Integers */
	"JSON", /* NT_JSON */
	"LBRACE", /* NT_LBRACE */
	"LBRACKET", /* NT_LBRACKET */
	"Members", /* NT_Members */
	"NUL", /* NT_NUL */
	"Number", /* NT_Number */
	"Object", /* NT_Object */
	"OptElem", /* NT_OptElem */
	"OptExp", /* NT_OptExp */
	"OptFrac", /* NT_OptFrac */
	"OptMems", /* NT_OptMems */
	"OptNeg", /* NT_OptNeg */
	"OptPM", /* NT_OptPM */
	"Pair", /* NT_Pair */
	"PlusORMinus", /* NT_PlusORMinus */
	"RBRACE", /* NT_RBRACE */
	"RBRACKET", /* NT_RBRACKET */
	"RepComPair0x", /* NT_RepComPair0x */
	"RepComVal0x", /* NT_RepComVal0x */
	"String", /* NT_String */
	"TRUE", /* NT_TRUE */
	"Value", /* NT_Value */
	"WS", /* NT_WS */ 
}

var tToString = []string { 
	"+", /* T_0 */
	",", /* T_1 */
	"-", /* T_2 */
	".", /* T_3 */
	"0", /* T_4 */
	":", /* T_5 */
	"[", /* T_6 */
	"]", /* T_7 */
	"block_comment", /* T_8 */
	"eE", /* T_9 */
	"escCharSpace", /* T_10 */
	"false", /* T_11 */
	"hex", /* T_12 */
	"line_comment", /* T_13 */
	"nonZero", /* T_14 */
	"null", /* T_15 */
	"repNum1x", /* T_16 */
	"string_ns", /* T_17 */
	"true", /* T_18 */
	"{", /* T_19 */
	"}", /* T_20 */ 
}

var stringNT = map[string]NT{ 
	"Array":NT_Array,
	"COLON":NT_COLON,
	"COMMA":NT_COMMA,
	"ComPair":NT_ComPair,
	"ComVal":NT_ComVal,
	"EXP":NT_EXP,
	"Elements":NT_Elements,
	"EscOrComment":NT_EscOrComment,
	"FALSE":NT_FALSE,
	"FRAC":NT_FRAC,
	"INT":NT_INT,
	"Integers":NT_Integers,
	"JSON":NT_JSON,
	"LBRACE":NT_LBRACE,
	"LBRACKET":NT_LBRACKET,
	"Members":NT_Members,
	"NUL":NT_NUL,
	"Number":NT_Number,
	"Object":NT_Object,
	"OptElem":NT_OptElem,
	"OptExp":NT_OptExp,
	"OptFrac":NT_OptFrac,
	"OptMems":NT_OptMems,
	"OptNeg":NT_OptNeg,
	"OptPM":NT_OptPM,
	"Pair":NT_Pair,
	"PlusORMinus":NT_PlusORMinus,
	"RBRACE":NT_RBRACE,
	"RBRACKET":NT_RBRACKET,
	"RepComPair0x":NT_RepComPair0x,
	"RepComVal0x":NT_RepComVal0x,
	"String":NT_String,
	"TRUE":NT_TRUE,
	"Value":NT_Value,
	"WS":NT_WS,
}

var leftRec = map[NT]NTs { 
	NT_Array: NTs {  NT_LBRACKET,  },
	NT_COLON: NTs {  },
	NT_COMMA: NTs {  },
	NT_ComPair: NTs {  NT_COMMA,  },
	NT_ComVal: NTs {  NT_COMMA,  },
	NT_EXP: NTs {  },
	NT_Elements: NTs {  NT_NUL,  NT_Value,  NT_OptNeg,  NT_Object,  NT_LBRACE,  NT_Array,  NT_LBRACKET,  NT_TRUE,  NT_INT,  NT_Integers,  NT_String,  NT_Number,  NT_FALSE,  },
	NT_EscOrComment: NTs {  },
	NT_FALSE: NTs {  },
	NT_FRAC: NTs {  },
	NT_INT: NTs {  NT_OptNeg,  NT_Integers,  },
	NT_Integers: NTs {  },
	NT_JSON: NTs {  NT_WS,  NT_EscOrComment,  NT_Object,  NT_LBRACE,  },
	NT_LBRACE: NTs {  },
	NT_LBRACKET: NTs {  },
	NT_Members: NTs {  NT_Pair,  NT_String,  },
	NT_NUL: NTs {  },
	NT_Number: NTs {  NT_INT,  NT_OptNeg,  NT_Integers,  },
	NT_Object: NTs {  NT_LBRACE,  },
	NT_OptElem: NTs {  NT_Number,  NT_Value,  NT_OptNeg,  NT_Elements,  NT_FALSE,  NT_LBRACKET,  NT_TRUE,  NT_Array,  NT_Object,  NT_INT,  NT_String,  NT_NUL,  NT_Integers,  NT_LBRACE,  },
	NT_OptExp: NTs {  NT_EXP,  },
	NT_OptFrac: NTs {  NT_FRAC,  },
	NT_OptMems: NTs {  NT_Members,  NT_Pair,  NT_String,  },
	NT_OptNeg: NTs {  },
	NT_OptPM: NTs {  NT_PlusORMinus,  },
	NT_Pair: NTs {  NT_String,  },
	NT_PlusORMinus: NTs {  },
	NT_RBRACE: NTs {  },
	NT_RBRACKET: NTs {  },
	NT_RepComPair0x: NTs {  NT_ComPair,  NT_COMMA,  },
	NT_RepComVal0x: NTs {  NT_ComVal,  NT_COMMA,  },
	NT_String: NTs {  },
	NT_TRUE: NTs {  },
	NT_Value: NTs {  NT_Object,  NT_LBRACE,  NT_Array,  NT_String,  NT_Number,  NT_INT,  NT_OptNeg,  NT_Integers,  NT_LBRACKET,  NT_TRUE,  NT_FALSE,  NT_NUL,  },
	NT_WS: NTs {  NT_EscOrComment,  NT_WS,  },
}
