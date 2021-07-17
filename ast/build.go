/*
Copyright 2021 Aaron Moss
Copyright 2020 Marius Ackerman

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Builds the Abstract Syntax Tree from a disambiguated parse BSR forest.

package ast

import (
	"fmt"
	"os"

	"github.com/bruceiv/pegll/lexer"
	"github.com/bruceiv/pegll/parser/bsr"
	"github.com/bruceiv/pegll/parser/symbols"
	"github.com/bruceiv/pegll/token"
	"github.com/bruceiv/pegll/util/runeset"
	"github.com/goccmack/goutil/stringset"
)

type builder struct {
	lex          *lexer.Lexer
	charLiterals *stringset.StringSet
}

// builder rule
type brule interface {
	isBrule()
}

func (*LexRule) isBrule()    {}
func (*SyntaxRule) isBrule() {}

// Build builds an AST from the BSR root. `root` is the root of a disambiguated BSR forest
//
func Build(root bsr.BSR, l *lexer.Lexer) *GoGLL {
	bld := &builder{
		lex:          l,
		charLiterals: stringset.New(),
	}
	gogll := bld.gogll(root)
	gogll.NonTerminals = bld.nonTerminals(gogll.SyntaxRules)
	gogll.StringLiterals = bld.getStringLiterals(gogll.SyntaxRules)
	gogll.Terminals = bld.terminals(gogll, gogll.GetStringLiterals())
	gogll.Lookaheads = bld.lookaheads(gogll)
	return gogll
}

// GoGLL : Package Rules ;
func (bld *builder) gogll(b bsr.BSR) *GoGLL {
	gogll := &GoGLL{
		Package: bld.packge(b.GetNTChild(symbols.NT_Package, 0)),
	}
	for _, rule := range bld.rules(b.GetNTChildI(1)) {
		switch r := rule.(type) {
		case *LexRule:
			bld.addLexRule(r, gogll)
		case *SyntaxRule:
			bld.addSyntaxRule(r, gogll)
		default:
			panic(fmt.Sprintf("Invalid %T", rule))
		}
	}
	return gogll
}

// Package : "package" string_lit ;
func (bld *builder) packge(b bsr.BSR) *Package {
	return &Package{
		tok: b.GetTChildI(1),
	}
}

// Rule : LexRule | SyntaxRule ;
func (bld *builder) rule(b bsr.BSR) brule {
	// fmt.Printf("build.rule: %s\n", b)
	if b.Alternate() == 0 {
		return bld.lexRule(b.GetNTChildI(0))
	}
	return bld.syntaxRule(b.GetNTChildI(0))
}

// Rules
//     :   Rule
//     |   Rule Rules
//     ;
func (bld *builder) rules(b bsr.BSR) []brule {
	rules := []brule{bld.rule(b.GetNTChildI(0))}
	if b.Alternate() == 1 {
		rules = append(rules, bld.rules(b.GetNTChildI(1))...)
	}
	return rules
}

func (bld *builder) nonTerminals(rules []*SyntaxRule) *stringset.StringSet {
	nts := stringset.New()
	for _, r := range rules {
		if nts.Contain(r.Head.ID()) {
			bld.fail(fmt.Errorf("Duplicate rule %s", r.Head.ID()), r.Head.Lext())
		} else {
			nts.Add(r.Head.ID())
		}
	}
	return nts
}

func (bld *builder) terminals(g *GoGLL, stringLiterals []string) *stringset.StringSet {
	terminals := bld.getLexRuleIDs(g.LexRules)
	terminals.Add(stringLiterals...)
	return terminals
}

func (bld *builder) lookaheads(g *GoGLL) *stringset.StringSet {
	ls := stringset.New()
	for _, r := range g.SyntaxRules {
		for _, a := range r.Alternates {
			for _, s := range a.Symbols {
				if l, ok := s.(*Lookahead); ok {
					ls.Add(l.ID())
				}
			}
		}
	}
	return ls
}

func (bld *builder) getLexRuleIDs(rules []*LexRule) *stringset.StringSet {
	terminals := stringset.New()
	for _, rule := range rules {
		if terminals.Contain(rule.ID()) {
			bld.fail(fmt.Errorf("duplicate lex rule %s", rule.ID()), rule.Lext())
		}
		terminals.Add(rule.ID())
	}
	return terminals
}

func (bld *builder) getStringLiterals(rules []*SyntaxRule) map[string]*StringLit {
	slits := make(map[string]*StringLit)
	for _, r := range rules {
		for _, a := range r.Alternates {
			for _, s := range a.Symbols {
				if sl, ok := s.(*StringLit); ok {
					slits[sl.ID()] = sl
				}
			}
		}
	}
	return slits
}

/*** Lex Rules ***/

