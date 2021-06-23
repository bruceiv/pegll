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


JSON            : WS Object                     ;

Object          : LBRACE Members Mems1x RBRACE  ;
        Mems1x  : Members Mems1x
                / empty                         ;

Members         : Pair ComPair0x                ;
      ComPair0x : ComPair ComPair0x
                / empty                         ; 
        ComPair  : COMMA Pair                   ;

Pair            : String COLON Value            ;

Array           : LBRACKET OptElem RBRACKET     ;
        OptElem : Elements 
                / empty                         ;

Elements        : Value ComVal0x                ;
       ComVal0x : ComVal ComVal0x
                / empty                         ; 
        ComVal  : COMMA Value                   ;

Value           : String 
                | Number 
                | Object 
                | Array 
                | TRUE 
                | FALSE 
                | NUL                           ;
```  
#### ***String and Character Literals***
```
String          : dQuote Close WS       ;
        Close   : dQuote
                / CHAR Close            ;
    dQuote      : any "\""              ;

CHAR            : carrot 
                | bSlash CharCode       ;  
        bSlash  : '\\'                  ;
       CharCode : esc
                | "u" HEX HEX HEX HEX   ;
        esc     : any "\\\"/bfnrt"      ;
        carrot : any "^\\"              ;        
```
#### ***Numeric Literals***
```
 HEX            :   NumberHEX                                   ;
      NumberHEX : Number aA_fF 
                | empty                                         ;
        aA_fF   : any "abcdefABCDEF"                            ; 
        
Number          : INT OptFrac OptExp WS                         ;
        OptFrac : frac
                | empty                                         ;
        OptExp  : exp
                | empty                                         ;

INT             : optNeg Integers                               ;
       Integers : integer
                / zero                                          ;
        zero    : any "0"                                       ;
        integer : any "123456789" { < any "0123456789" > }      ;
        optNeg  : [ '-' ]                                       ;
                       
frac            : any "." < any "0123456789" >                  ;

exp             : any "eE" [ any "+-" ] < any "0123456789" >    ;  

```
#### ***Operators and Special Characters***
```
TRUE            : "true"   WS           ;
FALSE           : "false"  WS           ;
NUL             : "null"   WS           ;
COMMA           : ","      WS           ;
COLON           : ":"      WS           ;
LBRACE          : "{"      WS           ;
RBRACE          : "}"      WS           ;
LBRACKET        : "["      WS           ;              
RBRACKET        : "]"      WS           ;
```
#### ***Whitespace and Escape Sequences***
```
WS              : EscOrComment WS
                | empty                 ;

EscOrComment    : escChar 
                | LineOrBlock           ;
                
escCharSpace    : < any " \t\r\n" >     ;

escChar         : any "\t\r\n"          ; 

LineOrBlock     : line_comment 
                | block_comment         ;

!line_comment   : '/' '/' { not "\n" }  ;

!block_comment  : '/''*' 
                { not "*" 
                | '*' not "/" 
                } '*''/'                ;

newLine         : any "\r\n"            ;                 
```
#
### **COPYRIGHT AND LICENSING INFORMATION**
**Copyright 2021 Brynn Harrington and Emily Hoppe**

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.

You may obtain a copy of the License [here](http://www.apache.org/licenses/LICENSE-2.0) or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.