
//! Module token is generated by GoGLL. Do not edit

extern crate lazy_static;

use std::rc::Rc;
use std::fmt;
use lazy_static::lazy_static;
use std::collections::HashMap;

/// Token is returned by the lexer for every scanned lexical token
pub struct Token {
	pub typ: Type,
	pub lext: usize, 
	pub rext: usize,
	input: Rc<Vec<char>>,
}

#[derive(PartialEq, Eq, Hash, Clone, Copy)]
pub enum Type {	
	Error, // "Error"
	EOF, // "$"
	T_0, // "ip"
	T_1, // "name"
	T_2, // "number1"
	T_3, // "sap"
	T_4, // "string"
	T_5, // "timestamp"
}

/**
New returns a new token.  
lext is the left extent and rext the right extent of the token in the input.  
input is the input slice scanned by the lexer.
*/
pub fn new<'a>(t: Type, lext: usize, rext: usize, input: &Rc<Vec<char>>) -> Rc<Token> {
	Rc::new(Token{
		typ:   t,
		lext:  lext,
		rext:  rext,
		input: input.clone(),
	})
}

impl Token {
	/// get_line_column returns the (line, column) of the left extent of the token
	pub fn get_line_column(&self) -> (usize, usize) {
		let mut line = 1;
		let mut col = 1;
		let mut j = 0;
		while j < self.lext {
			match self.input[j] {
			'\n' => {
				line += 1;
				col = 1
			},
			'\t' => col += 4,
			_ => col += 1
			}
			j += 1
		}
		(line, col)
	}
	
	/// literal returns the literal runes of t scanned by the lexer
	pub fn literal(&self) -> Vec<char> {
		self.input[self.lext..self.rext].to_vec()
	}
	
    /// literal_string returs the literal string of t scanned by the lexer
    #[allow(dead_code)]
	pub fn literal_string(&self) -> String {
		self.literal().iter().collect::<String>()
    }

} // impl Token

impl <'a>fmt::Display for Token {
	fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
		let (ln, col) = self.get_line_column();
		write!(f, "({}, ({},{}) {})", 
			self.typ, ln, col, self.literal().iter().collect::<String>())
	}

}

impl <'a>Type {
	/// id returns the token type ID of token Type t
	#[allow(dead_code)]
	pub fn id(&self) -> &'a str {
		TYPE_TO_STRING[self]
	}
	
}

impl <'a>fmt::Display for Type {
	fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
		write!(f, "{}", TYPE_TO_STRING[self])
	}

}

lazy_static! {
    static ref TYPE_TO_STRING: HashMap<Type, &'static str> = {
        let mut m = HashMap::new(); 
		m.insert(Type::Error, "Error");
		m.insert(Type::EOF, "$");
		m.insert(Type::T_0, "ip");
		m.insert(Type::T_1, "name");
		m.insert(Type::T_2, "number1");
		m.insert(Type::T_3, "sap");
		m.insert(Type::T_4, "string");
		m.insert(Type::T_5, "timestamp");
        m
    };
}

lazy_static! {
	static ref STRING_TO_TYPE: HashMap<&'static str, Type> = {
		let mut m = HashMap::new(); 
		m.insert("Error", Type::Error); 
		m.insert("EOF", Type::EOF); 
		m.insert("T_0", Type::T_0); 
		m.insert("T_1", Type::T_1); 
		m.insert("T_2", Type::T_2); 
		m.insert("T_3", Type::T_3); 
		m.insert("T_4", Type::T_4); 
		m.insert("T_5", Type::T_5); 
		m
	};
}
