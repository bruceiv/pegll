// Package bsr is generated by gogll. Do not edit.

/*
Package bsr implements a Binary Subtree Representation set as defined in

	Scott et al
	Derivation representation using binary subtree sets,
	Science of Computer Programming 175 (2019)
*/
package bsr

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"

	"Optional/lexer"
	"Optional/parser/slot"
	"Optional/parser/symbols"
	"Optional/token"
)

type bsr interface {
	LeftExtent() int
	RightExtent() int
	Pivot() int
}

/*
Set contains the set of Binary Subtree Representations (BSR).
*/
type Set struct {
	slotEntries   map[BSR]bool
	ntSlotEntries map[ntSlot][]BSR
	stringEntries map[stringBSR]bool
	rightExtent   int
	lex           *lexer.Lexer

	startSym symbols.NT
}

type ntSlot struct {
	nt          symbols.NT
	leftExtent  int
	rightExtent int
}

// BSR is the binary subtree representation of a parsed nonterminal
type BSR struct {
	Label       slot.Label
	leftExtent  int
	pivot       int
	rightExtent int
	set         *Set
}

type stringBSR struct {
	Label       slot.Label
	leftExtent  int
	pivot       int
	rightExtent int
	set         *Set
}

// New returns a new initialised BSR Set
func New(startSymbol symbols.NT, l *lexer.Lexer) *Set {
	return &Set{
		slotEntries:   make(map[BSR]bool),
		ntSlotEntries: make(map[ntSlot][]BSR),
		stringEntries: make(map[stringBSR]bool),
		rightExtent:   0,
		lex:           l,
		startSym:      startSymbol,
	}
}

/*
Add a bsr to the set. (i,j) is the extent. k is the pivot.
*/
func (s *Set) Add(l slot.Label, i, k, j int) {
	// fmt.Printf("bsr.Add(%s,%d,%d,%d)\n", l,i,k,j)
	if l.EoR() {
		s.insert(BSR{l, i, k, j, s})
	} else {
		if l.Pos() > 1 {
			s.insert(stringBSR{l, i, k, j, s})
		}
	}
}

// AddEmpty adds a grammar slot: X : ϵ•
func (s *Set) AddEmpty(l slot.Label, i int) {
	s.insert(BSR{l, i, i, i, s})
}

/*
Contain returns true iff the BSR Set contains the NT symbol with left and
right extent.
*/
func (s *Set) Contain(nt symbols.NT, left, right int) bool {
	// fmt.Printf("bsr.Contain(%s,%d,%d)\n",nt,left,right)
	for e := range s.slotEntries {
		// fmt.Printf("  (%s,%d,%d)\n",e.Label.Head(),e.leftExtent,e.rightExtent)
		if e.Label.Head() == nt && e.leftExtent == left && e.rightExtent == right {
			// fmt.Println("  true")
			return true
		}
	}
	// fmt.Println("  false")
	return false
}

// Dump prints the NT symbols of the parse forest.
func (s *Set) Dump() {
	for _, root := range s.GetRoots() {
		s.dump(root, 0)
	}
}

// Dumps all BSR nodes in a set, not just those reachable from roots.
func (s *Set) FlatDump() {
	for b := range s.slotEntries {
		fmt.Println(b)
	}
}

func (s *Set) dump(b BSR, level int) {
	fmt.Print(indent(level, " "))
	fmt.Println(b)
	for _, cn := range b.GetAllNTChildren() {
		for _, c := range cn {
			s.dump(c, level+1)
		}
	}
}

func indent(n int, c string) string {
	buf := new(bytes.Buffer)
	for i := 0; i < 4*n; i++ {
		fmt.Fprint(buf, c)
	}
	return buf.String()
}

// GetAll returns all BSR grammar slot entries
func (s *Set) GetAll() (bsrs []BSR) {
	for b := range s.slotEntries {
		bsrs = append(bsrs, b)
	}
	return
}

// GetRightExtent returns the right extent of the BSR set
func (s *Set) GetRightExtent() int {
	return s.rightExtent
}

// GetRoot returns the root of the parse tree of an unambiguous parse.
// GetRoot fails if the parse was ambiguous. Use GetRoots() for ambiguous parses.
func (s *Set) GetRoot() BSR {
	rts := s.GetRoots()
	if len(rts) != 1 {
		failf("%d parse trees exist for start symbol %s", len(rts), s.startSym)
	}
	return rts[0]
}

