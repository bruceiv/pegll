package bsr

import (
	"bytes"
	"gogll/cfg"
	"gogll/goutil/ioutil"
	"text/template"
)

func Gen(bsrFile string) {
	tmpl, err := template.New("BSR").Parse(bsrTmpl)
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	if err = tmpl.Execute(buf, cfg.Package); err != nil {
		panic(err)
	}
	if err = ioutil.WriteFile(bsrFile, buf.Bytes()); err != nil {
		panic(err)
	}
}

const bsrTmpl = `/*
Package bsr is generated by gogll. Do not edit.
It implements a Binary Subtree Representation set as defined in

	Scott et al
	Derivation representation using binary subtree sets,
	Science of Computer Programming 175 (2019)

*/
package bsr

import (
	"fmt"
	"sort"
	"strings"
	
	"{{.}}/parser/slot"
)

type BSR interface {
	LeftExtent() int
	RightExtent() int
	InputPos() int
}

var Set = newSet()

type BSRSet struct {
	slotEntries   map[Slot]bool
	stringEntries map[String]bool
}

type Slot struct {
	Label       slot.Label
	leftExtent  int
	inputPos    int
	rightExtent int
}

type String struct {
	Label slot.Label
	leftExtent  int
	inputPos    int
	rightExtent int
}

func newSet() *BSRSet {
	return &BSRSet{
		slotEntries:   make(map[Slot]bool),
		stringEntries: make(map[String]bool),
	}
}

/*
Add a BSR to the set. (i,j) is the extent. k is the pivot.
*/
func Add(l slot.Label, i, k, j int) {
	fmt.Printf("bsr.Add(%s,%d,%d,%d)\n", l,i,k,j)
	if l.EoR() {	
		insert(Slot{l, i, k, j})
	} else {
		if l.Pos() > 0 {
			insert(String{l, i, k, j})
		}
	}
}

func AddEmpty(l slot.Label, i int) {
	insert(String{l, i, i, i})
}

func Contain(nt string, left, right int) bool {
	for e, _ := range Set.stringEntries {
		if e.Label.Slot().NT == nt && e.LeftExtent() == left && e.RightExtent() == right {
			return true
		}
	}
	return false
}

func Init() {
	Set = newSet()
}

func insert(bsr BSR) {
	switch s := bsr.(type) {
	case Slot:
		Set.slotEntries[s] = true
	case String:
		Set.stringEntries[s] = true
	default:
		panic(fmt.Sprintf("Invalid type %T", bsr))
	}
}

func (s Slot) LeftExtent() int {
	return s.leftExtent
}

func (s Slot) RightExtent() int {
	return s.rightExtent
}

func (s Slot) InputPos() int {
	return s.inputPos
}

func (s Slot) String() string {
	return fmt.Sprintf("%s,%d,%d,%d", s.Label, s.leftExtent, s.inputPos, s.rightExtent)
}

func (s String) LeftExtent() int {
	return s.leftExtent
}

func (s String) RightExtent() int {
	return s.rightExtent
}

func (s String) InputPos() int {
	return s.inputPos
}

func (s String) Empty() bool {
	return s.leftExtent == s.inputPos && s.inputPos == s.rightExtent
}

func (s String) String() string {
	fmt.Printf("bsr.String.String(): %s, %d, %d, %d\n",
		s.Label.Symbols(), s.leftExtent, s.inputPos, s.rightExtent)
	ss := s.Label.Symbols()[s.leftExtent:s.RightExtent()]
	str := strings.Join(ss, " ")
	return fmt.Sprintf("%s,%d,%d,%d", str, s.leftExtent, s.inputPos,
		s.rightExtent)
}

func Dump() {
	DumpSlots()
	DumpStrings()
}

func DumpSlots() {
	fmt.Printf("Slots (%d)\n", len(GetSlots()))
	for _, s := range GetSlots() {
		DumpSlot(s)
	}
}

func DumpSlot(s Slot) {
	fmt.Println(s)
}

func DumpStrings() {
	fmt.Printf("Strings(%d)\n", len(GetStrings()))
	for _, s := range GetStrings() {
		DumpString(s)
	}
}

func DumpString(s String) {
	fmt.Println(s)
}

func GetSlots() (slots []Slot) {
	for s := range Set.slotEntries {
		slots = append(slots, s)
	}
	sort.Slice(slots,
		func(i, j int) bool {
			return slots[i].Label < slots[j].Label
		})
	return
}

func GetStrings() (strings []String) {
	for s := range Set.stringEntries {
		strings = append(strings, s)
	}
	sort.Slice(strings,
		func(i, j int) bool {
			return strings[i].Label < strings[j].Label
		})
	return
}
`
