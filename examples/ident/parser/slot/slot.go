
// Package slot is generated by gogll. Do not edit. 
package slot

import(
	"bytes"
	"fmt"
	
	"ident/parser/symbols"
	"ident/token"
)

type Label int

const(
	IdChar0R0 Label = iota
	IdChar0R1
	IdChar1M0
	IdChar1R0
	IdChar1R1
	IdChar2M0
	IdChar2R0
	IdChar2R1
	IdChar3M0
	IdChar3R0
	IdChar3R1
	IdChar4M0
	IdChar4R0
	IdChar4R1
	IdChar5F0
	Ident0R0
	Ident0R1
	Ident0R2
	Ident0R3
	Ident1F0
	Keyword0R0
	Keyword0R1
	Keyword0R2
	Keyword1M0
	Keyword1R0
	Keyword1R1
	Keyword1R2
	Keyword1R3
	Keyword2F0
	RepidChar0x0R0
	RepidChar0x0R1
	RepidChar0x0R2
	RepidChar0x1R0
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

func (l Label) IsLookahead() bool {
	s := l.Slot()
	return s.Pos > 0 && s.Symbols[s.Pos-1].IsLookahead()
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
			fmt.Fprintf(buf, "???")
		}
		fmt.Fprintf(buf, "%s ", sym)
	}
	if s.Pos >= len(s.Symbols) {
		fmt.Fprintf(buf, "???")
	}
	return buf.String()
}

var slots = map[Label]*Slot{ 
	IdChar0R0: {
		symbols.NT_IdChar, 0, 0, 
		symbols.Symbols{  
			symbols.T_1,
		}, 
		IdChar0R0, 
	},
	IdChar0R1: {
		symbols.NT_IdChar, 0, 1, 
		symbols.Symbols{  
			symbols.T_1,
		}, 
		IdChar0R1, 
	},
	IdChar1M0: {
		symbols.NT_IdChar, 1, 0, 
		symbols.Symbols{  
			symbols.T_0,
		}, 
		IdChar1M0, 
	},
	IdChar1R0: {
		symbols.NT_IdChar, 1, 0, 
		symbols.Symbols{  
			symbols.T_0,
		}, 
		IdChar1R0, 
	},
	IdChar1R1: {
		symbols.NT_IdChar, 1, 1, 
		symbols.Symbols{  
			symbols.T_0,
		}, 
		IdChar1R1, 
	},
	IdChar2M0: {
		symbols.NT_IdChar, 2, 0, 
		symbols.Symbols{  
			symbols.T_3,
		}, 
		IdChar2M0, 
	},
	IdChar2R0: {
		symbols.NT_IdChar, 2, 0, 
		symbols.Symbols{  
			symbols.T_3,
		}, 
		IdChar2R0, 
	},
	IdChar2R1: {
		symbols.NT_IdChar, 2, 1, 
		symbols.Symbols{  
			symbols.T_3,
		}, 
		IdChar2R1, 
	},
	IdChar3M0: {
		symbols.NT_IdChar, 3, 0, 
		symbols.Symbols{  
			symbols.T_4,
		}, 
		IdChar3M0, 
	},
	IdChar3R0: {
		symbols.NT_IdChar, 3, 0, 
		symbols.Symbols{  
			symbols.T_4,
		}, 
		IdChar3R0, 
	},
	IdChar3R1: {
		symbols.NT_IdChar, 3, 1, 
		symbols.Symbols{  
			symbols.T_4,
		}, 
		IdChar3R1, 
	},
	IdChar4M0: {
		symbols.NT_IdChar, 4, 0, 
		symbols.Symbols{  
			symbols.T_2,
		}, 
		IdChar4M0, 
	},
	IdChar4R0: {
		symbols.NT_IdChar, 4, 0, 
		symbols.Symbols{  
			symbols.T_2,
		}, 
		IdChar4R0, 
	},
	IdChar4R1: {
		symbols.NT_IdChar, 4, 1, 
		symbols.Symbols{  
			symbols.T_2,
		}, 
		IdChar4R1, 
	},
	IdChar5F0: {
		symbols.NT_IdChar, 5, 0, 
		symbols.Symbols{ 
		}, 
		IdChar5F0, 
	},
	Ident0R0: {
		symbols.NT_Ident, 0, 0, 
		symbols.Symbols{  
			symbols.LN_NT_Keyword, 
			symbols.NT_IdChar, 
			symbols.NT_RepidChar0x,
		}, 
		Ident0R0, 
	},
	Ident0R1: {
		symbols.NT_Ident, 0, 1, 
		symbols.Symbols{  
			symbols.LN_NT_Keyword, 
			symbols.NT_IdChar, 
			symbols.NT_RepidChar0x,
		}, 
		Ident0R1, 
	},
	Ident0R2: {
		symbols.NT_Ident, 0, 2, 
		symbols.Symbols{  
			symbols.LN_NT_Keyword, 
			symbols.NT_IdChar, 
			symbols.NT_RepidChar0x,
		}, 
		Ident0R2, 
	},
	Ident0R3: {
		symbols.NT_Ident, 0, 3, 
		symbols.Symbols{  
			symbols.LN_NT_Keyword, 
			symbols.NT_IdChar, 
			symbols.NT_RepidChar0x,
		}, 
		Ident0R3, 
	},
	Ident1F0: {
		symbols.NT_Ident, 1, 0, 
		symbols.Symbols{ 
		}, 
		Ident1F0, 
	},
	Keyword0R0: {
		symbols.NT_Keyword, 0, 0, 
		symbols.Symbols{  
			symbols.T_1, 
			symbols.T_0,
		}, 
		Keyword0R0, 
	},
	Keyword0R1: {
		symbols.NT_Keyword, 0, 1, 
		symbols.Symbols{  
			symbols.T_1, 
			symbols.T_0,
		}, 
		Keyword0R1, 
	},
	Keyword0R2: {
		symbols.NT_Keyword, 0, 2, 
		symbols.Symbols{  
			symbols.T_1, 
			symbols.T_0,
		}, 
		Keyword0R2, 
	},
	Keyword1M0: {
		symbols.NT_Keyword, 1, 0, 
		symbols.Symbols{  
			symbols.T_0, 
			symbols.T_3, 
			symbols.T_4,
		}, 
		Keyword1M0, 
	},
	Keyword1R0: {
		symbols.NT_Keyword, 1, 0, 
		symbols.Symbols{  
			symbols.T_0, 
			symbols.T_3, 
			symbols.T_4,
		}, 
		Keyword1R0, 
	},
	Keyword1R1: {
		symbols.NT_Keyword, 1, 1, 
		symbols.Symbols{  
			symbols.T_0, 
			symbols.T_3, 
			symbols.T_4,
		}, 
		Keyword1R1, 
	},
	Keyword1R2: {
		symbols.NT_Keyword, 1, 2, 
		symbols.Symbols{  
			symbols.T_0, 
			symbols.T_3, 
			symbols.T_4,
		}, 
		Keyword1R2, 
	},
	Keyword1R3: {
		symbols.NT_Keyword, 1, 3, 
		symbols.Symbols{  
			symbols.T_0, 
			symbols.T_3, 
			symbols.T_4,
		}, 
		Keyword1R3, 
	},
	Keyword2F0: {
		symbols.NT_Keyword, 2, 0, 
		symbols.Symbols{ 
		}, 
		Keyword2F0, 
	},
	RepidChar0x0R0: {
		symbols.NT_RepidChar0x, 0, 0, 
		symbols.Symbols{  
			symbols.NT_IdChar, 
			symbols.NT_RepidChar0x,
		}, 
		RepidChar0x0R0, 
	},
	RepidChar0x0R1: {
		symbols.NT_RepidChar0x, 0, 1, 
		symbols.Symbols{  
			symbols.NT_IdChar, 
			symbols.NT_RepidChar0x,
		}, 
		RepidChar0x0R1, 
	},
	RepidChar0x0R2: {
		symbols.NT_RepidChar0x, 0, 2, 
		symbols.Symbols{  
			symbols.NT_IdChar, 
			symbols.NT_RepidChar0x,
		}, 
		RepidChar0x0R2, 
	},
	RepidChar0x1R0: {
		symbols.NT_RepidChar0x, 1, 0, 
		symbols.Symbols{ 
		}, 
		RepidChar0x1R0, 
	},
}

