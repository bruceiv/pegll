// Package parser is generated by gogll. Do not edit.
package parser

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"axbc/lexer"
	"axbc/parser/bsr"
	"axbc/parser/slot"
	"axbc/parser/symbols"
	"axbc/token"
)

type parser struct {
	cI int

	R *descriptors
	U *descriptors

	popped   map[poppedNode]bool
	crf_m    map[clusterNode][]*crfNode
	crf_f    map[clusterNode][]*crfNode
	crfNodes map[crfNode]*crfNode

	lex         *lexer.Lexer
	parseErrors []*Error

	bsrSet *bsr.Set
}

// index used for non-matches
const failInd = -1

func newParser(l *lexer.Lexer) *parser {
	return &parser{
		cI:     0,
		lex:    l,
		R:      &descriptors{},
		U:      &descriptors{},
		popped: make(map[poppedNode]bool),
		crf_m: map[clusterNode][]*crfNode{
			{symbols.NT_AxBC, 0}: {},
		},
		crf_f:       map[clusterNode][]*crfNode{},
		crfNodes:    map[crfNode]*crfNode{},
		bsrSet:      bsr.New(symbols.NT_AxBC, l),
		parseErrors: nil,
	}
}

// Parse returns the BSR set containing the parse forest.
// If the parse was successfull []*Error is nil
func Parse(l *lexer.Lexer) (*bsr.Set, []*Error) {
	return newParser(l).parse()
}

func (p *parser) parse() (*bsr.Set, []*Error) {
	var L slot.Label
	m, cU := len(p.lex.I), 0
	p.ntAdd(symbols.NT_AxBC, 0)
	// p.DumpDescriptors()
	for !p.R.empty() {
		L, cU, p.cI = p.R.remove()
		tokens := p.lex.Tokens(p.cI)
		origTokens := tokens
		var rext int
		var ok bool

		// fmt.Println()
		// fmt.Printf("L:%s, cI:%d, I[p.cI]:%s, cU:%d\n", L, p.cI, p.lex.Tokens[p.cI], cU)
		// p.DumpDescriptors()

		for {
			switch L {
			case slot.AorB0R0: // AorB : ∙Repa0x
				rext, ok = p.testSelect(slot.AorB0R0, tokens)
				if !ok {
					p.parseError(slot.AorB0R0, p.cI, tokens, first[slot.AorB0R0])
					L, p.cI = slot.AorB1R0, cU
					goto nextSlot
				}
				p.call(slot.AorB0R1, slot.AorB1R0, symbols.NT_Repa0x, cU, p.cI)
			case slot.AorB0R1: // AorB : Repa0x ∙

				p.rtn(symbols.NT_AorB, cU, p.cI)
			case slot.AorB1R0: // AorB : ∙a b
				rext, ok := p.testSelect(slot.AorB1R0, tokens)
				if !ok {
					p.parseError(slot.AorB1R0, p.cI, tokens, first[slot.AorB1R0])
					L, p.cI = fail_AorB, cU
					goto nextSlot
				}
				p.bsrSet.Add(slot.AorB1R1, cU, p.cI, rext)
				p.cI = rext
				tokens = p.lex.Tokens(p.cI)
				rext, ok = p.testSelect(slot.AorB1R1, tokens)
				if !ok {
					p.parseError(slot.AorB1R1, p.cI, tokens, first[slot.AorB1R1])
					L, p.cI = fail_AorB, cU
					tokens = origTokens
					goto nextSlot
				}
				p.bsrSet.Add(slot.AorB1R2, cU, p.cI, rext)
				p.cI = rext
				p.rtn(symbols.NT_AorB, cU, p.cI)
			case slot.AorB2F0: // AorB failure case
				p.rtn(symbols.NT_AorB, cU, failInd)
			case slot.AxBC0R0: // AxBC : ∙AorB c
				rext, ok = p.testSelect(slot.AxBC0R0, tokens)
				if !ok {
					p.parseError(slot.AxBC0R0, p.cI, tokens, first[slot.AxBC0R0])
					L, p.cI = fail_AxBC, cU
					goto nextSlot
				}
				p.call(slot.AxBC0R1, slot.AxBC1F0, symbols.NT_AorB, cU, p.cI)
			case slot.AxBC0R1: // AxBC : AorB ∙c
				rext, ok = p.testSelect(slot.AxBC0R1, tokens)
				if !ok {
					p.parseError(slot.AxBC0R1, p.cI, tokens, first[slot.AxBC0R1])
					L, p.cI = fail_AxBC, cU
					goto nextSlot
				}
				p.bsrSet.Add(slot.AxBC0R2, cU, p.cI, rext)
				p.cI = rext
				p.rtn(symbols.NT_AxBC, cU, p.cI)
			case slot.AxBC1F0: // AxBC failure case
				p.rtn(symbols.NT_AxBC, cU, failInd)
			case slot.Repa0x0R0: // Repa0x : ∙a Repa0x
				rext, ok = p.testSelect(slot.Repa0x0R0, tokens)
				if !ok {
					p.parseError(slot.Repa0x0R0, p.cI, tokens, first[slot.Repa0x0R0])
					L, p.cI = slot.Repa0x1R0, cU
					goto nextSlot
				}
				p.bsrSet.Add(slot.Repa0x0R1, cU, p.cI, rext)
				p.cI = rext
				tokens = p.lex.Tokens(p.cI)
				rext, ok = p.testSelect(slot.Repa0x0R1, tokens)
				if !ok {
					p.parseError(slot.Repa0x0R1, p.cI, tokens, first[slot.Repa0x0R1])
					L, p.cI = slot.Repa0x1R0, cU
					tokens = origTokens
					goto nextSlot
				}
				p.call(slot.Repa0x0R2, slot.Repa0x1R0, symbols.NT_Repa0x, cU, p.cI)
			case slot.Repa0x0R2: // Repa0x : a Repa0x ∙

				p.rtn(symbols.NT_Repa0x, cU, p.cI)
			case slot.Repa0x1R0: // Repa0x : ∙
				p.bsrSet.AddEmpty(slot.Repa0x1R0, p.cI)
				p.rtn(symbols.NT_Repa0x, cU, p.cI)

			default:
				panic("This must not happen")
			}
			// if exit switch normally, also exit loop and proceed to next
			// descriptor; if exit with goto nextSlot, repeat switch at next
			// slot
			break
		nextSlot:
		}
	}
	if !p.bsrSet.Contain(symbols.NT_AxBC, 0, m) {
		p.sortParseErrors()
		return nil, p.parseErrors
	}
	return p.bsrSet, nil
}

