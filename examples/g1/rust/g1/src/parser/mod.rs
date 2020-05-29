//! Module parser is generated by GoGLL. Do not edit.

extern crate lazy_static;

pub mod bsr;
mod slot;
mod symbols;

use crate::lexer;
use crate::token;
use slot::{Label};
use symbols::{NT,Symbol};

use lazy_static::lazy_static;
use std::collections::{HashMap, HashSet};
use std::fmt;
use std::rc::Rc;

struct Parser {
	c_i: usize,

	r: Vec<Rc<Descriptor>>,
	u: Vec<Rc<Descriptor>>,

	popped:    HashSet<Box<PoppedNode>>,
	crf:       HashMap<ClusterNode, HashSet<CRFNode>>,
	crf_nodes: HashSet<CRFNode>,

	lex:    Rc<lexer::Lexer>,
	errors: Vec<Box<ParseError>>,

	bsr_set: Box<bsr::Set>,
}

// ToDo: delete
// struct Descriptors {
// 	set: Vec<Box<Descriptor>>,
// }

#[derive(Hash,Eq,PartialEq,Debug)]
struct Descriptor {
	l: Label,
	k: usize,
	i: usize,
}

/**
Error is returned by Parse at every point at which the parser fails to parse
a grammar production. For non-LL-1 grammars there will be an error for each
alternate attempted by the parser.

The errors are sorted in descending order of input position (index of token in
the stream of tokens).

Normally the error of interest is the one that has parsed the largest number of
tokens.
*/
pub struct Error {
	/// Index of token that caused the error.
	pub c_i: usize,

	/// Grammar slot at which the error occured.
	pub slot: Label,

	/// The token at which the error occurred.
	pub token: Rc<token::Token>,

	/// The tokens expected at the point where the error occurred
    pub expected: Box<HashSet<token::Type>>,
    
    /// The line in the input where the error occurred
    pub line: usize,

    /// The column on the line where the error occurred
    pub column: usize,
}

// ParseErrors are generated during the parse. After a failed parse they 
// are translated to Errors, which are returned to the user.
struct ParseError {
    c_i: usize,
    slot: Label,
    token: Rc<token::Token>,
    expected: Expected,
}

// Expected indicates whether to use the First or Follow set for the exported error.
enum Expected {
    First,
    Follow(NT)
}

#[derive(Hash, Eq, PartialEq, Debug, Clone, Copy)]
struct PoppedNode {
	x: NT,
    k: usize,
    j: usize,
}

#[derive(Hash, Eq, PartialEq, Debug, Clone, Copy)]
struct ClusterNode {
	x: symbols::NT,
	k: usize,
}

// Call return forest node
#[derive(Hash, Eq, PartialEq, Debug, Clone, Copy)]
struct CRFNode {
	l: Label,
	i: usize,
}

/// Parse returns the BSR set containing the parse forest.
/// If the parse was successfull the length of Vec\<Box\<Error\>\> = 0.
#[allow(dead_code)]
pub fn parse(l: Rc<lexer::Lexer>) -> (Box<bsr::Set>, Vec<Box<Error>>) {
    let mut p = Parser::new(l.clone());
    p.parse();
    if !p.bsr_set.contain(&NT::Exp, 0, l.tokens.len()-1) {
        let errors = p.export_errors();
        (p.bsr_set, errors)
    } else {
        (p.bsr_set, vec![])
    }
}

impl Parser {
    fn new(l: Rc<lexer::Lexer>) -> Box<Parser> {
        let mut p = Box::new(Parser{
            c_i:         0,
            lex:         l.clone(),
            r:           Vec::with_capacity(1024),
            u:           Vec::with_capacity(1024),
            popped:      HashSet::with_capacity(1024),
            crf:         HashMap::with_capacity(1024),
            crf_nodes:   HashSet::with_capacity(1024),
            bsr_set:     bsr::Set::new(NT::Exp, l.clone()),
            errors:      Vec::with_capacity(1024),
        });
        p.crf.insert(ClusterNode::new(NT::Exp, 0), HashSet::with_capacity(128));
        p
    }

