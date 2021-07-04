
// Package slot is generated by gogll. Do not edit. 
package slot

import(
	"bytes"
	"fmt"
	
	"JSON/parser/symbols"
	"JSON/token"
)

type Label int

const(
	Array0R0 Label = iota
	Array0R1
	Array0R2
	Array0R3
	COLON0R0
	COLON0R1
	COLON0R2
	COMMA0R0
	COMMA0R1
	COMMA0R2
	ComPair0R0
	ComPair0R1
	ComPair0R2
	ComVal0R0
	ComVal0R1
	ComVal0R2
	EXP0R0
	EXP0R1
	EXP0R2
	EXP0R3
	Elements0R0
	Elements0R1
	Elements0R2
	EscOrComment0R0
	EscOrComment0R1
	EscOrComment1R0
	EscOrComment1R1
	EscOrComment2R0
	EscOrComment2R1
	EscOrComment3R0
	FALSE0R0
	FALSE0R1
	FALSE0R2
	FRAC0R0
	FRAC0R1
	FRAC0R2
	INT0R0
	INT0R1
	INT0R2
	Integers0R0
	Integers0R1
	Integers1R0
	Integers1R1
	JSON0R0
	JSON0R1
	JSON0R2
	LBRACE0R0
	LBRACE0R1
	LBRACE0R2
	LBRACKET0R0
	LBRACKET0R1
	LBRACKET0R2
	Members0R0
	Members0R1
	Members0R2
	NUL0R0
	NUL0R1
	NUL0R2
	Number0R0
	Number0R1
	Number0R2
	Number0R3
	Number0R4
	Object0R0
	Object0R1
	Object0R2
	Object0R3
	OptElem0R0
	OptElem0R1
	OptElem1R0
	OptExp0R0
	OptExp0R1
	OptExp1R0
	OptFrac0R0
	OptFrac0R1
	OptFrac1R0
	OptMems0R0
	OptMems0R1
	OptMems1R0
	OptNeg0R0
	OptNeg0R1
	OptNeg1R0
	OptPM0R0
	OptPM0R1
	OptPM1R0
	Pair0R0
	Pair0R1
	Pair0R2
	Pair0R3
	PlusORMinus0R0
	PlusORMinus0R1
	PlusORMinus1R0
	PlusORMinus1R1
	RBRACE0R0
	RBRACE0R1
	RBRACE0R2
	RBRACKET0R0
	RBRACKET0R1
	RBRACKET0R2
	RepComPair0x0R0
	RepComPair0x0R1
	RepComPair0x0R2
	RepComPair0x1R0
	RepComVal0x0R0
	RepComVal0x0R1
	RepComVal0x0R2
	RepComVal0x1R0
	String0R0
	String0R1
	String0R2
	TRUE0R0
	TRUE0R1
	TRUE0R2
	Value0R0
	Value0R1
	Value1R0
	Value1R1
	Value2R0
	Value2R1
	Value3R0
	Value3R1
	Value4R0
	Value4R1
	Value5R0
	Value5R1
	Value6R0
	Value6R1
	WS0R0
	WS0R1
	WS0R2
	WS1R0
)

type Slot struct {
	NT      symbols.NT
	Alt     int
	Pos     int
	Symbols symbols.Symbols
	Label 	Label
}

type Index struct {
	NT      symbols.NT
	Alt     int
	Pos     int
}

func GetAlternates(nt symbols.NT) []Label {
	alts, exist := alternates[nt]
	if !exist {
		panic(fmt.Sprintf("Invalid NT %s", nt))
	}
	return alts
}

func GetLabel(nt symbols.NT, alt, pos int) Label {
	l, exist := slotIndex[Index{nt,alt,pos}]
	if exist {
		return l
	}
	panic(fmt.Sprintf("Error: no slot label for NT=%s, alt=%d, pos=%d", nt, alt, pos))
}

func (l Label) EoR() bool {
	return l.Slot().EoR()
}

func (l Label) Head() symbols.NT {
	return l.Slot().NT
}

func (l Label) Index() Index {
	s := l.Slot()
	return Index{s.NT, s.Alt, s.Pos}
}

func (l Label) Alternate() int {
	return l.Slot().Alt
}

func (l Label) Pos() int {
	return l.Slot().Pos
}

func (l Label) Slot() *Slot {
	s, exist := slots[l]
	if !exist {
		panic(fmt.Sprintf("Invalid slot label %d", l))
	}
	return s
}

func (l Label) String() string {
	return l.Slot().String()
}

func (l Label) Symbols() symbols.Symbols {
	return l.Slot().Symbols
}

func (l Label) IsNullable() bool {
	return nullable[l]
}

func (l Label) FirstContains(typ token.Type) bool {
	return firstT[l][typ]
}

func (s *Slot) EoR() bool {
	return s.Pos >= len(s.Symbols)
}

func (s *Slot) Successor() *Slot {
	if s.EoR() {
		return nil
	} else {
		// TODO try slots[s.Label + 1]
		return slots[slotIndex[Index{s.NT,s.Alt,s.Pos+1}]]
	}
}

func (s *Slot) String() string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%s : ", s.NT)
	for i, sym := range s.Symbols {
		if i == s.Pos {
			fmt.Fprintf(buf, "∙")
		}
		fmt.Fprintf(buf, "%s ", sym)
	}
	if s.Pos >= len(s.Symbols) {
		fmt.Fprintf(buf, "∙")
	}
	return buf.String()
}

