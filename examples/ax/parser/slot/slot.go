
// Package slot is generated by gogll. Do not edit. 
package slot

import(
	"bytes"
	"fmt"
	
	"ax/parser/symbols"
	"ax/token"
)

type Label int

const(
	AStar0R0 Label = iota
	AStar0R1
	AStar1F0
	Suffa0R0
	Suffa0R1
	Suffa0R2
	Suffa1R0
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
	AStar0R0: {
		symbols.NT_AStar, 0, 0, 
		symbols.Symbols{  
			symbols.NT_Suffa,
		}, 
		AStar0R0, 
	},
	AStar0R1: {
		symbols.NT_AStar, 0, 1, 
		symbols.Symbols{  
			symbols.NT_Suffa,
		}, 
		AStar0R1, 
	},
	AStar1F0: {
		symbols.NT_AStar, 1, 0, 
		symbols.Symbols{ 
		}, 
		AStar1F0, 
	},
	Suffa0R0: {
		symbols.NT_Suffa, 0, 0, 
		symbols.Symbols{  
			symbols.T_0, 
			symbols.NT_Suffa,
		}, 
		Suffa0R0, 
	},
	Suffa0R1: {
		symbols.NT_Suffa, 0, 1, 
		symbols.Symbols{  
			symbols.T_0, 
			symbols.NT_Suffa,
		}, 
		Suffa0R1, 
	},
	Suffa0R2: {
		symbols.NT_Suffa, 0, 2, 
		symbols.Symbols{  
			symbols.T_0, 
			symbols.NT_Suffa,
		}, 
		Suffa0R2, 
	},
	Suffa1R0: {
		symbols.NT_Suffa, 1, 0, 
		symbols.Symbols{ 
		}, 
		Suffa1R0, 
	},
}

var slotIndex = map[Index]Label { 
	Index{ symbols.NT_AStar,0,0 }: AStar0R0,
	Index{ symbols.NT_AStar,0,1 }: AStar0R1,
	Index{ symbols.NT_AStar,1,0 }: AStar1F0,
	Index{ symbols.NT_Suffa,0,0 }: Suffa0R0,
	Index{ symbols.NT_Suffa,0,1 }: Suffa0R1,
	Index{ symbols.NT_Suffa,0,2 }: Suffa0R2,
	Index{ symbols.NT_Suffa,1,0 }: Suffa1R0,
}

var alternates = map[symbols.NT][]Label{ 
	symbols.NT_AStar:[]Label{ AStar0R0 },
	symbols.NT_Suffa:[]Label{ Suffa0R0,Suffa1R0 },
}

var nullable = []bool { 
	true, // AStar0R0 
	true, // AStar0R1 
	false, // AStar1F0 
	false, // Suffa0R0 
	true, // Suffa0R1 
	true, // Suffa0R2 
	true, // Suffa1R0 
}

var firstT = []map[token.Type]bool { 
	{  token.T_0: true,  }, // AStar0R0 
	{  }, // AStar0R1 
	{  }, // AStar1F0 
	{  token.T_0: true,  }, // Suffa0R0 
	{  token.T_0: true,  }, // Suffa0R1 
	{  }, // Suffa0R2 
	{  }, // Suffa1R0 
}
