# **`JSON` Grammar**
### **AUTHORSHIP INFORMATION**
#### *Authors :* Brynn Harrington and Emily Hoppe Copyright (C) 2021
#### *Adapted from :* Aaron Moss's [`JSON` Egg Grammar](https://github.com/bruceiv/egg/blob/deriv/grammars/JSON-u.egg)
#### *Creation Date :* June 11, 2021 
#### *Last Modified :* June 18, 2021
#### *Copyright and Licensing Information :* See end of file.

###  **GENERAL DESCRIPTION**
An originally Egg Parsing grammar created by Aaron Moss ported into the GoGLL grammar for the `JSON` language tests. Modification of `JSON` grammar from [Egg](https://github.com/bruceiv/egg/blob/deriv/grammars/JSON-u.egg) to test `JSON` input files under the parser generated.
### **STATUS ON GRAMMAR**
#### *Markdown File Creation:* Not working 
#### *Parser Generated :* Incomplete
#### *Test File Creation:* Incomplete
#### *Testing Results:* Unknown
### **`JSON` Grammar Guide**
ERRORS
- COME BACK TO FINISH OPTIONAL PORTION OF HEX
- FIX ESCAPES UNDER STRING/CHAR LITS
```
package "JSON" 
```  
#### ***String and Character Literals***
```


```
#### ***Numeric Literals***
```
 HEX:   NumberHEX;
        NumberHEX : Number aA_fF 
                | empty ;
        aA_fF   : any "abcdefABCDEF"    ; 
        
Number          : INT optFrac optExp WS ;
        optFrac : [ any "." < any "0123456789" >  ] ;
        optExp  : [ any "eE" 
                [ any "+-" ] 
                < any "0123456789" > ] ;
INT     : neg Integers ;
        Integers: integer
                | zero ;
        zero    : any "0";
        integer    : any "123456789" { < any "0123456789" > } ;
        neg     : [ '-' ]               ;
FRAC    : "." numbers1x         ;
EXP     : eE plusMinus numbers1x;  
        numbers1x : < any "0123456789" > ;
        plusMinus : [ any "+-" ] ;
        eE      : any "eE" ;
```
#### ***Operators and Special Characters***
```
TRUE            : "true" WS ;
FALSE           : "false" WS ;
NUL             : "null" WS ;
COMMA           : "," WS ;
COLON           : ":" WS ;
LBRACE          : "{" WS ;
RBRACE          : "}" WS ;
LBRACKET        : "[" WS ;
RBRACKET        : "]" WS ;
```
#### ***Whitespace and Escape Sequences***
```
WS              : EscOrComment WS
                | empty ;

EscOrComment    : escChar 
                | LineOrBlock ;
escCharSpace    : < any " \t\r\n" > ;
escChar         : any "\t\r\n" ; 

LineOrBlock     : line_comment 
                | block_comment ;
!line_comment   : '/' '/' { not "\n" } ;
!block_comment  : '/''*' 
                { not "*" 
                | '*' not "/" 
                } '*''/' ;

newLine         : any "\r\n" ;                 
```
#### ORIGINAL GRAMMAR
#### ***Object Creation***
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
#### ***String and Character Literals***
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
#### ***Numeric Literals***
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
#### ***Operators and Special Characters***
TRUE            : "true" _              ;
FALSE           : "false" _             ;
NUL             : "null" _              ;
COMMA           : ',' _                 ;
COLON           : ':' _                 ;
LBRACE          : '{' _                 ;
RBRACE          : '}' _                 ;
LBRACKET        : '[' _                 ;
RBRACKET        : ']' _                 ;
#### ***Whitespace and Escape Sequences***
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
#### PARTIALLY WORKING GRAMMAR
 HEX:   NumberHEX;
        NumberHEX : Number aA_fF 
                | empty ;
        aA_fF   : any "abcdefABCDEF"    ; 
Number          : INT optFrac optExp WS ;
        optFrac : [ any "." < any "0123456789" >  ] ;
        optExp  : [ any "eE" [ any "+-" ] < any "0123456789" > ] ;
INT     : neg Integers ;
        Integers: integer
                | zero ;
        zero    : any "0";
        integer    : any "123456789" { < any "0123456789" > } ;
        neg     : [ '-' ]               ;
FRAC    : "." numbers1x         ;
EXP     : eE plusMinus numbers1x;  
        optNumbers : { < any "0123456789" > } ;
        numbers1x : < any "0123456789" > ;
        plusMinus : [ any "+-" ] ;
        eE      : any "eE" ;

TRUE            : "true" WS ;
FALSE           : "false" WS ;
NUL             : "null" WS ;
COMMA           : "," WS ;
COLON           : ":" WS ;
LBRACE          : "{" WS ;
RBRACE          : "}" WS ;
LBRACKET        : "[" WS ;
RBRACKET        : "]" WS ;

WS              : EscOrComment WS
                | empty ;

EscOrComment    : escChar 
                | LineOrBlock ;
escCharSpace    : < any " \t\r\n" > ;
escChar         : any "\t\r\n" ; 

LineOrBlock     : line_comment 
                | block_comment ;
!line_comment   : '/' '/' { not "\n" } ;
!block_comment  : '/''*' 
                { not "*" 
                | '*' not "/" 
                } '*''/' ;

newLine         : any "\r\n" ;  
#
### **COPYRIGHT AND LICENSING INFORMATION**
**Copyright 2021 Brynn Harrington and Emily Hoppe**

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.

You may obtain a copy of the License [here](http://www.apache.org/licenses/LICENSE-2.0) or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.