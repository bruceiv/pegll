// Introduces PEG-supporting operations over BSR sets

package bsr

import (
	"axbc/parser/slot"
	sym "axbc/parser/symbols"
)

// index key for count of slots with given extents
type slotKey struct {
	label slot.Label
	left  int
	right int
}

// index key for nonterminal-match lookup
type ntKey struct {
	nt   sym.NT
	left int
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
	nSlots     map[slotKey]int
	nts        map[ntKey][]BSR
	pivots     map[int][]BSR
	successors map[succKey][]BSR
	s *Set
}

// System-dependent maximum int value
// Based off https://stackoverflow.com/questions/6878590/the-maximum-value-for-an-int-type-in-go
const noAlt = int(^uint(0) >> 1)

// Sets up indices used in PEG construction
func (s *Set) buildPeg() *pegBsr {
	p := &pegBsr{
		nSlots: make(map[slotKey]int),
		nts: make(map[ntKey][]BSR),
		pivots: make(map[int][]BSR),
		successors: make(map[succKey][]BSR),
		s: s,
	}

	for b := range s.slotEntries {
		x := b.Label.Slot()
		// count of slots with same left/right extent
		p.nSlots[slotKey{b.Label, b.leftExtent, b.rightExtent}] += 1
		if x.EoR() {
			// nonterminal match at end-of-rule
			nKey := ntKey{x.NT, b.leftExtent}
			p.nts[nKey] = append(p.nts[nKey], b)
		} else {
			// internal nodes with the given pivot
			p.pivots[b.pivot] = append(p.pivots[b.pivot], b)
		}
		// slot successors
		sKey := succKey{b.Label, b.leftExtent, b.pivot}
		p.successors[sKey] = append(p.successors[sKey], b)
	}

	return p;
}

// Deletes a BSR from its set, returning false if it is the last copy of 
// that slot with that extent.
// Does not remove from nts, pivots, or successors indices
func (p *pegBsr) deleteNode(b BSR) bool {
	delete(p.s.slotEntries, b)
	sKey := slotKey{b.Label,b.leftExtent,b.rightExtent}
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
	succK := succKey{ slot.GetLabel(x.NT, x.Alt, x.Pos + 1), b.leftExtent, b.rightExtent }
	for _, b := range p.successors[succK] {
		p.deleteNodeAndSuccessors(b)
	}
	delete(p.successors, succK)
}

// Filters slots out of a BSR set that have had their nonterminal at 
// position i deleted, as well as any successors. 
func (p *pegBsr) filterMissingSlots(i int) {
	// loop over all nodes with the given pivot
	for _, b := range p.pivots[i] {
		// skip deleted nodes
		if ! p.s.slotEntries[b] {
			continue
		}
		x := b.Label.Slot()

		// get nonterminal after slot, skipping otherwise
		nt, ok := x.Symbols[x.Pos].(sym.NT)
		if ! ok {
			continue
		}

		// check there is still a match for this nonterminal with the same
		// right extent, deleting slot if none found
		found := false
		for _, nb := range p.nts[ntKey{nt, i}] {
			if p.s.slotEntries[nb] && nb.rightExtent == b.rightExtent {
				found = true
				break
			}
		}
		if ! found {
			p.deleteNodeAndSuccessors(b)
		}
	}
	// TODO could reset pivots[i] here, but I don't think it's used again
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
		finished := make(map[sym.NT]bool)
		for len(finished) < sym.NumNTs {
			for ni := 0; ni < sym.NumNTs; ni++ {
				nt := sym.NT(ni)
				// skip nonterminals that are not ready yet
				if ! allFinished(nt, finished) {
					continue
				}

				// find best key(s) by ordered choice on alternatives
				// TODO maybe break ties by furthest right extent?
				nKey := ntKey{nt, i}
				mins := make([]BSR, 0, 1)
				minAlt := noAlt
				for _, b := range p.nts[nKey] {
					// skip deleted nodes
					if ! p.s.slotEntries[b] {
						continue
					}

					bAlt := b.Label.Alternate()
					if bAlt < minAlt {
						// delete old minimum alt(s)
						for _, old := range mins {
							p.deleteNode(old)
						}

						// set new
						mins = []BSR{ b }
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

				finished[nt] = true;
			}
		}

		// eliminate slots broken by the removed nonterminals
		p.filterMissingSlots(i)
	}
}

// checks that all the nonterminals in the first list of nt are in the 
// finished map
func allFinished(nt sym.NT, finished map[sym.NT]bool) bool {
	for _, f := range nt.StartsWith() {
		if ! finished[f] {
			return false
		}
	}
	return true
}