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

import "github.com/goccmack/gogll/token"

// The syntax part of the AST
type SynOptional struct { //Where do we get it to connect to the '?' ????  --> similar to Lext function in lex.go??
	tok *token.Token //I think contains the ?
	//Possibly needs: strLit *token.Token -- might contain the 'rule?' structure
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

func (*SynOptional) isSyntaxSymbol() {} // I believe this gets it included in the list

func (*NT) isSyntaxSymbol() {}

// Terminals
func (*TokID) isSyntaxSymbol()     {}
func (*StringLit) isSyntaxSymbol() {}

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
