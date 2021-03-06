
// Package lexer is generated by GoGLL. Do not edit.
package lexer

import (
	// "fmt"
	"io/ioutil"
	"strings"
	"unicode"

	"axbc/token"
)

type state int

const nullState state = -1
const noMatch int = -1

// TokenSet represents a set of tokens which may match at an implied input
// index. The slice can be considered a map from the token ID to the length
// of the possible match at that ID, or -1 for no such match.
//
// Unlike a traditional CFG lexer, the PEGLL lexer returns a slice of
// *all* possibly-matching tokens, rather than picking the single maximal
// munch. If a single token matches at multiple lengths, the longest such
// is taken.
type TokenSet []int

// Lexer contains both the input slice of runes and the slice of possible
// tokens parsed from the input
type Lexer struct {
	// I is the input slice of runes
	I []rune

	// Tokens is the slice of possible tokens constructed by the lexer from I
	// A nil TokenSet means that the lexer has not been run at that position 
	// yet
	tokens []*TokenSet
}

/*
NewFile constructs a Lexer created from the input file, fname.

If the input file is a markdown file NewFile process treats all text outside
code blocks as whitespace. All text inside code blocks are treated as input text.

If the input file is a normal text file NewFile treats all text in the inputfile
as input text.
*/
func NewFile(fname string) *Lexer {
	buf, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	input := []rune(string(buf))
	if strings.HasSuffix(fname, ".md") {
		loadMd(input)
	}
	return New(input)
}

func loadMd(input []rune) {
	i := 0
	text := true
	for i < len(input) {
		if i <= len(input)-3 && input[i] == '`' && input[i+1] == '`' && input[i+2] == '`' {
			text = !text
			for j := 0; j < 3; j++ {
				input[i+j] = ' '
			}
			i += 3
		}
		if i < len(input) {
			if text {
				if input[i] == '\n' {
					input[i] = '\n'
				} else {
					input[i] = ' '
				}
			}
			i += 1
		}
	}
}

/*
New constructs a Lexer from a slice of runes.

All contents of the input slice are treated as input text.
*/
func New(input []rune) *Lexer {
	// initialize data structure with nil tokens
	tLen := len(input) + 1
	lex := &Lexer{
		I:      input,
		tokens: make([]*TokenSet, tLen, tLen),
	}
	// set up non-nil EOF token
	lex.tokens[tLen-1] = initEmptyTokenSet()
	(*lex.tokens[tLen-1])[token.EOF] = 0
	return lex
}

// returns token set at location i
func (l *Lexer) Tokens(i int) *TokenSet {
	i = l.skipWhitespace(i)
	// lex if needed
	if l.tokens[i] == nil {
		l.scan(i)
		// TODO also consider suppression
	}
	// return tokens
	return l.tokens[i]
}

// returns i updated to skip whitespace
func (l *Lexer) skipWhitespace(i int) int {
	for i < len(l.I) && unicode.IsSpace(l.I[i]) {
		i++
	}
	return i
}

// Returns empty token set
func initEmptyTokenSet() *TokenSet {
	nTokens := len(token.TypeToID)
	tokens := make([]int, nTokens, nTokens)
	for j, _ := range tokens {
		tokens[j] = noMatch
	}
	return (*TokenSet)(&tokens)
}

func (l *Lexer) scan(i int) {
	// set up empty token array
	l.tokens[i] = initEmptyTokenSet()

	// loop until no further tokens
	s := state(0) // current state
	rext := i     // current character
	for s != nullState {
		// check for found tokens
		for _, typ := range accept[s] {
			(*l.tokens[i])[typ] = rext
		}
		// check for end-of-string
		if rext >= len(l.I) {
			return
		}
		// advance to next state
		s = nextState[s](l.I[rext])
		rext++
	}
}

func escape(r rune) string {
	switch r {
	case '"':
		return "\""
	case '\\':
		return "\\\\"
	case '\r':
		return "\\r"
	case '\n':
		return "\\n"
	case '\t':
		return "\\t"
	}
	return string(r)
}

// RightExtent gets the right extent of the token tok at index i, -1 for none such
func (l *Lexer) RightExtent(tok token.Type, i int) int {
	i = l.skipWhitespace(i)
	if l.tokens[i] == nil {
		return -1
	}
	return (*l.tokens[i])[tok]
}

// GetLineColumn returns the line and column of rune[i] in the input
func (l *Lexer) GetLineColumn(i int) (line, col int) {
	line, col = 1, 1
	for j := 0; j < i; j++ {
		switch l.I[j] {
		case '\n':
			line++
			col = 1
		case '\t':
			col += 4
		default:
			col++
		}
	}
	return
}

// GetString returns the input string between the given two indices,
// inclusive, or empty string if range empty
func (l *Lexer) GetString(lext, rext int) string {
	if rext < lext {
		return ""
	}
	return string(l.I[lext:rext])
}

func any(r rune, set []rune) bool {
	for _, r1 := range set {
		if r == r1 {
			return true
		}
	}
	return false
}

func not(r rune, set []rune) bool {
	for _, r1 := range set {
		if r == r1 {
			return false
		}
	}
	return true
}

var accept = [][]token.Type{ 
	{  token.T_0,  }, 
	{  token.T_0,  }, 
	{  token.T_2,  }, 
	{  token.T_0,  }, 
	{  token.T_1,  }, 
}

var nextState = []func(r rune) state{ 
	// Set0
	func(r rune) state {
		switch { 
		case r == 'a':
			return 1 
		case r == 'c':
			return 2 
		}
		return nullState
	}, 
	// Set1
	func(r rune) state {
		switch { 
		case r == 'a':
			return 3 
		case r == 'b':
			return 4 
		}
		return nullState
	}, 
	// Set2
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
	// Set3
	func(r rune) state {
		switch { 
		case r == 'a':
			return 3 
		}
		return nullState
	}, 
	// Set4
	func(r rune) state {
		switch { 
		}
		return nullState
	}, 
}