var slots = map[Label]*Slot{ 
	Array0R0: {
		symbols.NT_Array, 0, 0, 
		symbols.Symbols{  
			symbols.NT_LBRACKET, 
			symbols.NT_OptElem, 
			symbols.NT_RBRACKET,
		}, 
		Array0R0, 
	},
	Array0R1: {
		symbols.NT_Array, 0, 1, 
		symbols.Symbols{  
			symbols.NT_LBRACKET, 
			symbols.NT_OptElem, 
			symbols.NT_RBRACKET,
		}, 
		Array0R1, 
	},
	Array0R2: {
		symbols.NT_Array, 0, 2, 
		symbols.Symbols{  
			symbols.NT_LBRACKET, 
			symbols.NT_OptElem, 
			symbols.NT_RBRACKET,
		}, 
		Array0R2, 
	},
	Array0R3: {
		symbols.NT_Array, 0, 3, 
		symbols.Symbols{  
			symbols.NT_LBRACKET, 
			symbols.NT_OptElem, 
			symbols.NT_RBRACKET,
		}, 
		Array0R3, 
	},
	COLON0R0: {
		symbols.NT_COLON, 0, 0, 
		symbols.Symbols{  
			symbols.T_5, 
			symbols.NT_WS,
		}, 
		COLON0R0, 
	},
	COLON0R1: {
		symbols.NT_COLON, 0, 1, 
		symbols.Symbols{  
			symbols.T_5, 
			symbols.NT_WS,
		}, 
		COLON0R1, 
	},
	COLON0R2: {
		symbols.NT_COLON, 0, 2, 
		symbols.Symbols{  
			symbols.T_5, 
			symbols.NT_WS,
		}, 
		COLON0R2, 
	},
	COMMA0R0: {
		symbols.NT_COMMA, 0, 0, 
		symbols.Symbols{  
			symbols.T_1, 
			symbols.NT_WS,
		}, 
		COMMA0R0, 
	},
	COMMA0R1: {
		symbols.NT_COMMA, 0, 1, 
		symbols.Symbols{  
			symbols.T_1, 
			symbols.NT_WS,
		}, 
		COMMA0R1, 
	},
	COMMA0R2: {
		symbols.NT_COMMA, 0, 2, 
		symbols.Symbols{  
			symbols.T_1, 
			symbols.NT_WS,
		}, 
		COMMA0R2, 
	},
	ComPair0R0: {
		symbols.NT_ComPair, 0, 0, 
		symbols.Symbols{  
			symbols.NT_COMMA, 
			symbols.NT_Pair,
		}, 
		ComPair0R0, 
	},
	ComPair0R1: {
		symbols.NT_ComPair, 0, 1, 
		symbols.Symbols{  
			symbols.NT_COMMA, 
			symbols.NT_Pair,
		}, 
		ComPair0R1, 
	},
	ComPair0R2: {
		symbols.NT_ComPair, 0, 2, 
		symbols.Symbols{  
			symbols.NT_COMMA, 
			symbols.NT_Pair,
		}, 
		ComPair0R2, 
	},
	ComVal0R0: {
		symbols.NT_ComVal, 0, 0, 
		symbols.Symbols{  
			symbols.NT_COMMA, 
			symbols.NT_Value,
		}, 
		ComVal0R0, 
	},
	ComVal0R1: {
		symbols.NT_ComVal, 0, 1, 
		symbols.Symbols{  
			symbols.NT_COMMA, 
			symbols.NT_Value,
		}, 
		ComVal0R1, 
	},
	ComVal0R2: {
		symbols.NT_ComVal, 0, 2, 
		symbols.Symbols{  
			symbols.NT_COMMA, 
			symbols.NT_Value,
		}, 
		ComVal0R2, 
	},
	EXP0R0: {
		symbols.NT_EXP, 0, 0, 
		symbols.Symbols{  
			symbols.T_9, 
			symbols.NT_OptPM, 
			symbols.T_16,
		}, 
		EXP0R0, 
	},
	EXP0R1: {
		symbols.NT_EXP, 0, 1, 
		symbols.Symbols{  
			symbols.T_9, 
			symbols.NT_OptPM, 
			symbols.T_16,
		}, 
		EXP0R1, 
	},
	EXP0R2: {
		symbols.NT_EXP, 0, 2, 
		symbols.Symbols{  
			symbols.T_9, 
			symbols.NT_OptPM, 
			symbols.T_16,
		}, 
		EXP0R2, 
	},
	EXP0R3: {
		symbols.NT_EXP, 0, 3, 
		symbols.Symbols{  
			symbols.T_9, 
			symbols.NT_OptPM, 
			symbols.T_16,
		}, 
		EXP0R3, 
	},
	Elements0R0: {
		symbols.NT_Elements, 0, 0, 
		symbols.Symbols{  
			symbols.NT_Value, 
			symbols.NT_RepComVal0x,
		}, 
		Elements0R0, 
	},
	Elements0R1: {
		symbols.NT_Elements, 0, 1, 
		symbols.Symbols{  
			symbols.NT_Value, 
			symbols.NT_RepComVal0x,
		}, 
		Elements0R1, 
	},
	Elements0R2: {
		symbols.NT_Elements, 0, 2, 
		symbols.Symbols{  
			symbols.NT_Value, 
			symbols.NT_RepComVal0x,
		}, 
		Elements0R2, 
	},
	EscOrComment0R0: {
		symbols.NT_EscOrComment, 0, 0, 
		symbols.Symbols{  
			symbols.T_10,
		}, 
		EscOrComment0R0, 
	},
	EscOrComment0R1: {
		symbols.NT_EscOrComment, 0, 1, 
		symbols.Symbols{  
			symbols.T_10,
		}, 
		EscOrComment0R1, 
	},
	EscOrComment1R0: {
		symbols.NT_EscOrComment, 1, 0, 
		symbols.Symbols{  
			symbols.T_13,
		}, 
		EscOrComment1R0, 
	},
	EscOrComment1R1: {
		symbols.NT_EscOrComment, 1, 1, 
		symbols.Symbols{  
			symbols.T_13,
		}, 
		EscOrComment1R1, 
	},
	EscOrComment2R0: {
		symbols.NT_EscOrComment, 2, 0, 
		symbols.Symbols{  
			symbols.T_8,
		}, 
		EscOrComment2R0, 
	},
	EscOrComment2R1: {
		symbols.NT_EscOrComment, 2, 1, 
		symbols.Symbols{  
			symbols.T_8,
		}, 
		EscOrComment2R1, 
	},
	EscOrComment3R0: {
		symbols.NT_EscOrComment, 3, 0, 
		symbols.Symbols{ 
		}, 
		EscOrComment3R0, 
	},
	FALSE0R0: {
		symbols.NT_FALSE, 0, 0, 
		symbols.Symbols{  
			symbols.T_11, 
			symbols.NT_WS,
		}, 
		FALSE0R0, 
	},
	FALSE0R1: {
		symbols.NT_FALSE, 0, 1, 
		symbols.Symbols{  
			symbols.T_11, 
			symbols.NT_WS,
		}, 
		FALSE0R1, 
	},
	FALSE0R2: {
		symbols.NT_FALSE, 0, 2, 
		symbols.Symbols{  
			symbols.T_11, 
			symbols.NT_WS,
		}, 
		FALSE0R2, 
	},
	FRAC0R0: {
		symbols.NT_FRAC, 0, 0, 
		symbols.Symbols{  
			symbols.T_3, 
			symbols.T_16,
		}, 
		FRAC0R0, 
	},
	FRAC0R1: {
		symbols.NT_FRAC, 0, 1, 
		symbols.Symbols{  
			symbols.T_3, 
			symbols.T_16,
		}, 
		FRAC0R1, 
	},
	FRAC0R2: {
		symbols.NT_FRAC, 0, 2, 
		symbols.Symbols{  
			symbols.T_3, 
			symbols.T_16,
		}, 
		FRAC0R2, 
	},
	INT0R0: {
		symbols.NT_INT, 0, 0, 
		symbols.Symbols{  
			symbols.NT_OptNeg, 
			symbols.NT_Integers,
		}, 
		INT0R0, 
	},
	INT0R1: {
		symbols.NT_INT, 0, 1, 
		symbols.Symbols{  
			symbols.NT_OptNeg, 
			symbols.NT_Integers,
		}, 
		INT0R1, 
	},
	INT0R2: {
		symbols.NT_INT, 0, 2, 
		symbols.Symbols{  
			symbols.NT_OptNeg, 
			symbols.NT_Integers,
		}, 
		INT0R2, 
	},
	Integers0R0: {
		symbols.NT_Integers, 0, 0, 
		symbols.Symbols{  
			symbols.T_14,
		}, 
		Integers0R0, 
	},
	Integers0R1: {
		symbols.NT_Integers, 0, 1, 
		symbols.Symbols{  
			symbols.T_14,
		}, 
		Integers0R1, 
	},
	Integers1R0: {
		symbols.NT_Integers, 1, 0, 
		symbols.Symbols{  
			symbols.T_4,
		}, 
		Integers1R0, 
	},
	Integers1R1: {
		symbols.NT_Integers, 1, 1, 
		symbols.Symbols{  
			symbols.T_4,
		}, 
		Integers1R1, 
	},
	JSON0R0: {
		symbols.NT_JSON, 0, 0, 
		symbols.Symbols{  
			symbols.NT_WS, 
			symbols.NT_Object,
		}, 
		JSON0R0, 
	},
	JSON0R1: {
		symbols.NT_JSON, 0, 1, 
		symbols.Symbols{  
			symbols.NT_WS, 
			symbols.NT_Object,
		}, 
		JSON0R1, 
	},
	JSON0R2: {
		symbols.NT_JSON, 0, 2, 
		symbols.Symbols{  
			symbols.NT_WS, 
			symbols.NT_Object,
		}, 
		JSON0R2, 
	},
	LBRACE0R0: {
		symbols.NT_LBRACE, 0, 0, 
		symbols.Symbols{  
			symbols.T_19, 
			symbols.NT_WS,
		}, 
		LBRACE0R0, 
	},
	LBRACE0R1: {
		symbols.NT_LBRACE, 0, 1, 
		symbols.Symbols{  
			symbols.T_19, 
			symbols.NT_WS,
		}, 
		LBRACE0R1, 
	},
	LBRACE0R2: {
		symbols.NT_LBRACE, 0, 2, 
		symbols.Symbols{  
			symbols.T_19, 
			symbols.NT_WS,
		}, 
		LBRACE0R2, 
	},
	LBRACKET0R0: {
		symbols.NT_LBRACKET, 0, 0, 
		symbols.Symbols{  
			symbols.T_6, 
			symbols.NT_WS,
		}, 
		LBRACKET0R0, 
	},
	LBRACKET0R1: {
		symbols.NT_LBRACKET, 0, 1, 
		symbols.Symbols{  
			symbols.T_6, 
			symbols.NT_WS,
		}, 
		LBRACKET0R1, 
	},
	LBRACKET0R2: {
		symbols.NT_LBRACKET, 0, 2, 
		symbols.Symbols{  
			symbols.T_6, 
			symbols.NT_WS,
		}, 
		LBRACKET0R2, 
	},
	Members0R0: {
		symbols.NT_Members, 0, 0, 
		symbols.Symbols{  
			symbols.NT_Pair, 
			symbols.NT_RepComPair0x,
		}, 
		Members0R0, 
	},
	Members0R1: {
		symbols.NT_Members, 0, 1, 
		symbols.Symbols{  
			symbols.NT_Pair, 
			symbols.NT_RepComPair0x,
		}, 
		Members0R1, 
	},
	Members0R2: {
		symbols.NT_Members, 0, 2, 
		symbols.Symbols{  
			symbols.NT_Pair, 
			symbols.NT_RepComPair0x,
		}, 
		Members0R2, 
	},
	NUL0R0: {
		symbols.NT_NUL, 0, 0, 
		symbols.Symbols{  
			symbols.T_15, 
			symbols.NT_WS,
		}, 
		NUL0R0, 
	},
	NUL0R1: {
		symbols.NT_NUL, 0, 1, 
		symbols.Symbols{  
			symbols.T_15, 
			symbols.NT_WS,
		}, 
		NUL0R1, 
	},
	NUL0R2: {
		symbols.NT_NUL, 0, 2, 
		symbols.Symbols{  
			symbols.T_15, 
			symbols.NT_WS,
		}, 
		NUL0R2, 
	},
	Number0R0: {
		symbols.NT_Number, 0, 0, 
		symbols.Symbols{  
			symbols.NT_INT, 
			symbols.NT_OptFrac, 
			symbols.NT_OptExp, 
			symbols.NT_WS,
		}, 
		Number0R0, 
	},
	Number0R1: {
		symbols.NT_Number, 0, 1, 
		symbols.Symbols{  
			symbols.NT_INT, 
			symbols.NT_OptFrac, 
			symbols.NT_OptExp, 
			symbols.NT_WS,
		}, 
		Number0R1, 
	},
	Number0R2: {
		symbols.NT_Number, 0, 2, 
		symbols.Symbols{  
			symbols.NT_INT, 
			symbols.NT_OptFrac, 
			symbols.NT_OptExp, 
			symbols.NT_WS,
		}, 
		Number0R2, 
	},
	Number0R3: {
		symbols.NT_Number, 0, 3, 
		symbols.Symbols{  
			symbols.NT_INT, 
			symbols.NT_OptFrac, 
			symbols.NT_OptExp, 
			symbols.NT_WS,
		}, 
		Number0R3, 
	},
	Number0R4: {
		symbols.NT_Number, 0, 4, 
		symbols.Symbols{  
			symbols.NT_INT, 
			symbols.NT_OptFrac, 
			symbols.NT_OptExp, 
			symbols.NT_WS,
		}, 
		Number0R4, 
	},
	Object0R0: {
		symbols.NT_Object, 0, 0, 
		symbols.Symbols{  
			symbols.NT_LBRACE, 
			symbols.NT_OptMems, 
			symbols.NT_RBRACE,
		}, 
		Object0R0, 
	},
	Object0R1: {
		symbols.NT_Object, 0, 1, 
		symbols.Symbols{  
			symbols.NT_LBRACE, 
			symbols.NT_OptMems, 
			symbols.NT_RBRACE,
		}, 
		Object0R1, 
	},
	Object0R2: {
		symbols.NT_Object, 0, 2, 
		symbols.Symbols{  
			symbols.NT_LBRACE, 
			symbols.NT_OptMems, 
			symbols.NT_RBRACE,
		}, 
		Object0R2, 
	},
	Object0R3: {
		symbols.NT_Object, 0, 3, 
		symbols.Symbols{  
			symbols.NT_LBRACE, 
			symbols.NT_OptMems, 
			symbols.NT_RBRACE,
		}, 
		Object0R3, 
	},
	OptElem0R0: {
		symbols.NT_OptElem, 0, 0, 
		symbols.Symbols{  
			symbols.NT_Elements,
		}, 
		OptElem0R0, 
	},
	OptElem0R1: {
		symbols.NT_OptElem, 0, 1, 
		symbols.Symbols{  
			symbols.NT_Elements,
		}, 
		OptElem0R1, 
	},
	OptElem1R0: {
		symbols.NT_OptElem, 1, 0, 
		symbols.Symbols{ 
		}, 
		OptElem1R0, 
	},
	OptExp0R0: {
		symbols.NT_OptExp, 0, 0, 
		symbols.Symbols{  
			symbols.NT_EXP,
		}, 
		OptExp0R0, 
	},
	OptExp0R1: {
		symbols.NT_OptExp, 0, 1, 
		symbols.Symbols{  
			symbols.NT_EXP,
		}, 
		OptExp0R1, 
	},
	OptExp1R0: {
		symbols.NT_OptExp, 1, 0, 
		symbols.Symbols{ 
		}, 
		OptExp1R0, 
	},
	OptFrac0R0: {
		symbols.NT_OptFrac, 0, 0, 
		symbols.Symbols{  
			symbols.NT_FRAC,
		}, 
		OptFrac0R0, 
	},
	OptFrac0R1: {
		symbols.NT_OptFrac, 0, 1, 
		symbols.Symbols{  
			symbols.NT_FRAC,
		}, 
		OptFrac0R1, 
	},
	OptFrac1R0: {
		symbols.NT_OptFrac, 1, 0, 
		symbols.Symbols{ 
		}, 
		OptFrac1R0, 
	},
	OptMems0R0: {
		symbols.NT_OptMems, 0, 0, 
		symbols.Symbols{  
			symbols.NT_Members,
		}, 
		OptMems0R0, 
	},
	OptMems0R1: {
		symbols.NT_OptMems, 0, 1, 
		symbols.Symbols{  
			symbols.NT_Members,
		}, 
		OptMems0R1, 
	},
	OptMems1R0: {
		symbols.NT_OptMems, 1, 0, 
		symbols.Symbols{ 
		}, 
		OptMems1R0, 
	},
	OptNeg0R0: {
		symbols.NT_OptNeg, 0, 0, 
		symbols.Symbols{  
			symbols.T_2,
		}, 
		OptNeg0R0, 
	},
	OptNeg0R1: {
		symbols.NT_OptNeg, 0, 1, 
		symbols.Symbols{  
			symbols.T_2,
		}, 
		OptNeg0R1, 
	},
	OptNeg1R0: {
		symbols.NT_OptNeg, 1, 0, 
		symbols.Symbols{ 
		}, 
		OptNeg1R0, 
	},
	OptPM0R0: {
		symbols.NT_OptPM, 0, 0, 
		symbols.Symbols{  
			symbols.NT_PlusORMinus,
		}, 
		OptPM0R0, 
	},
	OptPM0R1: {
		symbols.NT_OptPM, 0, 1, 
		symbols.Symbols{  
			symbols.NT_PlusORMinus,
		}, 
		OptPM0R1, 
	},
	OptPM1R0: {
		symbols.NT_OptPM, 1, 0, 
		symbols.Symbols{ 
		}, 
		OptPM1R0, 
	},
	Pair0R0: {
		symbols.NT_Pair, 0, 0, 
		symbols.Symbols{  
			symbols.NT_String, 
			symbols.NT_COLON, 
			symbols.NT_Value,
		}, 
		Pair0R0, 
	},
	Pair0R1: {
		symbols.NT_Pair, 0, 1, 
		symbols.Symbols{  
			symbols.NT_String, 
			symbols.NT_COLON, 
			symbols.NT_Value,
		}, 
		Pair0R1, 
	},
	Pair0R2: {
		symbols.NT_Pair, 0, 2, 
		symbols.Symbols{  
			symbols.NT_String, 
			symbols.NT_COLON, 
			symbols.NT_Value,
		}, 
		Pair0R2, 
	},
	Pair0R3: {
		symbols.NT_Pair, 0, 3, 
		symbols.Symbols{  
			symbols.NT_String, 
			symbols.NT_COLON, 
			symbols.NT_Value,
		}, 
		Pair0R3, 
	},
	PlusORMinus0R0: {
		symbols.NT_PlusORMinus, 0, 0, 
		symbols.Symbols{  
			symbols.T_0,
		}, 
		PlusORMinus0R0, 
	},
	PlusORMinus0R1: {
		symbols.NT_PlusORMinus, 0, 1, 
		symbols.Symbols{  
			symbols.T_0,
		}, 
		PlusORMinus0R1, 
	},
	PlusORMinus1R0: {
		symbols.NT_PlusORMinus, 1, 0, 
		symbols.Symbols{  
			symbols.T_2,
		}, 
		PlusORMinus1R0, 
	},
	PlusORMinus1R1: {
		symbols.NT_PlusORMinus, 1, 1, 
		symbols.Symbols{  
			symbols.T_2,
		}, 
		PlusORMinus1R1, 
	},
	RBRACE0R0: {
		symbols.NT_RBRACE, 0, 0, 
		symbols.Symbols{  
			symbols.T_20, 
			symbols.NT_WS,
		}, 
		RBRACE0R0, 
	},
	RBRACE0R1: {
		symbols.NT_RBRACE, 0, 1, 
		symbols.Symbols{  
			symbols.T_20, 
			symbols.NT_WS,
		}, 
		RBRACE0R1, 
	},
	RBRACE0R2: {
		symbols.NT_RBRACE, 0, 2, 
		symbols.Symbols{  
			symbols.T_20, 
			symbols.NT_WS,
		}, 
		RBRACE0R2, 
	},
	RBRACKET0R0: {
		symbols.NT_RBRACKET, 0, 0, 
		symbols.Symbols{  
			symbols.T_7, 
			symbols.NT_WS,
		}, 
		RBRACKET0R0, 
	},
	RBRACKET0R1: {
		symbols.NT_RBRACKET, 0, 1, 
		symbols.Symbols{  
			symbols.T_7, 
			symbols.NT_WS,
		}, 
		RBRACKET0R1, 
	},
	RBRACKET0R2: {
		symbols.NT_RBRACKET, 0, 2, 
		symbols.Symbols{  
			symbols.T_7, 
			symbols.NT_WS,
		}, 
		RBRACKET0R2, 
	},
	RepComPair0x0R0: {
		symbols.NT_RepComPair0x, 0, 0, 
		symbols.Symbols{  
			symbols.NT_ComPair, 
			symbols.NT_RepComPair0x,
		}, 
		RepComPair0x0R0, 
	},
	RepComPair0x0R1: {
		symbols.NT_RepComPair0x, 0, 1, 
		symbols.Symbols{  
			symbols.NT_ComPair, 
			symbols.NT_RepComPair0x,
		}, 
		RepComPair0x0R1, 
	},
	RepComPair0x0R2: {
		symbols.NT_RepComPair0x, 0, 2, 
		symbols.Symbols{  
			symbols.NT_ComPair, 
			symbols.NT_RepComPair0x,
		}, 
		RepComPair0x0R2, 
	},
	RepComPair0x1R0: {
		symbols.NT_RepComPair0x, 1, 0, 
		symbols.Symbols{ 
		}, 
		RepComPair0x1R0, 
	},
	RepComVal0x0R0: {
		symbols.NT_RepComVal0x, 0, 0, 
		symbols.Symbols{  
			symbols.NT_ComVal, 
			symbols.NT_RepComVal0x,
		}, 
		RepComVal0x0R0, 
	},
	RepComVal0x0R1: {
		symbols.NT_RepComVal0x, 0, 1, 
		symbols.Symbols{  
			symbols.NT_ComVal, 
			symbols.NT_RepComVal0x,
		}, 
		RepComVal0x0R1, 
	},
	RepComVal0x0R2: {
		symbols.NT_RepComVal0x, 0, 2, 
		symbols.Symbols{  
			symbols.NT_ComVal, 
			symbols.NT_RepComVal0x,
		}, 
		RepComVal0x0R2, 
	},
	RepComVal0x1R0: {
		symbols.NT_RepComVal0x, 1, 0, 
		symbols.Symbols{ 
		}, 
		RepComVal0x1R0, 
	},
	String0R0: {
		symbols.NT_String, 0, 0, 
		symbols.Symbols{  
			symbols.T_17, 
			symbols.NT_WS,
		}, 
		String0R0, 
	},
	String0R1: {
		symbols.NT_String, 0, 1, 
		symbols.Symbols{  
			symbols.T_17, 
			symbols.NT_WS,
		}, 
		String0R1, 
	},
	String0R2: {
		symbols.NT_String, 0, 2, 
		symbols.Symbols{  
			symbols.T_17, 
			symbols.NT_WS,
		}, 
		String0R2, 
	},
	TRUE0R0: {
		symbols.NT_TRUE, 0, 0, 
		symbols.Symbols{  
			symbols.T_18, 
			symbols.NT_WS,
		}, 
		TRUE0R0, 
	},
	TRUE0R1: {
		symbols.NT_TRUE, 0, 1, 
		symbols.Symbols{  
			symbols.T_18, 
			symbols.NT_WS,
		}, 
		TRUE0R1, 
	},
	TRUE0R2: {
		symbols.NT_TRUE, 0, 2, 
		symbols.Symbols{  
			symbols.T_18, 
			symbols.NT_WS,
		}, 
		TRUE0R2, 
	},
	Value0R0: {
		symbols.NT_Value, 0, 0, 
		symbols.Symbols{  
			symbols.NT_String,
		}, 
		Value0R0, 
	},
	Value0R1: {
		symbols.NT_Value, 0, 1, 
		symbols.Symbols{  
			symbols.NT_String,
		}, 
		Value0R1, 
	},
	Value1R0: {
		symbols.NT_Value, 1, 0, 
		symbols.Symbols{  
			symbols.NT_Number,
		}, 
		Value1R0, 
	},
	Value1R1: {
		symbols.NT_Value, 1, 1, 
		symbols.Symbols{  
			symbols.NT_Number,
		}, 
		Value1R1, 
	},
	Value2R0: {
		symbols.NT_Value, 2, 0, 
		symbols.Symbols{  
			symbols.NT_Object,
		}, 
		Value2R0, 
	},
	Value2R1: {
		symbols.NT_Value, 2, 1, 
		symbols.Symbols{  
			symbols.NT_Object,
		}, 
		Value2R1, 
	},
	Value3R0: {
		symbols.NT_Value, 3, 0, 
		symbols.Symbols{  
			symbols.NT_Array,
		}, 
		Value3R0, 
	},
	Value3R1: {
		symbols.NT_Value, 3, 1, 
		symbols.Symbols{  
			symbols.NT_Array,
		}, 
		Value3R1, 
	},
	Value4R0: {
		symbols.NT_Value, 4, 0, 
		symbols.Symbols{  
			symbols.NT_TRUE,
		}, 
		Value4R0, 
	},
	Value4R1: {
		symbols.NT_Value, 4, 1, 
		symbols.Symbols{  
			symbols.NT_TRUE,
		}, 
		Value4R1, 
	},
	Value5R0: {
		symbols.NT_Value, 5, 0, 
		symbols.Symbols{  
			symbols.NT_FALSE,
		}, 
		Value5R0, 
	},
	Value5R1: {
		symbols.NT_Value, 5, 1, 
		symbols.Symbols{  
			symbols.NT_FALSE,
		}, 
		Value5R1, 
	},
	Value6R0: {
		symbols.NT_Value, 6, 0, 
		symbols.Symbols{  
			symbols.NT_NUL,
		}, 
		Value6R0, 
	},
	Value6R1: {
		symbols.NT_Value, 6, 1, 
		symbols.Symbols{  
			symbols.NT_NUL,
		}, 
		Value6R1, 
	},
	WS0R0: {
		symbols.NT_WS, 0, 0, 
		symbols.Symbols{  
			symbols.NT_EscOrComment, 
			symbols.NT_WS,
		}, 
		WS0R0, 
	},
	WS0R1: {
		symbols.NT_WS, 0, 1, 
		symbols.Symbols{  
			symbols.NT_EscOrComment, 
			symbols.NT_WS,
		}, 
		WS0R1, 
	},
	WS0R2: {
		symbols.NT_WS, 0, 2, 
		symbols.Symbols{  
			symbols.NT_EscOrComment, 
			symbols.NT_WS,
		}, 
		WS0R2, 
	},
	WS1R0: {
		symbols.NT_WS, 1, 0, 
		symbols.Symbols{ 
		}, 
		WS1R0, 
	},
}

