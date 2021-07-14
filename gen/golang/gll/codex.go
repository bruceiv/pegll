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

package gll

import (
	"bytes"
	"text/template"

	"github.com/bruceiv/pegll/ast"
	"github.com/bruceiv/pegll/gslot"
)

func (g *gen) genAlternatesCode() string {
	buf := new(bytes.Buffer)
	for _, nt := range g.g.NonTerminals.ElementsSorted() {
		buf.WriteString(g.getRuleCode(nt))
	}
	return buf.String()
}

const (
	ordered   = true
	unordered = false
)

func (g *gen) getRuleCode(nt string) string {
	buf := new(bytes.Buffer)
	rule := g.g.GetSyntaxRule(nt)
	if rule.IsOrdered || len(rule.Alternates) == 1 {
		// try alternates one-at-a-time if ordered or only one
		for i, alt := range rule.Alternates {
			buf.WriteString(g.getAlternateCode(nt, alt, i, len(rule.Alternates), ordered))
		}
	} else {
		// double-generate alternates of unordered rules to encode
		// match status in rule labels
		for i, alt := range rule.Alternates {
			buf.WriteString(g.getAlternateCode(nt, alt, i, len(rule.Alternates), unordered))
		}
	}
	// finish with failure alternate for all nonterminals that don't end with an empty alternate
	if !rule.AlwaysMatches() {
		buf.WriteString(g.getFailAlternate(nt, len(rule.Alternates)))
	}
	return buf.String()
}

type AltData struct {
	NT           string
	AltLabel     string
	PassAltLabel string
	AltComment   string
	HasPass      bool
	Slots        []*SlotData
	LastSlot     *SlotData
}

func (g *gen) getAlternateCode(nt string, alt *ast.SyntaxAlternate, altI, altN int, isOrdered bool) string {
	// fmt.Printf("codes.genAlternateCode: %s altI %d\n", nt, altI)
	var tmpl *template.Template
	var err error
	if alt.Empty() {
		tmpl, err = template.New("EmptyAlternate").Parse(eAltCodeTmpl)
	} else if isOrdered {
		tmpl, err = template.New("OrderedAlternate").Parse(oAltCodeTmpl)
	} else {
		tmpl, err = template.New("UnorderedAlternate").Parse(uAltCodeTmpl)
	}
	if err != nil {
		panic(err)
	}
	buf, data := new(bytes.Buffer), g.getAltData(nt, alt, altI, altN, isOrdered)
	if err = tmpl.Execute(buf, data); err != nil {
		panic(err)
	}
	return buf.String()
}

func (g *gen) getAltData(nt string, alt *ast.SyntaxAlternate, altI, altN int, isOrdered bool) *AltData {
	// fmt.Printf("codex.getAltData %s[%d]\n", nt, altI)
	L := gslot.NewLabel(nt, altI, 0, gslot.Unknown, g.gs, g.ff)
	d := &AltData{
		NT:           nt,
		AltLabel:     L.Label(),
		PassAltLabel: gslot.LabelFor(nt, altI, 0, gslot.Match),
		AltComment:   L.String(),
		HasPass:      altN > 1 && altI > 0,
	}
	if !alt.Empty() {
		d.Slots = g.getSlotsData(nt, alt, altI, altN, isOrdered)
		d.LastSlot = d.Slots[len(d.Slots)-1]
	}
	return d
}

func (g *gen) getSlotsData(nt string, alt *ast.SyntaxAlternate, altI, altN int, isOrdered bool) (data []*SlotData) {
	for i, sym := range alt.Symbols {
		// fmt.Printf("getSlotsData(%s) %s\n", nt, getSlotData(nt, altI, sym, i))
		data = append(data, g.getSlotData(nt, altI, altN, sym.String(), i, isOrdered))
	}
	return
}