    fn parse(&mut self) {
        // let mut c_u = 0;
        self.nt_add(NT::Exp, 0);
        // let mut slotNo = 0;
        while self.r.len() > 0 {
            let (l, c_u, c_i) = self.r_remove();
            self.c_i = c_i;

            // println!("{no}:{l} i {i} u {u} tok {t}", 
            //     no=slotNo, l=l, i=c_i, u=c_u, t=self.lex.tokens[c_i]);
            // slotNo += 1;

            // for d in self.r.iter() {
            //     println!("  {}", d);
            // }

            (|| {
                match l { 
                    // Exp : ∙Exp & Exp 
                    Label::Exp0R0 => { 
                        self.call(Label::Exp0R1, c_u, self.c_i);
                    },
                    // Exp : Exp ∙& Exp 
                    Label::Exp0R1 => {
                        if !self.test_select(Label::Exp0R1){ 
                            self.error_first(Label::Exp0R1, self.c_i);
                            return; 
                        }
                        self.bsr_set.add(Label::Exp0R2, c_u, self.c_i, self.c_i+1);
                        self.c_i += 1; 
                        if !self.test_select(Label::Exp0R2){ 
                            self.error_first(Label::Exp0R2, self.c_i);
                            return; 
                        }
                        self.call(Label::Exp0R3, c_u, self.c_i);
                    },
                    // Exp : Exp & Exp ∙
                    Label::Exp0R3 => {
                        if self.follow(NT::Exp) {
                            self.rtn(NT::Exp, c_u, self.c_i)
                        } else { 
                            self.error_follow(Label::Exp0R0, self.c_i, NT::Exp)
                        }
                    }, 
                    // Exp : ∙Exp | Exp 
                    Label::Exp1R0 => { 
                        self.call(Label::Exp1R1, c_u, self.c_i);
                    },
                    // Exp : Exp ∙| Exp 
                    Label::Exp1R1 => {
                        if !self.test_select(Label::Exp1R1){ 
                            self.error_first(Label::Exp1R1, self.c_i);
                            return; 
                        }
                        self.bsr_set.add(Label::Exp1R2, c_u, self.c_i, self.c_i+1);
                        self.c_i += 1; 
                        if !self.test_select(Label::Exp1R2){ 
                            self.error_first(Label::Exp1R2, self.c_i);
                            return; 
                        }
                        self.call(Label::Exp1R3, c_u, self.c_i);
                    },
                    // Exp : Exp | Exp ∙
                    Label::Exp1R3 => {
                        if self.follow(NT::Exp) {
                            self.rtn(NT::Exp, c_u, self.c_i)
                        } else { 
                            self.error_follow(Label::Exp1R0, self.c_i, NT::Exp)
                        }
                    }, 
                    // Exp : ∙id 
                    Label::Exp2R0 => { 
                        self.bsr_set.add(Label::Exp2R1, c_u, self.c_i, self.c_i+1);
                        self.c_i += 1; 
                        if self.follow(NT::Exp) {
                            self.rtn(NT::Exp, c_u, self.c_i)
                        } else { 
                            self.error_follow(Label::Exp2R0, self.c_i, NT::Exp)
                        }
                    }, 
                    _ => unimplemented!()
                };
            })();
        };
    }
    
    fn nt_add(&mut self, nt: NT, j: usize) {
        // println!("nt_add({},{}", nt, j);

        let mut failed = true;
        let mut expected: HashSet<token::Type> = HashSet::with_capacity(128);
        for l in slot::get_alternates(&nt).iter() {
            if self.test_select(*l) {
                self.dsc_add(*l, j, j);
                failed = false
            } else {
                for tok in FIRST[l].iter() {
                    expected.insert(tok.clone());
                }
            }
        }
        if failed {
            for l in slot::get_alternates(&nt) {
                self.error_first(*l, j)
            }
        }
    }

