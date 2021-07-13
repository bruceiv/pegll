package states

import (
	"bytes"
	"fmt"

	lr0items "github.com/bruceiv/pegll/lr1/items"
)

type ConfigGroupSet struct {
	list    []*ConfigGroup
	contain map[string]*ConfigGroup
}

func NewConfigGroupSet() *ConfigGroupSet {
	return &ConfigGroupSet{
		list:    make([]*ConfigGroup, 0, 4),
		contain: make(map[string]*ConfigGroup),
	}
}

func (this *ConfigGroupSet) Add(c ...*ConfigGroup) *ConfigGroupSet {
	for _, cfgrp := range c {
		if cg, exist := this.contain[cfgrp.hashKey]; !exist {
			this.list = append(this.list, cfgrp)
			this.contain[cfgrp.hashKey] = cfgrp
		} else {
			cg.AddContext(cfgrp.ContextSet.List...)
		}
	}
	return this
}

func (this *ConfigGroupSet) Clone() *ConfigGroupSet {
	clone := &ConfigGroupSet{
		list:    make([]*ConfigGroup, 0, len(this.list)),
		contain: make(map[string]*ConfigGroup),
	}
	for _, cfg := range this.list {
		clone.Add(NewConfigGroup(cfg.Item, cfg.ContextSet.Clone().List...))
	}
	return clone
}

func (this *ConfigGroupSet) CloneFromCore() *ConfigGroupSet {
	clone := &ConfigGroupSet{
		list:    make([]*ConfigGroup, 0, len(this.list)),
		contain: make(map[string]*ConfigGroup),
	}
	for _, cfg := range this.list {
		clone.Add(NewConfigGroup(cfg.Item))
	}
	return clone
}

/*
Return true iff this contains a config group with the same core as cfgrp.
*/
func (this *ConfigGroupSet) ContainCore(cfgrp *ConfigGroup) bool {
	_, exist := this.contain[cfgrp.hashKey]
	return exist
}

/*
Return the difference in context of each config group in both sets.
*/
func (this *ConfigGroupSet) ContextDiff(that *ConfigGroupSet) *ConfigGroupSet {
	diff := NewConfigGroupSet()
	for _, thisGrp := range this.List() {
		thatGrp := that.GetGroup(thisGrp)
		contextDiff := thisGrp.ContextSet.Diff(thatGrp.ContextSet)
		if contextDiff.Size() > 0 {
			diff.Add(NewConfigGroup(thisGrp.Item, contextDiff.List...))
		}
	}
	return diff
}

func (this *ConfigGroupSet) Equal(that *ConfigGroupSet) bool {
	if len(this.list) != len(that.list) {
		return false
	}
	for key, grp := range this.contain {
		thatGrp, exist := that.contain[key]
		if !exist {
			return false
		}
		if !grp.Equal(thatGrp) {
			return false
		}
	}
	return true
}

func (this *ConfigGroupSet) String() string {
	w := new(bytes.Buffer)
	for _, cfg := range this.list {
		fmt.Fprintf(w, "%s\n", cfg)
	}
	return w.String()
}

func (this *ConfigGroupSet) GetGroup(cg *ConfigGroup) *ConfigGroup {
	return this.contain[cg.hashKey]
}

func (this *ConfigGroupSet) GetGroupByCore(item *lr0items.Item) *ConfigGroup {
	if item == nil {
		return nil
	}
	return this.contain[item.HashKey()]
}

func (this *ConfigGroupSet) Core() []*lr0items.Item {
	core := make([]*lr0items.Item, 0, 4)
	for _, cfg := range this.List() {
		core = append(core, cfg.Core())
	}
	return core
}

func (this *ConfigGroupSet) List() []*ConfigGroup {
	return this.list
}

/*
Select from this all config groups with item matching those in configGroups.
*/
func (this *ConfigGroupSet) SelectNextSubset(configGroups ...*ConfigGroup) (selectedSet *ConfigGroupSet) {
	selectedSet = NewConfigGroupSet()
	for _, thisGrp := range this.list {
		for _, thatGrp := range configGroups {
			thisNext := thisGrp.Item.NextItem
			if thisNext.Equal(thatGrp.Item) {
				selectedSet.Add(NewConfigGroup(thisNext, thisGrp.ContextSet.List...))
			}
		}
	}
	return
}

func (this *ConfigGroupSet) CoreEqual(that *ConfigGroupSet) bool {
	if len(this.list) != len(that.list) {
		return false
	}
	for key := range this.contain {
		if _, exist := that.contain[key]; !exist {
			return false
		}
	}
	return true
}

/*
sort.Interface
*/

// Len is the number of elements in the collection.
func (this *ConfigGroupSet) Len() int {
	return len(this.list)
}

// Less returns whether the element with index i should sort
// before the element with index j.
func (this *ConfigGroupSet) Less(i, j int) bool {
	return this.list[i].Compare(this.list[j]) < 0
}

// Swap swaps the elements with indexes i and j.
func (this *ConfigGroupSet) Swap(i, j int) {
	j1 := this.list[j]
	this.list[j] = this.list[i]
	this.list[i] = j1
}
