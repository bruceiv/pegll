// Package parser is generated by gogll. Do not edit.
//
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
package parser

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/goccmack/gogll/examples/ambiguous1/goutil/bsr"
	"github.com/goccmack/gogll/examples/ambiguous1/goutil/md"
	"github.com/goccmack/gogll/examples/ambiguous1/parser/slot"

	"io/ioutil"
)

func ParseFile(fname string) (error, []*ParseError) {
	var buf []byte
	var err error
	if strings.HasSuffix(fname, ".md") {
		var str string
		str, err = md.GetSource(fname)
		if err != nil {
			parseErrorError(err)
		}
		buf = []byte(str)
	} else {
		buf, err = ioutil.ReadFile(fname)
		if err != nil {
			parseErrorError(err)
		}
	}
	return Parse(buf)
}

var (
	cI    = 0
	sz    = 0
	nextI = ""
	r     rune

	R *descriptors
	U *descriptors

	popped   map[poppedNode]bool
	crf      map[clusterNode][]*crfNode
	crfNodes map[crfNode]*crfNode

	input       []byte
	parseErrors []*ParseError
)

func initParser(I []byte) {
	input = I
	cI, nextI, sz = 0, "", 0
	R, U = &descriptors{}, &descriptors{}
	popped = make(map[poppedNode]bool)
	crf = map[clusterNode][]*crfNode{
		{"S", 0}: {},
	}
	crfNodes = map[crfNode]*crfNode{}
	bsr.Init("S", I)
	parseErrors = nil
}

func Parse(I []byte) (error, []*ParseError) {
	initParser(I)
	var L slot.Label
	m, cU := len(I), 0
	nextI, r, sz = decodeRune(I[cI:])
	ntAdd("S", 0)
	// DumpDescriptors()
	for !R.empty() {
		L, cU, cI = R.remove()
		nextI, r, sz = decodeRune(I[cI:])

		// fmt.Println()
		// fmt.Printf("L:%s, cI:%d, I[cI]:%s, cU:%d\n", L, cI, nextI, cU)
		// DumpDescriptors()

		switch L {
		case slot.A0R0: // A : ∙a

			bsr.Add(slot.A0R1, cU, cI, cI+sz)
			cI += sz
			nextI, r, sz = decodeRune(I[cI:])
			if follow["A"]() {
				rtn("A", cU, cI)
			}
		case slot.B0R0: // B : ∙letter

			bsr.Add(slot.B0R1, cU, cI, cI+sz)
			cI += sz
			nextI, r, sz = decodeRune(I[cI:])
			if follow["B"]() {
				rtn("B", cU, cI)
			}
		case slot.S0R0: // S : ∙A S

			call(slot.S0R1, cU, cI)
		case slot.S0R1: // S : A ∙S

			if !testSelect[slot.S0R1]() {
				parseError(slot.S0R1, cI)
				break
			}

			call(slot.S0R2, cU, cI)
		case slot.S0R2: // S : A S ∙

			if follow["S"]() {
				rtn("S", cU, cI)
			}
		case slot.S1R0: // S : ∙B S

			call(slot.S1R1, cU, cI)
		case slot.S1R1: // S : B ∙S

			if !testSelect[slot.S1R1]() {
				parseError(slot.S1R1, cI)
				break
			}

			call(slot.S1R2, cU, cI)
		case slot.S1R2: // S : B S ∙

			if follow["S"]() {
				rtn("S", cU, cI)
			}
		case slot.S2R0: // S : ∙
			bsr.AddEmpty(slot.S2R0, cI)

			if follow["S"]() {
				rtn("S", cU, cI)
			}

		default:
			panic("This must not happen")
		}
	}
	if !bsr.Contain("S", 0, m) {
		sortParseErrors(I)
		err := fmt.Errorf("Error: Parse Failed right extent=%d, m=%d",
			bsr.GetRightExtent(), len(I))
		return err, parseErrors
	}
	return nil, nil
}

func ntAdd(nt string, j int) {
	// fmt.Printf("ntAdd(%s, %d)\n", nt, j)
	failed := true
	for _, l := range slot.GetAlternates(nt) {
		if testSelect[l]() {
			dscAdd(l, j, j)
		} else {
			failed = false
		}
	}
	if failed {
		for _, l := range slot.GetAlternates(nt) {
			parseError(l, j)
		}
	}
}