// GetRoots returns all the roots of parse trees of the start symbol of the grammar.
func (s *Set) GetRoots() (roots []BSR) {
	for b := range s.slotEntries {
		if b.Label.Head() == s.startSym && b.leftExtent == 0 && s.rightExtent == b.rightExtent {
			roots = append(roots, b)
		}
	}
	return
}

func (s *Set) getString(l slot.Label, leftExtent, rightExtent int) stringBSR {
	for str := range s.stringEntries {
		if str.Label == l && str.leftExtent == leftExtent && str.rightExtent == rightExtent {
			return str
		}
	}
	fmt.Printf("Error: no string %s left extent=%d right extent=%d pos=%d\n",
		strings.Join(l.Symbols()[:l.Pos()].Strings(), " "), leftExtent, rightExtent, l.Pos())
	panic("must not happen")
}

func (s *Set) insert(bsr bsr) {
	if bsr.RightExtent() > s.rightExtent {
		s.rightExtent = bsr.RightExtent()
	}
	switch b := bsr.(type) {
	case BSR:
		s.slotEntries[b] = true
		nt := ntSlot{b.Label.Head(), b.leftExtent, b.rightExtent}
		s.ntSlotEntries[nt] = append(s.ntSlotEntries[nt], b)
	case stringBSR:
		s.stringEntries[b] = true
	default:
		panic(fmt.Sprintf("Invalid type %T", bsr))
	}
}

// Alternate returns the index of the grammar rule alternate.
func (b BSR) Alternate() int {
	return b.Label.Alternate()
}

// GetAllNTChildren returns all the NT Children of b. If an NT child of b has
// ambiguous parses then all parses of that child are returned.
func (b BSR) GetAllNTChildren() [][]BSR {
	children := [][]BSR{}
	for i, s := range b.Label.Symbols() {
		if s.IsNonTerminal() {
			sChildren := b.GetNTChildrenI(i)
			children = append(children, sChildren)
		}
	}
	return children
}

// GetNTChild returns the BSR of occurrence i of nt in s.
// GetNTChild fails if s has ambiguous subtrees of occurrence i of nt.
func (b BSR) GetNTChild(nt symbols.NT, i int) BSR {
	bsrs := b.GetNTChildren(nt, i)
	if len(bsrs) != 1 {
		ambiguousSlots := []string{}
		for _, c := range bsrs {
			ambiguousSlots = append(ambiguousSlots, c.String())
		}
		b.set.fail(b, "%s is ambiguous in %s\n  %s", nt, b, strings.Join(ambiguousSlots, "\n  "))
	}
	return bsrs[0]
}

// GetNTChildI returns the BSR of NT symbol[i] in the BSR set.
// GetNTChildI fails if the BSR set has ambiguous subtrees of NT i.
func (b BSR) GetNTChildI(i int) BSR {
	bsrs := b.GetNTChildrenI(i)
	if len(bsrs) != 1 {
		b.set.fail(b, "NT %d is ambiguous in %s", i, b)
	}
	return bsrs[0]
}

// GetNTChildren returns all the BSRs of occurrence i of nt in s
func (b BSR) GetNTChildren(nt symbols.NT, i int) []BSR {
	// fmt.Printf("GetNTChild(%s,%d) %s\n", nt, i, b)
	positions := []int{}
	for j, s := range b.Label.Symbols() {
		if s == nt {
			positions = append(positions, j)
		}
	}
	if len(positions) == 0 {
		b.set.fail(b, "Error: %s has no NT %s", b, nt)
	}
	return b.GetNTChildrenI(positions[i])
}

// GetNTChildrenI returns all the BSRs of NT symbol[i] in s
func (b BSR) GetNTChildrenI(i int) []BSR {
	// fmt.Printf("bsr.GetNTChildI(%d) %s\n", i, b)
	if i >= len(b.Label.Symbols()) {
		b.set.fail(b, "Error: cannot get NT child %d of %s", i, b)
	}
	if len(b.Label.Symbols()) == 1 {
		return b.set.getNTSlot(b.Label.Symbols()[i], b.pivot, b.rightExtent)
	}
	if len(b.Label.Symbols()) == 2 {
		if i == 0 {
			return b.set.getNTSlot(b.Label.Symbols()[i], b.leftExtent, b.pivot)
		}
		return b.set.getNTSlot(b.Label.Symbols()[i], b.pivot, b.rightExtent)
	}
	idx := b.Label.Index()
	str := stringBSR{b.Label, b.leftExtent, b.pivot, b.rightExtent, b.set}
	for idx.Pos > i+1 && idx.Pos > 2 {
		idx.Pos--
		str = b.set.getString(slot.GetLabel(idx.NT, idx.Alt, idx.Pos), str.leftExtent, str.pivot)
		// fmt.Printf("  %s\n", str)
	}
	if i == 0 {
		return b.set.getNTSlot(b.Label.Symbols()[i], str.leftExtent, str.pivot)
	}
	return b.set.getNTSlot(b.Label.Symbols()[i], str.pivot, str.rightExtent)
}