    /*
    suppose that L is Y ::=αX ·β
    if there is no CRF node labelled (L,i)
        create one let u be the CRF node labelled (L,i)
    if there is no CRF node labelled (X, j) {
        create a CRF node v labelled (X, j)
        create an edge from v to u
        nt_add(X, j)
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
    fn call(&mut self, l: Label, i: usize, j: usize) {
        let u = CRFNode::new(l, i);
        if let None = self.crf_nodes.get(&u) {
            self.crf_nodes.insert(u);
        }
        let x = match l.symbols()[l.pos()-1]{
            Symbol::NT(x) => x,
            _ => panic!("Symbol::T is invalid"),
        };
        let nd_v = ClusterNode::new(x, j);
        match self.crf.get_mut(&nd_v) {
            None => {
                let mut m: HashSet<CRFNode> = HashSet::with_capacity(128);
                m.insert(u);
                self.crf.insert(nd_v, m);
                self.nt_add(x, j);
            },
            Some(v) => {
                if !v.contains(&u) {
                    v.insert(u);
                    let mut descs: Vec<Rc<Descriptor>> = Vec::new();
                    for pnd in self.popped.iter() {
                        if pnd.x == x && pnd.k == j {
                            descs.push(Descriptor::new(l, i, pnd.j));
                            self.bsr_set.add(l, i, j, pnd.j);
                        }
                    }
                    for d in descs.into_iter() {
                        self.dsc_add(d.l, d.k, d.i)
                    }
                }
            }
        }
    }
    
    fn rtn(&mut self, x: NT, k: usize, j: usize) {
        let pn = PoppedNode::new(x, k, j);
        if !self.popped.contains(&pn) {
            self.popped.insert(pn);
            for nd in self.crf[&ClusterNode::new(x, k)].clone() {
                self.dsc_add(nd.l, nd.i, j);
                self.bsr_set.add(nd.l, nd.i, k, j);
            }
        }
    }
    
    fn dsc_add(&mut self, l: Label, k: usize, i: usize) {
        let d = Descriptor::new(l, k, i);
        if !self.u.contains(&d) {
            self.r.push(d.clone());
            self.u.push(d.clone());
        }
    }
    
    fn r_remove(&mut self) -> (Label, usize, usize) {
        match self.r.pop() {
            Some(d) => return (d.l, d.k, d.i),
            None => panic!("empty")
        }
    }

    fn error_first(&mut self, l: Label, i: usize) {
        self.errors.push(
            Box::new(ParseError{
                c_i: i, 
                slot: l, 
                token: self.lex.tokens[i].clone(), 
                expected: Expected::First,
            })
        );
    }

    fn error_follow(&mut self, l: Label, i: usize, nt: NT) {
        self.errors.push(
            Box::new(ParseError{
                c_i: i, 
                slot: l, 
                token: self.lex.tokens[i].clone(), 
                expected: Expected::Follow(nt),
            })
        );
    }

    fn export_errors(&mut self) -> Vec<Box<Error>> {
        let mut errs: Vec<Box<Error>> = Vec::new();
        self.errors.sort_by(|a,b| a.token.lext.cmp(&b.token.lext));
        for err in self.errors.iter() {
            let (ln, col) = self.lex.get_line_column(err.token.lext);
            errs.push(Box::new(Error{
                c_i: err.c_i,
                slot: err.slot,
                token: err.token.clone(),
                expected: match err.expected {
                    Expected::First => FIRST[&err.slot].clone(),
                    Expected::Follow(nt) => FOLLOW[&nt].clone(),
                },
                line: ln,
                column: col,
            }));
        }
        errs
    }
    
    fn test_select(&self, l: Label) -> bool {
        FIRST[&l].contains(&self.lex.tokens[self.c_i].typ)
    }

    fn follow(&self, nt: NT) -> bool {
        FOLLOW[&nt].contains(&self.lex.tokens[self.c_i].typ)
    }
    
} /*** impl Parser ***/

impl ClusterNode {
    fn new(nt: NT, k: usize) -> ClusterNode {
        ClusterNode{
            x: nt,
            k: k,
        }
    }
}

impl CRFNode {
    fn new(l: Label, i: usize) -> CRFNode {
        CRFNode{
            l: l,
            i: i,
        }
    }
}

impl Descriptor {
    fn new(l: Label, k: usize, i: usize) -> Rc<Descriptor> {
        Rc::new(Descriptor{
            l: l,
            k: k,
            i: i,
        })
    }
}

impl PoppedNode {
    fn new(x: NT, k: usize, j: usize) -> Box<PoppedNode> {
        Box::new(PoppedNode{
            x: x,
            k: k,
            j: j,
        })
    }
}

impl fmt::Display for Descriptor {    
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "l={l},k={k},i={i}", 
            l=self.l,
            k=self.k,
            i=self.i,
        )
    }
}
    