func (p *parser) ntAdd(nt symbols.NT, j int) {
	l := slot.GetAlternates(nt)[0]
	p.dscAdd(l, j, j)
}

/*** Call Return Forest ***/

type poppedNode struct {
	X    symbols.NT
	k, j int
}

type clusterNode struct {
	X symbols.NT
	k int
}

type crfNode struct {
	L slot.Label
	i int
}

func (p *parser) call(Lm, Lf slot.Label, X symbols.NT, i, j int) {
	// fmt.Printf("p.call(%s,%d,%d)\n", L,i,j)
	um, exist := p.crfNodes[crfNode{Lm, i}]
	// fmt.Printf("  u exist=%t\n", exist)
	if !exist {
		um = &crfNode{Lm, i}
		p.crfNodes[*um] = um
	}
	uf, exist := p.crfNodes[crfNode{Lf, i}]
	if !exist {
		uf = &crfNode{Lf, i}
		p.crfNodes[*uf] = uf
	}

	ndV := clusterNode{X, j}
	vm, existm := p.crf_m[ndV]
	vf, existf := p.crf_f[ndV]
	if !existm && !existf {
		// fmt.Println("  v !exist")
		p.crf_m[ndV] = []*crfNode{um}
		p.crf_f[ndV] = []*crfNode{uf}
		p.ntAdd(X, j)
	} else {
		// fmt.Println("  v exist")
		if !existEdge(vm, um) {
			// fmt.Printf("  !existEdge(%v)\n", u)
			p.crf_m[ndV] = append(vm, um)
			// fmt.Printf("|popped|=%d\n", len(popped))
			for pnd := range p.popped {
				if pnd.X == X && pnd.k == j && pnd.j != failInd {
					p.addMatch(Lm, i, j, pnd.j)
				}
			}
		}
		if !existEdge(vf, uf) {
			p.crf_f[ndV] = append(vf, uf)
			for pnd := range p.popped {
				if pnd.X == X && pnd.k == j && pnd.j == failInd {
					p.addFail(Lf, i, j)
				}
			}
		}
	}
}