// func (b BSR) GetString() string {
// 	return set.lex.GetString(b.LeftExtent(),b.RightExtent())
// }

// GetTChildI returns the terminal symbol at position i in b.
// GetTChildI panics if symbol i is not a valid terminal
func (b BSR) GetTChildI(i int) *token.Token {
	symbols := b.Label.Symbols()

	if i >= len(symbols) {
		panic(fmt.Sprintf("%s has no T child %d", b, i))
	}
	if symbols[i].IsNonTerminal() {
		panic(fmt.Sprintf("symbol %d in %s is an NT", i, b))
	}

	lext := b.leftExtent
	for j := 0; j < i; j++ {
		if symbols[j].IsNonTerminal() {
			nt := b.GetNTChildI(j)
			lext += nt.rightExtent - nt.leftExtent
		} else {
			lext++
		}
	}
	return b.set.lex.Tokens[lext]
}

func deleteNTSlotEntry(b BSR) {
	// fmt.Printf("deletNTSlotEntry(%s)\n", b)
	nts := ntSlot{b.Label.Head(), b.leftExtent, b.rightExtent}
	slots := b.set.ntSlotEntries[nts]
	slots1 := make([]BSR, 0, len(slots))
	bi := -1
	for i, s := range slots {
		if s == b && bi != -1 {
			panic(fmt.Sprintf("Duplicate slot entries: %d and %d", bi, i))
		} else {
			slots1 = append(slots1, s)
		}
	}
	b.set.ntSlotEntries[nts] = slots1
}

// LeftExtent returns the left extent of the BSR
func (b BSR) LeftExtent() int {
	return b.leftExtent
}

// RightExtent returns the right extent of the BSR
func (b BSR) RightExtent() int {
	return b.rightExtent
}

// Pivot returns the pivot of the BSR
func (b BSR) Pivot() int {
	return b.pivot
}

func (b BSR) String() string {
	return fmt.Sprintf("%s,%d,%d,%d - %s", b.Label, b.leftExtent, b.pivot, b.rightExtent,
		b.set.lex.GetString(b.LeftExtent(), b.RightExtent()-1))
}

func (s stringBSR) LeftExtent() int {
	return s.leftExtent
}

func (s stringBSR) RightExtent() int {
	return s.rightExtent
}

func (s stringBSR) Pivot() int {
	return s.pivot
}

func (s stringBSR) Empty() bool {
	return s.leftExtent == s.pivot && s.pivot == s.rightExtent
}

// String returns a string representation of s
func (s stringBSR) String() string {
	// fmt.Printf("bsr.stringBSR.stringBSR(): %s, %d, %d, %d\n",
	// 	s.Label.Symbols(), s.leftExtent, s.pivot, s.rightExtent)
	ss := s.Label.Symbols()[:s.Label.Pos()].Strings()
	str := strings.Join(ss, " ")
	return fmt.Sprintf("%s,%d,%d,%d - %s", str, s.leftExtent, s.pivot,
		s.rightExtent, s.set.lex.GetString(s.LeftExtent(), s.RightExtent()))
}

func (s *Set) getNTSlot(sym symbols.Symbol, leftExtent, rightExtent int) (bsrs []BSR) {
	nt, ok := sym.(symbols.NT)
	if !ok {
		line, col := s.getLineColumn(leftExtent)
		failf("%s is not an NT at line %d col %d", sym, line, col)
	}
	return s.ntSlotEntries[ntSlot{nt, leftExtent, rightExtent}]
}

func (s *Set) fail(b BSR, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	line, col := s.getLineColumn(b.LeftExtent())
	panic(fmt.Sprintf("Error in BSR: %s at line %d col %d\n", msg, line, col))
}

func failf(format string, args ...interface{}) {
	panic(fmt.Sprintf("Error in BSR: %s\n", fmt.Sprintf(format, args...)))
}

func decodeRune(str []byte) (string, rune, int) {
	if len(str) == 0 {
		return "$", '$', 0
	}
	r, sz := utf8.DecodeRune(str)
	if r == utf8.RuneError {
		panic(fmt.Sprintf("Rune error: %s", str))
	}
	switch r {
	case '\t', ' ':
		return "space", r, sz
	case '\n':
		return "\\n", r, sz
	}
	return string(str[:sz]), r, sz
}