var slotIndex = map[Index]Label { 
	Index{ symbols.NT_Array,0,0 }: Array0R0,
	Index{ symbols.NT_Array,0,1 }: Array0R1,
	Index{ symbols.NT_Array,0,2 }: Array0R2,
	Index{ symbols.NT_Array,0,3 }: Array0R3,
	Index{ symbols.NT_COLON,0,0 }: COLON0R0,
	Index{ symbols.NT_COLON,0,1 }: COLON0R1,
	Index{ symbols.NT_COLON,0,2 }: COLON0R2,
	Index{ symbols.NT_COMMA,0,0 }: COMMA0R0,
	Index{ symbols.NT_COMMA,0,1 }: COMMA0R1,
	Index{ symbols.NT_COMMA,0,2 }: COMMA0R2,
	Index{ symbols.NT_ComPair,0,0 }: ComPair0R0,
	Index{ symbols.NT_ComPair,0,1 }: ComPair0R1,
	Index{ symbols.NT_ComPair,0,2 }: ComPair0R2,
	Index{ symbols.NT_ComVal,0,0 }: ComVal0R0,
	Index{ symbols.NT_ComVal,0,1 }: ComVal0R1,
	Index{ symbols.NT_ComVal,0,2 }: ComVal0R2,
	Index{ symbols.NT_EXP,0,0 }: EXP0R0,
	Index{ symbols.NT_EXP,0,1 }: EXP0R1,
	Index{ symbols.NT_EXP,0,2 }: EXP0R2,
	Index{ symbols.NT_EXP,0,3 }: EXP0R3,
	Index{ symbols.NT_Elements,0,0 }: Elements0R0,
	Index{ symbols.NT_Elements,0,1 }: Elements0R1,
	Index{ symbols.NT_Elements,0,2 }: Elements0R2,
	Index{ symbols.NT_EscOrComment,0,0 }: EscOrComment0R0,
	Index{ symbols.NT_EscOrComment,0,1 }: EscOrComment0R1,
	Index{ symbols.NT_EscOrComment,1,0 }: EscOrComment1R0,
	Index{ symbols.NT_EscOrComment,1,1 }: EscOrComment1R1,
	Index{ symbols.NT_EscOrComment,2,0 }: EscOrComment2R0,
	Index{ symbols.NT_EscOrComment,2,1 }: EscOrComment2R1,
	Index{ symbols.NT_EscOrComment,3,0 }: EscOrComment3R0,
	Index{ symbols.NT_FALSE,0,0 }: FALSE0R0,
	Index{ symbols.NT_FALSE,0,1 }: FALSE0R1,
	Index{ symbols.NT_FALSE,0,2 }: FALSE0R2,
	Index{ symbols.NT_FRAC,0,0 }: FRAC0R0,
	Index{ symbols.NT_FRAC,0,1 }: FRAC0R1,
	Index{ symbols.NT_FRAC,0,2 }: FRAC0R2,
	Index{ symbols.NT_INT,0,0 }: INT0R0,
	Index{ symbols.NT_INT,0,1 }: INT0R1,
	Index{ symbols.NT_INT,0,2 }: INT0R2,
	Index{ symbols.NT_Integers,0,0 }: Integers0R0,
	Index{ symbols.NT_Integers,0,1 }: Integers0R1,
	Index{ symbols.NT_Integers,1,0 }: Integers1R0,
	Index{ symbols.NT_Integers,1,1 }: Integers1R1,
	Index{ symbols.NT_JSON,0,0 }: JSON0R0,
	Index{ symbols.NT_JSON,0,1 }: JSON0R1,
	Index{ symbols.NT_JSON,0,2 }: JSON0R2,
	Index{ symbols.NT_LBRACE,0,0 }: LBRACE0R0,
	Index{ symbols.NT_LBRACE,0,1 }: LBRACE0R1,
	Index{ symbols.NT_LBRACE,0,2 }: LBRACE0R2,
	Index{ symbols.NT_LBRACKET,0,0 }: LBRACKET0R0,
	Index{ symbols.NT_LBRACKET,0,1 }: LBRACKET0R1,
	Index{ symbols.NT_LBRACKET,0,2 }: LBRACKET0R2,
	Index{ symbols.NT_Members,0,0 }: Members0R0,
	Index{ symbols.NT_Members,0,1 }: Members0R1,
	Index{ symbols.NT_Members,0,2 }: Members0R2,
	Index{ symbols.NT_NUL,0,0 }: NUL0R0,
	Index{ symbols.NT_NUL,0,1 }: NUL0R1,
	Index{ symbols.NT_NUL,0,2 }: NUL0R2,
	Index{ symbols.NT_Number,0,0 }: Number0R0,
	Index{ symbols.NT_Number,0,1 }: Number0R1,
	Index{ symbols.NT_Number,0,2 }: Number0R2,
	Index{ symbols.NT_Number,0,3 }: Number0R3,
	Index{ symbols.NT_Number,0,4 }: Number0R4,
	Index{ symbols.NT_Object,0,0 }: Object0R0,
	Index{ symbols.NT_Object,0,1 }: Object0R1,
	Index{ symbols.NT_Object,0,2 }: Object0R2,
	Index{ symbols.NT_Object,0,3 }: Object0R3,
	Index{ symbols.NT_OptElem,0,0 }: OptElem0R0,
	Index{ symbols.NT_OptElem,0,1 }: OptElem0R1,
	Index{ symbols.NT_OptElem,1,0 }: OptElem1R0,
	Index{ symbols.NT_OptExp,0,0 }: OptExp0R0,
	Index{ symbols.NT_OptExp,0,1 }: OptExp0R1,
	Index{ symbols.NT_OptExp,1,0 }: OptExp1R0,
	Index{ symbols.NT_OptFrac,0,0 }: OptFrac0R0,
	Index{ symbols.NT_OptFrac,0,1 }: OptFrac0R1,
	Index{ symbols.NT_OptFrac,1,0 }: OptFrac1R0,
	Index{ symbols.NT_OptMems,0,0 }: OptMems0R0,
	Index{ symbols.NT_OptMems,0,1 }: OptMems0R1,
	Index{ symbols.NT_OptMems,1,0 }: OptMems1R0,
	Index{ symbols.NT_OptNeg,0,0 }: OptNeg0R0,
	Index{ symbols.NT_OptNeg,0,1 }: OptNeg0R1,
	Index{ symbols.NT_OptNeg,1,0 }: OptNeg1R0,
	Index{ symbols.NT_OptPM,0,0 }: OptPM0R0,
	Index{ symbols.NT_OptPM,0,1 }: OptPM0R1,
	Index{ symbols.NT_OptPM,1,0 }: OptPM1R0,
	Index{ symbols.NT_Pair,0,0 }: Pair0R0,
	Index{ symbols.NT_Pair,0,1 }: Pair0R1,
	Index{ symbols.NT_Pair,0,2 }: Pair0R2,
	Index{ symbols.NT_Pair,0,3 }: Pair0R3,
	Index{ symbols.NT_PlusORMinus,0,0 }: PlusORMinus0R0,
	Index{ symbols.NT_PlusORMinus,0,1 }: PlusORMinus0R1,
	Index{ symbols.NT_PlusORMinus,1,0 }: PlusORMinus1R0,
	Index{ symbols.NT_PlusORMinus,1,1 }: PlusORMinus1R1,
	Index{ symbols.NT_RBRACE,0,0 }: RBRACE0R0,
	Index{ symbols.NT_RBRACE,0,1 }: RBRACE0R1,
	Index{ symbols.NT_RBRACE,0,2 }: RBRACE0R2,
	Index{ symbols.NT_RBRACKET,0,0 }: RBRACKET0R0,
	Index{ symbols.NT_RBRACKET,0,1 }: RBRACKET0R1,
	Index{ symbols.NT_RBRACKET,0,2 }: RBRACKET0R2,
	Index{ symbols.NT_RepComPair0x,0,0 }: RepComPair0x0R0,
	Index{ symbols.NT_RepComPair0x,0,1 }: RepComPair0x0R1,
	Index{ symbols.NT_RepComPair0x,0,2 }: RepComPair0x0R2,
	Index{ symbols.NT_RepComPair0x,1,0 }: RepComPair0x1R0,
	Index{ symbols.NT_RepComVal0x,0,0 }: RepComVal0x0R0,
	Index{ symbols.NT_RepComVal0x,0,1 }: RepComVal0x0R1,
	Index{ symbols.NT_RepComVal0x,0,2 }: RepComVal0x0R2,
	Index{ symbols.NT_RepComVal0x,1,0 }: RepComVal0x1R0,
	Index{ symbols.NT_String,0,0 }: String0R0,
	Index{ symbols.NT_String,0,1 }: String0R1,
	Index{ symbols.NT_String,0,2 }: String0R2,
	Index{ symbols.NT_TRUE,0,0 }: TRUE0R0,
	Index{ symbols.NT_TRUE,0,1 }: TRUE0R1,
	Index{ symbols.NT_TRUE,0,2 }: TRUE0R2,
	Index{ symbols.NT_Value,0,0 }: Value0R0,
	Index{ symbols.NT_Value,0,1 }: Value0R1,
	Index{ symbols.NT_Value,1,0 }: Value1R0,
	Index{ symbols.NT_Value,1,1 }: Value1R1,
	Index{ symbols.NT_Value,2,0 }: Value2R0,
	Index{ symbols.NT_Value,2,1 }: Value2R1,
	Index{ symbols.NT_Value,3,0 }: Value3R0,
	Index{ symbols.NT_Value,3,1 }: Value3R1,
	Index{ symbols.NT_Value,4,0 }: Value4R0,
	Index{ symbols.NT_Value,4,1 }: Value4R1,
	Index{ symbols.NT_Value,5,0 }: Value5R0,
	Index{ symbols.NT_Value,5,1 }: Value5R1,
	Index{ symbols.NT_Value,6,0 }: Value6R0,
	Index{ symbols.NT_Value,6,1 }: Value6R1,
	Index{ symbols.NT_WS,0,0 }: WS0R0,
	Index{ symbols.NT_WS,0,1 }: WS0R1,
	Index{ symbols.NT_WS,0,2 }: WS0R2,
	Index{ symbols.NT_WS,1,0 }: WS1R0,
}

