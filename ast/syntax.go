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

// The syntax part of the AST
type SynOptional struct { //Where do we get it to connect to the '?' ????  --> similar to Lext function in lex.go??
	Tok  *token.Token //I think contains the ?
	Expr SyntaxSymbol //Contains the rule that is being made optional (we think)
}

type SyntaxAlternate struct {
	Symbols []SyntaxSymbol
}

type SyntaxRule struct {
	Head       *NT
	Alternates []*SyntaxAlternate
	IsOrdered  bool
}

type SyntaxSymbol interface {
	isSyntaxSymbol()
	// Lext returns the left extent of SyntaxSymbol in the input string
	Lext() int

	// The ID of the symbol, which is the literal string of a LexRule, SyntaxRule
	// or StringLit.
	ID() string

	String() string
}

// A lookahead expression
type Lookahead struct {
	// operator for expression
	Op *token.Token
	// operator subexpression. (should not be lookahead)
	Expr SyntaxSymbol
}

func (*NT) isSyntaxSymbol()        {}
func (*Lookahead) isSyntaxSymbol() {}

////////////////////////////////////////////////////////////////////////////////////////////////
func (*SynOptional) isSyntaxSymbol() {}


func (opt *SynOptional) Lext() int {
	return opt.Tok.Lext()
}
func (opt *SynOptional) ID() string {
	return opt.Tok.LiteralString() + opt.Expr.ID()
}
func (opt *SynOptional) String() string {
	return opt.Tok.LiteralString() + opt.Expr.String()
}

////////////////////////////////////////////////////////////////////////////////////////////////
// Terminals
func (*TokID) isSyntaxSymbol()     {}
func (*StringLit) isSyntaxSymbol() {}

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

func (a *SyntaxAlternate) GetSymbols() []string {
	symbols := make([]string, len(a.Symbols))
	for i, s := range a.Symbols {
		symbols[i] = s.ID()
	}
	return symbols
}

func (a *SyntaxAlternate) Empty() bool {
	return len(a.Symbols) == 0
}

// ID returns the head of rule r
func (r *SyntaxRule) ID() string {
	return r.Head.ID()
}

func (r *SyntaxRule) Lext() int {
	return r.Head.Lext()
}