func existEdge(nds []*crfNode, nd *crfNode) bool {
	for _, nd1 := range nds {
		if nd1 == nd {
			return true
		}
	}
	return false
}

func (p *parser) rtn(X symbols.NT, k, j int) {
	// fmt.Printf("p.rtn(%s,%d,%d)\n", X,k,j)
	pn := poppedNode{X, k, j}
	if _, exist := p.popped[pn]; !exist {
		p.popped[pn] = true
		if j != failInd {
			for _, nd := range p.crf_m[clusterNode{X, k}] {
				p.addMatch(nd.L, nd.i, k, j)
			}
		} else {
			for _, nd := range p.crf_f[clusterNode{X, k}] {
				p.addFail(nd.L, nd.i, k)
			}
		}
	}
}

func (p *parser) addMatch(L slot.Label, i, k, j int) {
	p.bsrSet.Add(L, i, k, j)
	if L.IsLookahead() {
		p.dscAdd(L, i, k)
	} else {
		p.dscAdd(L, i, j)
	}
}

func (p *parser) addFail(L slot.Label, i, k int) {
	if L.IsLookahead() {
		p.dscAdd(L, i, k)
	} else {
		p.dscAdd(L, i, i)
	}
}

// func CRFString() string {
// 	buf := new(bytes.Buffer)
// 	buf.WriteString("CRF: {")
// 	for cn, nds := range crf{
// 		for _, nd := range nds {
// 			fmt.Fprintf(buf, "%s->%s, ", cn, nd)
// 		}
// 	}
// 	buf.WriteString("}")
// 	return buf.String()
// }

func (cn clusterNode) String() string {
	return fmt.Sprintf("(%s,%d)", cn.X, cn.k)
}

func (n crfNode) String() string {
	return fmt.Sprintf("(%s,%d)", n.L.String(), n.i)
}

// func PoppedString() string {
// 	buf := new(bytes.Buffer)
// 	buf.WriteString("Popped: {")
// 	for p, _ := range popped {
// 		fmt.Fprintf(buf, "(%s,%d,%d) ", p.X, p.k, p.j)
// 	}
// 	buf.WriteString("}")
// 	return buf.String()
// }

/*** descriptors ***/

type descriptors struct {
	set []*descriptor
}

func (ds *descriptors) contain(d *descriptor) bool {
	for _, d1 := range ds.set {
		if d1 == d {
			return true
		}
	}
	return false
}

func (ds *descriptors) empty() bool {
	return len(ds.set) == 0
}

func (ds *descriptors) String() string {
	buf := new(bytes.Buffer)
	buf.WriteString("{")
	for i, d := range ds.set {
		if i > 0 {
			buf.WriteString("; ")
		}
		fmt.Fprintf(buf, "%s", d)
	}
	buf.WriteString("}")
	return buf.String()
}

type descriptor struct {
	L slot.Label
	k int
	i int
}

func (d *descriptor) String() string {
	return fmt.Sprintf("%s,%d,%d", d.L, d.k, d.i)
}

func (p *parser) dscAdd(L slot.Label, k, i int) {
	// fmt.Printf("p.dscAdd(%s,%d,%d)\n", L, k, i)
	d := &descriptor{L, k, i}
	if !p.U.contain(d) {
		p.R.set = append(p.R.set, d)
		p.U.set = append(p.U.set, d)
	}
}

func (ds *descriptors) remove() (L slot.Label, k, i int) {
	d := ds.set[len(ds.set)-1]
	ds.set = ds.set[:len(ds.set)-1]
	// fmt.Printf("remove: %s,%d,%d\n", d.L, d.k, d.i)
	return d.L, d.k, d.i
}

func (p *parser) DumpDescriptors() {
	p.DumpR()
	p.DumpU()
}

func (p *parser) DumpR() {
	fmt.Println("R:")
	for _, d := range p.R.set {
		fmt.Printf(" %s\n", d)
	}
}

func (p *parser) DumpU() {
	fmt.Println("U:")
	for _, d := range p.U.set {
		fmt.Printf(" %s\n", d)
	}
}

/*** TestSelect ***/

// func (p *parser) follow(nt symbols.NT) bool {
// 	// TODO update for new lexer
// 	_, exist := followSets[nt][p.lex.Tokens[p.cI].Type()]
// 	return exist
// }