var alternates = map[symbols.NT][]Label{ 
	symbols.NT_JSON:[]Label{ JSON0R0 },
	symbols.NT_Object:[]Label{ Object0R0 },
	symbols.NT_OptMems:[]Label{ OptMems0R0,OptMems1R0 },
	symbols.NT_Members:[]Label{ Members0R0 },
	symbols.NT_RepComPair0x:[]Label{ RepComPair0x0R0,RepComPair0x1R0 },
	symbols.NT_ComPair:[]Label{ ComPair0R0 },
	symbols.NT_Pair:[]Label{ Pair0R0 },
	symbols.NT_Array:[]Label{ Array0R0 },
	symbols.NT_OptElem:[]Label{ OptElem0R0,OptElem1R0 },
	symbols.NT_Elements:[]Label{ Elements0R0 },
	symbols.NT_RepComVal0x:[]Label{ RepComVal0x0R0,RepComVal0x1R0 },
	symbols.NT_ComVal:[]Label{ ComVal0R0 },
	symbols.NT_Value:[]Label{ Value0R0,Value1R0,Value2R0,Value3R0,Value4R0,Value5R0,Value6R0 },
	symbols.NT_String:[]Label{ String0R0 },
	symbols.NT_Number:[]Label{ Number0R0 },
	symbols.NT_OptFrac:[]Label{ OptFrac0R0,OptFrac1R0 },
	symbols.NT_OptExp:[]Label{ OptExp0R0,OptExp1R0 },
	symbols.NT_INT:[]Label{ INT0R0 },
	symbols.NT_Integers:[]Label{ Integers0R0,Integers1R0 },
	symbols.NT_OptNeg:[]Label{ OptNeg0R0,OptNeg1R0 },
	symbols.NT_FRAC:[]Label{ FRAC0R0 },
	symbols.NT_EXP:[]Label{ EXP0R0 },
	symbols.NT_OptPM:[]Label{ OptPM0R0,OptPM1R0 },
	symbols.NT_PlusORMinus:[]Label{ PlusORMinus0R0,PlusORMinus1R0 },
	symbols.NT_TRUE:[]Label{ TRUE0R0 },
	symbols.NT_FALSE:[]Label{ FALSE0R0 },
	symbols.NT_NUL:[]Label{ NUL0R0 },
	symbols.NT_COMMA:[]Label{ COMMA0R0 },
	symbols.NT_COLON:[]Label{ COLON0R0 },
	symbols.NT_LBRACE:[]Label{ LBRACE0R0 },
	symbols.NT_RBRACE:[]Label{ RBRACE0R0 },
	symbols.NT_LBRACKET:[]Label{ LBRACKET0R0 },
	symbols.NT_RBRACKET:[]Label{ RBRACKET0R0 },
	symbols.NT_WS:[]Label{ WS0R0,WS1R0 },
	symbols.NT_EscOrComment:[]Label{ EscOrComment0R0,EscOrComment1R0,EscOrComment2R0,EscOrComment3R0 },
}

