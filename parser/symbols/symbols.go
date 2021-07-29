
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
	NT_GoGLL NT = iota
	NT_LexAlternates 
	NT_LexBracket 
	NT_LexGroup 
	NT_LexOneOrMore 
	NT_LexOptional 
	NT_LexRule 
	NT_LexSymbol 
	NT_LexZeroOrMore 
	NT_OrderedAlternates 
	NT_Package 
	NT_RegExp 
	NT_Rule 
	NT_Rules 
	NT_SyntaxAlternate 
	NT_SyntaxAlternates 
	NT_SyntaxAtom 
	NT_SyntaxRule 
	NT_SyntaxSuffix 
	NT_SyntaxSymbol 
	NT_SyntaxSymbols 
	NT_UnicodeClass 
	NT_UnorderedAlternates 
)

const NumNTs = 23

type NTs []NT

// T is the type of terminals symbols
type T int
const( 
	T_0 T = iota // ! 
	T_1  // & 
	T_2  // ( 
	T_3  // ) 
	T_4  // * 
	T_5  // + 
	T_6  // . 
	T_7  // / 
	T_8  // : 
	T_9  // ; 
	T_10  // < 
	T_11  // > 
	T_12  // ? 
	T_13  // [ 
	T_14  // ] 
	T_15  // any 
	T_16  // char_lit 
	T_17  // empty 
	T_18  // letter 
	T_19  // lowcase 
	T_20  // not 
	T_21  // nt 
	T_22  // number 
	T_23  // package 
	T_24  // string_lit 
	T_25  // tokid 
	T_26  // upcase 
	T_27  // { 
	T_28  // | 
	T_29  // } 
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
	"GoGLL", /* NT_GoGLL */
	"LexAlternates", /* NT_LexAlternates */
	"LexBracket", /* NT_LexBracket */
	"LexGroup", /* NT_LexGroup */
	"LexOneOrMore", /* NT_LexOneOrMore */
	"LexOptional", /* NT_LexOptional */
	"LexRule", /* NT_LexRule */
	"LexSymbol", /* NT_LexSymbol */
	"LexZeroOrMore", /* NT_LexZeroOrMore */
	"OrderedAlternates", /* NT_OrderedAlternates */
	"Package", /* NT_Package */
	"RegExp", /* NT_RegExp */
	"Rule", /* NT_Rule */
	"Rules", /* NT_Rules */
	"SyntaxAlternate", /* NT_SyntaxAlternate */
	"SyntaxAlternates", /* NT_SyntaxAlternates */
	"SyntaxAtom", /* NT_SyntaxAtom */
	"SyntaxRule", /* NT_SyntaxRule */
	"SyntaxSuffix", /* NT_SyntaxSuffix */
	"SyntaxSymbol", /* NT_SyntaxSymbol */
	"SyntaxSymbols", /* NT_SyntaxSymbols */
	"UnicodeClass", /* NT_UnicodeClass */
	"UnorderedAlternates", /* NT_UnorderedAlternates */ 
}

var tToString = []string { 
	"!", /* T_0 */
	"&", /* T_1 */
	"(", /* T_2 */
	")", /* T_3 */
	"*", /* T_4 */
	"+", /* T_5 */
	".", /* T_6 */
	"/", /* T_7 */
	":", /* T_8 */
	";", /* T_9 */
	"<", /* T_10 */
	">", /* T_11 */
	"?", /* T_12 */
	"[", /* T_13 */
	"]", /* T_14 */
	"any", /* T_15 */
	"char_lit", /* T_16 */
	"empty", /* T_17 */
	"letter", /* T_18 */
	"lowcase", /* T_19 */
	"not", /* T_20 */
	"nt", /* T_21 */
	"number", /* T_22 */
	"package", /* T_23 */
	"string_lit", /* T_24 */
	"tokid", /* T_25 */
	"upcase", /* T_26 */
	"{", /* T_27 */
	"|", /* T_28 */
	"}", /* T_29 */ 
}

