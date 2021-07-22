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
/* build.go modifications (these are backwards sorry)
 * line 489
 	- don't think we need to add SynOptional into the
	syntaxSymbols function because it's being built in SyntaxSymbols
 * line 350
	- changed the struct so it has syntax symbols so we can traverse
	through it and see if there is a "?" - see Moss's answer in Teams
 * line 364
	- the way that it is built in SyntaxSymbols we need to traverse
	through the symbols, find the question mark, and then see if the
	token exists - since we already have it built to return empty in
	the synOptional function, we're simply checking for the existence of
	the token and, if it's there, adding it to the Expr part of the struct


	we need to figure out how to add NT_SynOptional
*/

package ast

import (
	"github.com/bruceiv/pegll/parser/bsr"
	"github.com/bruceiv/pegll/token"
)

// The syntax part of the AST
///////////////////////////// trying something different - I don't think
//////////////////////////// we need the token - I think we need the NT
/* type SynOptional struct { //Where do we get it to connect to the '?' ????  --> similar to Lext function in lex.go??
	Tok  *token.Token //I think contains the ?
	Expr SyntaxSymbol //Contains the rule that is being made optional (we think)
} */

type SynOptional struct {
	// expression made optional
	Expr     SyntaxSymbol
	ExprNode bsr.BSR
	Tok      *token.Token
}

// Line 126 in build.go --> do to we need to add the symbol to a set? Do we need to do this????

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

func (SynOptional) isSyntaxSymbol() {}

func (opt *SynOptional) ID() string {
	return opt.Expr.ID()
	// i think we only need the ID and not the literal string
	//return opt.Expr.ID()
	//return "?"
}
func (opt *SynOptional) Lext() int {
	// return opt.Tok.Lext()
	return opt.Expr.Lext()
}
func (opt *SynOptional) String() string {
	//return opt.Expr.String() + opt.Tok.LiteralString()
	return opt.Expr.ID()
}

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

// true if always matches; false if unable to guarantee always matches
func (r *SyntaxRule) AlwaysMatches() bool {
	return r.Alternates[len(r.Alternates)-1].Empty()
}
