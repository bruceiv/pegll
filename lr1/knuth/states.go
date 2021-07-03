package knuth

import (
	"github.com/bruceiv/pegll/lr1/first"
	"github.com/bruceiv/pegll/lr1/items"
	"github.com/bruceiv/pegll/lr1/states"
	"github.com/bruceiv/pegll/lr1/symbolsuccessors"
)

func States(symbols []string, lr0items *items.Items, first *first.First) *states.States {
	s := &states.States{
		List: make([]*states.State, 0, 64),
	}
	s.NewState0(symbols, lr0items, first, states.NewConfigGroup(lr0items.List[0], "$"))
	symSuccessors := symbolsuccessors.NewSymbolSuccessors()
	for si := 0; si < len(s.List); si++ {
		st_trans := make([]states.Transition, 0, 4)
		for _, trans := range s.List[si].Next().List() {
			newState := true
			for _, snum := range symSuccessors[trans.Sym] {
				if s.List[snum].Equal(trans.State) {
					st_trans = append(st_trans, states.Transition{trans.Sym, s.List[snum]})
					newState = false
				}
			}
			if newState {
				st_trans = append(st_trans, trans)
				s.List = append(s.List, trans.State)
				trans.State.Number = len(s.List) - 1
				symSuccessors[trans.Sym] = append(symSuccessors[trans.Sym], trans.State.Number)
			}
		}
		s.List[si].Transitions = states.NewTransitionsList(st_trans, symbols)
	}
	return s
}
