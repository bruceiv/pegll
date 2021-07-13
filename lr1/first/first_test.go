package first

import (
	"testing"

	"github.com/bruceiv/pegll/ast"
	"github.com/bruceiv/pegll/ast/rewrite"
	"github.com/bruceiv/pegll/frontend/lexer"
	"github.com/bruceiv/pegll/frontend/parser"
	"github.com/bruceiv/pegll/parser/symbols"
)

const src = `
A : B | "a";
B : ["b"] ;
C : B "c" ;
`

func TestFirst1(t *testing.T) {
	gram, err := parser.NewParser().Parse(lexer.NewLexer([]byte(src)))
	if err != nil {
		panic(err)
	}
	g := gram.(*ast.Grammar)
	basicProds := rewrite.BasicProds(g.SyntaxPart.ProdList)
	sym, serr := symbols.NewSymbols(g.LexPart, basicProds)
	if serr != nil {
		for _, err := range serr {
			t.Logf("%s", err)
		}
	}
	first := New(sym, basicProds.List())
	exp := FirstSet{"b", "x", "y"}
	f := first.FirstString([]string{"B"}, "x", "y")
	if !f.Equal(exp) {
		t.Errorf("Expected %s, got %s", exp, f)
	}

	exp = FirstSet{"x", "y"}
	f = first.FirstString(nil, "x", "y")
	if !f.Equal(exp) {
		t.Errorf("Expected %s, got %s", exp, f)
	}
}
