
// Package slot is generated by gogll. Do not edit. 
package slot

import(
	"bytes"
	"fmt"
	
	"JSON/parser/symbols"
)

type Label int

const(
	Array0R0 Label = iota
	Array0R1
	Array0R2
	Array0R3
	CHAR0R0
	CHAR0R1
	CHAR1R0
	CHAR1R1
	CHAR1R2
	COLON0R0
	COLON0R1
	COLON0R2
	COMMA0R0
	COMMA0R1
	COMMA0R2
	CharCode0R0
	CharCode0R1
	CharCode1R0
	CharCode1R1
	CharCode1R2
	CharCode1R3
	CharCode1R4
	CharCode1R5
	Close0R0
	Close0R1
	Close1R0
	Close1R1
	Close1R2
	ComPair0R0
	ComPair0R1
	ComPair0R2
	ComPair0x0R0
	ComPair0x0R1
	ComPair0x0R2
	ComPair0x1R0
	ComVal0R0
	ComVal0R1
	ComVal0R2
	ComVal0x0R0
	ComVal0x0R1
	ComVal0x0R2
	ComVal0x1R0
	Elements0R0
	Elements0R1
	Elements0R2
	EscOrComment0R0
	EscOrComment0R1
	EscOrComment1R0
	EscOrComment1R1
	FALSE0R0
	FALSE0R1
	FALSE0R2
	HEX0R0
	HEX0R1
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
	LineOrBlock0R0
	LineOrBlock0R1
	LineOrBlock1R0
	LineOrBlock1R1
	Members0R0
	Members0R1
	Members0R2
	Mems1x0R0
	Mems1x0R1
	Mems1x0R2
	Mems1x1R0
	NUL0R0
	NUL0R1
	NUL0R2
	Number0R0
	Number0R1
	Number0R2
	Number0R3
	Number0R4
	NumberHEX0R0
	NumberHEX0R1
	NumberHEX0R2
	NumberHEX1R0
	Object0R0
	Object0R1
	Object0R2
	Object0R3
	Object0R4
	OptElem0R0
	OptElem0R1
	OptElem1R0
	OptExp0R0
	OptExp0R1
	OptExp1R0
	OptFrac0R0
	OptFrac0R1
	OptFrac1R0
	Pair0R0
	Pair0R1
	Pair0R2
	Pair0R3
	RBRACE0R0
	RBRACE0R1
	RBRACE0R2
	RBRACKET0R0
	RBRACKET0R1
	RBRACKET0R2
	String0R0
	String0R1
	String0R2
	String0R3
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

func (s *Slot) EoR() bool {
	return s.Pos >= len(s.Symbols)
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
	CHAR0R0: {
		symbols.NT_CHAR, 0, 0, 
		symbols.Symbols{  
			symbols.T_7,
		}, 
		CHAR0R0, 
	},
	CHAR0R1: {
		symbols.NT_CHAR, 0, 1, 
		symbols.Symbols{  
			symbols.T_7,
		}, 
		CHAR0R1, 
	},
	CHAR1R0: {
		symbols.NT_CHAR, 1, 0, 
		symbols.Symbols{  
			symbols.T_5, 
			symbols.NT_CharCode,
		}, 
		CHAR1R0, 
	},
	CHAR1R1: {
		symbols.NT_CHAR, 1, 1, 
		symbols.Symbols{  
			symbols.T_5, 
			symbols.NT_CharCode,
		}, 
		CHAR1R1, 
	},
	CHAR1R2: {
		symbols.NT_CHAR, 1, 2, 
		symbols.Symbols{  
			symbols.T_5, 
			symbols.NT_CharCode,
		}, 
		CHAR1R2, 
	},
	COLON0R0: {
		symbols.NT_COLON, 0, 0, 
		symbols.Symbols{  
			symbols.T_1, 
			symbols.NT_WS,
		}, 
		COLON0R0, 
	},
	COLON0R1: {
		symbols.NT_COLON, 0, 1, 
		symbols.Symbols{  
			symbols.T_1, 
			symbols.NT_WS,
		}, 
		COLON0R1, 
	},
	COLON0R2: {
		symbols.NT_COLON, 0, 2, 
		symbols.Symbols{  
			symbols.T_1, 
			symbols.NT_WS,
		}, 
		COLON0R2, 
	},
	COMMA0R0: {
		symbols.NT_COMMA, 0, 0, 
		symbols.Symbols{  
			symbols.T_0, 
			symbols.NT_WS,
		}, 
		COMMA0R0, 
	},
	COMMA0R1: {
		symbols.NT_COMMA, 0, 1, 
		symbols.Symbols{  
			symbols.T_0, 
			symbols.NT_WS,
		}, 
		COMMA0R1, 
	},
	COMMA0R2: {
		symbols.NT_COMMA, 0, 2, 
		symbols.Symbols{  
			symbols.T_0, 
			symbols.NT_WS,
		}, 
		COMMA0R2, 
	},
	CharCode0R0: {
		symbols.NT_CharCode, 0, 0, 
		symbols.Symbols{  
			symbols.T_9,
		}, 
		CharCode0R0, 
	},
	CharCode0R1: {
		symbols.NT_CharCode, 0, 1, 
		symbols.Symbols{  
			symbols.T_9,
		}, 
		CharCode0R1, 
	},
	CharCode1R0: {
		symbols.NT_CharCode, 1, 0, 
		symbols.Symbols{  
			symbols.T_21, 
			symbols.NT_HEX, 
			symbols.NT_HEX, 
			symbols.NT_HEX, 
			symbols.NT_HEX,
		}, 
		CharCode1R0, 
	},
	CharCode1R1: {
		symbols.NT_CharCode, 1, 1, 
		symbols.Symbols{  
			symbols.T_21, 
			symbols.NT_HEX, 
			symbols.NT_HEX, 
			symbols.NT_HEX, 
			symbols.NT_HEX,
		}, 
		CharCode1R1, 
	},
	CharCode1R2: {
		symbols.NT_CharCode, 1, 2, 
		symbols.Symbols{  
			symbols.T_21, 
			symbols.NT_HEX, 
			symbols.NT_HEX, 
			symbols.NT_HEX, 
			symbols.NT_HEX,
		}, 
		CharCode1R2, 
	},
	CharCode1R3: {
		symbols.NT_CharCode, 1, 3, 
		symbols.Symbols{  
			symbols.T_21, 
			symbols.NT_HEX, 
			symbols.NT_HEX, 
			symbols.NT_HEX, 
			symbols.NT_HEX,
		}, 
		CharCode1R3, 
	},
	CharCode1R4: {
		symbols.NT_CharCode, 1, 4, 
		symbols.Symbols{  
			symbols.T_21, 
			symbols.NT_HEX, 
			symbols.NT_HEX, 
			symbols.NT_HEX, 
			symbols.NT_HEX,
		}, 
		CharCode1R4, 
	},
	CharCode1R5: {
		symbols.NT_CharCode, 1, 5, 
		symbols.Symbols{  
			symbols.T_21, 
			symbols.NT_HEX, 
			symbols.NT_HEX, 
			symbols.NT_HEX, 
			symbols.NT_HEX,
		}, 
		CharCode1R5, 
	},
	Close0R0: {
		symbols.NT_Close, 0, 0, 
		symbols.Symbols{  
			symbols.T_8,
		}, 
		Close0R0, 
	},
	Close0R1: {
		symbols.NT_Close, 0, 1, 
		symbols.Symbols{  
			symbols.T_8,
		}, 
		Close0R1, 
	},
	Close1R0: {
		symbols.NT_Close, 1, 0, 
		symbols.Symbols{  
			symbols.NT_CHAR, 
			symbols.NT_Close,
		}, 
		Close1R0, 
	},
	Close1R1: {
		symbols.NT_Close, 1, 1, 
		symbols.Symbols{  
			symbols.NT_CHAR, 
			symbols.NT_Close,
		}, 
		Close1R1, 
	},
	Close1R2: {
		symbols.NT_Close, 1, 2, 
		symbols.Symbols{  
			symbols.NT_CHAR, 
			symbols.NT_Close,
		}, 
		Close1R2, 
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
	ComPair0x0R0: {
		symbols.NT_ComPair0x, 0, 0, 
		symbols.Symbols{  
			symbols.NT_ComPair, 
			symbols.NT_ComPair0x,
		}, 
		ComPair0x0R0, 
	},
	ComPair0x0R1: {
		symbols.NT_ComPair0x, 0, 1, 
		symbols.Symbols{  
			symbols.NT_ComPair, 
			symbols.NT_ComPair0x,
		}, 
		ComPair0x0R1, 
	},
	ComPair0x0R2: {
		symbols.NT_ComPair0x, 0, 2, 
		symbols.Symbols{  
			symbols.NT_ComPair, 
			symbols.NT_ComPair0x,
		}, 
		ComPair0x0R2, 
	},
	ComPair0x1R0: {
		symbols.NT_ComPair0x, 1, 0, 
		symbols.Symbols{ 
		}, 
		ComPair0x1R0, 
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
	ComVal0x0R0: {
		symbols.NT_ComVal0x, 0, 0, 
		symbols.Symbols{  
			symbols.NT_ComVal, 
			symbols.NT_ComVal0x,
		}, 
		ComVal0x0R0, 
	},
	ComVal0x0R1: {
		symbols.NT_ComVal0x, 0, 1, 
		symbols.Symbols{  
			symbols.NT_ComVal, 
			symbols.NT_ComVal0x,
		}, 
		ComVal0x0R1, 
	},
	ComVal0x0R2: {
		symbols.NT_ComVal0x, 0, 2, 
		symbols.Symbols{  
			symbols.NT_ComVal, 
			symbols.NT_ComVal0x,
		}, 
		ComVal0x0R2, 
	},
	ComVal0x1R0: {
		symbols.NT_ComVal0x, 1, 0, 
		symbols.Symbols{ 
		}, 
		ComVal0x1R0, 
	},
	Elements0R0: {
		symbols.NT_Elements, 0, 0, 
		symbols.Symbols{  
			symbols.NT_Value, 
			symbols.NT_ComVal0x,
		}, 
		Elements0R0, 
	},
	Elements0R1: {
		symbols.NT_Elements, 0, 1, 
		symbols.Symbols{  
			symbols.NT_Value, 
			symbols.NT_ComVal0x,
		}, 
		Elements0R1, 
	},
	Elements0R2: {
		symbols.NT_Elements, 0, 2, 
		symbols.Symbols{  
			symbols.NT_Value, 
			symbols.NT_ComVal0x,
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
			symbols.NT_LineOrBlock,
		}, 
		EscOrComment1R0, 
	},
	EscOrComment1R1: {
		symbols.NT_EscOrComment, 1, 1, 
		symbols.Symbols{  
			symbols.NT_LineOrBlock,
		}, 
		EscOrComment1R1, 
	},
	FALSE0R0: {
		symbols.NT_FALSE, 0, 0, 
		symbols.Symbols{  
			symbols.T_13, 
			symbols.NT_WS,
		}, 
		FALSE0R0, 
	},
	FALSE0R1: {
		symbols.NT_FALSE, 0, 1, 
		symbols.Symbols{  
			symbols.T_13, 
			symbols.NT_WS,
		}, 
		FALSE0R1, 
	},
	FALSE0R2: {
		symbols.NT_FALSE, 0, 2, 
		symbols.Symbols{  
			symbols.T_13, 
			symbols.NT_WS,
		}, 
		FALSE0R2, 
	},
	HEX0R0: {
		symbols.NT_HEX, 0, 0, 
		symbols.Symbols{  
			symbols.NT_NumberHEX,
		}, 
		HEX0R0, 
	},
	HEX0R1: {
		symbols.NT_HEX, 0, 1, 
		symbols.Symbols{  
			symbols.NT_NumberHEX,
		}, 
		HEX0R1, 
	},
	INT0R0: {
		symbols.NT_INT, 0, 0, 
		symbols.Symbols{  
			symbols.T_19, 
			symbols.NT_Integers,
		}, 
		INT0R0, 
	},
	INT0R1: {
		symbols.NT_INT, 0, 1, 
		symbols.Symbols{  
			symbols.T_19, 
			symbols.NT_Integers,
		}, 
		INT0R1, 
	},
	INT0R2: {
		symbols.NT_INT, 0, 2, 
		symbols.Symbols{  
			symbols.T_19, 
			symbols.NT_Integers,
		}, 
		INT0R2, 
	},
	Integers0R0: {
		symbols.NT_Integers, 0, 0, 
		symbols.Symbols{  
			symbols.T_15,
		}, 
		Integers0R0, 
	},
	Integers0R1: {
		symbols.NT_Integers, 0, 1, 
		symbols.Symbols{  
			symbols.T_15,
		}, 
		Integers0R1, 
	},
	Integers1R0: {
		symbols.NT_Integers, 1, 0, 
		symbols.Symbols{  
			symbols.T_22,
		}, 
		Integers1R0, 
	},
	Integers1R1: {
		symbols.NT_Integers, 1, 1, 
		symbols.Symbols{  
			symbols.T_22,
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
			symbols.T_23, 
			symbols.NT_WS,
		}, 
		LBRACE0R0, 
	},
	LBRACE0R1: {
		symbols.NT_LBRACE, 0, 1, 
		symbols.Symbols{  
			symbols.T_23, 
			symbols.NT_WS,
		}, 
		LBRACE0R1, 
	},
	LBRACE0R2: {
		symbols.NT_LBRACE, 0, 2, 
		symbols.Symbols{  
			symbols.T_23, 
			symbols.NT_WS,
		}, 
		LBRACE0R2, 
	},
	LBRACKET0R0: {
		symbols.NT_LBRACKET, 0, 0, 
		symbols.Symbols{  
			symbols.T_2, 
			symbols.NT_WS,
		}, 
		LBRACKET0R0, 
	},
	LBRACKET0R1: {
		symbols.NT_LBRACKET, 0, 1, 
		symbols.Symbols{  
			symbols.T_2, 
			symbols.NT_WS,
		}, 
		LBRACKET0R1, 
	},
	LBRACKET0R2: {
		symbols.NT_LBRACKET, 0, 2, 
		symbols.Symbols{  
			symbols.T_2, 
			symbols.NT_WS,
		}, 
		LBRACKET0R2, 
	},
	LineOrBlock0R0: {
		symbols.NT_LineOrBlock, 0, 0, 
		symbols.Symbols{  
			symbols.T_16,
		}, 
		LineOrBlock0R0, 
	},
	LineOrBlock0R1: {
		symbols.NT_LineOrBlock, 0, 1, 
		symbols.Symbols{  
			symbols.T_16,
		}, 
		LineOrBlock0R1, 
	},
	LineOrBlock1R0: {
		symbols.NT_LineOrBlock, 1, 0, 
		symbols.Symbols{  
			symbols.T_6,
		}, 
		LineOrBlock1R0, 
	},
	LineOrBlock1R1: {
		symbols.NT_LineOrBlock, 1, 1, 
		symbols.Symbols{  
			symbols.T_6,
		}, 
		LineOrBlock1R1, 
	},
	Members0R0: {
		symbols.NT_Members, 0, 0, 
		symbols.Symbols{  
			symbols.NT_Pair, 
			symbols.NT_ComPair0x,
		}, 
		Members0R0, 
	},
	Members0R1: {
		symbols.NT_Members, 0, 1, 
		symbols.Symbols{  
			symbols.NT_Pair, 
			symbols.NT_ComPair0x,
		}, 
		Members0R1, 
	},
	Members0R2: {
		symbols.NT_Members, 0, 2, 
		symbols.Symbols{  
			symbols.NT_Pair, 
			symbols.NT_ComPair0x,
		}, 
		Members0R2, 
	},
	Mems1x0R0: {
		symbols.NT_Mems1x, 0, 0, 
		symbols.Symbols{  
			symbols.NT_Members, 
			symbols.NT_Mems1x,
		}, 
		Mems1x0R0, 
	},
	Mems1x0R1: {
		symbols.NT_Mems1x, 0, 1, 
		symbols.Symbols{  
			symbols.NT_Members, 
			symbols.NT_Mems1x,
		}, 
		Mems1x0R1, 
	},
	Mems1x0R2: {
		symbols.NT_Mems1x, 0, 2, 
		symbols.Symbols{  
			symbols.NT_Members, 
			symbols.NT_Mems1x,
		}, 
		Mems1x0R2, 
	},
	Mems1x1R0: {
		symbols.NT_Mems1x, 1, 0, 
		symbols.Symbols{ 
		}, 
		Mems1x1R0, 
	},
	NUL0R0: {
		symbols.NT_NUL, 0, 0, 
		symbols.Symbols{  
			symbols.T_18, 
			symbols.NT_WS,
		}, 
		NUL0R0, 
	},
	NUL0R1: {
		symbols.NT_NUL, 0, 1, 
		symbols.Symbols{  
			symbols.T_18, 
			symbols.NT_WS,
		}, 
		NUL0R1, 
	},
	NUL0R2: {
		symbols.NT_NUL, 0, 2, 
		symbols.Symbols{  
			symbols.T_18, 
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
	NumberHEX0R0: {
		symbols.NT_NumberHEX, 0, 0, 
		symbols.Symbols{  
			symbols.NT_Number, 
			symbols.T_4,
		}, 
		NumberHEX0R0, 
	},
	NumberHEX0R1: {
		symbols.NT_NumberHEX, 0, 1, 
		symbols.Symbols{  
			symbols.NT_Number, 
			symbols.T_4,
		}, 
		NumberHEX0R1, 
	},
	NumberHEX0R2: {
		symbols.NT_NumberHEX, 0, 2, 
		symbols.Symbols{  
			symbols.NT_Number, 
			symbols.T_4,
		}, 
		NumberHEX0R2, 
	},
	NumberHEX1R0: {
		symbols.NT_NumberHEX, 1, 0, 
		symbols.Symbols{ 
		}, 
		NumberHEX1R0, 
	},
	Object0R0: {
		symbols.NT_Object, 0, 0, 
		symbols.Symbols{  
			symbols.NT_LBRACE, 
			symbols.NT_Members, 
			symbols.NT_Mems1x, 
			symbols.NT_RBRACE,
		}, 
		Object0R0, 
	},
	Object0R1: {
		symbols.NT_Object, 0, 1, 
		symbols.Symbols{  
			symbols.NT_LBRACE, 
			symbols.NT_Members, 
			symbols.NT_Mems1x, 
			symbols.NT_RBRACE,
		}, 
		Object0R1, 
	},
	Object0R2: {
		symbols.NT_Object, 0, 2, 
		symbols.Symbols{  
			symbols.NT_LBRACE, 
			symbols.NT_Members, 
			symbols.NT_Mems1x, 
			symbols.NT_RBRACE,
		}, 
		Object0R2, 
	},
	Object0R3: {
		symbols.NT_Object, 0, 3, 
		symbols.Symbols{  
			symbols.NT_LBRACE, 
			symbols.NT_Members, 
			symbols.NT_Mems1x, 
			symbols.NT_RBRACE,
		}, 
		Object0R3, 
	},
	Object0R4: {
		symbols.NT_Object, 0, 4, 
		symbols.Symbols{  
			symbols.NT_LBRACE, 
			symbols.NT_Members, 
			symbols.NT_Mems1x, 
			symbols.NT_RBRACE,
		}, 
		Object0R4, 
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
			symbols.T_12,
		}, 
		OptExp0R0, 
	},
	OptExp0R1: {
		symbols.NT_OptExp, 0, 1, 
		symbols.Symbols{  
			symbols.T_12,
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
			symbols.T_14,
		}, 
		OptFrac0R0, 
	},
	OptFrac0R1: {
		symbols.NT_OptFrac, 0, 1, 
		symbols.Symbols{  
			symbols.T_14,
		}, 
		OptFrac0R1, 
	},
	OptFrac1R0: {
		symbols.NT_OptFrac, 1, 0, 
		symbols.Symbols{ 
		}, 
		OptFrac1R0, 
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
	RBRACE0R0: {
		symbols.NT_RBRACE, 0, 0, 
		symbols.Symbols{  
			symbols.T_24, 
			symbols.NT_WS,
		}, 
		RBRACE0R0, 
	},
	RBRACE0R1: {
		symbols.NT_RBRACE, 0, 1, 
		symbols.Symbols{  
			symbols.T_24, 
			symbols.NT_WS,
		}, 
		RBRACE0R1, 
	},
	RBRACE0R2: {
		symbols.NT_RBRACE, 0, 2, 
		symbols.Symbols{  
			symbols.T_24, 
			symbols.NT_WS,
		}, 
		RBRACE0R2, 
	},
	RBRACKET0R0: {
		symbols.NT_RBRACKET, 0, 0, 
		symbols.Symbols{  
			symbols.T_3, 
			symbols.NT_WS,
		}, 
		RBRACKET0R0, 
	},
	RBRACKET0R1: {
		symbols.NT_RBRACKET, 0, 1, 
		symbols.Symbols{  
			symbols.T_3, 
			symbols.NT_WS,
		}, 
		RBRACKET0R1, 
	},
	RBRACKET0R2: {
		symbols.NT_RBRACKET, 0, 2, 
		symbols.Symbols{  
			symbols.T_3, 
			symbols.NT_WS,
		}, 
		RBRACKET0R2, 
	},
	String0R0: {
		symbols.NT_String, 0, 0, 
		symbols.Symbols{  
			symbols.T_8, 
			symbols.NT_Close, 
			symbols.NT_WS,
		}, 
		String0R0, 
	},
	String0R1: {
		symbols.NT_String, 0, 1, 
		symbols.Symbols{  
			symbols.T_8, 
			symbols.NT_Close, 
			symbols.NT_WS,
		}, 
		String0R1, 
	},
	String0R2: {
		symbols.NT_String, 0, 2, 
		symbols.Symbols{  
			symbols.T_8, 
			symbols.NT_Close, 
			symbols.NT_WS,
		}, 
		String0R2, 
	},
	String0R3: {
		symbols.NT_String, 0, 3, 
		symbols.Symbols{  
			symbols.T_8, 
			symbols.NT_Close, 
			symbols.NT_WS,
		}, 
		String0R3, 
	},
	TRUE0R0: {
		symbols.NT_TRUE, 0, 0, 
		symbols.Symbols{  
			symbols.T_20, 
			symbols.NT_WS,
		}, 
		TRUE0R0, 
	},
	TRUE0R1: {
		symbols.NT_TRUE, 0, 1, 
		symbols.Symbols{  
			symbols.T_20, 
			symbols.NT_WS,
		}, 
		TRUE0R1, 
	},
	TRUE0R2: {
		symbols.NT_TRUE, 0, 2, 
		symbols.Symbols{  
			symbols.T_20, 
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
	Index{ symbols.NT_CHAR,0,0 }: CHAR0R0,
	Index{ symbols.NT_CHAR,0,1 }: CHAR0R1,
	Index{ symbols.NT_CHAR,1,0 }: CHAR1R0,
	Index{ symbols.NT_CHAR,1,1 }: CHAR1R1,
	Index{ symbols.NT_CHAR,1,2 }: CHAR1R2,
	Index{ symbols.NT_COLON,0,0 }: COLON0R0,
	Index{ symbols.NT_COLON,0,1 }: COLON0R1,
	Index{ symbols.NT_COLON,0,2 }: COLON0R2,
	Index{ symbols.NT_COMMA,0,0 }: COMMA0R0,
	Index{ symbols.NT_COMMA,0,1 }: COMMA0R1,
	Index{ symbols.NT_COMMA,0,2 }: COMMA0R2,
	Index{ symbols.NT_CharCode,0,0 }: CharCode0R0,
	Index{ symbols.NT_CharCode,0,1 }: CharCode0R1,
	Index{ symbols.NT_CharCode,1,0 }: CharCode1R0,
	Index{ symbols.NT_CharCode,1,1 }: CharCode1R1,
	Index{ symbols.NT_CharCode,1,2 }: CharCode1R2,
	Index{ symbols.NT_CharCode,1,3 }: CharCode1R3,
	Index{ symbols.NT_CharCode,1,4 }: CharCode1R4,
	Index{ symbols.NT_CharCode,1,5 }: CharCode1R5,
	Index{ symbols.NT_Close,0,0 }: Close0R0,
	Index{ symbols.NT_Close,0,1 }: Close0R1,
	Index{ symbols.NT_Close,1,0 }: Close1R0,
	Index{ symbols.NT_Close,1,1 }: Close1R1,
	Index{ symbols.NT_Close,1,2 }: Close1R2,
	Index{ symbols.NT_ComPair,0,0 }: ComPair0R0,
	Index{ symbols.NT_ComPair,0,1 }: ComPair0R1,
	Index{ symbols.NT_ComPair,0,2 }: ComPair0R2,
	Index{ symbols.NT_ComPair0x,0,0 }: ComPair0x0R0,
	Index{ symbols.NT_ComPair0x,0,1 }: ComPair0x0R1,
	Index{ symbols.NT_ComPair0x,0,2 }: ComPair0x0R2,
	Index{ symbols.NT_ComPair0x,1,0 }: ComPair0x1R0,
	Index{ symbols.NT_ComVal,0,0 }: ComVal0R0,
	Index{ symbols.NT_ComVal,0,1 }: ComVal0R1,
	Index{ symbols.NT_ComVal,0,2 }: ComVal0R2,
	Index{ symbols.NT_ComVal0x,0,0 }: ComVal0x0R0,
	Index{ symbols.NT_ComVal0x,0,1 }: ComVal0x0R1,
	Index{ symbols.NT_ComVal0x,0,2 }: ComVal0x0R2,
	Index{ symbols.NT_ComVal0x,1,0 }: ComVal0x1R0,
	Index{ symbols.NT_Elements,0,0 }: Elements0R0,
	Index{ symbols.NT_Elements,0,1 }: Elements0R1,
	Index{ symbols.NT_Elements,0,2 }: Elements0R2,
	Index{ symbols.NT_EscOrComment,0,0 }: EscOrComment0R0,
	Index{ symbols.NT_EscOrComment,0,1 }: EscOrComment0R1,
	Index{ symbols.NT_EscOrComment,1,0 }: EscOrComment1R0,
	Index{ symbols.NT_EscOrComment,1,1 }: EscOrComment1R1,
	Index{ symbols.NT_FALSE,0,0 }: FALSE0R0,
	Index{ symbols.NT_FALSE,0,1 }: FALSE0R1,
	Index{ symbols.NT_FALSE,0,2 }: FALSE0R2,
	Index{ symbols.NT_HEX,0,0 }: HEX0R0,
	Index{ symbols.NT_HEX,0,1 }: HEX0R1,
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
	Index{ symbols.NT_LineOrBlock,0,0 }: LineOrBlock0R0,
	Index{ symbols.NT_LineOrBlock,0,1 }: LineOrBlock0R1,
	Index{ symbols.NT_LineOrBlock,1,0 }: LineOrBlock1R0,
	Index{ symbols.NT_LineOrBlock,1,1 }: LineOrBlock1R1,
	Index{ symbols.NT_Members,0,0 }: Members0R0,
	Index{ symbols.NT_Members,0,1 }: Members0R1,
	Index{ symbols.NT_Members,0,2 }: Members0R2,
	Index{ symbols.NT_Mems1x,0,0 }: Mems1x0R0,
	Index{ symbols.NT_Mems1x,0,1 }: Mems1x0R1,
	Index{ symbols.NT_Mems1x,0,2 }: Mems1x0R2,
	Index{ symbols.NT_Mems1x,1,0 }: Mems1x1R0,
	Index{ symbols.NT_NUL,0,0 }: NUL0R0,
	Index{ symbols.NT_NUL,0,1 }: NUL0R1,
	Index{ symbols.NT_NUL,0,2 }: NUL0R2,
	Index{ symbols.NT_Number,0,0 }: Number0R0,
	Index{ symbols.NT_Number,0,1 }: Number0R1,
	Index{ symbols.NT_Number,0,2 }: Number0R2,
	Index{ symbols.NT_Number,0,3 }: Number0R3,
	Index{ symbols.NT_Number,0,4 }: Number0R4,
	Index{ symbols.NT_NumberHEX,0,0 }: NumberHEX0R0,
	Index{ symbols.NT_NumberHEX,0,1 }: NumberHEX0R1,
	Index{ symbols.NT_NumberHEX,0,2 }: NumberHEX0R2,
	Index{ symbols.NT_NumberHEX,1,0 }: NumberHEX1R0,
	Index{ symbols.NT_Object,0,0 }: Object0R0,
	Index{ symbols.NT_Object,0,1 }: Object0R1,
	Index{ symbols.NT_Object,0,2 }: Object0R2,
	Index{ symbols.NT_Object,0,3 }: Object0R3,
	Index{ symbols.NT_Object,0,4 }: Object0R4,
	Index{ symbols.NT_OptElem,0,0 }: OptElem0R0,
	Index{ symbols.NT_OptElem,0,1 }: OptElem0R1,
	Index{ symbols.NT_OptElem,1,0 }: OptElem1R0,
	Index{ symbols.NT_OptExp,0,0 }: OptExp0R0,
	Index{ symbols.NT_OptExp,0,1 }: OptExp0R1,
	Index{ symbols.NT_OptExp,1,0 }: OptExp1R0,
	Index{ symbols.NT_OptFrac,0,0 }: OptFrac0R0,
	Index{ symbols.NT_OptFrac,0,1 }: OptFrac0R1,
	Index{ symbols.NT_OptFrac,1,0 }: OptFrac1R0,
	Index{ symbols.NT_Pair,0,0 }: Pair0R0,
	Index{ symbols.NT_Pair,0,1 }: Pair0R1,
	Index{ symbols.NT_Pair,0,2 }: Pair0R2,
	Index{ symbols.NT_Pair,0,3 }: Pair0R3,
	Index{ symbols.NT_RBRACE,0,0 }: RBRACE0R0,
	Index{ symbols.NT_RBRACE,0,1 }: RBRACE0R1,
	Index{ symbols.NT_RBRACE,0,2 }: RBRACE0R2,
	Index{ symbols.NT_RBRACKET,0,0 }: RBRACKET0R0,
	Index{ symbols.NT_RBRACKET,0,1 }: RBRACKET0R1,
	Index{ symbols.NT_RBRACKET,0,2 }: RBRACKET0R2,
	Index{ symbols.NT_String,0,0 }: String0R0,
	Index{ symbols.NT_String,0,1 }: String0R1,
	Index{ symbols.NT_String,0,2 }: String0R2,
	Index{ symbols.NT_String,0,3 }: String0R3,
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
	symbols.NT_Mems1x:[]Label{ Mems1x0R0,Mems1x1R0 },
	symbols.NT_Members:[]Label{ Members0R0 },
	symbols.NT_ComPair0x:[]Label{ ComPair0x0R0,ComPair0x1R0 },
	symbols.NT_ComPair:[]Label{ ComPair0R0 },
	symbols.NT_Pair:[]Label{ Pair0R0 },
	symbols.NT_Array:[]Label{ Array0R0 },
	symbols.NT_OptElem:[]Label{ OptElem0R0,OptElem1R0 },
	symbols.NT_Elements:[]Label{ Elements0R0 },
	symbols.NT_ComVal0x:[]Label{ ComVal0x0R0,ComVal0x1R0 },
	symbols.NT_ComVal:[]Label{ ComVal0R0 },
	symbols.NT_Value:[]Label{ Value0R0,Value1R0,Value2R0,Value3R0,Value4R0,Value5R0,Value6R0 },
	symbols.NT_String:[]Label{ String0R0 },
	symbols.NT_Close:[]Label{ Close0R0,Close1R0 },
	symbols.NT_CHAR:[]Label{ CHAR0R0,CHAR1R0 },
	symbols.NT_CharCode:[]Label{ CharCode0R0,CharCode1R0 },
	symbols.NT_HEX:[]Label{ HEX0R0 },
	symbols.NT_NumberHEX:[]Label{ NumberHEX0R0,NumberHEX1R0 },
	symbols.NT_Number:[]Label{ Number0R0 },
	symbols.NT_OptFrac:[]Label{ OptFrac0R0,OptFrac1R0 },
	symbols.NT_OptExp:[]Label{ OptExp0R0,OptExp1R0 },
	symbols.NT_INT:[]Label{ INT0R0 },
	symbols.NT_Integers:[]Label{ Integers0R0,Integers1R0 },
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
	symbols.NT_EscOrComment:[]Label{ EscOrComment0R0,EscOrComment1R0 },
	symbols.NT_LineOrBlock:[]Label{ LineOrBlock0R0,LineOrBlock1R0 },
}

