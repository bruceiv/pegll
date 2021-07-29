
// Package token is generated by GoGLL. Do not edit
package token

import(
    "fmt"
)

// Token is returned by the lexer for every scanned lexical token
type Token struct {
    typ        Type
    lext, rext int
    input      []rune
}

/*
New returns a new token.
lext is the left extent and rext the right extent of the token in the input.
input is the input slice scanned by the lexer.
*/
func New(t Type, lext, rext int, input []rune) *Token {
    return &Token{
        typ:   t,
        lext:  lext,
        rext:  rext,
        input: input,
    }
}

// GetLineColumn returns the line and column of the left extent of t
func (t *Token) GetLineColumn() (line, col int) {
    line, col = 1, 1
    for j := 0; j < t.lext; j++ {
        switch t.input[j] {
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

// GetInput returns the input from which t was parsed.
func (t *Token) GetInput() []rune {
    return t.input
}

// Lext returns the left extent of t
func (t *Token) Lext() int {
    return t.lext
}

// Literal returs the literal runes of t scanned by the lexer
func (t *Token) Literal() []rune {
    return t.input[t.lext:t.rext]
}

// LiteralString returns string(t.Literal())
func (t *Token) LiteralString() string {
    return string(t.Literal())
}

// Rext returns the right extent of t in the input
func (t *Token) Rext() int {
    return t.rext
}

func (t *Token) String() string {
    return fmt.Sprintf("%s (%d,%d) %s",
        t.TypeID(), t.lext, t.rext, t.LiteralString())
}

// Suppress returns true iff t is suppressed by the lexer
func (t *Token) Suppress() bool {
	return Suppress[t.typ]
}

// Type returns the token Type of t
func (t *Token) Type() Type {
    return t.typ
}

// TypeID returns the token Type ID of t. 
// This may be different from the literal of token t.
func (t *Token) TypeID() string {
    return t.Type().ID()
}

// Type is the token type
type Type int

func (t Type) String() string {
    return TypeToString[t]
}

// ID returns the token type ID of token Type t
func (t Type) ID() string {
    return TypeToID[t]
}


const(
    Error  Type = iota  // Error 
    EOF  // $ 
    T_0  // + 
    T_1  // , 
    T_2  // - 
    T_3  // . 
    T_4  // 0 
    T_5  // : 
    T_6  // [ 
    T_7  // ] 
    T_8  // bSlash 
    T_9  // block_comment 
    T_10  // char 
    T_11  // dQuote 
    T_12  // eE 
    T_13  // esc 
    T_14  // escCharSpace 
    T_15  // false 
    T_16  // hex 
    T_17  // line_comment 
    T_18  // nonZero 
    T_19  // null 
    T_20  // num 
    T_21  // true 
    T_22  // u 
    T_23  // { 
    T_24  // } 
)

var TypeToString = []string{ 
    "Error",
    "EOF",
    "T_0",
    "T_1",
    "T_2",
    "T_3",
    "T_4",
    "T_5",
    "T_6",
    "T_7",
    "T_8",
    "T_9",
    "T_10",
    "T_11",
    "T_12",
    "T_13",
    "T_14",
    "T_15",
    "T_16",
    "T_17",
    "T_18",
    "T_19",
    "T_20",
    "T_21",
    "T_22",
    "T_23",
    "T_24",
}

var StringToType = map[string] Type { 
    "Error" : Error, 
    "EOF" : EOF, 
    "T_0" : T_0, 
    "T_1" : T_1, 
    "T_2" : T_2, 
    "T_3" : T_3, 
    "T_4" : T_4, 
    "T_5" : T_5, 
    "T_6" : T_6, 
    "T_7" : T_7, 
    "T_8" : T_8, 
    "T_9" : T_9, 
    "T_10" : T_10, 
    "T_11" : T_11, 
    "T_12" : T_12, 
    "T_13" : T_13, 
    "T_14" : T_14, 
    "T_15" : T_15, 
    "T_16" : T_16, 
    "T_17" : T_17, 
    "T_18" : T_18, 
    "T_19" : T_19, 
    "T_20" : T_20, 
    "T_21" : T_21, 
    "T_22" : T_22, 
    "T_23" : T_23, 
    "T_24" : T_24, 
}

var TypeToID = []string { 
    "Error", 
    "$", 
    "+", 
    ",", 
    "-", 
    ".", 
    "0", 
    ":", 
    "[", 
    "]", 
    "bSlash", 
    "block_comment", 
    "char", 
    "dQuote", 
    "eE", 
    "esc", 
    "escCharSpace", 
    "false", 
    "hex", 
    "line_comment", 
    "nonZero", 
    "null", 
    "num", 
    "true", 
    "u", 
    "{", 
    "}", 
}

var Suppress = []bool { 
    false, 
    false, 
    false, 
    false, 
    false, 
    false, 
    false, 
    false, 
    false, 
    false, 
    false, 
    true, 
    false, 
    false, 
    false, 
    false, 
    false, 
    false, 
    false, 
    true, 
    false, 
    false, 
    false, 
    false, 
    false, 
    false, 
    false, 
}

