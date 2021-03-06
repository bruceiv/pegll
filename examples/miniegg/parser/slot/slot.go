
// Package slot is generated by gogll. Do not edit. 
package slot

import(
	"bytes"
	"fmt"
	
	"miniegg/parser/symbols"
	"miniegg/token"
)

type Label int

const(
	Expr0R0 Label = iota
	Expr0R1
	Expr0R2
	Grammar0R0
	Grammar0R1
	Grammar0R2
	Grammar0R3
	RepExpr0x0R0
	RepExpr0x0R1
	RepExpr0x0R2
	RepExpr0x1R0
	RepRule0x0R0
	RepRule0x0R1
	RepRule0x0R2
	RepRule0x1R0
	Rule0R0
	Rule0R1
	Rule0R2
	Rule0R3
	Rule0R4
	Rule0R5
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
	Expr0R0: {
		symbols.NT_Expr, 0, 0, 
		symbols.Symbols{  
			symbols.T_2, 
			symbols.T_3,
		}, 
		Expr0R0, 
	},
	Expr0R1: {
		symbols.NT_Expr, 0, 1, 
		symbols.Symbols{  
			symbols.T_2, 
			symbols.T_3,
		}, 
		Expr0R1, 
	},
	Expr0R2: {
		symbols.NT_Expr, 0, 2, 
		symbols.Symbols{  
			symbols.T_2, 
			symbols.T_3,
		}, 
		Expr0R2, 
	},
	Grammar0R0: {
		symbols.NT_Grammar, 0, 0, 
		symbols.Symbols{  
			symbols.T_0, 
			symbols.NT_Rule, 
			symbols.NT_RepRule0x,
		}, 
		Grammar0R0, 
	},
	Grammar0R1: {
		symbols.NT_Grammar, 0, 1, 
		symbols.Symbols{  
			symbols.T_0, 
			symbols.NT_Rule, 
			symbols.NT_RepRule0x,
		}, 
		Grammar0R1, 
	},
	Grammar0R2: {
		symbols.NT_Grammar, 0, 2, 
		symbols.Symbols{  
			symbols.T_0, 
			symbols.NT_Rule, 
			symbols.NT_RepRule0x,
		}, 
		Grammar0R2, 
	},
	Grammar0R3: {
		symbols.NT_Grammar, 0, 3, 
		symbols.Symbols{  
			symbols.T_0, 
			symbols.NT_Rule, 
			symbols.NT_RepRule0x,
		}, 
		Grammar0R3, 
	},
	RepExpr0x0R0: {
		symbols.NT_RepExpr0x, 0, 0, 
		symbols.Symbols{  
			symbols.NT_Expr, 
			symbols.NT_RepExpr0x,
		}, 
		RepExpr0x0R0, 
	},
	RepExpr0x0R1: {
		symbols.NT_RepExpr0x, 0, 1, 
		symbols.Symbols{  
			symbols.NT_Expr, 
			symbols.NT_RepExpr0x,
		}, 
		RepExpr0x0R1, 
	},
	RepExpr0x0R2: {
		symbols.NT_RepExpr0x, 0, 2, 
		symbols.Symbols{  
			symbols.NT_Expr, 
			symbols.NT_RepExpr0x,
		}, 
		RepExpr0x0R2, 
	},
	RepExpr0x1R0: {
		symbols.NT_RepExpr0x, 1, 0, 
		symbols.Symbols{ 
		}, 
		RepExpr0x1R0, 
	},
	RepRule0x0R0: {
		symbols.NT_RepRule0x, 0, 0, 
		symbols.Symbols{  
			symbols.NT_Rule, 
			symbols.NT_RepRule0x,
		}, 
		RepRule0x0R0, 
	},
	RepRule0x0R1: {
		symbols.NT_RepRule0x, 0, 1, 
		symbols.Symbols{  
			symbols.NT_Rule, 
			symbols.NT_RepRule0x,
		}, 
		RepRule0x0R1, 
	},
	RepRule0x0R2: {
		symbols.NT_RepRule0x, 0, 2, 
		symbols.Symbols{  
			symbols.NT_Rule, 
			symbols.NT_RepRule0x,
		}, 
		RepRule0x0R2, 
	},
	RepRule0x1R0: {
		symbols.NT_RepRule0x, 1, 0, 
		symbols.Symbols{ 
		}, 
		RepRule0x1R0, 
	},
	Rule0R0: {
		symbols.NT_Rule, 0, 0, 
		symbols.Symbols{  
			symbols.T_2, 
			symbols.T_1, 
			symbols.T_0, 
			symbols.NT_Expr, 
			symbols.NT_RepExpr0x,
		}, 
		Rule0R0, 
	},
	Rule0R1: {
		symbols.NT_Rule, 0, 1, 
		symbols.Symbols{  
			symbols.T_2, 
			symbols.T_1, 
			symbols.T_0, 
			symbols.NT_Expr, 
			symbols.NT_RepExpr0x,
		}, 
		Rule0R1, 
	},
	Rule0R2: {
		symbols.NT_Rule, 0, 2, 
		symbols.Symbols{  
			symbols.T_2, 
			symbols.T_1, 
			symbols.T_0, 
			symbols.NT_Expr, 
			symbols.NT_RepExpr0x,
		}, 
		Rule0R2, 
	},
	Rule0R3: {
		symbols.NT_Rule, 0, 3, 
		symbols.Symbols{  
			symbols.T_2, 
			symbols.T_1, 
			symbols.T_0, 
			symbols.NT_Expr, 
			symbols.NT_RepExpr0x,
		}, 
		Rule0R3, 
	},
	Rule0R4: {
		symbols.NT_Rule, 0, 4, 
		symbols.Symbols{  
			symbols.T_2, 
			symbols.T_1, 
			symbols.T_0, 
			symbols.NT_Expr, 
			symbols.NT_RepExpr0x,
		}, 
		Rule0R4, 
	},
	Rule0R5: {
		symbols.NT_Rule, 0, 5, 
		symbols.Symbols{  
			symbols.T_2, 
			symbols.T_1, 
			symbols.T_0, 
			symbols.NT_Expr, 
			symbols.NT_RepExpr0x,
		}, 
		Rule0R5, 
	},
}

