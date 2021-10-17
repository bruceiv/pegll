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

package ast

import (
	"github.com/bruceiv/pegll/token"
)

// A syntax rule
type SyntaxRule struct {
	Head       *NT
	Alternates []*SyntaxAlternate
	IsOrdered  bool
}

// An alternate expression
type SyntaxAlternate struct {
	Symbols []SyntaxSymbol
}

// A syntax suffix operator
type SyntaxSuffix struct {
	// expression made optional
	Expr SyntaxSymbol
	// token for operator
	Tok *token.Token
	// signifies the type of suffix
	// 0: optional (?)
	// 1: repeat zero or more times (*)
	// 2: repeat one or more times (+)
	Type int
}

// A syntax symbol
type SyntaxSymbol interface {
	isSyntaxSymbol()
	// Lext returns the left extent of SyntaxSymbol in the input string
	Lext() int
	// The ID of the symbol
	// which is the literal string of a LexRule, SyntaxRule or StringLit.
	ID() string
	// The string of the symbol
	String() string
}

// A lookahead expression
type Lookahead struct {
	// operator for expression
	Op *token.Token
	// operator subexpression. (should not be lookahead)
	Expr SyntaxSymbol
}

// non-terminals
func (*NT) isSyntaxSymbol()          {}
func (*Lookahead) isSyntaxSymbol()   {}
func (SyntaxSuffix) isSyntaxSymbol() {}

// terminals
func (*TokID) isSyntaxSymbol()     {}
func (*StringLit) isSyntaxSymbol() {}

/* Syntax Rules */
// ID returns the head of rule r
func (r *SyntaxRule) ID() string {
	return r.Head.ID()
}
func (r *SyntaxRule) Lext() int {
	return r.Head.Lext()
}

// true if always matches; false if unable to guarantee always matches
func (r *SyntaxRule) AlwaysMatches() bool {
	return r.Alternates[len(r.Alternates)-1].Empty()
}
func (a *SyntaxAlternate) Empty() bool {
	return len(a.Symbols) == 0
}
func (a *SyntaxAlternate) GetSymbols() []string {
	symbols := make([]string, len(a.Symbols))
	for i, s := range a.Symbols {
		symbols[i] = s.ID()
	}
	return symbols
}

/* Syntax Suffix */
func (opt *SyntaxSuffix) ID() string {
	return opt.Expr.ID() + opt.Tok.LiteralString()
}
func (opt *SyntaxSuffix) Lext() int {
	return opt.Expr.Lext()
}
func (opt *SyntaxSuffix) String() string {
	return opt.Expr.String() + opt.Tok.LiteralString()
}

/* Lookahead */
func (e *Lookahead) Lext() int {
	return e.Op.Lext()
}
func (e *Lookahead) ID() string {
	return e.Op.LiteralString() + e.Expr.ID()
}
func (e *Lookahead) String() string {
	return e.Op.LiteralString() + e.Expr.String()
}

// true for positive (&) lookahead, false for negative (!) lookahead
func (e *Lookahead) Positive() bool {
	return e.Op.LiteralString() == "&"
}