func (p *parser) testSelect(l slot.Label, tokens *lexer.TokenSet) (int, bool) {
	// longest munch found so far; -1 for none such
	best := -1
	// check for nullable rule
	if l.IsNullable() {
		best = p.cI
	}
	// cycle through rules checking for valid tokens
	for tok, rext := range *tokens {
		// will exclude shorter than best matches as well as -1 for no match
		if rext <= best {
			continue
		}
		// keep matches for contained tokens
		if l.FirstContains(token.Type(tok)) {
			best = rext
		}
	}
	return best, best != -1
	// return l.IsNullable() || l.FirstContains(p.lex.Tokens[p.cI].Type())
	// _, exist := first[l][p.lex.Tokens[p.cI].Type()]
	// return exist
}

var first = []map[token.Type]string{
	// AorB : ∙Repa0x
	{
		token.T_0: "a",
		token.T_2: "c",
	},
	// AorB : Repa0x ∙
	{
		token.T_2: "c",
	},
	// AorB : ∙a b
	{
		token.T_0: "a",
	},
	// AorB : a ∙b
	{
		token.T_1: "b",
	},
	// AorB : a b ∙
	{
		token.T_2: "c",
	},
	// AorB : ∙
	{
		token.T_2: "c",
	},
	// AxBC : ∙AorB c
	{
		token.T_0: "a",
		token.T_2: "c",
	},
	// AxBC : AorB ∙c
	{
		token.T_2: "c",
	},
	// AxBC : AorB c ∙
	{
		token.EOF: "$",
	},
	// AxBC : ∙
	{
		token.EOF: "$",
	},
	// Repa0x : ∙a Repa0x
	{
		token.T_0: "a",
	},
	// Repa0x : a ∙Repa0x
	{
		token.T_0: "a",
		token.T_2: "c",
	},
	// Repa0x : a Repa0x ∙
	{
		token.T_2: "c",
	},
	// Repa0x : ∙
	{
		token.T_2: "c",
	},
}

var followSets = []map[token.Type]string{
	// AorB
	{
		token.T_2: "c",
	},
	// AxBC
	{
		token.EOF: "$",
	},
	// Repa0x
	{
		token.T_2: "c",
	},
}

/*** Errors ***/

/*
Error is returned by Parse at every point at which the parser fails to parse
a grammar production. For non-LL-1 grammars there will be an error for each
alternate attempted by the parser.

The errors are sorted in descending order of input position (index of token in
the stream of tokens).

Normally the error of interest is the one that has parsed the largest number of
tokens.
*/
type Error struct {
	// Input index of error.
	cI int

	// Grammar slot at which the error occured.
	Slot slot.Label

	// Lexer to recover values
	lex *lexer.Lexer

	// The tokens at which the error occurred.
	Tokens *lexer.TokenSet

	// The line and column in the input text at which the error occurred
	Line, Column int

	// The tokens expected at the point where the error occurred
	Expected map[token.Type]string
}

func (pe *Error) String() string {
	w := new(bytes.Buffer)
	fmt.Fprintf(w, "Parse Error: %s at line %d col %d\n",
		pe.Slot, pe.Line, pe.Column)

	fmt.Fprintf(w, "Got: [")
	isFirst := true
	for tok, rext := range *pe.Tokens {
		if rext == -1 {
			continue
		}
		if isFirst {
			isFirst = false
		} else {
			fmt.Fprintf(w, ", ")
		}
		fmt.Fprintf(w, "%s %q", (token.Type(tok)).ID(), pe.lex.I[pe.cI:rext])
	}
	fmt.Fprintf(w, "]\n")

	exp := []string{}
	for _, e := range pe.Expected {
		exp = append(exp, e)
	}
	fmt.Fprintf(w, "Expected one of: [%s]", strings.Join(exp, ", "))

	return w.String()
}

func (p *parser) parseError(slot slot.Label, i int, got *lexer.TokenSet, expected map[token.Type]string) {
	pe := &Error{cI: i, Slot: slot, lex: p.lex, Tokens: got, Expected: expected}
	p.parseErrors = append(p.parseErrors, pe)
}

func (p *parser) sortParseErrors() {
	sort.Slice(p.parseErrors,
		func(i, j int) bool {
			return p.parseErrors[j].cI < p.parseErrors[i].cI
		})
	for _, pe := range p.parseErrors {
		pe.Line, pe.Column = p.lex.GetLineColumn(pe.cI)
	}
}