var slotIndex = map[Index]Label { 
	Index{ symbols.NT_Expr,0,0 }: Expr0R0,
	Index{ symbols.NT_Expr,0,1 }: Expr0R1,
	Index{ symbols.NT_Expr,0,2 }: Expr0R2,
	Index{ symbols.NT_Grammar,0,0 }: Grammar0R0,
	Index{ symbols.NT_Grammar,0,1 }: Grammar0R1,
	Index{ symbols.NT_Grammar,0,2 }: Grammar0R2,
	Index{ symbols.NT_Grammar,0,3 }: Grammar0R3,
	Index{ symbols.NT_RepExpr0x,0,0 }: RepExpr0x0R0,
	Index{ symbols.NT_RepExpr0x,0,1 }: RepExpr0x0R1,
	Index{ symbols.NT_RepExpr0x,0,2 }: RepExpr0x0R2,
	Index{ symbols.NT_RepExpr0x,1,0 }: RepExpr0x1R0,
	Index{ symbols.NT_RepRule0x,0,0 }: RepRule0x0R0,
	Index{ symbols.NT_RepRule0x,0,1 }: RepRule0x0R1,
	Index{ symbols.NT_RepRule0x,0,2 }: RepRule0x0R2,
	Index{ symbols.NT_RepRule0x,1,0 }: RepRule0x1R0,
	Index{ symbols.NT_Rule,0,0 }: Rule0R0,
	Index{ symbols.NT_Rule,0,1 }: Rule0R1,
	Index{ symbols.NT_Rule,0,2 }: Rule0R2,
	Index{ symbols.NT_Rule,0,3 }: Rule0R3,
	Index{ symbols.NT_Rule,0,4 }: Rule0R4,
	Index{ symbols.NT_Rule,0,5 }: Rule0R5,
}

var alternates = map[symbols.NT][]Label{ 
	symbols.NT_Grammar:[]Label{ Grammar0R0 },
	symbols.NT_RepRule0x:[]Label{ RepRule0x0R0,RepRule0x1R0 },
	symbols.NT_Rule:[]Label{ Rule0R0 },
	symbols.NT_RepExpr0x:[]Label{ RepExpr0x0R0,RepExpr0x1R0 },
	symbols.NT_Expr:[]Label{ Expr0R0 },
}

var nullable = []bool { 
	false, // Expr0R0 
	false, // Expr0R1 
	true, // Expr0R2 
	false, // Grammar0R0 
	false, // Grammar0R1 
	true, // Grammar0R2 
	true, // Grammar0R3 
	false, // RepExpr0x0R0 
	true, // RepExpr0x0R1 
	true, // RepExpr0x0R2 
	true, // RepExpr0x1R0 
	false, // RepRule0x0R0 
	true, // RepRule0x0R1 
	true, // RepRule0x0R2 
	true, // RepRule0x1R0 
	false, // Rule0R0 
	false, // Rule0R1 
	false, // Rule0R2 
	false, // Rule0R3 
	true, // Rule0R4 
	true, // Rule0R5 
}

var firstT = []map[token.Type]bool { 
	{  token.T_2: true,  }, // Expr0R0 
	{  token.T_3: true,  }, // Expr0R1 
	{  }, // Expr0R2 
	{  token.T_0: true,  }, // Grammar0R0 
	{  token.T_2: true,  }, // Grammar0R1 
	{  token.T_2: true,  }, // Grammar0R2 
	{  }, // Grammar0R3 
	{  token.T_2: true,  }, // RepExpr0x0R0 
	{  token.T_2: true,  }, // RepExpr0x0R1 
	{  }, // RepExpr0x0R2 
	{  }, // RepExpr0x1R0 
	{  token.T_2: true,  }, // RepRule0x0R0 
	{  token.T_2: true,  }, // RepRule0x0R1 
	{  }, // RepRule0x0R2 
	{  }, // RepRule0x1R0 
	{  token.T_2: true,  }, // Rule0R0 
	{  token.T_1: true,  }, // Rule0R1 
	{  token.T_0: true,  }, // Rule0R2 
	{  token.T_2: true,  }, // Rule0R3 
	{  token.T_2: true,  }, // Rule0R4 
	{  }, // Rule0R5 
}
