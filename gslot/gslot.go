//  Copyright 2021 Aaron Moss
//  Copyright 2019 Marius Ackerman
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

// Package gslot implements grammar slots
package gslot

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/bruceiv/pegll/ast"
	"github.com/bruceiv/pegll/frstflw"
	"github.com/bruceiv/pegll/symbols"
)

const (
	// Slot mode for unknown parse status
	Unknown = 'R'
	// Slot mode for parse failure
	Fail = 'F'
	// Slot mode for matching parse
	Match = 'M'
)

type Label struct {
	Head      string
	Alternate int
	Pos       int
	Mode      rune
	gs        *GSlot
	ff        *frstflw.FF
}

type Slots []Label

type GSlot struct {
	g     *ast.GoGLL
	ff    *frstflw.FF
	slots map[Label]symbols.Symbols
}

func New(g *ast.GoGLL, ff *frstflw.FF) *GSlot {
	gs := &GSlot{
		g:     g,
		ff:    ff,
		slots: make(map[Label]symbols.Symbols),
	}
	gs.genSlots()
	return gs
}

func NewLabel(head string, alt, pos int, mode rune, gs *GSlot, ff *frstflw.FF) *Label {
	return &Label{
		Head:      head,
		Alternate: alt,
		Pos:       pos,
		Mode:      mode,
		gs:        gs,
		ff:        ff,
	}
}

func (gs *GSlot) Slots() Slots {
	res := make(Slots, 0, len(gs.slots))
	for l, _ := range gs.slots {
		res = append(res, l)
	}
	sort.Sort(res)
	return res
}

func (s Label) Label() string {
	return LabelFor(s.Head, s.Alternate, s.Pos, s.Mode)
}

func LabelFor(head string, alt, pos int, mode rune) string {
	return fmt.Sprintf("%s%d%c%d", head, alt, mode, pos)
}

func (s Label) IsEoR() bool {
	symbols := s.gs.slots[s]
	return s.Pos >= len(symbols)
}

func (s Label) IsFiR() bool {
	symbols := s.gs.slots[s]
	if s.Pos > 1 || len(symbols) <= 1 {
		return false
	}
	if s.ff.FirstOfSymbol(symbols[0].Literal()).Contain(frstflw.Empty) &&
		symbols[0] == symbols[1] {
		return false
	}
	return true
}

func (s Label) String() string {
	symbols := s.gs.slots[s]
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%s : ", s.Head)
	for i, sym := range symbols {
		if i == s.Pos {
			fmt.Fprintf(buf, "∙")
		}
		fmt.Fprintf(buf, "%s ", sym)
	}
	// fmt.Printf("slotLabel.String(): %s pos=%d len(symbols)=%d\n", s.Head, s.Pos, len(symbols))
	if s.Pos >= len(symbols) {
		fmt.Fprintf(buf, "∙")
	}
	// fmt.Println("  ", buf.String())
	return buf.String()
}

func (s Label) Symbols() symbols.Symbols {
	return s.gs.slots[s]
}

func (ss Slots) Labels() (labels []string) {
	for _, s := range ss {
		labels = append(labels, s.Label())
	}
	return
}

func (ss Slots) Len() int {
	return len(ss)
}

func (ss Slots) Less(i, j int) bool {
	if ss[i].Head < ss[j].Head {
		return true
	}

	if ss[i].Head == ss[j].Head {
		if ss[i].Alternate < ss[j].Alternate {
			return true
		}
		if ss[i].Alternate == ss[j].Alternate {
			if ss[i].Pos < ss[j].Pos {
				return true
			}

			if ss[i].Pos == ss[j].Pos {
				return ss[i].Mode < ss[j].Mode
			}
		}
	}
	return false
}

func (ss Slots) Swap(i, j int) {
	iTmp := ss[i]
	ss[i] = ss[j]
	ss[j] = iTmp
}

func (gs *GSlot) genSlots() {
	for _, rule := range gs.g.SyntaxRules {
		gs.genSlotsOfRule(rule)
	}
}

func (gs *GSlot) genSlotsOfRule(r *ast.SyntaxRule) {
	nt := r.Head.ID()
	// add normal rule slots
	for i, a := range r.Alternates {
		syms := getSymbols(a.GetSymbols())
		gs.genSlotsOfAlternate(nt, i, syms...)
	}
	if !r.AlwaysMatches() {
		// add fail slot to failable rules
		gs.genSlot(nt, len(r.Alternates), 0, Fail, []symbols.Symbol{}...)

		// add pass labels to failable unordered rules
		if !r.IsOrdered {
			for i, a := range r.Alternates {
				// no pass label for first alternate
				if i == 0 {
					continue
				}
				syms := getSymbols(a.GetSymbols())
				for j, sym := range syms {
					// first symbol always gets pass label
					if j == 0 {
						gs.genSlot(nt, i, j, Match, syms...)
					}
					// slots after nonterminal calls also get
					// pass label
					if sym.IsNonTerminal() || sym.IsLookahead() {
						gs.genSlot(nt, i, j+1, Match, syms...)
					}
				}
			}
		}
	}
}

func (gs *GSlot) genSlotsOfAlternate(nt string, altI int, symbls ...symbols.Symbol) {
	if len(symbls) == 0 {
		gs.genSlot(nt, altI, 0, Unknown, []symbols.Symbol{}...)
	} else {
		for pos := 0; pos <= len(symbls); pos++ {
			gs.genSlot(nt, altI, pos, Unknown, symbls...)
		}
	}
}

func (gs *GSlot) genSlot(nt string, altI, pos int, mode rune, symbols ...symbols.Symbol) {
	slot := Label{
		Head:      nt,
		Alternate: altI,
		Pos:       pos,
		Mode:      mode,
		gs:        gs,
		ff:        gs.ff,
	}
	gs.slots[slot] = symbols
}

func (gs *GSlot) Len() int {
	return len(gs.slots)
}

// getSymbols translates AST symbol strings to symbols.Symbol
func getSymbols(astSymbols []string) []symbols.Symbol {
	symbls := make([]symbols.Symbol, len(astSymbols))
	for i, s := range astSymbols {
		symbls[i] = symbols.FromASTString(s)
	}
	return symbls
}
