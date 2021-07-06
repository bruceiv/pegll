
// Package slot is generated by gogll. Do not edit. 
package slot

import(
	"bytes"
	"fmt"
	
	"calc/parser/symbols"
	"calc/token"
)

type Label int

const(
	CLOSE0R0 Label = iota
	CLOSE0R1
	CLOSE0R2
	DIVIDE0R0
	DIVIDE0R1
	DIVIDE0R2
	ELEMENT0R0
	ELEMENT0R1
	ELEMENT0R2
	ELEMENT0R3
	ELEMENT1R0
	ELEMENT1R1
	EXPR0R0
	EXPR0R1
	EXPR0R2
	MINUS0R0
	MINUS0R1
	MINUS0R2
	Number0R0
	Number0R1
	Number0R2
	OPEN0R0
	OPEN0R1
	OPEN0R2
	PLUS0R0
	PLUS0R1
	PLUS0R2
	PLUSorMINUS0R0
	PLUSorMINUS0R1
	PLUSorMINUS0R2
	PLUSorMINUS1R0
	PLUSorMINUS1R1
	PLUSorMINUS1R2
	PRODUCT0R0
	PRODUCT0R1
	PRODUCT0R2
	RepPLUSorMINUS0x0R0
	RepPLUSorMINUS0x0R1
	RepPLUSorMINUS0x0R2
	RepPLUSorMINUS0x1R0
	RepTIMESorDIV0x0R0
	RepTIMESorDIV0x0R1
	RepTIMESorDIV0x0R2
	RepTIMESorDIV0x1R0
	SUM0R0
	SUM0R1
	SUM0R2
	TIMES0R0
	TIMES0R1
	TIMES0R2
	TIMESorDIVIDE0R0
	TIMESorDIVIDE0R1
	TIMESorDIVIDE0R2
	TIMESorDIVIDE1R0
	TIMESorDIVIDE1R1
	TIMESorDIVIDE1R2
	WS0R0
	WS0R1
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
	CLOSE0R0: {
		symbols.NT_CLOSE, 0, 0, 
		symbols.Symbols{  
			symbols.T_1, 
			symbols.NT_WS,
		}, 
		CLOSE0R0, 
	},
	CLOSE0R1: {
		symbols.NT_CLOSE, 0, 1, 
		symbols.Symbols{  
			symbols.T_1, 
			symbols.NT_WS,
		}, 
		CLOSE0R1, 
	},
	CLOSE0R2: {
		symbols.NT_CLOSE, 0, 2, 
		symbols.Symbols{  
			symbols.T_1, 
			symbols.NT_WS,
		}, 
		CLOSE0R2, 
	},
	DIVIDE0R0: {
		symbols.NT_DIVIDE, 0, 0, 
		symbols.Symbols{  
			symbols.T_5, 
			symbols.NT_WS,
		}, 
		DIVIDE0R0, 
	},
	DIVIDE0R1: {
		symbols.NT_DIVIDE, 0, 1, 
		symbols.Symbols{  
			symbols.T_5, 
			symbols.NT_WS,
		}, 
		DIVIDE0R1, 
	},
	DIVIDE0R2: {
		symbols.NT_DIVIDE, 0, 2, 
		symbols.Symbols{  
			symbols.T_5, 
			symbols.NT_WS,
		}, 
		DIVIDE0R2, 
	},
	ELEMENT0R0: {
		symbols.NT_ELEMENT, 0, 0, 
		symbols.Symbols{  
			symbols.NT_OPEN, 
			symbols.NT_SUM, 
			symbols.NT_CLOSE,
		}, 
		ELEMENT0R0, 
	},
	ELEMENT0R1: {
		symbols.NT_ELEMENT, 0, 1, 
		symbols.Symbols{  
			symbols.NT_OPEN, 
			symbols.NT_SUM, 
			symbols.NT_CLOSE,
		}, 
		ELEMENT0R1, 
	},
	ELEMENT0R2: {
		symbols.NT_ELEMENT, 0, 2, 
		symbols.Symbols{  
			symbols.NT_OPEN, 
			symbols.NT_SUM, 
			symbols.NT_CLOSE,
		}, 
		ELEMENT0R2, 
	},
	ELEMENT0R3: {
		symbols.NT_ELEMENT, 0, 3, 
		symbols.Symbols{  
			symbols.NT_OPEN, 
			symbols.NT_SUM, 
			symbols.NT_CLOSE,
		}, 
		ELEMENT0R3, 
	},
	ELEMENT1R0: {
		symbols.NT_ELEMENT, 1, 0, 
		symbols.Symbols{  
			symbols.NT_Number,
		}, 
		ELEMENT1R0, 
	},
	ELEMENT1R1: {
		symbols.NT_ELEMENT, 1, 1, 
		symbols.Symbols{  
			symbols.NT_Number,
		}, 
		ELEMENT1R1, 
	},
	EXPR0R0: {
		symbols.NT_EXPR, 0, 0, 
		symbols.Symbols{  
			symbols.NT_WS, 
			symbols.NT_SUM,
		}, 
		EXPR0R0, 
	},
	EXPR0R1: {
		symbols.NT_EXPR, 0, 1, 
		symbols.Symbols{  
			symbols.NT_WS, 
			symbols.NT_SUM,
		}, 
		EXPR0R1, 
	},
	EXPR0R2: {
		symbols.NT_EXPR, 0, 2, 
		symbols.Symbols{  
			symbols.NT_WS, 
			symbols.NT_SUM,
		}, 
		EXPR0R2, 
	},
	MINUS0R0: {
		symbols.NT_MINUS, 0, 0, 
		symbols.Symbols{  
			symbols.T_4, 
			symbols.NT_WS,
		}, 
		MINUS0R0, 
	},
	MINUS0R1: {
		symbols.NT_MINUS, 0, 1, 
		symbols.Symbols{  
			symbols.T_4, 
			symbols.NT_WS,
		}, 
		MINUS0R1, 
	},
	MINUS0R2: {
		symbols.NT_MINUS, 0, 2, 
		symbols.Symbols{  
			symbols.T_4, 
			symbols.NT_WS,
		}, 
		MINUS0R2, 
	},
	Number0R0: {
		symbols.NT_Number, 0, 0, 
		symbols.Symbols{  
			symbols.T_6, 
			symbols.NT_WS,
		}, 
		Number0R0, 
	},
	Number0R1: {
		symbols.NT_Number, 0, 1, 
		symbols.Symbols{  
			symbols.T_6, 
			symbols.NT_WS,
		}, 
		Number0R1, 
	},
	Number0R2: {
		symbols.NT_Number, 0, 2, 
		symbols.Symbols{  
			symbols.T_6, 
			symbols.NT_WS,
		}, 
		Number0R2, 
	},
	OPEN0R0: {
		symbols.NT_OPEN, 0, 0, 
		symbols.Symbols{  
			symbols.T_0, 
			symbols.NT_WS,
		}, 
		OPEN0R0, 
	},
	OPEN0R1: {
		symbols.NT_OPEN, 0, 1, 
		symbols.Symbols{  
			symbols.T_0, 
			symbols.NT_WS,
		}, 
		OPEN0R1, 
	},
	OPEN0R2: {
		symbols.NT_OPEN, 0, 2, 
		symbols.Symbols{  
			symbols.T_0, 
			symbols.NT_WS,
		}, 
		OPEN0R2, 
	},
	PLUS0R0: {
		symbols.NT_PLUS, 0, 0, 
		symbols.Symbols{  
			symbols.T_3, 
			symbols.NT_WS,
		}, 
		PLUS0R0, 
	},
	PLUS0R1: {
		symbols.NT_PLUS, 0, 1, 
		symbols.Symbols{  
			symbols.T_3, 
			symbols.NT_WS,
		}, 
		PLUS0R1, 
	},
	PLUS0R2: {
		symbols.NT_PLUS, 0, 2, 
		symbols.Symbols{  
			symbols.T_3, 
			symbols.NT_WS,
		}, 
		PLUS0R2, 
	},
	PLUSorMINUS0R0: {
		symbols.NT_PLUSorMINUS, 0, 0, 
		symbols.Symbols{  
			symbols.NT_PLUS, 
			symbols.NT_PRODUCT,
		}, 
		PLUSorMINUS0R0, 
	},
	PLUSorMINUS0R1: {
		symbols.NT_PLUSorMINUS, 0, 1, 
		symbols.Symbols{  
			symbols.NT_PLUS, 
			symbols.NT_PRODUCT,
		}, 
		PLUSorMINUS0R1, 
	},
	PLUSorMINUS0R2: {
		symbols.NT_PLUSorMINUS, 0, 2, 
		symbols.Symbols{  
			symbols.NT_PLUS, 
			symbols.NT_PRODUCT,
		}, 
		PLUSorMINUS0R2, 
	},
	PLUSorMINUS1R0: {
		symbols.NT_PLUSorMINUS, 1, 0, 
		symbols.Symbols{  
			symbols.NT_MINUS, 
			symbols.NT_PRODUCT,
		}, 
		PLUSorMINUS1R0, 
	},
	PLUSorMINUS1R1: {
		symbols.NT_PLUSorMINUS, 1, 1, 
		symbols.Symbols{  
			symbols.NT_MINUS, 
			symbols.NT_PRODUCT,
		}, 
		PLUSorMINUS1R1, 
	},
	PLUSorMINUS1R2: {
		symbols.NT_PLUSorMINUS, 1, 2, 
		symbols.Symbols{  
			symbols.NT_MINUS, 
			symbols.NT_PRODUCT,
		}, 
		PLUSorMINUS1R2, 
	},
	PRODUCT0R0: {
		symbols.NT_PRODUCT, 0, 0, 
		symbols.Symbols{  
			symbols.NT_ELEMENT, 
			symbols.NT_RepTIMESorDIV0x,
		}, 
		PRODUCT0R0, 
	},
	PRODUCT0R1: {
		symbols.NT_PRODUCT, 0, 1, 
		symbols.Symbols{  
			symbols.NT_ELEMENT, 
			symbols.NT_RepTIMESorDIV0x,
		}, 
		PRODUCT0R1, 
	},
	PRODUCT0R2: {
		symbols.NT_PRODUCT, 0, 2, 
		symbols.Symbols{  
			symbols.NT_ELEMENT, 
			symbols.NT_RepTIMESorDIV0x,
		}, 
		PRODUCT0R2, 
	},
	RepPLUSorMINUS0x0R0: {
		symbols.NT_RepPLUSorMINUS0x, 0, 0, 
		symbols.Symbols{  
			symbols.NT_PLUSorMINUS, 
			symbols.NT_RepPLUSorMINUS0x,
		}, 
		RepPLUSorMINUS0x0R0, 
	},
	RepPLUSorMINUS0x0R1: {
		symbols.NT_RepPLUSorMINUS0x, 0, 1, 
		symbols.Symbols{  
			symbols.NT_PLUSorMINUS, 
			symbols.NT_RepPLUSorMINUS0x,
		}, 
		RepPLUSorMINUS0x0R1, 
	},
	RepPLUSorMINUS0x0R2: {
		symbols.NT_RepPLUSorMINUS0x, 0, 2, 
		symbols.Symbols{  
			symbols.NT_PLUSorMINUS, 
			symbols.NT_RepPLUSorMINUS0x,
		}, 
		RepPLUSorMINUS0x0R2, 
	},
	RepPLUSorMINUS0x1R0: {
		symbols.NT_RepPLUSorMINUS0x, 1, 0, 
		symbols.Symbols{ 
		}, 
		RepPLUSorMINUS0x1R0, 
	},
	RepTIMESorDIV0x0R0: {
		symbols.NT_RepTIMESorDIV0x, 0, 0, 
		symbols.Symbols{  
			symbols.NT_TIMESorDIVIDE, 
			symbols.NT_RepTIMESorDIV0x,
		}, 
		RepTIMESorDIV0x0R0, 
	},
	RepTIMESorDIV0x0R1: {
		symbols.NT_RepTIMESorDIV0x, 0, 1, 
		symbols.Symbols{  
			symbols.NT_TIMESorDIVIDE, 
			symbols.NT_RepTIMESorDIV0x,
		}, 
		RepTIMESorDIV0x0R1, 
	},
	RepTIMESorDIV0x0R2: {
		symbols.NT_RepTIMESorDIV0x, 0, 2, 
		symbols.Symbols{  
			symbols.NT_TIMESorDIVIDE, 
			symbols.NT_RepTIMESorDIV0x,
		}, 
		RepTIMESorDIV0x0R2, 
	},
	RepTIMESorDIV0x1R0: {
		symbols.NT_RepTIMESorDIV0x, 1, 0, 
		symbols.Symbols{ 
		}, 
		RepTIMESorDIV0x1R0, 
	},
	SUM0R0: {
		symbols.NT_SUM, 0, 0, 
		symbols.Symbols{  
			symbols.NT_PRODUCT, 
			symbols.NT_RepPLUSorMINUS0x,
		}, 
		SUM0R0, 
	},
	SUM0R1: {
		symbols.NT_SUM, 0, 1, 
		symbols.Symbols{  
			symbols.NT_PRODUCT, 
			symbols.NT_RepPLUSorMINUS0x,
		}, 
		SUM0R1, 
	},
	SUM0R2: {
		symbols.NT_SUM, 0, 2, 
		symbols.Symbols{  
			symbols.NT_PRODUCT, 
			symbols.NT_RepPLUSorMINUS0x,
		}, 
		SUM0R2, 
	},
	TIMES0R0: {
		symbols.NT_TIMES, 0, 0, 
		symbols.Symbols{  
			symbols.T_2, 
			symbols.NT_WS,
		}, 
		TIMES0R0, 
	},
	TIMES0R1: {
		symbols.NT_TIMES, 0, 1, 
		symbols.Symbols{  
			symbols.T_2, 
			symbols.NT_WS,
		}, 
		TIMES0R1, 
	},
	TIMES0R2: {
		symbols.NT_TIMES, 0, 2, 
		symbols.Symbols{  
			symbols.T_2, 
			symbols.NT_WS,
		}, 
		TIMES0R2, 
	},
	TIMESorDIVIDE0R0: {
		symbols.NT_TIMESorDIVIDE, 0, 0, 
		symbols.Symbols{  
			symbols.NT_TIMES, 
			symbols.NT_ELEMENT,
		}, 
		TIMESorDIVIDE0R0, 
	},
	TIMESorDIVIDE0R1: {
		symbols.NT_TIMESorDIVIDE, 0, 1, 
		symbols.Symbols{  
			symbols.NT_TIMES, 
			symbols.NT_ELEMENT,
		}, 
		TIMESorDIVIDE0R1, 
	},
	TIMESorDIVIDE0R2: {
		symbols.NT_TIMESorDIVIDE, 0, 2, 
		symbols.Symbols{  
			symbols.NT_TIMES, 
			symbols.NT_ELEMENT,
		}, 
		TIMESorDIVIDE0R2, 
	},
	TIMESorDIVIDE1R0: {
		symbols.NT_TIMESorDIVIDE, 1, 0, 
		symbols.Symbols{  
			symbols.NT_DIVIDE, 
			symbols.NT_ELEMENT,
		}, 
		TIMESorDIVIDE1R0, 
	},
	TIMESorDIVIDE1R1: {
		symbols.NT_TIMESorDIVIDE, 1, 1, 
		symbols.Symbols{  
			symbols.NT_DIVIDE, 
			symbols.NT_ELEMENT,
		}, 
		TIMESorDIVIDE1R1, 
	},
	TIMESorDIVIDE1R2: {
		symbols.NT_TIMESorDIVIDE, 1, 2, 
		symbols.Symbols{  
			symbols.NT_DIVIDE, 
			symbols.NT_ELEMENT,
		}, 
		TIMESorDIVIDE1R2, 
	},
	WS0R0: {
		symbols.NT_WS, 0, 0, 
		symbols.Symbols{  
			symbols.T_7,
		}, 
		WS0R0, 
	},
	WS0R1: {
		symbols.NT_WS, 0, 1, 
		symbols.Symbols{  
			symbols.T_7,
		}, 
		WS0R1, 
	},
	WS1R0: {
		symbols.NT_WS, 1, 0, 
		symbols.Symbols{ 
		}, 
		WS1R0, 
	},
}