// LexRule
//     : tokid ":" RegExp ";"
//     | "!" tokid ":" RegExp ";"
//     ;
func (bld *builder) lexRule(b bsr.BSR) *LexRule {
	if b.Alternate() == 0 {
		return &LexRule{
			TokID:  bld.tokID(b.GetTChildI(0)),
			RegExp: bld.regexp(b.GetNTChildI(2)),
		}
	}
	return &LexRule{
		Suppress: true,
		TokID:    bld.tokID(b.GetTChildI(1)),
		RegExp:   bld.regexp(b.GetNTChildI(3)),
	}
}

// RegExp : LexSymbol | LexSymbol RegExp ;
func (bld *builder) regexp(b bsr.BSR) *RegExp {
	re := &RegExp{
		Symbols: []LexSymbol{bld.lexSymbol(b.GetNTChildI(0))},
	}
	if b.Alternate() == 1 {
		re1 := bld.regexp(b.GetNTChildI(1))
		re.Symbols = append(re.Symbols, re1.Symbols...)
	}
	return re
}

// LexSymbol : "." | any string_lit | char_lit | LexBracket | not string_lit | UnicodeClass ;
func (bld *builder) lexSymbol(b bsr.BSR) LexSymbol {
	switch b.Alternate() {
	case 0:
		return bld.any(b.GetTChildI(0))
	case 1:
		return bld.anyOf(b.GetTChildI(0), b.GetTChildI(1))
	case 2:
		return bld.charLiteral(b.GetTChildI(0))
	case 3:
		return bld.lexBracket(b.GetNTChildI(0))
	case 4:
		return bld.not(b.GetTChildI(0), b.GetTChildI(1))
	case 5:
		return bld.unicodeClass(b.GetNTChildI(0))
	}
	panic(fmt.Sprintf("Invalid case %d", b.Alternate()))
}

// Any : "." ;
func (bld *builder) any(t *token.Token) *Any {
	return &Any{
		tok: t,
	}
}

func (bld *builder) anyOf(any, strLit *token.Token) *AnyOf {
	return &AnyOf{
		any:    any,
		strLit: strLit,
		Set:    bld.parseStringSet(strLit),
	}
}

func (bld *builder) charLiteral(tok *token.Token) *CharLiteral {
	switch tok.Literal()[1] {
	case '\\':
		if tok.Literal()[2] == '\'' {
			bld.charLiterals.Add("'")
		} else {
			bld.charLiterals.Add(string(tok.Literal()[1:3]))
		}
	case '"':
		bld.charLiterals.Add("\\\"")
	default:
		bld.charLiterals.Add(string(tok.Literal()[1:2]))
	}
	return NewCharLiteral(tok, tok.Literal())
}

// LexBracket : LexGroup | LexOptional | LexZeroOrMore | LexOneOrMore ;
func (bld *builder) lexBracket(b bsr.BSR) *LexBracket {
	switch b.Alternate() {
	case 0:
		return bld.lexGroup(b.GetNTChildI(0))
	case 1:
		return bld.lexOptional(b.GetNTChildI(0))
	case 2:
		return bld.lexZeroOrMore(b.GetNTChildI(0))
	case 3:
		return bld.lexOneOrMore(b.GetNTChildI(0))
	}
	panic("implement")
}

