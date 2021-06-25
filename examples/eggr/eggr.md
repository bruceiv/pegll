# **`eggr` GRAMMAR**

### **AUTHORSHIP INFORMATION**
#### *Author :* Brynn Harrington and Emily Hoppe Copyright (C) 2021
#### *Adapted from :* Aaron Moss's [`eggr` Egg Grammar](https://github.com/bruceiv/egg/blob/deriv/grammars/eggr.egg)
#### *Creation Date :* June 11, 2021 
#### *Last Modified :* June 24, 2021
#### *Copyright and Licensing Information :* See end of file.

###  **GENERAL DESCRIPTION**
A modification of `eggr`[Egg](https://github.com/bruceiv/egg/blob/deriv/grammars/eggr.egg) parsing grammar ported into GoGLL a given structure.

### **STATUS ON GRAMMAR**
#### *Markdown File Creation:* Working
#### *Parser Generated :* Complete
#### *Test File Creation:* Incomplete
#### *Testing Results:* Unknown
### **`eggr` GRAMMAR GUIDE**
The following grammar takes a given structure and tests the grammars capability to match the structure of `Egg` parsing grammar.
```
package "eggr"
```
The following sections handles the grammar, rules, and choices:
- `Grammar` is matched by at least one rule;
- `Rule` is an identifier equal to a choice;
- `Choice` is zero or more piped sequences. 
```
Grammar                 : WS Rule RepRule0x             ;
        RepRule0x       : Rule RepRule0x
                        / empty                         ; 
Rule                    : Identifier EQUAL Choice ;

Choice                  : Sequence RepPipedSeq0x        ;
     RepPipedSeq0x      : PIPE Sequence RepPipedSeq0x
                        / empty                         ; 
```
The following sections handles sequences and expressions, where:
- `Sequence` is one or more expressions;
- `Expression` uses operators of the grammar to determine rules. 
```
Sequence                : Expression RepExpr0x          ;
        RepExpr0x       : Expression RepExpr0x
                        / empty                         ;

Expression              : AND Primary 
                        | NOT Primary 
                        | Primary OptStarPlus           ;                     
    OptStarPlus         : OPT 
                        | STAR
                        | PLUS
                        | empty                         ; 
```
The following sections handles `Primary` and `Identifier`, in which :
- `Primary` is matched by an `Identifier` followed by `NEQUAL`, a `Choice` enclosed in parentheses, `StringLiteral`, `CharLiteral`, `CharClass`, `ANY`, or `EMPTY`;
- `Identifier` is matched by a letter followed by whitespace and then a letter or number combination repeated zero or more times.
```
Primary                 : Identifier NEQUAL
                        | OPEN Choice CLOSE
                        | StringLiteral
                        | CharLiteral
                        | CharClass
                        | ANY
                        | EMPTY                         ;

Identifier              : LetWS LetOrNum0x WS           ;
        LetWS           : let
                        / WS                            ;
     LetOrNum0x         : LetOrNum LetOrNum0x
                        / EMPTY                         ;     
       LetOrNum         : let
                        / num
                        / WS                            ;
        num             : number                        ;
        let             : letter                        ;
```
The following sections handles `StringLiteral`, `CharLiteral`, `CharClass`, and `Character` where:
- `StringLiteral` is any string literal matched by a character repeated zero or more times surrounded in double quotes and then followed by a whitespace;
- `CharLiteral` is any character literal, surrounded by single quotes, followed by whitespace;
- `CharClass` is the character class consisting of `UnclosedChars`, any character other than ']', repeated zero of more times; 
- `Character` is the semantic rule to represent a character in the `eggr` grammar. 
```
StringLiteral           : dQuote String dQuote WS       ;
        dQuote          : any "\""                      ;
        String          : Character String
                        / empty                         ;

CharLiteral             : "'" Character "'"             ;

CharClass               : "[" UnclosedChars "]" WS      ;
        UnclosedChars   : UnclosedChar UnclosedChars 
                        / empty                         ;
        UnclosedChar    : notSqBk Character             ;
        notSqBk         : not "]"                       ;

Character               : notQuotesEsc
                        | esc                           ;  
   notQuotesEsc         : not "'\"\\"                   ;
        esc             :'\\' any "nrt'\"\\"            ;  
```
The following section handles `EQUAL`, `NEQUAL` `PIPE`, `and`, `not`, `OPT`, `STAR`, `PLUS`, `OPEN`, `CLOSE`, `ANY`, and `EMPTY`, where:
- `EQUAL` is a semantic rule representing the character '=' followed by whitespace;
- `NEQUAL` is a semantic rule representing any character except '=' followed by whitespace;
- `PIPE` is a semantic rule representing the character '|' followed by whitespace;
- `AND` is a semantic rule representing the character '&' followed by whitespace;
- `NOT` is a semantic rule representing the character '!' followed by whitespace;
- `OPT` is a semantic rule representing the character '?' followed by whitespace;
- `STAR` is a semantic rule representing the character '*' followed by whitespace;
- `PLUS` is a semantic rule representing the character '+' followed by whitespace;
- `OPEN` is a semantic rule representing the character '(' followed by whitespace;
- `CLOSE` is a semantic rule representing the character ')' followed by whitespace;
- `ANY` is a semantic rule representing the character '.' followed by whitespace;
- `EMPTY` is a semantic rule representing the character ';' followed by whitespace.
```
EQUAL                   : "="   WS                      ;
NEQUAL                  : neq   WS                      ;
        neq             : not "="                       ;
PIPE                    : "|"   WS                      ;
AND                     : "&"   WS                      ;
NOT                     : "!"   WS                      ;
OPT                     : "?"   WS                      ;
STAR                    : "*"   WS                      ;
PLUS                    : "+"   WS                      ;
OPEN                    : "("   WS                      ;
CLOSE                   : "("   WS                      ;
ANY                     : "."   WS                      ;
EMPTY                   : ";"   WS                      ;      
```     
The following section handles `WS`, `SpaceOrComment`, `space` and `endOfLine` where: 
- `WS` is a semantic rule matching any kind of whitespace - a space/comment sequence or an EMPTY;
- `SpaceOrComment` is a semantic rule where either a `space` or `LineOrBlock` comment can be matched; 
- `space` is a lexical rule for escape characters representing a set of spaces or the end of line;
- `endOfLine` is a lexical rule for the escape characters that signify the end of the line. Any of these characters will indicate the end of line has been reached in `eggr`. 
```
WS                      : SpaceOrComment WS
                        / EMPTY                         ;
SpaceOrComment          : space
                        | LineOrBlock                   ;
space                   : any " \t\r\n"                 ;
endOfLine               : any "\n\r"                    ;  
```
The following section handles `LineOrBlock`, `lineComment`, and `blockComment` where:
- `LineOrBlock` represents the semantic rule for either a line or a block comment;
- `!lineComment` is a lexical rule representing a C-style line comment. Everything from the first slash to the end of line is a comment;
- `!blockComment` is a lexical rule representing a C-style block comment. Everything between and including `/*` and `*/` is a comment;
- The `!` in front of `!lineComment` and `!blockComment` instructs the lexer to suppress those tokens. See the [grammar for details.](../../gogll.md) 
#### *Note:* `!lineComment` and `!blockComment` were taken from [comments.md.](https://github.com/bruceiv/pegll/tree/main/examples/comments) 
```
LineOrBlock             : lineComment 
                        | blockComment                  ;
!lineComment            : '/' '/' { not "\n" }          ;
!blockComment           : '/''*' 
                        { not "*" 
                        | '*' not "/" 
                        } '*''/'                        ;       
```
#
### **COPYRIGHT AND LICENSING INFORMATION**
**Copyright 2021 Brynn Harrington and Emily Hoppe**

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.

You may obtain a copy of the License [here](http://www.apache.org/licenses/LICENSE-2.0) or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.