var nullable = []bool { 
	false, // Array0R0 
	false, // Array0R1 
	false, // Array0R2 
	true, // Array0R3 
	false, // COLON0R0 
	true, // COLON0R1 
	true, // COLON0R2 
	false, // COMMA0R0 
	true, // COMMA0R1 
	true, // COMMA0R2 
	false, // ComPair0R0 
	false, // ComPair0R1 
	true, // ComPair0R2 
	false, // ComVal0R0 
	false, // ComVal0R1 
	true, // ComVal0R2 
	false, // EXP0R0 
	false, // EXP0R1 
	false, // EXP0R2 
	true, // EXP0R3 
	false, // Elements0R0 
	true, // Elements0R1 
	true, // Elements0R2 
	false, // EscOrComment0R0 
	true, // EscOrComment0R1 
	false, // EscOrComment1R0 
	true, // EscOrComment1R1 
	false, // EscOrComment2R0 
	true, // EscOrComment2R1 
	true, // EscOrComment3R0 
	false, // FALSE0R0 
	true, // FALSE0R1 
	true, // FALSE0R2 
	false, // FRAC0R0 
	false, // FRAC0R1 
	true, // FRAC0R2 
	false, // INT0R0 
	false, // INT0R1 
	true, // INT0R2 
	false, // Integers0R0 
	true, // Integers0R1 
	false, // Integers1R0 
	true, // Integers1R1 
	false, // JSON0R0 
	false, // JSON0R1 
	true, // JSON0R2 
	false, // LBRACE0R0 
	true, // LBRACE0R1 
	true, // LBRACE0R2 
	false, // LBRACKET0R0 
	true, // LBRACKET0R1 
	true, // LBRACKET0R2 
	false, // Members0R0 
	true, // Members0R1 
	true, // Members0R2 
	false, // NUL0R0 
	true, // NUL0R1 
	true, // NUL0R2 
	false, // Number0R0 
	true, // Number0R1 
	true, // Number0R2 
	true, // Number0R3 
	true, // Number0R4 
	false, // Object0R0 
	false, // Object0R1 
	false, // Object0R2 
	true, // Object0R3 
	false, // OptElem0R0 
	true, // OptElem0R1 
	true, // OptElem1R0 
	false, // OptExp0R0 
	true, // OptExp0R1 
	true, // OptExp1R0 
	false, // OptFrac0R0 
	true, // OptFrac0R1 
	true, // OptFrac1R0 
	false, // OptMems0R0 
	true, // OptMems0R1 
	true, // OptMems1R0 
	false, // OptNeg0R0 
	true, // OptNeg0R1 
	true, // OptNeg1R0 
	false, // OptPM0R0 
	true, // OptPM0R1 
	true, // OptPM1R0 
	false, // Pair0R0 
	false, // Pair0R1 
	false, // Pair0R2 
	true, // Pair0R3 
	false, // PlusORMinus0R0 
	true, // PlusORMinus0R1 
	false, // PlusORMinus1R0 
	true, // PlusORMinus1R1 
	false, // RBRACE0R0 
	true, // RBRACE0R1 
	true, // RBRACE0R2 
	false, // RBRACKET0R0 
	true, // RBRACKET0R1 
	true, // RBRACKET0R2 
	false, // RepComPair0x0R0 
	true, // RepComPair0x0R1 
	true, // RepComPair0x0R2 
	true, // RepComPair0x1R0 
	false, // RepComVal0x0R0 
	true, // RepComVal0x0R1 
	true, // RepComVal0x0R2 
	true, // RepComVal0x1R0 
	false, // String0R0 
	true, // String0R1 
	true, // String0R2 
	false, // TRUE0R0 
	true, // TRUE0R1 
	true, // TRUE0R2 
	false, // Value0R0 
	true, // Value0R1 
	false, // Value1R0 
	true, // Value1R1 
	false, // Value2R0 
	true, // Value2R1 
	false, // Value3R0 
	true, // Value3R1 
	false, // Value4R0 
	true, // Value4R1 
	false, // Value5R0 
	true, // Value5R1 
	false, // Value6R0 
	true, // Value6R1 
	true, // WS0R0 
	true, // WS0R1 
	true, // WS0R2 
	true, // WS1R0 
}

