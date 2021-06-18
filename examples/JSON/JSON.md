## JSON grammar
Modification of `JSON` grammar from [Egg](https://github.com/bruceiv/egg/blob/deriv/grammars/JSON-u.egg) to test JSON gramar.

```
//-------------------------------------------------------------
//  GENERAL DESCRIPTION
//  An originally Egg Parsing grammar created by Dr. Aaron Moss
//  ported into the GoGLL grammar for the JSON language tests.
//
//  EGG PARSING GRAMMAR
// @ Author : Aaron Moss Copyright (C) 2017
//
//  GOGLL PARSING GRAMMAR
// @ Authors : Dr. Aaron Moss, Brynn Harrington, and Emily Hoppe
//  Creation Date : June 12, 2021 
//  Last Modification: June 14, 2021
//
//  COPYRIGHT AND LICENSING INFORMATION
//  This is free software; you can redistribute and/or modify
//  it under the terms of the GNU Library General Public License
//  as published by the Free Software Foundation; either version 2
//  of the License or (at your option) any later version.
//
//  This file is distributed in the hope that it will be useful,
//  but WITHOUT ANY WARRANTY; without even the implied warranty of
//  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
//
//  For more details, see the GNU Library General Public License
//  at http://www.fsf.org/copyleft/gpl.html.
//-------------------------------------------------------------

package "JSON"                          ;
//-------------------------------------------------------------
//  Object Creation ??
//-------------------------------------------------------------
JSON            : _ Object              ;

Object          : LBRACE OptMem RBRACE  ;
        OptMem  : < Members >           ;

Members         : Pair ComPair          ;
        ComPair : {COMMA Pair}          ;

Pair            : String COLON Value    ;

Array           : LBRACKET OptElem RBRACKET ;
        OptElem : [ Elements ]          ;

Elements        : Value ComVal          ;
        ComVal  : {COMMA Value}         ;

Value           : String 
                | Number 
                | Object 
                | Array 
                | TRUE 
                | FALSE 
                | NUL                   ;

//-------------------------------------------------------------
// Literals
//-------------------------------------------------------------
String          : '\"' Close _          ;
        Close   : '\"' 
                / CHAR Close            ;

CHAR            : UpSlash | '\\' ChCode ;
        ChCode  : Escs 
                | "u" HEX HEX HEX HEX   ;
        Escs    : '\\' 
                | '\"' 
                | '/' 
                | 'b' 
                | 'f' 
                | 'n' 
                | 'r' 
                | 't'                   ;
        UpSlash : '^' 
                | '\\'                  ;

HEX             : < Number aA-fF >      ;
        aA-fF   : any "abcdefABCDEF"    ;  

Number          : INT OptFrac OptExp _  ;
        OptFrac : [ FRAC ]              ;
        OptExp  : [ EXP ]               ;

INT             : Neg Ints              ;
        Ints    : ( NotZero OptNums 
                | '0' )                 ; 
        Neg     : [ '-' ]               ;

FRAC            : '.' Numbers1x         ;

EXP             : eE PlusMinus Numbers1x;
        NotZero : not '0' number        ;
        OptNums : { Numbers1x }         ;
      Numbers1x : < number >            ;
      PlusMinus : [ '+' | '-' ]         ;
        eE      : 'e' | 'E'             ;

//-------------------------------------------------------------
// Separators and Operators
//-------------------------------------------------------------
TRUE            : "true" _              ;
FALSE           : "false" _             ;
NUL             : "null" _              ;
COMMA           : ',' _                 ;
COLON           : ':' _                 ;
LBRACE          : '{' _                 ;
RBRACE          : '}' _                 ;
LBRACKET        : '[' _                 ;
RBRACKET        : ']' _                 ;

//-------------------------------------------------------------
//  Escape Character and Whitespace Sequences
//-------------------------------------------------------------
_               : { EscChar 
                | BlockComment 
                | Comment }             ;

EscCharSpace    : < ' ' 
                | EscChar >              ;

EscChar         : '\t' 
                | newLine               ; 

BlockComment    : "*/" 
                / "/*" BlockComment     ;

Comment         : newLine 
                / "//" Comment          ; 

newLine         : '\r' 
                | '\n'                  ;

```