func (g *gen) getSlotData(nt string, altI, altN int, symbol string, pos int, isOrdered bool) *SlotData {
	preLabel := gslot.NewLabel(nt, altI, pos, gslot.Unknown, g.gs, g.ff)
	postLabel := gslot.NewLabel(nt, altI, pos+1, gslot.Unknown, g.gs, g.ff)
	var failLabel string
	var passLabel string
	if altI+1 < altN {
		failLabel = `slot.` + gslot.LabelFor(nt, altI+1, 0, gslot.Unknown)
		passLabel = `slot.` + gslot.LabelFor(nt, altI+1, 0, gslot.Match)
	} else {
		failLabel = `slot.` + gslot.LabelFor(nt, altN, 0, gslot.Fail)
		passLabel = `failInd`
	}
	sd := &SlotData{
		AltLabel:      gslot.LabelFor(nt, altI, 0, gslot.Unknown),
		PreLabel:      preLabel.Label(),
		PostLabel:     postLabel.Label(),
		FailLabel:     failLabel,
		PassLabel:     passLabel,
		PassPostLabel: gslot.LabelFor(nt, altI, pos, gslot.Match),
		Comment:       postLabel.String(),
		NotLastAlt:    altI+1 < altN,
		IsNT:          false,
		IsPLook:       false,
		IsNLook:       false,
		Head:          nt,
	}
	if g.g.Terminals.Contain(symbol) {
		sd.CallNT = "<error: not NT>"
	} else if g.g.Lookaheads.Contain(symbol) {
		switch symbol[0] {
		case '&':
			sd.IsPLook = true
		case '!':
			sd.IsNLook = true
		default:
			panic("Invalid lookahead symbol: " + symbol)
		}
		sd.CallNT = symbol[1:]
	} else {
		sd.IsNT = true
		sd.CallNT = symbol
	}
	// fmt.Printf("getSlotData: altlabel:%s, pre:%s, post:%s\n",
	// 	sd.AltLabel, sd.PreLabel, sd.PostLabel)
	return sd
}

func (g *gen) getFailAlternate(nt string, nAlt int) string {
	return `		case slot.` + gslot.LabelFor(nt, nAlt, 0, gslot.Fail) + `: // ` + nt + ` failure case
			p.rtn(symbols.NT_` + nt + `, cU, failInd)
	`
}

type SlotData struct {
	AltLabel      string
	PreLabel      string
	PostLabel     string
	FailLabel     string
	PassLabel     string
	PassPostLabel string
	Comment       string
	IsNT          bool
	IsPLook       bool
	IsNLook       bool
	CallNT        string
	NotLastAlt    bool
	Head          string
}

const eAltCodeTmpl = `case slot.{{.AltLabel}}: // {{.AltComment}}
			p.bsrSet.AddEmpty(slot.{{.AltLabel}},p.cI)
			p.rtn(symbols.NT_{{.NT}}, cU, p.cI)
	`

const oAltCodeTmpl = `		case slot.{{.AltLabel}}: // {{.AltComment}}
		{{range $i, $slot := .Slots}}
			if !p.testSelect(slot.{{$slot.PreLabel}}){ 
				p.parseError(slot.{{$slot.PreLabel}}, p.cI, first[slot.{{$slot.PreLabel}}])
				L, p.cI = {{$slot.FailLabel}}, cU
				goto nextSlot
			}
			{{if $slot.IsNT}}p.call(slot.{{$slot.PostLabel}}, {{$slot.FailLabel}}, symbols.NT_{{$slot.CallNT}}, cU, p.cI)
		case slot.{{$slot.PostLabel}}: // {{$slot.Comment}} 
			{{else if $slot.IsPLook}}p.call(slot.{{$slot.PostLabel}}, {{$slot.FailLabel}}, symbols.NT_{{$slot.CallNT}}, cU, p.cI)
		case slot.{{$slot.PostLabel}}: // {{$slot.Comment}}
			{{else if $slot.IsNLook}}p.call({{$slot.FailLabel}}, slot.{{$slot.PostLabel}}, symbols.NT_{{$slot.CallNT}}, cU, p.cI)
		case slot.{{$slot.PostLabel}}: // {{$slot.Comment}}
			{{else}}p.bsrSet.Add(slot.{{$slot.PostLabel}}, cU, p.cI, p.cI+1)
			p.cI++ {{end}}{{end}}
			p.rtn(symbols.NT_{{.NT}}, cU, p.cI)
	`

