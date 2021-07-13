//  Copyright 2020 Marius Ackerman
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

/*
Package parser controls the generation of all Rust LR(1) parser-related code.
*/
package parser

import (
	"github.com/bruceiv/pegll/cfg"
	"github.com/bruceiv/pegll/lr1/action"
	"github.com/bruceiv/pegll/lr1/basicprod"
	"github.com/bruceiv/pegll/lr1/states"
)

func Gen(pkg string, bprods []*basicprod.Production, states *states.States, actions action.Actions) {
	genActionTable(pkg, cfg.BaseDir, bprods, states, actions)
	genGotoTable(cfg.BaseDir, states)
	genParser(pkg, bprods, states)
	genProductionsTable(pkg, bprods, states)

	return
}
