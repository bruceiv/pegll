# **`JSON` Grammar**
### **AUTHORSHIP INFORMATION**
#### *Authors :* Brynn Harrington and Emily Hoppe Copyright (C) 2021
#### *Adapted from :* Aaron Moss's [`JSON` Egg Grammar](https://github.com/bruceiv/egg/blob/deriv/grammars/JSON-u.egg)
#### *Creation Date :* June 11, 2021 
#### *Last Modified :* June 18, 2021
#### *Copyright and Licensing Information :* See end of file.

###  **GENERAL DESCRIPTION**
An originally Egg Parsing grammar created by Dr. Aaron Moss ported into the GoGLL grammar for the `JSON` language tests. Modification of `JSON` grammar from [Egg](https://github.com/bruceiv/egg/blob/deriv/grammars/JSON-u.egg) to test `JSON` input files under the parser generated.
### **`JSON` Grammar Guide**
NEED TO FINISH ONE GRAMMAR IS WORKING 
#### ***Object Creation***
```
package "JSON"                          ;
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
```
#### ***String and Character Literals***
```
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
```
#### ***Numeric Literals***
```
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
```
#### ***Operators and Special Characters***
```
TRUE            : "true" _              ;
FALSE           : "false" _             ;
NUL             : "null" _              ;
COMMA           : ',' _                 ;
COLON           : ':' _                 ;
LBRACE          : '{' _                 ;
RBRACE          : '}' _                 ;
LBRACKET        : '[' _                 ;
RBRACKET        : ']' _                 ;
```
#### ***Whitespace and Escape Sequences***
```
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
#
### **COPYRIGHT AND LICENSING INFORMATION**
**Copyright 2021 Brynn Harrington and Emily Hoppe**

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.

You may obtain a copy of the License [here](http://www.apache.org/licenses/LICENSE-2.0) or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.