var slotIndex = map[Index]Label { 
	Index{ symbols.NT_CLOSE,0,0 }: CLOSE0R0,
	Index{ symbols.NT_CLOSE,0,1 }: CLOSE0R1,
	Index{ symbols.NT_CLOSE,0,2 }: CLOSE0R2,
	Index{ symbols.NT_DIVIDE,0,0 }: DIVIDE0R0,
	Index{ symbols.NT_DIVIDE,0,1 }: DIVIDE0R1,
	Index{ symbols.NT_DIVIDE,0,2 }: DIVIDE0R2,
	Index{ symbols.NT_ELEMENT,0,0 }: ELEMENT0R0,
	Index{ symbols.NT_ELEMENT,0,1 }: ELEMENT0R1,
	Index{ symbols.NT_ELEMENT,0,2 }: ELEMENT0R2,
	Index{ symbols.NT_ELEMENT,0,3 }: ELEMENT0R3,
	Index{ symbols.NT_ELEMENT,1,0 }: ELEMENT1R0,
	Index{ symbols.NT_ELEMENT,1,1 }: ELEMENT1R1,
	Index{ symbols.NT_EXPR,0,0 }: EXPR0R0,
	Index{ symbols.NT_EXPR,0,1 }: EXPR0R1,
	Index{ symbols.NT_EXPR,0,2 }: EXPR0R2,
	Index{ symbols.NT_MINUS,0,0 }: MINUS0R0,
	Index{ symbols.NT_MINUS,0,1 }: MINUS0R1,
	Index{ symbols.NT_MINUS,0,2 }: MINUS0R2,
	Index{ symbols.NT_Number,0,0 }: Number0R0,
	Index{ symbols.NT_Number,0,1 }: Number0R1,
	Index{ symbols.NT_Number,0,2 }: Number0R2,
	Index{ symbols.NT_OPEN,0,0 }: OPEN0R0,
	Index{ symbols.NT_OPEN,0,1 }: OPEN0R1,
	Index{ symbols.NT_OPEN,0,2 }: OPEN0R2,
	Index{ symbols.NT_PLUS,0,0 }: PLUS0R0,
	Index{ symbols.NT_PLUS,0,1 }: PLUS0R1,
	Index{ symbols.NT_PLUS,0,2 }: PLUS0R2,
	Index{ symbols.NT_PLUSorMINUS,0,0 }: PLUSorMINUS0R0,
	Index{ symbols.NT_PLUSorMINUS,0,1 }: PLUSorMINUS0R1,
	Index{ symbols.NT_PLUSorMINUS,0,2 }: PLUSorMINUS0R2,
	Index{ symbols.NT_PLUSorMINUS,1,0 }: PLUSorMINUS1R0,
	Index{ symbols.NT_PLUSorMINUS,1,1 }: PLUSorMINUS1R1,
	Index{ symbols.NT_PLUSorMINUS,1,2 }: PLUSorMINUS1R2,
	Index{ symbols.NT_PRODUCT,0,0 }: PRODUCT0R0,
	Index{ symbols.NT_PRODUCT,0,1 }: PRODUCT0R1,
	Index{ symbols.NT_PRODUCT,0,2 }: PRODUCT0R2,
	Index{ symbols.NT_RepPLUSorMINUS0x,0,0 }: RepPLUSorMINUS0x0R0,
	Index{ symbols.NT_RepPLUSorMINUS0x,0,1 }: RepPLUSorMINUS0x0R1,
	Index{ symbols.NT_RepPLUSorMINUS0x,0,2 }: RepPLUSorMINUS0x0R2,
	Index{ symbols.NT_RepPLUSorMINUS0x,1,0 }: RepPLUSorMINUS0x1R0,
	Index{ symbols.NT_RepTIMESorDIV0x,0,0 }: RepTIMESorDIV0x0R0,
	Index{ symbols.NT_RepTIMESorDIV0x,0,1 }: RepTIMESorDIV0x0R1,
	Index{ symbols.NT_RepTIMESorDIV0x,0,2 }: RepTIMESorDIV0x0R2,
	Index{ symbols.NT_RepTIMESorDIV0x,1,0 }: RepTIMESorDIV0x1R0,
	Index{ symbols.NT_SUM,0,0 }: SUM0R0,
	Index{ symbols.NT_SUM,0,1 }: SUM0R1,
	Index{ symbols.NT_SUM,0,2 }: SUM0R2,
	Index{ symbols.NT_TIMES,0,0 }: TIMES0R0,
	Index{ symbols.NT_TIMES,0,1 }: TIMES0R1,
	Index{ symbols.NT_TIMES,0,2 }: TIMES0R2,
	Index{ symbols.NT_TIMESorDIVIDE,0,0 }: TIMESorDIVIDE0R0,
	Index{ symbols.NT_TIMESorDIVIDE,0,1 }: TIMESorDIVIDE0R1,
	Index{ symbols.NT_TIMESorDIVIDE,0,2 }: TIMESorDIVIDE0R2,
	Index{ symbols.NT_TIMESorDIVIDE,1,0 }: TIMESorDIVIDE1R0,
	Index{ symbols.NT_TIMESorDIVIDE,1,1 }: TIMESorDIVIDE1R1,
	Index{ symbols.NT_TIMESorDIVIDE,1,2 }: TIMESorDIVIDE1R2,
	Index{ symbols.NT_WS,0,0 }: WS0R0,
	Index{ symbols.NT_WS,0,1 }: WS0R1,
	Index{ symbols.NT_WS,1,0 }: WS1R0,
}

