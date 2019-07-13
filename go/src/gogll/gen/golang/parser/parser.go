package parser

import (
	"bytes"
	"fmt"
	"go/format"
	"gogll/ast"
	"gogll/cfg"
	"gogll/gen/golang/parser/slots"
	"gogll/goutil/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

/*** Main parser section ***/

func Gen(parserDir string) {
	genParser(parserDir)
	slots.Gen(filepath.Join(parserDir, "slot", "slot.go"))
}

func genParser(parserDir string) {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("Parser Main Template").Parse(mainTemplate)
	if err != nil {
		parseErrorError(err)
	}
	data := getData(parserDir)
	if err = tmpl.Execute(buf, data); err != nil {
		parseErrorError(err)
	}
	fmtSrc, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Printf("Error formatting generated parsers: %s\n", err)
		fmtSrc = buf.Bytes()
	}
	fname := path.Join(parserDir, "parser.go")
	if err := ioutil.WriteFile(fname, []byte(fmtSrc)); err != nil {
		parseErrorError(err)
	}
}

type Data struct {
	Package     string
	Imports     []string
	StartSymbol string
	CodeX       string
	TestSelect  string
}

func getData(baseDir string) *Data {
	data := &Data{
		Package:     cfg.Package,
		Imports:     getImports(baseDir),
		StartSymbol: ast.GetStartSymbol(),
		CodeX:       genAlternatesCode(),
		TestSelect:  genTestSelect(),
	}
	return data
}

func getImports(baseDir string) []string {
	return []string{
		"io/ioutil",
	}
}

func getPackage(baseDir string) string {
	if ast.GetPackage() == "" {
		pl := strings.Split(filepath.Clean(filepath.ToSlash(baseDir)), "/")
		pkg := pl[len(pl)-1]
		return pkg
	}
	return ast.GetPackage()
}

func parseErrorError(err error) {
	fmt.Printf("Error generating parser: %s\n", err)
	panic("fix me")
	os.Exit(1)
}

