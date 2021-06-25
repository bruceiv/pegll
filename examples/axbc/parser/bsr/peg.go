// Introduces PEG-supporting operations over BSR sets

package bsr

// System-dependent maximum int value
// Based off https://stackoverflow.com/questions/6878590/the-maximum-value-for-an-int-type-in-go
const noAlt = int(^uint(0) >> 1)

// GetOrderedRoot returns the root of a parse tree using PEG semantics.
// In particular, it may match a prefix of the input (right extent is not
// string length), and possible matches are disambiguated using order of 
// the alternatives. Returns the root on success, nil on no match for the 
// root rule.
// NOTE: this does NOT check that all children of the root respect 
// ordered choice, there needs to be a bottom-up pass for that...
// might be able to generate a automatic AST?
func (s *Set) GetOrderedRoot() *BSR {
	var minAlt = noAlt
	var root *BSR = nil
	for b := range s.slotEntries {
		if b.Label.Head() == s.startSym && b.leftExtent == 0 && b.Label.Alternate() < minAlt {
			minAlt = b.Label.Alternate()
			root = &b
		}
	}
	return root
}