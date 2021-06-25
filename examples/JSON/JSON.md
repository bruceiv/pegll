# **`JSON` GRAMMAR**
### **AUTHORSHIP INFORMATION**
#### *Authors :* Brynn Harrington and Emily Hoppe Copyright (C) 2021
#### *Adapted from :* Aaron Moss's [`JSON` Egg Grammar](https://github.com/bruceiv/egg/blob/deriv/grammars/JSON-u.egg)
#### *Creation Date :* June 11, 2021 
#### *Last Modified :* June 24, 2021
#### *Copyright and Licensing Information :* See end of file.

###  **GENERAL DESCRIPTION**
Aa modification of `JSON` grammar from [Egg](https://github.com/bruceiv/egg/blob/deriv/grammars/JSON-u.egg) parsing grammar ported into GoGLL to test `JSON` input files under the parser generated.
### **STATUS ON GRAMMAR**
#### *Markdown File Creation:* Working 
#### *Parser Generated :* Complete
#### *Test File Creation:* Incomplete
#### *Testing Results:* Unknown
### **`JSON` GRAMMAR GUIDE**
The following grammar tests input files from the `JSON` language.
```
package "JSON" 
```
#### ***Higher-Level Language Structures***
The following are the GoGLL representations of the higher level JSON components.
```
JSON            : WS Object                             ;

Object          : LBRACE Members RepMems0x RBRACE       ;
     RepMems0x  : Members RepMems0x
                / empty                                 ;

Members         : Pair RepComPair0x                     ;
   RepComPair0x : COMMA Pair RepComPair0x  
                / empty                                 ; 

Pair            : String COLON Value                    ;

Array           : LBRACKET OptElem RBRACKET             ;
        OptElem : Elements 
                / empty                                 ;

Elements        : Value RepComVal0x                     ;
    RepComVal0x : COMMA Value RepComVal0x
                / empty                                 ; 

Value           : String 
                | Number 
                | Object 
                | Array 
                | TRUE 
                | FALSE 
                | NUL                                   ;
```  
#### ***String and Character Literals***
The following are the GoGLL representations of the JSON string and character literals.
```
String          : dQuote Close WS                       ;
        Close   : dQuote
                / CHAR Close                            ;
    dQuote      : any "\""                              ;

CHAR            : carrotSlash 
                | bSlash CharCode                       ;  
        bSlash  : '\\'                                  ;
       CharCode : esc
                | "u" HEX HEX HEX HEX                   ;
        esc     : any "\\\"/bfnrt"                      ;
    carrotSlash : any "^\\"                             ;        
```
#### ***Numeric Literals***
The following are the GoGLL representations of the JSON numeric literals.
```
 HEX            : Number aA_fF 
                | empty                                 ;
        aA_fF   : any "abcdefABCDEF"                    ; 
        
Number          : INT OptFrac OptExp WS                 ;
        OptFrac : frac
                | empty                                 ;
        OptExp  : exp
                | empty                                 ;

INT             : optNeg Integers                       ;
       Integers : integer
                / zero                                  ;
        zero    : any "0"                               ;
        integer : any "123456789" { < number > }        ;
        optNeg  : [ '-' ]                               ;
                       
frac            : any "." < number >                    ;

exp             : any "eE" [ any "+-" ] < number >      ;  

```
#### ***Operators and Special Characters***
The following are the GoGLL representations of the JSON operators and special characters.
```
TRUE            : "true"   WS                           ;
FALSE           : "false"  WS                           ;
NUL             : "null"   WS                           ;
COMMA           : ","      WS                           ;
COLON           : ":"      WS                           ;
LBRACE          : "{"      WS                           ;
RBRACE          : "}"      WS                           ;
LBRACKET        : "["      WS                           ;              
RBRACKET        : "]"      WS                           ;
```
#### ***Whitespace and Escape Sequences***
The following are the GoGLL representations of the JSON whitespace and escape sequences.
###### *Note:* `!line_comment` and `!block_comment` were taken from Ackerman's [comments.md.](https://github.com/bruceiv/pegll/tree/main/examples/comments) 
```
WS              : EscOrComment WS
                / empty                                 ;

EscOrComment    : escChar 
                | LineOrBlock                           ;
                
escCharSpace    : < any " \t\r\n" >                     ;

escChar         : any "\t\r\n"                          ; 

LineOrBlock     : line_comment 
                | block_comment                         ;

!line_comment   : '/' '/' { not "\n" }                  ;               

!block_comment  : '/''*' 
                { not "*" 
                | '*' not "/" 
                } '*''/'                                ;

newLine         : any "\r\n"                            ;                 
```
#
### **COPYRIGHT AND LICENSING INFORMATION**
**Copyright 2021 Brynn Harrington and Emily Hoppe**

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.

You may obtain a copy of the License [here](http://www.apache.org/licenses/LICENSE-2.0) or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.