func (s *Set) getLineColumn(cI int) (line, col int) {
	return s.lex.GetLineColumnOfToken(cI)
}

// ReportAmbiguous lists the ambiguous subtrees of the parse forest
func (s *Set) ReportAmbiguous() {
	fmt.Println("Ambiguous BSR Subtrees:")
	rts := s.GetRoots()
	if len(rts) != 1 {
		fmt.Printf("BSR has %d ambigous roots\n", len(rts))
	}
	for i, b := range s.GetRoots() {
		fmt.Println("In root", i)
		if !s.report(b) {
			fmt.Println("No ambiguous BSRs")
		}
	}
}

// report return true iff at least one ambigous BSR was found
func (s *Set) report(b BSR) bool {
	ambiguous := false
	for i, sym := range b.Label.Symbols() {
		ln, col := s.getLineColumn(b.LeftExtent())
		if sym.IsNonTerminal() {
			if len(b.GetNTChildrenI(i)) != 1 {
				ambiguous = true
				fmt.Printf("  Ambigous: in %s: NT %s (%d) at line %d col %d \n",
					b, sym, i, ln, col)
				fmt.Println("   Children:")
				for _, c := range b.GetNTChildrenI(i) {
					fmt.Printf("     %s\n", c)
				}
			}
			for _, b1 := range b.GetNTChildrenI(i) {
				s.report(b1)
			}
		}
	}
	return ambiguous
}

// IsAmbiguous returns true if the BSR set does not have exactly one root, or
// if any BSR in the set has an NT symbol, which does not have exactly one
// sub-tree.
func (s *Set) IsAmbiguous() bool {
	if len(s.GetRoots()) != 1 {
		return true
	}
	return isAmbiguous(s.GetRoot())
}

// isAmbiguous returns true if b or any of its NT children is ambiguous.
// A BSR is ambigous if any of its NT symbols does not have exactly one
// subtrees (children).
func isAmbiguous(b BSR) bool {
	for i, s := range b.Label.Symbols() {
		if s.IsNonTerminal() {
			if len(b.GetNTChildrenI(i)) != 1 {
				return true
			}
			for _, b1 := range b.GetNTChildrenI(i) {
				if isAmbiguous(b1) {
					return true
				}
			}
		}
	}
	return false
}

// index key for count of slots with given extents
type slotKey struct {
	label slot.Label
	left  int
	right int
}

// index key for nonterminal-match lookup
type ntKey struct {
	nt   symbols.NT
	left int
}

// The information necessary to check if a BSR node's dependencies are still valid
type bsrDepend = struct {
	b     BSR        // the BSR node with the dependency
	nt    symbols.NT // the nonterminal to look for
	right int        // the right extent of that match (left given by map key)
}

// index key for successor lookup
type succKey struct {
	label slot.Label
	left  int
	pivot int
}

// BSR set augmented with additional indices necessary for PEG filtering.
// Fields map keys to lists of BSRs in s that match
// TODO maybe more efficient to store s as []BSR, build indices as []int
type pegBsr struct {
	nSlots       map[slotKey]int
	nts          map[ntKey][]BSR
	dependencies map[int][]bsrDepend
	successors   map[succKey][]BSR
	s            *Set
}

// System-dependent maximum int value
// Based off https://stackoverflow.com/questions/6878590/the-maximum-value-for-an-int-type-in-go
const noAlt = int(^uint(0) >> 1)

// Sets up indices used in PEG construction
func (s *Set) buildPeg() *pegBsr {
	p := &pegBsr{
		nSlots:       make(map[slotKey]int),
		nts:          make(map[ntKey][]BSR),
		dependencies: make(map[int][]bsrDepend),
		successors:   make(map[succKey][]BSR),
		s:            s,
	}

	for b := range s.slotEntries {
		x := b.Label.Slot()
		// count of slots with same left/right extent
		p.nSlots[slotKey{b.Label, b.leftExtent, b.rightExtent}] += 1
		if x.EoR() {
			// nonterminal match at end-of-rule
			nKey := ntKey{x.NT, b.leftExtent}
			p.nts[nKey] = append(p.nts[nKey], b)
		}
		if x.Pos != 0 {
			// nodes with a nonterminal after the given pivot
			if nt, ok := x.Symbols[x.Pos-1].(symbols.NT); ok {
				p.dependencies[b.pivot] = append(p.dependencies[b.pivot], bsrDepend{b, nt, b.rightExtent})
			} else if x.Pos > 1 {
				// only check nonterminals before the pivot if they exist and the symbol
				// afterward is a terminal (if both symbols are nonterminals, there will
				// be another BSR node with pivot before the first nonterminal)
				if nt, ok := x.Symbols[x.Pos-2].(symbols.NT); ok {
					p.dependencies[b.leftExtent] = append(p.dependencies[b.leftExtent], bsrDepend{b, nt, b.pivot})
				}
			}
		}
		// slot successors
		sKey := succKey{b.Label, b.leftExtent, b.pivot}
		p.successors[sKey] = append(p.successors[sKey], b)
	}

	return p
}