var stringNT = map[string]NT{ 
	"GoGLL":NT_GoGLL,
	"LexAlternates":NT_LexAlternates,
	"LexBracket":NT_LexBracket,
	"LexGroup":NT_LexGroup,
	"LexOneOrMore":NT_LexOneOrMore,
	"LexOptional":NT_LexOptional,
	"LexRule":NT_LexRule,
	"LexSymbol":NT_LexSymbol,
	"LexZeroOrMore":NT_LexZeroOrMore,
	"OrderedAlternates":NT_OrderedAlternates,
	"Package":NT_Package,
	"RegExp":NT_RegExp,
	"Rule":NT_Rule,
	"Rules":NT_Rules,
	"SyntaxAlternate":NT_SyntaxAlternate,
	"SyntaxAlternates":NT_SyntaxAlternates,
	"SyntaxAtom":NT_SyntaxAtom,
	"SyntaxRule":NT_SyntaxRule,
	"SyntaxSuffix":NT_SyntaxSuffix,
	"SyntaxSymbol":NT_SyntaxSymbol,
	"SyntaxSymbols":NT_SyntaxSymbols,
	"UnicodeClass":NT_UnicodeClass,
	"UnorderedAlternates":NT_UnorderedAlternates,
}

var leftRec = map[NT]NTs { 
	NT_GoGLL: NTs {  NT_Package,  },
	NT_LexAlternates: NTs {  NT_RegExp,  NT_LexBracket,  NT_LexGroup,  NT_LexOptional,  NT_LexZeroOrMore,  NT_LexOneOrMore,  NT_UnicodeClass,  NT_LexSymbol,  },
	NT_LexBracket: NTs {  NT_LexGroup,  NT_LexOptional,  NT_LexZeroOrMore,  NT_LexOneOrMore,  },
	NT_LexGroup: NTs {  },
	NT_LexOneOrMore: NTs {  },
	NT_LexOptional: NTs {  },
	NT_LexRule: NTs {  },
	NT_LexSymbol: NTs {  NT_LexOneOrMore,  NT_UnicodeClass,  NT_LexBracket,  NT_LexGroup,  NT_LexOptional,  NT_LexZeroOrMore,  },
	NT_LexZeroOrMore: NTs {  },
	NT_OrderedAlternates: NTs {  NT_SyntaxAlternate,  NT_SyntaxSymbol,  NT_SyntaxAtom,  NT_SyntaxSymbols,  NT_SyntaxSuffix,  },
	NT_Package: NTs {  },
	NT_RegExp: NTs {  NT_UnicodeClass,  NT_LexSymbol,  NT_LexBracket,  NT_LexGroup,  NT_LexOptional,  NT_LexZeroOrMore,  NT_LexOneOrMore,  },
	NT_Rule: NTs {  NT_SyntaxRule,  NT_LexRule,  },
	NT_Rules: NTs {  NT_Rule,  NT_LexRule,  NT_SyntaxRule,  },
	NT_SyntaxAlternate: NTs {  NT_SyntaxSymbol,  NT_SyntaxAtom,  NT_SyntaxSymbols,  NT_SyntaxSuffix,  },
	NT_SyntaxAlternates: NTs {  NT_SyntaxAlternate,  NT_SyntaxSuffix,  NT_SyntaxSymbol,  NT_SyntaxAtom,  NT_SyntaxSymbols,  },
	NT_SyntaxAtom: NTs {  },
	NT_SyntaxRule: NTs {  },
	NT_SyntaxSuffix: NTs {  NT_SyntaxAtom,  },
	NT_SyntaxSymbol: NTs {  NT_SyntaxSuffix,  NT_SyntaxAtom,  },
	NT_SyntaxSymbols: NTs {  NT_SyntaxSymbol,  NT_SyntaxAtom,  NT_SyntaxSuffix,  },
	NT_UnicodeClass: NTs {  },
	NT_UnorderedAlternates: NTs {  NT_SyntaxSymbols,  NT_SyntaxSuffix,  NT_SyntaxSymbol,  NT_SyntaxAtom,  NT_SyntaxAlternate,  },
}
