# **`eggr` Grammar**

### **AUTHORSHIP INFORMATION**
#### *Author :* Brynn Harrington and Emily Hoppe Copyright (C) 2021
#### *Adapted from :* Aaron Moss's [`eggr` Egg Grammar](https://github.com/bruceiv/egg/blob/deriv/grammars/eggr.egg)
#### *Creation Date :* June 11, 2021 
#### *Last Modified :* June 23, 2021
#### *Copyright and Licensing Information :* See end of file.

###  **GENERAL DESCRIPTION**
An originally Egg Parsing grammar created by Aaron Moss ported into the GoGLL grammar to test a given structure.Modification of `eggr` grammar from [Egg](https://github.com/bruceiv/egg/blob/deriv/grammars/eggr.egg) to test an example structure.
### **STATUS ON GRAMMAR**
#### *Markdown File Creation:* Working
#### *Parser Generated :* Complete
#### *Test File Creation:* Incomplete
#### *Testing Results:* Unknown
### **`eggr` Grammar Guide**
NEED TO FIX:
- double check translation of regular expressions in semantic rules

```
package "eggr"
```
The following sections handles the grammar, rules, and choices:
- `Grammar` is matched by at least one rule;
- `Rule` is an identifier equal to a choice;
- `Choice` is zero or more piped sequences. 
```
Grammar         : WS Rule Rules ;
        Rules   : Rule Rules
                / empty; 
Rule            : Identifier "=" Choice ;

Choice          : Sequence PipedSeq0x ;
     PipedSeq0x : PipedSeq PipedSeq0x
                / empty ; 
       PipedSeq : "|" Sequence ;
```
The following sections handles sequences and expressions, where:
- `Sequence` is one or more expressions;
- `Expression` uses operators of the grammar to determine rules. 
```
Sequence        : Expression Expr1x ;
        Expr1x  : Expression Expr1x
                / empty ;

Expression      : "&" WS Primary 
                | "!" WS Primary 
                | Primary OptStarPlus ;                     
    OptStarPlus : "?" WS 
                | "*" WS 
                | "+" WS 
                | empty ; 
```
The following sections handles primaries and identifers, in which :
- `Primary` is matched by an `Identifier` followed by anything but '=', a nested `Choice`, `StringLiteral`, `CharLiteral`, `CharClass`, any character, or a semicolon;
` `Identifier` is matched by a letter followed by whitespace and then a letter or number combination repeated zero or more times.
```
Primary         : Identifier neq
                | "(" Choice ")"
                | StringLiteral
                | CharLiteral
                | CharClass
                | "." WS
                | ";"  WS;
        neq     : not "=" ;

Identifier      : LetWS LetOrNum0x WS ;
        LetWS   : let
                / WS ;
     LetOrNum0x : LetOrNum LetOrNum0x
                / empty ;     
       LetOrNum : let
                / num
                / WS ;
        num     : number ;
        let     : letter ;
```
The following sections handles string and character literals where:
- `StringLiteral` is any string literal, which is made of a character repeated zero or more times surrounded in double quotes, followed by a whitespace. 
- `CharLiteral` is any character literal, surrounded by single quotes, followed by whitespace.
- `CharClass` is the character class consisting of unclosed characters `UnChars`, any character other than ']', repeated zero of more times. 
- `Character` is the semantic rule to represent a character in the `eggr` grammar. 
```
StringLiteral   : dQuote String dQuote WS ;
        dQuote  : any "\"" ;
        String  : Character String
                / empty ;

CharLiteral     : "'" Character "'" ;

CharClass       : "[" UnChars "]" WS;
        UnChars : UnChar UnChars 
                / empty ;
        UnChar : notSqBk Character ;
        notSqBk : not "]" ;

Character       : notQuotesEsc
                | escAny ;  
   notQuotesEsc : not "'\"\\" ;
        escAny  :'\\' any "nrt'\"\\" ;        
```
The following section handles whitespace, spacing/commenting, and escape character/end of line sequences where: 
- `WS` is a semantic rule matching any kind of whitespace - a space/comment sequence or an empty
- `SpaceOrComment` is a semantic rule where the option between a space or `LineOrBlock` comment can be matched. 
- `space` is a lexical rule for escape characters representing a set of spaces or the end of line.
- `end_of_line` is a lexical rule for the escape characters that signify the end of the line. Any of these characters will indicate the end of line has been reached in `eggr`. 
```
WS              : SpaceOrComment WS
                | empty ;
SpaceOrComment  : space
                | LineOrBlock ;
space           : any " \t\r\n" ;
end_of_line     : any "\n\r" ;  
```
The following section handles commenting where:
- `LineOrBlock` represents the semantic rule for either a line or a block comment. 
- `!line_comment` is a lexical rule representing a C-style line comment. Everything from the first slash to the end of line is a comment. 
- `!block_comment` is a lexical rule representing a C-style block comment. Everything between and including `/*` and `*/` is a comment. 
- The `!` in front of `!line_comment` and `!block_comment` instructs the lexer to suppress those tokens. See the [grammar for details.](../../gogll.md) 
#### *Note:* `!line_comment` and `!block_comment` were taken from [comments.md.](https://github.com/bruceiv/pegll/tree/main/examples/comments) 
```
LineOrBlock     : line_comment 
                | block_comment ;
!line_comment   : '/' '/' {not "\n"} ;
!block_comment  : '/''*' 
                { not "*" 
                | '*' not "/" 
                } '*''/' ;
```
#
### **COPYRIGHT AND LICENSING INFORMATION**
**Copyright 2021 Brynn Harrington and Emily Hoppe**

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.

You may obtain a copy of the License [here](http://www.apache.org/licenses/LICENSE-2.0) or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.