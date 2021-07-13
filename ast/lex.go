/*
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

// The lexer part of the AST

package ast

import (
	"bytes"
	"fmt"

	"github.com/bruceiv/pegll/token"
	"github.com/bruceiv/pegll/util/runeset"
)

// This file contains the AST components for lexical rules

// TriState has values: {Undefined, False, True}
type TriState int

const (
	// Undefined is a TriState value
	Undefined TriState = iota
	// False is a TriState value
	False
	// True is a TriState value
	True
)

type Any struct {
	tok *token.Token
}

type AnyOf struct {
	any    *token.Token
	strLit *token.Token
	Set    *runeset.RuneSet
}

type CharLiteral struct {
	tok     *token.Token
	Literal []rune
}

type LexBracket struct {
	leftBracket *token.Token
	Type        BracketType
	Alternates  []*RegExp
}

type BracketType int

const (
	LexGroup BracketType = iota
	LexOptional
	LexZeroOrMore
	LexOneOrMore
)

type LexBase interface {
	isLexBase()
	LexSymbol
	Equal(LexBase) bool
}

func (*Any) isLexBase()          {}
func (*AnyOf) isLexBase()        {}
func (*CharLiteral) isLexBase()  {}
func (*Not) isLexBase()          {}
func (*UnicodeClass) isLexBase() {}

type LexRule struct {
	Suppress bool
	TokID    *TokID
	RegExp   *RegExp
}

type LexSymbol interface {
	isLexSymbol()
	Lext() int
	String() string
}

func (*Any) isLexSymbol()          {}
func (*AnyOf) isLexSymbol()        {}
func (*CharLiteral) isLexSymbol()  {}
func (*LexBracket) isLexSymbol()   {}
func (*Not) isLexSymbol()          {}
func (*UnicodeClass) isLexSymbol() {}

type Not struct {
	not    *token.Token
	strLit *token.Token
	Set    *runeset.RuneSet
}

type RegExp struct {
	Symbols []LexSymbol
}

type StringLit struct {
	tok *token.Token
}

type UnicodeClass struct {
	tok  *token.Token
	Type UnicodeClassType
}

type UnicodeClassType int

const (
	Letter UnicodeClassType = iota
	Upcase
	Lowcase
	Number
	Space
)

/*** Methods ***/

func (*Any) Equal(other LexBase) bool {
	if other == nil {
		return false
	}
	_, ok := other.(*Any)
	return ok
}

func (a *Any) Lext() int {
	return a.tok.Lext()
}

func (ao *AnyOf) Equal(other LexBase) bool {
	if other == nil {
		return false
	}
	ao1, ok := other.(*AnyOf)
	if !ok {
		return false
	}
	return ao.Set.Equal(ao1.Set)
}

func (a *AnyOf) Lext() int {
	return a.any.Lext()
}

func NewCharLiteral(tok *token.Token, literal []rune) *CharLiteral {
	return &CharLiteral{
		tok:     tok,
		Literal: literal,
	}
}

func (c *CharLiteral) Char() rune {
	if c.Literal[1] == '\\' {
		switch c.tok.Literal()[2] {
		case '\'':
			return '\''
		case '\\':
			return '\\'
		case 't':
			return '\t'
		case 'n':
			return '\n'
		case 'r':
			return '\r'
		default:
			panic(fmt.Sprintf("invalid '%c'", c.Literal[2]))
		}
	} else {
		return c.Literal[1]
	}
}

func (c *CharLiteral) Equal(other LexBase) bool {
	if other == nil {
		return false
	}
	c1, ok := other.(*CharLiteral)
	if !ok {
		return false
	}
	// fmt.Printf("'%c'.Equal('%c') = %t\n", c.Char(), c1.Char(), c.Char() == c1.Char())
	return c.Char() == c1.Char()
}

func (c *CharLiteral) Lext() int {
	return c.tok.Lext()
}

func (l *LexBracket) LeftBracket() string {
	switch l.Type {
	case LexGroup:
		return "("
	case LexOptional:
		return "["
	case LexZeroOrMore:
		return "{"
	case LexOneOrMore:
		return "<"
	}
	panic("invalid")
}

func (l *LexBracket) RightBracket() string {
	switch l.Type {
	case LexGroup:
		return ")"
	case LexOptional:
		return "]"
	case LexZeroOrMore:
		return "}"
	case LexOneOrMore:
		return ">"
	}
	panic("invalid")
}

// Returns the id of the lex rule
func (l *LexRule) ID() string {
	return l.TokID.ID()
}

func (l *LexRule) Lext() int {
	return l.TokID.Lext()
}

func (l *LexRule) String() string {
	return fmt.Sprintf("%s : %s ;", l.ID(), l.RegExp)
}

func (b *LexBracket) Lext() int {
	return b.leftBracket.Lext()
}

func (n *Not) Equal(other LexBase) bool {
	if other == nil {
		return false
	}
	n1, ok := other.(*Not)
	if !ok {
		return false
	}
	return n.Set.Equal(n1.Set)
}

func (n *Not) Lext() int {
	return n.not.Lext()
}

func (re *RegExp) String() string {
	w := new(bytes.Buffer)
	for _, symbol := range re.Symbols {
		fmt.Fprint(w, symbol)
	}
	return w.String()
}

func (u *UnicodeClass) Equal(other LexBase) bool {
	if other == nil {
		return false
	}
	u1, ok := other.(*UnicodeClass)
	if !ok {
		return false
	}
	return u.Type == u1.Type
}

func (u *UnicodeClass) Lext() int {
	return u.Lext()
}

func (*Any) String() string {
	return "."
}

func (a *AnyOf) String() string {
	return fmt.Sprintf("any %s", string(a.strLit.Literal()))
}

func (c *CharLiteral) String() string {
	return string(c.Literal)
}

func (lb *LexBracket) String() string {
	w := new(bytes.Buffer)
	fmt.Fprint(w, lb.LeftBracket())
	for i, alt := range lb.Alternates {
		if i > 0 {
			fmt.Fprint(w, " | ")
		}
		fmt.Fprint(w, alt)
	}
	fmt.Fprint(w, lb.RightBracket())
	return w.String()
}

func (n *Not) String() string {
	return fmt.Sprintf("not %s", string(n.strLit.Literal()))
}

func (sl *StringLit) ID() string {
	lit := sl.tok.Literal()
	return string(lit[1 : len(lit)-1])
}

func (sl *StringLit) Literal() []rune {
	return sl.tok.Literal()
}

func (sl *StringLit) Value() []rune {
	return sl.tok.Literal()[1 : len(sl.tok.Literal())-1]
}

func (u *UnicodeClass) String() string {
	return string(u.tok.Literal())
}

/*** Utils ***/

// StringLitToTokID returns a dummy TokID with ID = id
func StringLitToTokID(id *StringLit) *TokID {
	return &TokID{
		token.New(token.StringToType["tokid"],
			id.tok.Lext()+1, id.tok.Rext()-1, id.tok.GetInput()),
	}
}

// CharLitFromStringLit returns a dummy CharLiteral with Literal sl.Literal[i]
func CharLitFromStringLit(sl *StringLit, i int) *CharLiteral {
	return NewCharLiteral(token.New(token.StringToType["char_lit"], sl.Lext()+i, sl.Lext()+i+1, sl.tok.GetInput()),
		[]rune{'\'', sl.Value()[i], '\''})
}