var alternates = map[symbols.NT][]Label{ 
	symbols.NT_EXPR:[]Label{ EXPR0R0 },
	symbols.NT_SUM:[]Label{ SUM0R0 },
	symbols.NT_RepPLUSorMINUS0x:[]Label{ RepPLUSorMINUS0x0R0,RepPLUSorMINUS0x1R0 },
	symbols.NT_PLUSorMINUS:[]Label{ PLUSorMINUS0R0,PLUSorMINUS1R0 },
	symbols.NT_PRODUCT:[]Label{ PRODUCT0R0 },
	symbols.NT_RepTIMESorDIV0x:[]Label{ RepTIMESorDIV0x0R0,RepTIMESorDIV0x1R0 },
	symbols.NT_TIMESorDIVIDE:[]Label{ TIMESorDIVIDE0R0,TIMESorDIVIDE1R0 },
	symbols.NT_ELEMENT:[]Label{ ELEMENT0R0,ELEMENT1R0 },
	symbols.NT_Number:[]Label{ Number0R0 },
	symbols.NT_PLUS:[]Label{ PLUS0R0 },
	symbols.NT_MINUS:[]Label{ MINUS0R0 },
	symbols.NT_TIMES:[]Label{ TIMES0R0 },
	symbols.NT_DIVIDE:[]Label{ DIVIDE0R0 },
	symbols.NT_OPEN:[]Label{ OPEN0R0 },
	symbols.NT_CLOSE:[]Label{ CLOSE0R0 },
	symbols.NT_WS:[]Label{ WS0R0,WS1R0 },
}