const mainTemplate = `
/* 
Package parser is generated by gogll. Do not edit.
*/
package parser

import(
	"bytes"
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"

	"{{.Package}}/goutil/bsr"
	"{{.Package}}/goutil/stringset"
	"{{.Package}}/parser/slot"
	{{range $i, $import := .Imports}}
	"{{$import}}" {{end}}
)

const (
	Dollar = "$"
	Empty = "empty"
)

func ParseFile(fname string) {
	buf, err := ioutil.ReadFile(fname)
	if err != nil {
		parseErrorError(err)
	}
	Parse(buf)
}

var (
	cI    = 0
	sz    = 0
	nextI = ""

	R *descriptors
	U *descriptors

	popped 		map[poppedNode]bool
	crf			map[clusterNode][]*crfNode
	crfNodes	map[crfNode]*crfNode
)

func initParser() {
	cI, nextI, sz = 0, "", 0
	R, U = &descriptors{}, &descriptors{}
	popped = make(map[poppedNode]bool)
	crf = map[clusterNode][]*crfNode{
		{"{{.StartSymbol}}", 0}:{},
	}
	crfNodes = map[crfNode]*crfNode{}
	bsr.Init()
}

func Parse(I []byte) error {
	initParser()
	var L slot.Label
	m, cU := len(I), 0
	nextI, _, sz = decodeRune(I[cI:])
	ntAdd("{{.StartSymbol}}", 0)
	fmt.Printf("R:%s\n", R)
	fmt.Printf("U:%s\n", U)
	for !R.empty() {
		L, cU, cI = R.remove()
		nextI, _, sz = decodeRune(I[cI:])

		fmt.Printf("L:%s, cI:%d, I[cI]:%s\n", L, cI, nextI)
		fmt.Printf("R:%s\n", R)
		fmt.Printf("U:%s\n", U)

		switch L { 
{{.CodeX}}

		default:
			panic("This must not happen")
		}
	}
	if !bsr.Contain("{{.StartSymbol}}",0,m) {
		return parseError()
	}
	return nil
}

func ntAdd(nt string, j int) {
	for _, l := range slot.GetAlternates(nt) {
		if testSelect[l](nextI) {
			dscAdd(l, j, j)
		} else {
			fmt.Println("testSelect == false")
		}
	}
}

/*** Call Return Forest ***/

type poppedNode struct {
	X    string
	k, j int
}

type clusterNode struct {
	X string
	k int
}

type crfNode struct {
	L slot.Label
	i int
}

/*
suppose that L is Y ::=αX ·β
if there is no CRF node labelled (L,i) 
	create one let u be the CRF node labelled (L,i)
if there is no CRF node labelled (X, j) { 
	create a CRF node v labelled (X, j) 
	create an edge from v to u 
	ntAdd(X, j) 
} else { 
	let v be the CRF node labelled (X, j) 
	if there is not an edge from v to u {
		create an edge from v to u 
		for all ((X, j,h)∈P) {
			dscAdd(L, i, h); 
			bsrAdd(L, i, j, h) 
		} 
	} 
}
*/
func call(L slot.Label, i, j int) {
	fmt.Printf("call(%s,%d,%d)\n", L,i,j)
	u, exist := crfNodes[crfNode{L, i}]
	fmt.Printf("  u exist=%t\n", exist)
	if !exist {
		u = &crfNode{L, i}
		crfNodes[*u] = u
	}
	X := L.Symbols()[L.Pos()-1]
	ndV := clusterNode{X, j}
	v, exist := crf[ndV]
	if !exist {
		fmt.Println("  v exist")
		crf[ndV] = []*crfNode{u}
		ntAdd(X, j)
	} else {
		fmt.Println("  v !exist")
		if !existEdge(v, u) {
			fmt.Printf("  !existEdge(%s)\n", u)
			crf[ndV] = append(v, u)
			fmt.Printf("|popped|=%d\n", len(popped))
			for pnd, _ := range popped {
				if pnd.X == X && pnd.k == j {
					dscAdd(L, i, pnd.j)
					bsr.Add(L, i, j, pnd.j)
				}
			}
		}
	}
}

func existEdge(nds []*crfNode, nd *crfNode) bool {
	for _, nd1 := range nds {
		if nd1 == nd {
			return true
		}
	}
	return false
}

func rtn(X string, k, j int) {
	fmt.Printf("rtn(%s,%d,%d)\n", X,k,j)
	p := poppedNode{X, k, j}
	if _, exist := popped[p]; !exist {
		popped[p] = true
		for _, nd := range crf[clusterNode{X, k}] {
			dscAdd(nd.L, nd.i, j)
			bsr.Add(nd.L, nd.i, k, j)
		}
	}
}
/*** descriptors ***/

type descriptors struct {
	set []*descriptor
}

func (ds *descriptors) contain(d *descriptor) bool {
	for _, d1 := range ds.set {
		if d1 == d {
			return true
		}
	}
	return false
}

func (ds *descriptors) empty() bool {
	return len(ds.set) == 0
}

func (ds *descriptors) String() string {
	buf := new(bytes.Buffer)
	buf.WriteString("{")
	for i, d := range ds.set {
		if i > 0 {
			buf.WriteString("; ")
		}
		fmt.Fprintf(buf, "%s", d)
	}
	buf.WriteString("}")
	return buf.String()
}

type descriptor struct {
	L slot.Label
	k int
	i int
}

func (d *descriptor) String() string {
	return fmt.Sprintf("%s,%d,%d", d.L, d.k, d.k)
}

func dscAdd(L slot.Label, k, i int) {
	fmt.Printf("dscAdd(%s,%d,%d)\n", L, k, i)
	d := &descriptor{L, k, i}
	if !U.contain(d) {
		R.set = append(R.set, d)
		U.set = append(U.set, d)
	}
}

func (ds *descriptors) remove() (L slot.Label, k, i int) {
	d := ds.set[len(ds.set)-1]
	ds.set = ds.set[:len(ds.set)-1]
	return d.L, k, d.i
}

/*** Rune decoding ***/
func decodeRune(str []byte) (string, rune, int) {
	if len(str) == 0 {
		return Dollar, -1, 0
	}
	r, sz := utf8.DecodeRune(str)
	if r == utf8.RuneError {
		panic(fmt.Sprintf("Rune error: %s", str))
	}
	if r == '\\' {
		r1, sz1 := utf8.DecodeRune(str[sz:])
		if r1 == utf8.RuneError {
			panic(fmt.Sprintf("Rune error: %s", str))
		}
		r, sz = r1, sz+sz1
	}
	chr := string(str[:sz])
	return chr, r, sz
}

func runeToString(r rune) string {
	buf := make([]byte, utf8.RuneLen(r))
	utf8.EncodeRune(buf, r)
	return string(buf)
}

/*** TestSelect ***/

{{.TestSelect}}

/*** Unicode functions ***/

func any(r rune) bool {
	return true
}
	
func letter(r rune) bool {
	return unicode.IsLetter(r)
}
	
func number(r rune) bool {
	return unicode.IsNumber(r)
}
	
func upcase(r rune) bool {
	return unicode.IsUpper(r)
}
	
func lowcase(r rune) bool {
	return unicode.IsLower(r)
}
	
func not(r rune, set string) bool {
	bs := []byte(set)
	for i := 0; i < len(set); {
		r1, sz := utf8.DecodeRune(bs[i:])
		if r1 == utf8.RuneError {
			panic(fmt.Sprintf("Rune error: %s", set))
		}
		if r1 == r {
			return false
		} 
		i += sz
	}
	return true
}
	
func space(r rune) bool {
	return unicode.IsSpace(r)
}
	
/*** Errors ***/

func parseError() error {
	return fmt.Errorf("parse Error")
}

func parseErrorError(err error) {
	fmt.Printf("Error: %s\n", err)
	os.Exit(1)
}
`