var slotIndex = map[Index]Label { 
	Index{ symbols.NT_IdChar,0,0 }: IdChar0R0,
	Index{ symbols.NT_IdChar,0,1 }: IdChar0R1,
	Index{ symbols.NT_IdChar,1,0 }: IdChar1M0,
	Index{ symbols.NT_IdChar,1,0 }: IdChar1R0,
	Index{ symbols.NT_IdChar,1,1 }: IdChar1R1,
	Index{ symbols.NT_IdChar,2,0 }: IdChar2M0,
	Index{ symbols.NT_IdChar,2,0 }: IdChar2R0,
	Index{ symbols.NT_IdChar,2,1 }: IdChar2R1,
	Index{ symbols.NT_IdChar,3,0 }: IdChar3M0,
	Index{ symbols.NT_IdChar,3,0 }: IdChar3R0,
	Index{ symbols.NT_IdChar,3,1 }: IdChar3R1,
	Index{ symbols.NT_IdChar,4,0 }: IdChar4M0,
	Index{ symbols.NT_IdChar,4,0 }: IdChar4R0,
	Index{ symbols.NT_IdChar,4,1 }: IdChar4R1,
	Index{ symbols.NT_IdChar,5,0 }: IdChar5F0,
	Index{ symbols.NT_Ident,0,0 }: Ident0R0,
	Index{ symbols.NT_Ident,0,1 }: Ident0R1,
	Index{ symbols.NT_Ident,0,2 }: Ident0R2,
	Index{ symbols.NT_Ident,0,3 }: Ident0R3,
	Index{ symbols.NT_Ident,1,0 }: Ident1F0,
	Index{ symbols.NT_Keyword,0,0 }: Keyword0R0,
	Index{ symbols.NT_Keyword,0,1 }: Keyword0R1,
	Index{ symbols.NT_Keyword,0,2 }: Keyword0R2,
	Index{ symbols.NT_Keyword,1,0 }: Keyword1M0,
	Index{ symbols.NT_Keyword,1,0 }: Keyword1R0,
	Index{ symbols.NT_Keyword,1,1 }: Keyword1R1,
	Index{ symbols.NT_Keyword,1,2 }: Keyword1R2,
	Index{ symbols.NT_Keyword,1,3 }: Keyword1R3,
	Index{ symbols.NT_Keyword,2,0 }: Keyword2F0,
	Index{ symbols.NT_RepidChar0x,0,0 }: RepidChar0x0R0,
	Index{ symbols.NT_RepidChar0x,0,1 }: RepidChar0x0R1,
	Index{ symbols.NT_RepidChar0x,0,2 }: RepidChar0x0R2,
	Index{ symbols.NT_RepidChar0x,1,0 }: RepidChar0x1R0,
}

