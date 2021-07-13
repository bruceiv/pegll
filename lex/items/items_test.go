package items

import (
	"testing"

	"github.com/bruceiv/pegll/ast"
	"github.com/bruceiv/pegll/lexer"
	"github.com/bruceiv/pegll/parser"
	"github.com/bruceiv/pegll/parser/bsr"
)

const src = `package "names"
qualifiedName : letter {letter|number|'_'} <'.' <letter|number|'_'>> ;
`

func Test1(t *testing.T) {
	lex := lexer.New([]rune(src))
	parser.Parse(lex)
	g := ast.Build(bsr.GetRoot(), lex)

	New(g)
}