impl fmt::Display for Error {    
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        let mut errs: Vec<String> = Vec::new();
        for tok in self.expected.iter() {
            errs.push(format!("{}",tok));
        };
        write!(f, "Error: {slot}, token {tok}, expected {{{exp}}} at line {ln} col {col}", 
            slot=self.slot,
            tok=self.token,
            exp=errs.join(","),
            ln=self.line,
            col=self.column,
        )
    }
}
    
    lazy_static! {
    static ref FIRST: HashMap<Label, Box<HashSet<token::Type>>> = {
        let mut fmap = HashMap::new(); 
        // Exp : ∙Exp & Exp 
            let mut hset = Box::new(HashSet::new()); 
            hset.insert(token::Type::Type1); // id 
            fmap.insert(Label::Exp0R0, hset);
        // Exp : Exp ∙& Exp 
            let mut hset = Box::new(HashSet::new()); 
            hset.insert(token::Type::Type0); // & 
            fmap.insert(Label::Exp0R1, hset);
        // Exp : Exp & ∙Exp 
            let mut hset = Box::new(HashSet::new()); 
            hset.insert(token::Type::Type1); // id 
            fmap.insert(Label::Exp0R2, hset);
        // Exp : Exp & Exp ∙
            let mut hset = Box::new(HashSet::new()); 
            hset.insert(token::Type::Type0); // & 
            hset.insert(token::Type::EOF); // EOF 
            hset.insert(token::Type::Type2); // | 
            fmap.insert(Label::Exp0R3, hset);
        // Exp : ∙Exp | Exp 
            let mut hset = Box::new(HashSet::new()); 
            hset.insert(token::Type::Type1); // id 
            fmap.insert(Label::Exp1R0, hset);
        // Exp : Exp ∙| Exp 
            let mut hset = Box::new(HashSet::new()); 
            hset.insert(token::Type::Type2); // | 
            fmap.insert(Label::Exp1R1, hset);
        // Exp : Exp | ∙Exp 
            let mut hset = Box::new(HashSet::new()); 
            hset.insert(token::Type::Type1); // id 
            fmap.insert(Label::Exp1R2, hset);
        // Exp : Exp | Exp ∙
            let mut hset = Box::new(HashSet::new()); 
            hset.insert(token::Type::Type0); // & 
            hset.insert(token::Type::EOF); // EOF 
            hset.insert(token::Type::Type2); // | 
            fmap.insert(Label::Exp1R3, hset);
        // Exp : ∙id 
            let mut hset = Box::new(HashSet::new()); 
            hset.insert(token::Type::Type1); // id 
            fmap.insert(Label::Exp2R0, hset);
        // Exp : id ∙
            let mut hset = Box::new(HashSet::new()); 
            hset.insert(token::Type::Type0); // & 
            hset.insert(token::Type::EOF); // EOF 
            hset.insert(token::Type::Type2); // | 
            fmap.insert(Label::Exp2R1, hset);
        fmap
    };

    static ref FOLLOW: HashMap<NT, Box<HashSet<token::Type>>> = {
        let mut fmap = HashMap::new(); 
        // Exp
            let mut hset = Box::new(HashSet::new()); 
            hset.insert(token::Type::Type0); // & 
            hset.insert(token::Type::EOF); // EOF 
            hset.insert(token::Type::Type2); // | 
            fmap.insert(NT::Exp, hset);
        fmap
    };
}