var alternates = map[symbols.NT][]Label{ 
	symbols.NT_Ident:[]Label{ Ident0R0 },
	symbols.NT_Keyword:[]Label{ Keyword0R0,Keyword1R0 },
	symbols.NT_IdChar:[]Label{ IdChar0R0,IdChar1R0,IdChar2R0,IdChar3R0,IdChar4R0 },
	symbols.NT_RepidChar0x:[]Label{ RepidChar0x0R0,RepidChar0x1R0 },
}

var nullable = []bool { 
	false, // IdChar0R0 
	true, // IdChar0R1 
	false, // IdChar1M0 
	false, // IdChar1R0 
	true, // IdChar1R1 
	false, // IdChar2M0 
	false, // IdChar2R0 
	true, // IdChar2R1 
	false, // IdChar3M0 
	false, // IdChar3R0 
	true, // IdChar3R1 
	false, // IdChar4M0 
	false, // IdChar4R0 
	true, // IdChar4R1 
	false, // IdChar5F0 
	false, // Ident0R0 
	false, // Ident0R1 
	true, // Ident0R2 
	true, // Ident0R3 
	false, // Ident1F0 
	false, // Keyword0R0 
	false, // Keyword0R1 
	true, // Keyword0R2 
	false, // Keyword1M0 
	false, // Keyword1R0 
	false, // Keyword1R1 
	false, // Keyword1R2 
	true, // Keyword1R3 
	false, // Keyword2F0 
	false, // RepidChar0x0R0 
	true, // RepidChar0x0R1 
	true, // RepidChar0x0R2 
	true, // RepidChar0x1R0 
}

var firstT = []map[token.Type]bool { 
	{  token.T_1: true,  }, // IdChar0R0 
	{  }, // IdChar0R1 
	{  token.T_0: true,  }, // IdChar1M0 
	{  token.T_0: true,  }, // IdChar1R0 
	{  }, // IdChar1R1 
	{  token.T_3: true,  }, // IdChar2M0 
	{  token.T_3: true,  }, // IdChar2R0 
	{  }, // IdChar2R1 
	{  token.T_4: true,  }, // IdChar3M0 
	{  token.T_4: true,  }, // IdChar3R0 
	{  }, // IdChar3R1 
	{  token.T_2: true,  }, // IdChar4M0 
	{  token.T_2: true,  }, // IdChar4R0 
	{  }, // IdChar4R1 
	{  }, // IdChar5F0 
	{  token.T_1: true,  token.T_0: true,  token.T_3: true,  token.T_4: true,  token.T_2: true,  }, // Ident0R0 
	{  token.T_0: true,  token.T_3: true,  token.T_4: true,  token.T_2: true,  token.T_1: true,  }, // Ident0R1 
	{  token.T_2: true,  token.T_1: true,  token.T_0: true,  token.T_3: true,  token.T_4: true,  }, // Ident0R2 
	{  }, // Ident0R3 
	{  }, // Ident1F0 
	{  token.T_1: true,  }, // Keyword0R0 
	{  token.T_0: true,  }, // Keyword0R1 
	{  }, // Keyword0R2 
	{  token.T_0: true,  }, // Keyword1M0 
	{  token.T_0: true,  }, // Keyword1R0 
	{  token.T_3: true,  }, // Keyword1R1 
	{  token.T_4: true,  }, // Keyword1R2 
	{  }, // Keyword1R3 
	{  }, // Keyword2F0 
	{  token.T_1: true,  token.T_0: true,  token.T_3: true,  token.T_4: true,  token.T_2: true,  }, // RepidChar0x0R0 
	{  token.T_4: true,  token.T_2: true,  token.T_1: true,  token.T_0: true,  token.T_3: true,  }, // RepidChar0x0R1 
	{  }, // RepidChar0x0R2 
	{  }, // RepidChar0x1R0 
}
