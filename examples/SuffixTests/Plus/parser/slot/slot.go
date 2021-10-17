
// Package slot is generated by gogll. Do not edit. 
package slot

import(
	"bytes"
	"fmt"
	
	"Plus/parser/symbols"
	"Plus/token"
)

type Label int

const(
	Base0R0 Label = iota
	Base0R1
	Base1F0
	Rep0R0
	Rep0R1
	Rep0R2
	Rep1F0
	Required0R0
	Required0R1
	Required1F0
	S10R0
	S10R1
	S10R2
	S11F0
	Suff1Base0R0
	Suff1Base0R1
	Suff1Base0R2
	Suff1Base1F0
	SuffBase0R0
	SuffBase0R1
	SuffBase0R2
	SuffBase1R0
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
	Base0R0: {
		symbols.NT_Base, 0, 0, 
		symbols.Symbols{  
			symbols.NT_Base,
		}, 
		Base0R0, 
	},
	Base0R1: {
		symbols.NT_Base, 0, 1, 
		symbols.Symbols{  
			symbols.NT_Base,
		}, 
		Base0R1, 
	},
	Base1F0: {
		symbols.NT_Base, 1, 0, 
		symbols.Symbols{ 
		}, 
		Base1F0, 
	},
	Rep0R0: {
		symbols.NT_Rep, 0, 0, 
		symbols.Symbols{  
			symbols.NT_Suff1Base, 
			symbols.NT_SuffBase,
		}, 
		Rep0R0, 
	},
	Rep0R1: {
		symbols.NT_Rep, 0, 1, 
		symbols.Symbols{  
			symbols.NT_Suff1Base, 
			symbols.NT_SuffBase,
		}, 
		Rep0R1, 
	},
	Rep0R2: {
		symbols.NT_Rep, 0, 2, 
		symbols.Symbols{  
			symbols.NT_Suff1Base, 
			symbols.NT_SuffBase,
		}, 
		Rep0R2, 
	},
	Rep1F0: {
		symbols.NT_Rep, 1, 0, 
		symbols.Symbols{ 
		}, 
		Rep1F0, 
	},
	Required0R0: {
		symbols.NT_Required, 0, 0, 
		symbols.Symbols{  
			symbols.NT_Required,
		}, 
		Required0R0, 
	},
	Required0R1: {
		symbols.NT_Required, 0, 1, 
		symbols.Symbols{  
			symbols.NT_Required,
		}, 
		Required0R1, 
	},
	Required1F0: {
		symbols.NT_Required, 1, 0, 
		symbols.Symbols{ 
		}, 
		Required1F0, 
	},
	S10R0: {
		symbols.NT_S1, 0, 0, 
		symbols.Symbols{  
			symbols.NT_Required, 
			symbols.NT_Rep,
		}, 
		S10R0, 
	},
	S10R1: {
		symbols.NT_S1, 0, 1, 
		symbols.Symbols{  
			symbols.NT_Required, 
			symbols.NT_Rep,
		}, 
		S10R1, 
	},
	S10R2: {
		symbols.NT_S1, 0, 2, 
		symbols.Symbols{  
			symbols.NT_Required, 
			symbols.NT_Rep,
		}, 
		S10R2, 
	},
	S11F0: {
		symbols.NT_S1, 1, 0, 
		symbols.Symbols{ 
		}, 
		S11F0, 
	},
	Suff1Base0R0: {
		symbols.NT_Suff1Base, 0, 0, 
		symbols.Symbols{  
			symbols.NT_Base, 
			symbols.NT_SuffBase,
		}, 
		Suff1Base0R0, 
	},
	Suff1Base0R1: {
		symbols.NT_Suff1Base, 0, 1, 
		symbols.Symbols{  
			symbols.NT_Base, 
			symbols.NT_SuffBase,
		}, 
		Suff1Base0R1, 
	},
	Suff1Base0R2: {
		symbols.NT_Suff1Base, 0, 2, 
		symbols.Symbols{  
			symbols.NT_Base, 
			symbols.NT_SuffBase,
		}, 
		Suff1Base0R2, 
	},
	Suff1Base1F0: {
		symbols.NT_Suff1Base, 1, 0, 
		symbols.Symbols{ 
		}, 
		Suff1Base1F0, 
	},
	SuffBase0R0: {
		symbols.NT_SuffBase, 0, 0, 
		symbols.Symbols{  
			symbols.NT_Base, 
			symbols.NT_SuffBase,
		}, 
		SuffBase0R0, 
	},
	SuffBase0R1: {
		symbols.NT_SuffBase, 0, 1, 
		symbols.Symbols{  
			symbols.NT_Base, 
			symbols.NT_SuffBase,
		}, 
		SuffBase0R1, 
	},
	SuffBase0R2: {
		symbols.NT_SuffBase, 0, 2, 
		symbols.Symbols{  
			symbols.NT_Base, 
			symbols.NT_SuffBase,
		}, 
		SuffBase0R2, 
	},
	SuffBase1R0: {
		symbols.NT_SuffBase, 1, 0, 
		symbols.Symbols{ 
		}, 
		SuffBase1R0, 
	},
}