// LexGroup : "(" LexAlternates ")" ;
func (bld *builder) lexGroup(b bsr.BSR) *LexBracket {
	return &LexBracket{
		leftBracket: b.GetTChildI(0),
		Type:        LexGroup,
		Alternates:  bld.lexAlternates(b.GetNTChildI(1)),
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////
// LexOptional : "[" LexAlternates "]" ;
func (bld *builder) lexOptional(b bsr.BSR) *LexBracket {
	return &LexBracket{
		leftBracket: b.GetTChildI(0),
		Type:        LexOptional,
		Alternates:  bld.lexAlternates(b.GetNTChildI(1)),
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////

// LexZeroOrMore : "{" LexAlternates "}" ;
func (bld *builder) lexZeroOrMore(b bsr.BSR) *LexBracket {
	return &LexBracket{
		leftBracket: b.GetTChildI(0),
		Type:        LexZeroOrMore,
		Alternates:  bld.lexAlternates(b.GetNTChildI(1)),
	}
}

// LexOneOrMore : "<" LexAlternates ">" ;
func (bld *builder) lexOneOrMore(b bsr.BSR) *LexBracket {
	return &LexBracket{
		leftBracket: b.GetTChildI(0),
		Type:        LexOneOrMore,
		Alternates:  bld.lexAlternates(b.GetNTChildI(1)),
	}
}

// LexAlternates : RegExp | RegExp "|" LexAlternates ;
func (bld *builder) lexAlternates(b bsr.BSR) []*RegExp {
	alts := []*RegExp{bld.regexp(b.GetNTChildI(0))}
	if b.Alternate() == 1 {
		alts = append(alts, bld.lexAlternates(b.GetNTChildI(2))...)
	}
	return alts
}

func (bld *builder) not(not, strLit *token.Token) *Not {
	return &Not{
		not:    not,
		strLit: strLit,
		Set:    bld.parseStringSet(strLit),
	}
}

// UnicodeClass : "letter" | "upcase" | "lowcase" | "number" | "space" ;
func (bld *builder) unicodeClass(b bsr.BSR) *UnicodeClass {
	class := &UnicodeClass{
		tok: b.GetTChildI(0),
	}
	switch b.Alternate() {
	case 0:
		class.Type = Letter
	case 1:
		class.Type = Upcase
	case 2:
		class.Type = Lowcase
	case 3:
		class.Type = Number
	case 4:
		class.Type = Space
	}
	return class
}

/*** Syntax Rules ***/

// SynOptional : SyntaxAtom "?" ;
/*
Rule : Opt_rule? ;
Rule_alt : Opt_rule?
		/ Other_rule ;
Rule_right : ( Opt_rule
		   / empty )
		   / Other_rule ;
//// i think we need to add an alternates struct to cover the chance it's empty
*/
func (bld *builder) synOptional(b bsr.BSR) *SynOptional {
	/* return SynOptional{
		Tok:  b.GetTChildI(1),
		Expr: bld.atom(b.GetNTChildI(0)),
	} */
	opt := &SynOptional{}
	if b.Alternate() == 0 {
		opt.Expr = bld.atom(b.GetNTChildI(0))
	} // if empty, SynOptional with be returned with empty atom
	return opt
}

// SyntaxAlternate
//     :   SyntaxSymbols
//     |   "empty"
//     ;
func (bld *builder) syntaxAlternate(b bsr.BSR) *SyntaxAlternate {
	alt := &SyntaxAlternate{}
	if b.Alternate() == 0 {
		alt.Symbols = bld.syntaxSymbols(b.GetNTChildI(0))
	} // if alt = empty return alt with empty symbols
	return alt
}

// SyntaxAlternates
//     :   SyntaxAlternate
//     |   SyntaxAlternate "|" UnorderedAlternates
//     |   SyntaxAlternate "/" OrderedAlternates
//     ;
// (boolean is true if an ordered alternate list)
func (bld *builder) syntaxAlternates(b bsr.BSR) ([]*SyntaxAlternate, bool) {
	alts := []*SyntaxAlternate{
		bld.syntaxAlternate(b.GetNTChild(symbols.NT_SyntaxAlternate, 0)),
	}
	switch b.Alternate() {
	case 0:
		return alts, false
	case 1:
		return append(alts, bld.unorderedAlternates(b.GetNTChild(symbols.NT_UnorderedAlternates, 0))...), false
	case 2:
		return append(alts, bld.orderedAlternates(b.GetNTChild(symbols.NT_OrderedAlternates, 0))...), true
	}
	panic("invalid SyntaxAlternates")
}

// UnorderedAlternates
// 		: SyntaxAlternate
// 		| SyntaxAlternate "|" UnorderedAlternates
// 		;
func (bld *builder) unorderedAlternates(b bsr.BSR) []*SyntaxAlternate {
	alts := []*SyntaxAlternate{
		bld.syntaxAlternate(b.GetNTChild(symbols.NT_SyntaxAlternate, 0)),
	}
	if b.Alternate() == 1 {
		alts = append(alts, bld.unorderedAlternates(b.GetNTChild(symbols.NT_UnorderedAlternates, 0))...)
	}
	return alts
}

// OrderedAlternates
// 		: SyntaxAlternate
// 		| SyntaxAlternate "/" OrderedAlternates
// 		;
func (bld *builder) orderedAlternates(b bsr.BSR) []*SyntaxAlternate {
	alts := []*SyntaxAlternate{
		bld.syntaxAlternate(b.GetNTChild(symbols.NT_SyntaxAlternate, 0)),
	}
	if b.Alternate() == 1 {
		alts = append(alts, bld.orderedAlternates(b.GetNTChild(symbols.NT_OrderedAlternates, 0))...)
	}
	return alts
}

// SyntaxRule : nt ":" SyntaxAlternates ";"  ;
func (bld *builder) syntaxRule(b bsr.BSR) brule {
	alts, ord := bld.syntaxAlternates(b.GetNTChild(symbols.NT_SyntaxAlternates, 0))
	return &SyntaxRule{
		Head:       bld.nt(b.GetTChildI(0)),
		Alternates: alts,
		IsOrdered:  ord,
	}
}

// SyntaxSymbol
//     : "&" SyntaxAtom
//     / "!" SyntaxAtom
//	   / SynOptional
//     / SyntaxAtom
//     ;
func (bld *builder) symbol(b bsr.BSR) SyntaxSymbol {
	switch b.Alternate() {
	case 0, 1:
		return &Lookahead{
			Op:   b.GetTChildI(0),
			Expr: bld.atom(b.GetNTChildI(1)),
		}
	case 2:
		return bld.synOptional(b.GetNTChildI(0))
	case 3:
		return bld.atom(b.GetNTChildI(0))
	}
	panic(fmt.Sprintf("invalid alternate %d", b.Alternate()))
}

////////////////////////////////////////////////////////////////////////////////////////////////
// SyntaxAtom : nt | tokid | string_lit ;
func (bld *builder) atom(b bsr.BSR) SyntaxSymbol {
	switch b.Alternate() {
	case 0:
		return bld.nt(b.GetTChildI(0))
	case 1:
		return bld.tokID(b.GetTChildI(0))
	case 2:
		return bld.stringLit(b.GetTChildI(0))
	}
	panic(fmt.Sprintf("invalid alternate %d", b.Alternate()))
}

////////////////////////////////////////////////////////////////////////////////////////////////

// SyntaxSymbols
//     :   SyntaxSymbol
//     |   SyntaxSymbol SyntaxSymbols
//     ;
func (bld *builder) syntaxSymbols(b bsr.BSR) []SyntaxSymbol {
	symbols := []SyntaxSymbol{bld.symbol(b.GetNTChildI(0))}
	//symbols = bld.addOptNode(symbols) //Add SynOptional Nodes
	if b.Alternate() == 1 {
		symbols = append(symbols, bld.syntaxSymbols(b.GetNTChildI(1))...)
	}
	return symbols
}

////////// this function will keep it from compiling right now
/* func (bld *builder) addOptNode(symbols []SyntaxSymbol) []SyntaxSymbol {
	//if most recent symbol is synOptional, then append(`/empty`)
	if symbols[(len(symbols)-1)].ID() == "?" {
		temp := SynOptional{}
		symbols = append(symbols, temp)
	}
	return symbols
} */

/*** Shared ***/

// NT : nt  ;
func (bld *builder) nt(tok *token.Token) *NT {
	return &NT{
		tok: tok,
	}
}

func (bld *builder) stringLit(tok *token.Token) *StringLit {
	return &StringLit{
		tok: tok,
	}
}

// TokID : id ;
func (bld *builder) tokID(tok *token.Token) *TokID {
	return &TokID{
		tok: tok,
	}
}

/*** Utils ***/

func (bld *builder) addLexRule(r *LexRule, gogll *GoGLL) {
	if nil != gogll.GetLexRule(r.ID()) {
		bld.fail(fmt.Errorf("duplicate lex rule %s", r.ID()), r.Lext())
	}
	gogll.LexRules = append(gogll.LexRules, r)
}

func (bld *builder) addSyntaxRule(r *SyntaxRule, gogll *GoGLL) {
	if nil != gogll.GetSyntaxRule(r.ID()) {
		bld.fail(fmt.Errorf("duplicate syntax rule %s", r.ID()), r.Lext())
	}
	gogll.SyntaxRules = append(gogll.SyntaxRules, r)
}

// parse the string set from tokens any or not
func (bld *builder) parseStringSet(strLit *token.Token) *runeset.RuneSet {
	rs := runeset.New()
	for i := 1; i < len(strLit.Literal())-1; i++ {
		if strLit.Literal()[i] == '\\' {
			i++
			switch strLit.Literal()[i] {
			case '\\':
				rs.Add('\\')
			case '"':
				rs.Add('"')
			case 'r':
				rs.Add('\r')
			case 'n':
				rs.Add('\n')
			case 't':
				rs.Add('\t')
			default:
				bld.fail(fmt.Errorf("invalid escape char"), strLit.Lext())
			}
		} else {
			rs.Add(strLit.Literal()[i])
		}
	}
	return rs
}

/*** Errors ***/

// i is the position of the failure in input slice of runes
func (bld *builder) fail(err error, i int) {
	ln, col := bld.lex.GetLineColumn(i)
	fmt.Printf("AST Error: %s at line %d col %d\n", err, ln, col)
	os.Exit(1)
}
