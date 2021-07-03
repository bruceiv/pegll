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

package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime/pprof"

	"github.com/bruceiv/pegll/ast"
	"github.com/bruceiv/pegll/cfg"
	"github.com/bruceiv/pegll/frstflw"
	genff "github.com/bruceiv/pegll/gen/firstfollow"
	gengogll "github.com/bruceiv/pegll/gen/golang/gll"
	gengolexer "github.com/bruceiv/pegll/gen/golang/lexer"
	gengolr1 "github.com/bruceiv/pegll/gen/golang/lr1"
	gengotoken "github.com/bruceiv/pegll/gen/golang/token"
	"github.com/bruceiv/pegll/gen/lexfsa"
	genrustgll "github.com/bruceiv/pegll/gen/rust/gll"
	genrustlexer "github.com/bruceiv/pegll/gen/rust/lexer"
	genrustlr1 "github.com/bruceiv/pegll/gen/rust/lr1"
	genrusttoken "github.com/bruceiv/pegll/gen/rust/token"
	"github.com/bruceiv/pegll/gen/slots"
	gensymbols "github.com/bruceiv/pegll/gen/symbols"
	"github.com/bruceiv/pegll/gslot"
	"github.com/bruceiv/pegll/lex/items"
	"github.com/bruceiv/pegll/lexer"
	"github.com/bruceiv/pegll/lr1"
	"github.com/bruceiv/pegll/parser"
	"github.com/bruceiv/pegll/sc"
	"github.com/bruceiv/pegll/symbols"
)

func main() {
	cfg.GetParams()
	if *cfg.CPUProfile {
		f, err := os.Create("cpu.prof")
		if err != nil {
			fmt.Println("could not create CPU profile: ", err)
			os.Exit(1)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			fmt.Println("could not start CPU profile: ", err)
			os.Exit(1)
		}
		defer pprof.StopCPUProfile()
	}
	lex := lexer.NewFile(cfg.SrcFile)
	bsrSet, errs := parser.Parse(lex)
	if errs != nil {
		parseErrors(errs)
	}

	if bsrSet.IsAmbiguous() {
		fmt.Println("Error: Ambiguous parse forest")
		bsrSet.ReportAmbiguous()
		os.Exit(1)
	}

	g := ast.Build(bsrSet.GetRoot(), lex)
	sc.Go(g, lex)
	symbols.Init(g)

	ff := frstflw.New(g)
	gs := gslot.New(g, ff)

	lexSets := items.New(g)
	if cfg.Verbose {
		gensymbols.Gen(g)
		genff.Gen(g, ff)
		slots.Gen(gs)
		lexfsa.Gen(filepath.Join(cfg.BaseDir, "lexfsa.txt"), lexSets)
	}

	switch {
	case *cfg.Go:
		gengolexer.Gen(g, lexSets)
		gengotoken.Gen(g)
		if len(g.SyntaxRules) > 0 {
			if *cfg.GLL {
				gengogll.Gen(g, gs, ff)
			} else {
				bprods, states, actions := lr1.Gen(g)
				gengolr1.Gen(g.Package.GetString(), bprods, states, actions)
			}
		}
	case *cfg.Rust:
		genrusttoken.Gen(filepath.Join(cfg.BaseDir, "src", "token", "mod.rs"))
		genrustlexer.Gen(path.Join(cfg.BaseDir, "src", "lexer", "mod.rs"), g, lexSets)
		if len(g.SyntaxRules) > 0 {
			if *cfg.GLL {
				genrustgll.Gen(path.Join(cfg.BaseDir, "src", "parser"), g, gs, ff)
			} else {
				bprods, states, actions := lr1.Gen(g)
				genrustlr1.Gen(g.Package.GetString(), bprods, states, actions)
			}
		}
	}

}

func fail(err error) {
	fmt.Printf("Error: %s\n", err)
	os.Exit(1)
}

func parseErrors(errs []*parser.Error) {
	fmt.Println("Parse Errors:")
	ln := errs[0].Line
	for _, err := range errs {
		if err.Line == ln {
			fmt.Println(err)
		}
	}
	os.Exit(1)
}