var nullable = []bool { 
	false, // CLOSE0R0 
	true, // CLOSE0R1 
	true, // CLOSE0R2 
	false, // DIVIDE0R0 
	true, // DIVIDE0R1 
	true, // DIVIDE0R2 
	false, // ELEMENT0R0 
	false, // ELEMENT0R1 
	false, // ELEMENT0R2 
	true, // ELEMENT0R3 
	false, // ELEMENT1R0 
	true, // ELEMENT1R1 
	false, // EXPR0R0 
	false, // EXPR0R1 
	true, // EXPR0R2 
	false, // MINUS0R0 
	true, // MINUS0R1 
	true, // MINUS0R2 
	false, // Number0R0 
	true, // Number0R1 
	true, // Number0R2 
	false, // OPEN0R0 
	true, // OPEN0R1 
	true, // OPEN0R2 
	false, // PLUS0R0 
	true, // PLUS0R1 
	true, // PLUS0R2 
	false, // PLUSorMINUS0R0 
	false, // PLUSorMINUS0R1 
	true, // PLUSorMINUS0R2 
	false, // PLUSorMINUS1R0 
	false, // PLUSorMINUS1R1 
	true, // PLUSorMINUS1R2 
	false, // PRODUCT0R0 
	true, // PRODUCT0R1 
	true, // PRODUCT0R2 
	false, // RepPLUSorMINUS0x0R0 
	true, // RepPLUSorMINUS0x0R1 
	true, // RepPLUSorMINUS0x0R2 
	true, // RepPLUSorMINUS0x1R0 
	false, // RepTIMESorDIV0x0R0 
	true, // RepTIMESorDIV0x0R1 
	true, // RepTIMESorDIV0x0R2 
	true, // RepTIMESorDIV0x1R0 
	false, // SUM0R0 
	true, // SUM0R1 
	true, // SUM0R2 
	false, // TIMES0R0 
	true, // TIMES0R1 
	true, // TIMES0R2 
	false, // TIMESorDIVIDE0R0 
	false, // TIMESorDIVIDE0R1 
	true, // TIMESorDIVIDE0R2 
	false, // TIMESorDIVIDE1R0 
	false, // TIMESorDIVIDE1R1 
	true, // TIMESorDIVIDE1R2 
	false, // WS0R0 
	true, // WS0R1 
	true, // WS1R0 
}

