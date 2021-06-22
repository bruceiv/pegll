# **`Java` Grammar**
### **AUTHORSHIP INFORMATION**
#### *Authors :* Brynn Harrington and Emily Hoppe Copyright (C) 2021
#### *Adapted from :* Aaron Moss"s [`Java` Egg Grammar](https:github.com/bruceiv/egg/blob/deriv/grammars/Java-u.egg) and Roman Reziejowski's [`Java` Mouse Parser-Generator](http://home.swipnet.se/redz/roman)
#### *Creation Date :* June 11, 2021 
#### *Last Modified :* June 22, 2021
#### *Copyright and Licensing Information :* See end of file.

###  **GENERAL DESCRIPTION**
An originally Egg Parsing grammar created by Aaron Moss ported into the GoGLL grammar for the `Java` language tests. Modification of `Java` grammar from [Egg](https:github.com/bruceiv/egg/blob/deriv/grammars/Java-u.egg) to test `Java` input files under the parser generated.

### **`Java` Grammar Guide**
NEED TO FINISH ONE GRAMMAR IS WORKING 

See the [grammar for details.](../../gogll.md)

### **STATUS ON GRAMMAR**
#### *Markdown File Creation:* Not working 
#### *Parser Generated :* Incomplete
#### *Test File Creation:* Incomplete
#### *Testing Results:* Unknown
```
package "Java"

```
#### ***Separators and Operators***
```

INC               :  "++"           _                 ;
LBRK              :  "["            _                 ;
LE                :  "<="           _                 ;
LPAR              :  "("            _                 ;
LPOINT            :  "<"            _                 ;
LT                :  "<"![=<]       _                 ;
LWING             :  "{"            _                 ;
MINUS             :  "-"![=\-]      _                 ;
MINUS_EQU         :  "-="           _                 ;
MOD               :  "%"!"="        _                 ;
MOD_EQU           :  "%="           _                 ;
NOT_EQUAL         :  "!="           _                 ;   
OR                :  "|"![=|]       _                 ;
OR_EQU            :  "|="           _                 ;
OR_OR             :  "||"           _                 ;
PLUS              :  "+"![=+]       _                 ;
PLUS_EQU          :  "+="           _                 ;


QUERY             :  "?"            WS                 ;
RBRK              :  "]"            WS                 ;
RPAR              :  ")"            WS                 ;
RPOINT            :  ">"            WS                 ;
RWING             :  "}"            WS                 ;
SEMI              :  ";"            WS                 ;
SL                :  "<<" nEq       WS                 ;
SL_EQU            :  "<<="          WS                 ;
SR                :  ">>" notEqCar  WS                 ;
SR_EQU            :  ">>="          WS                 ;
STAR              :  "*" nEq        WS                 ;
STAR_EQU          :  "*="           WS                 ;
TILDA             :  "~"            WS                 ;

notEqCar          :  not "=" not ">"                   ; 
nEq               : not "="                            ;


```
### ***Escape Characters/Sequences, Comments, and Spacing***
- Note: To match the -> operator in GoGLL, the following syntax is used:
    
    (Egg): XtoY : X -> Y
    
    (GoGLL): XtoY : Y / X XtoY;
    
```
WS                : EscOrLineOrBlock     
                  | empty                             ;
EscOrLineOrBlock  : line_comment 
                  | block_comment                     
                  | escCharSp                         ;
      escCharSp   : < any " \t\r\n" >                 ;
      
!line_comment : '/' '/' {not "\n"} ;
!block_comment : '/''*' {not "*" | '*' not "/"} '*''/' ;
      newline       : any "\r\n"                        ;

```
### **COPYRIGHT AND LICENSING INFORMATION**
**Copyright 2021 Brynn Harrington and Emily Hoppe**

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.

You may obtain a copy of the License [here](http://www.apache.org/licenses/LICENSE-2.0) or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.