/*** Call Return Forest ***/

type poppedNode struct {
	X    string
	k, j int
}

type clusterNode struct {
	X string
	k int
}

type crfNode struct {
	L slot.Label
	i int
}

/*
suppose that L is Y ::=αX ·β
if there is no CRF node labelled (L,i)
	create one let u be the CRF node labelled (L,i)
if there is no CRF node labelled (X, j) {
	create a CRF node v labelled (X, j)
	create an edge from v to u
	ntAdd(X, j)
} else {
	let v be the CRF node labelled (X, j)
	if there is not an edge from v to u {
		create an edge from v to u
		for all ((X, j,h)∈P) {
			dscAdd(L, i, h);
			bsrAdd(L, i, j, h)
		}
	}
}
*/
func call(L slot.Label, i, j int) {
	// fmt.Printf("call(%s,%d,%d)\n", L,i,j)
	u, exist := crfNodes[crfNode{L, i}]
	// fmt.Printf("  u exist=%t\n", exist)
	if !exist {
		u = &crfNode{L, i}
		crfNodes[*u] = u
	}
	X := L.Symbols()[L.Pos()-1]
	ndV := clusterNode{X, j}
	v, exist := crf[ndV]
	if !exist {
		// fmt.Println("  v !exist")
		crf[ndV] = []*crfNode{u}
		ntAdd(X, j)
	} else {
		// fmt.Println("  v exist")
		if !existEdge(v, u) {
			// fmt.Printf("  !existEdge(%v)\n", u)
			crf[ndV] = append(v, u)
			// fmt.Printf("|popped|=%d\n", len(popped))
			for pnd, _ := range popped {
				if pnd.X == X && pnd.k == j {
					dscAdd(L, i, pnd.j)
					bsr.Add(L, i, j, pnd.j)
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

func rtn(X string, k, j int) {
	// fmt.Printf("rtn(%s,%d,%d)\n", X,k,j)
	p := poppedNode{X, k, j}
	if _, exist := popped[p]; !exist {
		popped[p] = true
		for _, nd := range crf[clusterNode{X, k}] {
			dscAdd(nd.L, nd.i, j)
			bsr.Add(nd.L, nd.i, k, j)
		}
	}
}

func CRFString() string {
	buf := new(bytes.Buffer)
	buf.WriteString("CRF: {")
	for cn, nds := range crf {
		for _, nd := range nds {
			fmt.Fprintf(buf, "%s->%s, ", cn, nd)
		}
	}
	buf.WriteString("}")
	return buf.String()
}

func (cn clusterNode) String() string {
	return fmt.Sprintf("(%s,%d)", cn.X, cn.k)
}

func (n crfNode) String() string {
	return fmt.Sprintf("(%s,%d)", n.L.String(), n.i)
}

func PoppedString() string {
	buf := new(bytes.Buffer)
	buf.WriteString("Popped: {")
	for p, _ := range popped {
		fmt.Fprintf(buf, "(%s,%d,%d) ", p.X, p.k, p.j)
	}
	buf.WriteString("}")
	return buf.String()
}

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

func dscAdd(L slot.Label, k, i int) {
	// fmt.Printf("dscAdd(%s,%d,%d)\n", L, k, i)
	d := &descriptor{L, k, i}
	if !U.contain(d) {
		R.set = append(R.set, d)
		U.set = append(U.set, d)
	}
}

func (ds *descriptors) remove() (L slot.Label, k, i int) {
	d := ds.set[len(ds.set)-1]
	ds.set = ds.set[:len(ds.set)-1]
	// fmt.Printf("remove: %s,%d,%d\n", d.L, d.k, d.i)
	return d.L, d.k, d.i
}

func DumpDescriptors() {
	DumpR()
	DumpU()
}

func DumpR() {
	fmt.Println("R:")
	for _, d := range R.set {
		fmt.Printf(" %s\n", d)
	}
}

func DumpU() {
	fmt.Println("U:")
	for _, d := range U.set {
		fmt.Printf(" %s\n", d)
	}
}

/*** Rune decoding ***/
func decodeRune(str []byte) (string, rune, int) {
	if len(str) == 0 {
		return "$", '$', 0
	}
	r, sz := utf8.DecodeRune(str)
	if r == utf8.RuneError {
		panic(fmt.Sprintf("Rune error: %s", str))
	}
	switch r {
	case '\\':
		r, sz = utf8.DecodeRune(str)
		if r == utf8.RuneError {
			panic(fmt.Sprintf("Rune error: %s", str))
		}
		switch r {
		case '"':
			return "\"", r, sz
		case 'n':
			return "n", r, sz
		case 'r':
			return "r", r, sz
		case 't':
			return "t", r, sz
		case '\\':
			return "\\", r, sz
		case '\'':
			return "'", r, sz
		}
	case '\t', ' ':
		return "space", r, sz
	case '\n':
		return "\\n", r, sz
	}
	return string(str[:sz]), r, sz
}

func runeToString(r rune) string {
	buf := make([]byte, utf8.RuneLen(r))
	utf8.EncodeRune(buf, r)
	return string(buf)
}

/*** TestSelect ***/

var testSelect = map[slot.Label]func() bool{
	slot.A0R0: func() bool {
		return r == 'a'
	},

	slot.A0R1: func() bool {
		return r == '$' ||
			r == 'a' ||
			letter(r)
	},

	slot.B0R0: func() bool {
		return letter(r)
	},

	slot.B0R1: func() bool {
		return r == '$' ||
			r == 'a' ||
			letter(r)
	},

	slot.S0R0: func() bool {
		return r == 'a'
	},

	slot.S0R1: func() bool {
		return r == 'a' ||
			letter(r) ||
			r == '$'
	},

	slot.S0R2: func() bool {
		return r == '$'
	},

	slot.S1R0: func() bool {
		return letter(r)
	},

	slot.S1R1: func() bool {
		return r == 'a' ||
			letter(r) ||
			r == '$'
	},

	slot.S1R2: func() bool {
		return r == '$'
	},

	slot.S2R0: func() bool {
		return r == '$'
	},
}

var follow = map[string]func() bool{
	"A": func() bool {
		return r == '$' ||
			r == 'a' ||
			letter(r)
	},

	"B": func() bool {
		return r == '$' ||
			r == 'a' ||
			letter(r)
	},

	"S": func() bool {
		return r == '$'
	},
}

/*** Unicode functions ***/

func any(r rune) bool {
	return true
}

func anyof(r rune, set string) bool {
	return strings.ContainsRune(set, r)
}

func letter(r rune) bool {
	return unicode.IsLetter(r)
}

func number(r rune) bool {
	return unicode.IsNumber(r)
}

func upcase(r rune) bool {
	return unicode.IsUpper(r)
}

func lowcase(r rune) bool {
	return unicode.IsLower(r)
}

func not(r rune, set string) bool {
	bs := []byte(set)
	for i := 0; i < len(set); {
		r1, sz := utf8.DecodeRune(bs[i:])
		if r1 == utf8.RuneError {
			panic(fmt.Sprintf("Rune error: %s", set))
		}
		if r1 == r {
			return false
		}
		i += sz
	}
	return true
}

func space(r rune) bool {
	return unicode.IsSpace(r)
}

/*** Errors ***/

type ParseError struct {
	Slot         slot.Label
	InputPos     int
	Char         string
	Line, Column int
}

func (pe *ParseError) String() string {
	return fmt.Sprintf("Parse Error: %s cI=%d I[cI]=%s at line %d col %d",
		pe.Slot, pe.InputPos, pe.Char, pe.Line, pe.Column)
}

func parseError(slot slot.Label, I int) {
	pe := &ParseError{Slot: slot, InputPos: I, Char: nextI}
	parseErrors = append(parseErrors, pe)
}

func sortParseErrors(I []byte) {
	sort.Slice(parseErrors,
		func(i, j int) bool {
			return parseErrors[j].InputPos < parseErrors[i].InputPos
		})
	for _, pe := range parseErrors {
		pe.Line, pe.Column = GetLineColumn(pe.InputPos)
	}
}

func GetLineColumn(cI int) (line, col int) {
	line, col = 1, 1
	for j := 0; j < cI; {
		_, r, sz := decodeRune(input[j:])
		switch r {
		case '\n':
			line++
			col = 1
		case '\t':
			col += 4
		default:
			col++
		}
		j += sz
	}
	return
}

func parseErrorError(err error) {
	fmt.Printf("Error: %s\n", err)
	os.Exit(1)
}