const uAltCodeTmpl = `		case slot.{{.AltLabel}}: // {{.AltComment}}
		{{range $i, $slot := .Slots}}
			if !p.testSelect(slot.{{$slot.PreLabel}}){ 
				p.parseError(slot.{{$slot.PreLabel}}, p.cI, first[slot.{{$slot.PreLabel}}])
				L, p.cI = {{$slot.FailLabel}}, cU
				goto nextSlot
			}
			{{if $slot.IsNT}}p.call(slot.{{$slot.PostLabel}}, {{$slot.FailLabel}}, symbols.NT_{{$slot.CallNT}}, cU, p.cI)
		case slot.{{$slot.PostLabel}}: // {{$slot.Comment}} 
			{{else if $slot.IsPLook}}p.call(slot.{{$slot.PostLabel}}, {{$slot.FailLabel}}, symbols.NT_{{$slot.CallNT}}, cU, p.cI)
		case slot.{{$slot.PostLabel}}: // {{$slot.Comment}}
			{{else if $slot.IsNLook}}p.call({{$slot.FailLabel}}, slot.{{$slot.PostLabel}}, symbols.NT_{{$slot.CallNT}}, cU, p.cI)
		case slot.{{$slot.PostLabel}}: // {{$slot.Comment}}
			{{else}}p.bsrSet.Add(slot.{{$slot.PostLabel}}, cU, p.cI, p.cI+1)
			p.cI++ {{end}}
			p.rtn(symbols.NT_{{$slot.Head}}, cU, p.cI)
			{{if .NotLastAlt}}L, p.cI = {{$slot.PassLabel}}, cU
			goto nextSlot{{end}}{{end}}
		{{if .HasPass}}case slot.{{.PassAltLabel}}: // {{.AltComment}} [with previous match]
		{{range $i, $slot := .Slots}}
		if !p.testSelect(slot.{{$slot.PreLabel}}){ 
			p.parseError(slot.{{$slot.PreLabel}}, p.cI, first[slot.{{$slot.PreLabel}}])
			{{if .NotLastAlt}}L, p.cI = {{$slot.PassLabel}}, cU
			goto nextSlot{{end}}
		}
		{{if $slot.IsNT}}p.call(slot.{{$slot.PassPostLabel}}, {{$slot.PassLabel}}, symbols.NT_{{$slot.CallNT}}, cU, p.cI)
	case slot.{{$slot.PassPostLabel}}: // {{$slot.Comment}} 
		{{else if $slot.IsPLook}}p.call(slot.{{$slot.PostLabel}}, {{$slot.PassLabel}}, symbols.NT_{{$slot.CallNT}}, cU, p.cI)
	case slot.{{$slot.PostLabel}}: // {{$slot.Comment}}
		{{else if $slot.IsNLook}}p.call({{$slot.PassLabel}}, slot.{{$slot.PostLabel}}, symbols.NT_{{$slot.CallNT}}, cU, p.cI)
	case slot.{{$slot.PostLabel}}: // {{$slot.Comment}}
		{{else}}p.bsrSet.Add(slot.{{$slot.PostLabel}}, cU, p.cI, p.cI+1)
		p.cI++ {{end}}
		p.rtn(symbols.NT_{{$slot.Head}}, cU, p.cI)
		{{if .NotLastAlt}}L, p.cI = {{$slot.PassLabel}}, cU
		goto nextSlot{{end}}
		{{end}}{{end}}
	`
