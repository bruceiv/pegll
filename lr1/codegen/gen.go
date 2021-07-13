package codegen

import (
	"github.com/bruceiv/pegll/lr1/action"
	"github.com/bruceiv/pegll/lr1/basicprod"
	"github.com/bruceiv/pegll/lr1/codegen/golang"
	"github.com/bruceiv/pegll/lr1/states"
)

func Gen(pkg string, prods []*basicprod.Production, states *states.States, actions action.Actions) {
	golang.Gen(pkg, prods, states, actions)
}