var firstT = []map[token.Type]bool { 
	{  token.T_6: true,  }, // Array0R0 
	{  token.T_17: true,  token.T_6: true,  token.T_19: true,  token.T_4: true,  token.T_18: true,  token.T_15: true,  token.T_7: true,  token.T_2: true,  token.T_14: true,  token.T_11: true,  }, // Array0R1 
	{  token.T_7: true,  }, // Array0R2 
	{  }, // Array0R3 
	{  token.T_5: true,  }, // COLON0R0 
	{  token.T_10: true,  token.T_13: true,  token.T_8: true,  }, // COLON0R1 
	{  }, // COLON0R2 
	{  token.T_1: true,  }, // COMMA0R0 
	{  token.T_10: true,  token.T_13: true,  token.T_8: true,  }, // COMMA0R1 
	{  }, // COMMA0R2 
	{  token.T_1: true,  }, // ComPair0R0 
	{  token.T_17: true,  }, // ComPair0R1 
	{  }, // ComPair0R2 
	{  token.T_1: true,  }, // ComVal0R0 
	{  token.T_4: true,  token.T_14: true,  token.T_6: true,  token.T_11: true,  token.T_19: true,  token.T_2: true,  token.T_17: true,  token.T_18: true,  token.T_15: true,  }, // ComVal0R1 
	{  }, // ComVal0R2 
	{  token.T_9: true,  }, // EXP0R0 
	{  token.T_16: true,  token.T_0: true,  token.T_2: true,  }, // EXP0R1 
	{  token.T_16: true,  }, // EXP0R2 
	{  }, // EXP0R3 
	{  token.T_4: true,  token.T_19: true,  token.T_18: true,  token.T_11: true,  token.T_15: true,  token.T_6: true,  token.T_17: true,  token.T_2: true,  token.T_14: true,  }, // Elements0R0 
	{  token.T_1: true,  }, // Elements0R1 
	{  }, // Elements0R2 
	{  token.T_10: true,  }, // EscOrComment0R0 
	{  }, // EscOrComment0R1 
	{  token.T_13: true,  }, // EscOrComment1R0 
	{  }, // EscOrComment1R1 
	{  token.T_8: true,  }, // EscOrComment2R0 
	{  }, // EscOrComment2R1 
	{  }, // EscOrComment3R0 
	{  token.T_11: true,  }, // FALSE0R0 
	{  token.T_8: true,  token.T_10: true,  token.T_13: true,  }, // FALSE0R1 
	{  }, // FALSE0R2 
	{  token.T_3: true,  }, // FRAC0R0 
	{  token.T_16: true,  }, // FRAC0R1 
	{  }, // FRAC0R2 
	{  token.T_4: true,  token.T_14: true,  token.T_2: true,  }, // INT0R0 
	{  token.T_14: true,  token.T_4: true,  }, // INT0R1 
	{  }, // INT0R2 
	{  token.T_14: true,  }, // Integers0R0 
	{  }, // Integers0R1 
	{  token.T_4: true,  }, // Integers1R0 
	{  }, // Integers1R1 
	{  token.T_8: true,  token.T_10: true,  token.T_13: true,  token.T_19: true,  }, // JSON0R0 
	{  token.T_19: true,  }, // JSON0R1 
	{  }, // JSON0R2 
	{  token.T_19: true,  }, // LBRACE0R0 
	{  token.T_10: true,  token.T_13: true,  token.T_8: true,  }, // LBRACE0R1 
	{  }, // LBRACE0R2 
	{  token.T_6: true,  }, // LBRACKET0R0 
	{  token.T_10: true,  token.T_13: true,  token.T_8: true,  }, // LBRACKET0R1 
	{  }, // LBRACKET0R2 
	{  token.T_17: true,  }, // Members0R0 
	{  token.T_1: true,  }, // Members0R1 
	{  }, // Members0R2 
	{  token.T_15: true,  }, // NUL0R0 
	{  token.T_10: true,  token.T_13: true,  token.T_8: true,  }, // NUL0R1 
	{  }, // NUL0R2 
	{  token.T_2: true,  token.T_14: true,  token.T_4: true,  }, // Number0R0 
	{  token.T_3: true,  token.T_9: true,  token.T_8: true,  token.T_10: true,  token.T_13: true,  }, // Number0R1 
	{  token.T_13: true,  token.T_8: true,  token.T_9: true,  token.T_10: true,  }, // Number0R2 
	{  token.T_10: true,  token.T_13: true,  token.T_8: true,  }, // Number0R3 
	{  }, // Number0R4 
	{  token.T_19: true,  }, // Object0R0 
	{  token.T_17: true,  token.T_20: true,  }, // Object0R1 
	{  token.T_20: true,  }, // Object0R2 
	{  }, // Object0R3 
	{  token.T_6: true,  token.T_2: true,  token.T_4: true,  token.T_18: true,  token.T_15: true,  token.T_19: true,  token.T_14: true,  token.T_11: true,  token.T_17: true,  }, // OptElem0R0 
	{  }, // OptElem0R1 
	{  }, // OptElem1R0 
	{  token.T_9: true,  }, // OptExp0R0 
	{  }, // OptExp0R1 
	{  }, // OptExp1R0 
	{  token.T_3: true,  }, // OptFrac0R0 
	{  }, // OptFrac0R1 
	{  }, // OptFrac1R0 
	{  token.T_17: true,  }, // OptMems0R0 
	{  }, // OptMems0R1 
	{  }, // OptMems1R0 
	{  token.T_2: true,  }, // OptNeg0R0 
	{  }, // OptNeg0R1 
	{  }, // OptNeg1R0 
	{  token.T_0: true,  token.T_2: true,  }, // OptPM0R0 
	{  }, // OptPM0R1 
	{  }, // OptPM1R0 
	{  token.T_17: true,  }, // Pair0R0 
	{  token.T_5: true,  }, // Pair0R1 
	{  token.T_4: true,  token.T_14: true,  token.T_15: true,  token.T_6: true,  token.T_17: true,  token.T_2: true,  token.T_19: true,  token.T_18: true,  token.T_11: true,  }, // Pair0R2 
	{  }, // Pair0R3 
	{  token.T_0: true,  }, // PlusORMinus0R0 
	{  }, // PlusORMinus0R1 
	{  token.T_2: true,  }, // PlusORMinus1R0 
	{  }, // PlusORMinus1R1 
	{  token.T_20: true,  }, // RBRACE0R0 
	{  token.T_10: true,  token.T_13: true,  token.T_8: true,  }, // RBRACE0R1 
	{  }, // RBRACE0R2 
	{  token.T_7: true,  }, // RBRACKET0R0 
	{  token.T_13: true,  token.T_8: true,  token.T_10: true,  }, // RBRACKET0R1 
	{  }, // RBRACKET0R2 
	{  token.T_1: true,  }, // RepComPair0x0R0 
	{  token.T_1: true,  }, // RepComPair0x0R1 
	{  }, // RepComPair0x0R2 
	{  }, // RepComPair0x1R0 
	{  token.T_1: true,  }, // RepComVal0x0R0 
	{  token.T_1: true,  }, // RepComVal0x0R1 
	{  }, // RepComVal0x0R2 
	{  }, // RepComVal0x1R0 
	{  token.T_17: true,  }, // String0R0 
	{  token.T_10: true,  token.T_13: true,  token.T_8: true,  }, // String0R1 
	{  }, // String0R2 
	{  token.T_18: true,  }, // TRUE0R0 
	{  token.T_10: true,  token.T_13: true,  token.T_8: true,  }, // TRUE0R1 
	{  }, // TRUE0R2 
	{  token.T_17: true,  }, // Value0R0 
	{  }, // Value0R1 
	{  token.T_4: true,  token.T_2: true,  token.T_14: true,  }, // Value1R0 
	{  }, // Value1R1 
	{  token.T_19: true,  }, // Value2R0 
	{  }, // Value2R1 
	{  token.T_6: true,  }, // Value3R0 
	{  }, // Value3R1 
	{  token.T_18: true,  }, // Value4R0 
	{  }, // Value4R1 
	{  token.T_11: true,  }, // Value5R0 
	{  }, // Value5R1 
	{  token.T_15: true,  }, // Value6R0 
	{  }, // Value6R1 
	{  token.T_10: true,  token.T_13: true,  token.T_8: true,  }, // WS0R0 
	{  token.T_10: true,  token.T_13: true,  token.T_8: true,  }, // WS0R1 
	{  }, // WS0R2 
	{  }, // WS1R0 
}