var slotIndex = map[Index]Label { 
	Index{ symbols.NT_Base,0,0 }: Base0R0,
	Index{ symbols.NT_Base,0,1 }: Base0R1,
	Index{ symbols.NT_Base,1,0 }: Base1F0,
	Index{ symbols.NT_Rep,0,0 }: Rep0R0,
	Index{ symbols.NT_Rep,0,1 }: Rep0R1,
	Index{ symbols.NT_Rep,0,2 }: Rep0R2,
	Index{ symbols.NT_Rep,1,0 }: Rep1F0,
	Index{ symbols.NT_Required,0,0 }: Required0R0,
	Index{ symbols.NT_Required,0,1 }: Required0R1,
	Index{ symbols.NT_Required,1,0 }: Required1F0,
	Index{ symbols.NT_S1,0,0 }: S10R0,
	Index{ symbols.NT_S1,0,1 }: S10R1,
	Index{ symbols.NT_S1,0,2 }: S10R2,
	Index{ symbols.NT_S1,1,0 }: S11F0,
	Index{ symbols.NT_Suff1Base,0,0 }: Suff1Base0R0,
	Index{ symbols.NT_Suff1Base,0,1 }: Suff1Base0R1,
	Index{ symbols.NT_Suff1Base,0,2 }: Suff1Base0R2,
	Index{ symbols.NT_Suff1Base,1,0 }: Suff1Base1F0,
	Index{ symbols.NT_SuffBase,0,0 }: SuffBase0R0,
	Index{ symbols.NT_SuffBase,0,1 }: SuffBase0R1,
	Index{ symbols.NT_SuffBase,0,2 }: SuffBase0R2,
	Index{ symbols.NT_SuffBase,1,0 }: SuffBase1R0,
}

var alternates = map[symbols.NT][]Label{ 
	symbols.NT_S1:[]Label{ S10R0 },
	symbols.NT_Required:[]Label{ Required0R0 },
	symbols.NT_Rep:[]Label{ Rep0R0 },
	symbols.NT_Base:[]Label{ Base0R0 },
	symbols.NT_Suff1Base:[]Label{ Suff1Base0R0 },
	symbols.NT_SuffBase:[]Label{ SuffBase0R0,SuffBase1R0 },
}

var nullable = []bool { 
	false, // Base0R0 
	true, // Base0R1 
	false, // Base1F0 
	false, // Rep0R0 
	true, // Rep0R1 
	true, // Rep0R2 
	false, // Rep1F0 
	false, // Required0R0 
	true, // Required0R1 
	false, // Required1F0 
	false, // S10R0 
	false, // S10R1 
	true, // S10R2 
	false, // S11F0 
	false, // Suff1Base0R0 
	true, // Suff1Base0R1 
	true, // Suff1Base0R2 
	false, // Suff1Base1F0 
	false, // SuffBase0R0 
	true, // SuffBase0R1 
	true, // SuffBase0R2 
	true, // SuffBase1R0 
}

var firstT = []map[token.Type]bool { 
	{  token.T_0: true,  }, // Base0R0 
	{  }, // Base0R1 
	{  }, // Base1F0 
	{  token.T_0: true,  }, // Rep0R0 
	{  token.T_0: true,  }, // Rep0R1 
	{  }, // Rep0R2 
	{  }, // Rep1F0 
	{  token.T_1: true,  }, // Required0R0 
	{  }, // Required0R1 
	{  }, // Required1F0 
	{  token.T_1: true,  }, // S10R0 
	{  token.T_0: true,  }, // S10R1 
	{  }, // S10R2 
	{  }, // S11F0 
	{  token.T_0: true,  }, // Suff1Base0R0 
	{  token.T_0: true,  }, // Suff1Base0R1 
	{  }, // Suff1Base0R2 
	{  }, // Suff1Base1F0 
	{  token.T_0: true,  }, // SuffBase0R0 
	{  token.T_0: true,  }, // SuffBase0R1 
	{  }, // SuffBase0R2 
	{  }, // SuffBase1R0 
}