// Deletes a BSR from its set, returning false if it is the last copy of
// that slot with that extent.
// Does not remove from nts, pivots, or successors indices
func (p *pegBsr) deleteNode(b BSR) bool {
	delete(p.s.slotEntries, b)
	sKey := slotKey{b.Label, b.leftExtent, b.rightExtent}
	p.nSlots[sKey] -= 1
	return p.nSlots[sKey] <= 0
}

// Deletes a BSR from its set, as well as all its successors
// clears successor indices when finished deleting
func (p *pegBsr) deleteNodeAndSuccessors(b BSR) {
	// remove node itself, leaving successors alone if more copies of
	// this slot with the same left & right extents
	if p.deleteNode(b) {
		return
	}

	x := b.Label.Slot()
	// end early if no successor
	if x.EoR() {
		return
	}
	// otherwise look up successor nodes with continuation label,
	// same left extent, pivoting on right extent, and delete
	// TODO try succcessor label := b.Label + 1
	succK := succKey{slot.GetLabel(x.NT, x.Alt, x.Pos+1), b.leftExtent, b.rightExtent}
	for _, b := range p.successors[succK] {
		p.deleteNodeAndSuccessors(b)
	}
	delete(p.successors, succK)
}

// Filters slots out of a BSR set that have had their nonterminal at
// position i deleted, as well as any successors.
func (p *pegBsr) filterMissingSlots(i int) {
	// loop over all nodes with the given pivot
	for _, d := range p.dependencies[i] {
		// skip deleted nodes
		if !p.s.slotEntries[d.b] {
			continue
		}

		// check there is still a match for this nonterminal with the same
		// right extent, deleting slot if none found
		found := false
		for _, nb := range p.nts[ntKey{d.nt, i}] {
			if p.s.slotEntries[nb] && nb.rightExtent == d.right {
				found = true
				break
			}
		}
		if !found {
			p.deleteNodeAndSuccessors(d.b)
		}
	}
	// TODO could reset dependencies[i] here, but I don't think it's used again
}

// Removes all the BSR nodes that don't match the PEG ordered choice
// property
func (s *Set) FilterByOrderedChoice() {
	// construct necessary indices
	p := s.buildPeg()
	// for each input index, in reverse order
	for i := s.GetRightExtent(); i >= 0; i-- {
		// keep going over the list of nonterminals until all of them are
		// finished for this input index
		// TODO should pre-generate a topologically-sorted slice
		// TODO don't filter unordered nonterminals
		finished := make(map[symbols.NT]bool)
		for len(finished) < symbols.NumNTs {
			for ni := 0; ni < symbols.NumNTs; ni++ {
				nt := symbols.NT(ni)
				// skip nonterminals that are not ready yet
				if !allFinished(nt, finished) {
					continue
				}

				// find best key(s) by ordered choice on alternatives
				// TODO maybe break ties by furthest right extent?
				nKey := ntKey{nt, i}
				mins := make([]BSR, 0, 1)
				minAlt := noAlt
				for _, b := range p.nts[nKey] {
					// skip deleted nodes
					if !p.s.slotEntries[b] {
						continue
					}

					bAlt := b.Label.Alternate()
					if bAlt < minAlt {
						// delete old minimum alt(s)
						for _, old := range mins {
							p.deleteNode(old)
						}

						// set new
						mins = []BSR{b}
						minAlt = bAlt
					} else if bAlt == minAlt {
						// add equal-priority node to minimums
						mins = append(mins, b)
					} else {
						// delete lower-priority node
						p.deleteNode(b)
					}
				}
				// reset list of available matches for nt
				p.nts[nKey] = mins

				finished[nt] = true
			}
		}

		// eliminate slots broken by the removed nonterminals
		p.filterMissingSlots(i)
	}
}

// checks that all the nonterminals in the first list of nt are in the
// finished map
func allFinished(nt symbols.NT, finished map[symbols.NT]bool) bool {
	for _, f := range nt.LeftRec() {
		if !finished[f] {
			return false
		}
	}
	return true
}