var firstT = []map[token.Type]bool { 
	{  token.T_1: true,  }, // CLOSE0R0 
	{  token.T_7: true,  }, // CLOSE0R1 
	{  }, // CLOSE0R2 
	{  token.T_5: true,  }, // DIVIDE0R0 
	{  token.T_7: true,  }, // DIVIDE0R1 
	{  }, // DIVIDE0R2 
	{  token.T_0: true,  }, // ELEMENT0R0 
	{  token.T_0: true,  token.T_6: true,  }, // ELEMENT0R1 
	{  token.T_1: true,  }, // ELEMENT0R2 
	{  }, // ELEMENT0R3 
	{  token.T_6: true,  }, // ELEMENT1R0 
	{  }, // ELEMENT1R1 
	{  token.T_6: true,  token.T_0: true,  token.T_7: true,  }, // EXPR0R0 
	{  token.T_6: true,  token.T_0: true,  }, // EXPR0R1 
	{  }, // EXPR0R2 
	{  token.T_4: true,  }, // MINUS0R0 
	{  token.T_7: true,  }, // MINUS0R1 
	{  }, // MINUS0R2 
	{  token.T_6: true,  }, // Number0R0 
	{  token.T_7: true,  }, // Number0R1 
	{  }, // Number0R2 
	{  token.T_0: true,  }, // OPEN0R0 
	{  token.T_7: true,  }, // OPEN0R1 
	{  }, // OPEN0R2 
	{  token.T_3: true,  }, // PLUS0R0 
	{  token.T_7: true,  }, // PLUS0R1 
	{  }, // PLUS0R2 
	{  token.T_3: true,  }, // PLUSorMINUS0R0 
	{  token.T_6: true,  token.T_0: true,  }, // PLUSorMINUS0R1 
	{  }, // PLUSorMINUS0R2 
	{  token.T_4: true,  }, // PLUSorMINUS1R0 
	{  token.T_0: true,  token.T_6: true,  }, // PLUSorMINUS1R1 
	{  }, // PLUSorMINUS1R2 
	{  token.T_0: true,  token.T_6: true,  }, // PRODUCT0R0 
	{  token.T_5: true,  token.T_2: true,  }, // PRODUCT0R1 
	{  }, // PRODUCT0R2 
	{  token.T_3: true,  token.T_4: true,  }, // RepPLUSorMINUS0x0R0 
	{  token.T_4: true,  token.T_3: true,  }, // RepPLUSorMINUS0x0R1 
	{  }, // RepPLUSorMINUS0x0R2 
	{  }, // RepPLUSorMINUS0x1R0 
	{  token.T_5: true,  token.T_2: true,  }, // RepTIMESorDIV0x0R0 
	{  token.T_5: true,  token.T_2: true,  }, // RepTIMESorDIV0x0R1 
	{  }, // RepTIMESorDIV0x0R2 
	{  }, // RepTIMESorDIV0x1R0 
	{  token.T_0: true,  token.T_6: true,  }, // SUM0R0 
	{  token.T_4: true,  token.T_3: true,  }, // SUM0R1 
	{  }, // SUM0R2 
	{  token.T_2: true,  }, // TIMES0R0 
	{  token.T_7: true,  }, // TIMES0R1 
	{  }, // TIMES0R2 
	{  token.T_2: true,  }, // TIMESorDIVIDE0R0 
	{  token.T_6: true,  token.T_0: true,  }, // TIMESorDIVIDE0R1 
	{  }, // TIMESorDIVIDE0R2 
	{  token.T_5: true,  }, // TIMESorDIVIDE1R0 
	{  token.T_0: true,  token.T_6: true,  }, // TIMESorDIVIDE1R1 
	{  }, // TIMESorDIVIDE1R2 
	{  token.T_7: true,  }, // WS0R0 
	{  }, // WS0R1 
	{  }, // WS1R0 
}
