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
ERRORS: 
- decimal numeral not working correctly 

See the [grammar for details.](../../gogll.md)

### **STATUS ON GRAMMAR**
#### *Markdown File Creation:* Not working 
#### *Parser Generated :* Incomplete
#### *Test File Creation:* Incomplete
#### *Testing Results:* Unknown
```
package "Java"

```
#### ***General Numeric Literals***
- Note: In IntegerLiteral, OctalNumeral may prefix 
HexNumeral and DecimalNumeral may prefix OctalNumeral
```

FloatLiteral      : HexFloat 
                  | DecimalFloat                      ;

IntegerLiteral    : NumeralAlts optOneL               ;
      NumeralAlts : HexNumeral 
                  | OctalNumeral  
                  | DecimalNumeral                    ;
      optOneL     : [ any "1L" ]                      ;

DecimalFloat      :  repDig1x dot repDig0x optExpo fF_dD 
                  | dot RepDig1xExp 
                  | RepDig1xExp fF_dD
                  | RepDig1xOptExp fF_dD                   ;
   RepDig1xOptExp : repDig1x optExpo                       ;
      RepDig1xExp : repDig1x Exponent                      ;
      optExpo     : [ any "eE" { number } [ any "+\\-" ] ] ;

```

#### ***BASE-SIXTEEN AND BASE-EIGHT LITERALS***
Incomplete decimalnumeral!!!
```
DecimalNumeral  : ze 
                | 1thru9 repNumx0                   ;
    repNumx0    : number repNumx0 / empty           ;
    1thru9      : any "123456789"                   ;

HexFloat          : HexSignificand BinaryExponent fF_dD ;
      fF_dD       : [ any "fFdD" ]                      ; 

HexSignificand    : HexNumeral optDot 
                  | RepHex0xDot hexDigit RepHex0x   ;
      RepHex0xDot : zeroxX RepHex0x dot             ;
      optDot      : [ '.' ]                         ; 
      dot         : '.'                             ;

HexNumeral        : zeroxX hexDigit RepHex0x        ; 
      zeroxX      : any "0xX"                       ; 
      RepHex0x : hexDigit RepHex0x / empty          ;

hexDigit        : < number any "abcdefABCDEF" >     ;

OctalNumeral    : ze Int07 Rep07x1                  ; 
      ze        : '0'                               ;
      Rep07x1   : Int07 Rep07x1 
                / empty                             ;

OctalEscape     : Int03Two07
                / Two07
                / Int07                             ;
    Int03Two07  : int03 Two07                       ;
      Two07     : Int07 Int07                       ;
      Int07     : int03
                | any4567                           ;
      any4567   : any "4567"                        ;
      int03     : any "0123"                        ;
```
#### ***Exponent and Digital Literals***
Original Egg grammar had a NT "Digit", which is replaced here in GoGll by the reserved word "number".
```
Exponent        : eE optPSM repDig0x                    ;
      eE        : any "eE"                              ;
      repDig0x  : { number }                            ; 
      optPSM    : [ any "+\\-" ]                        ;

BinaryExponent  : pP psm repDig1x                       ;
    pP          : any "pP"                              ;
    psm         : any "+\\-"                            ;
    repDig1x    : < number >                            ;        

```   

#### ***Separators and Operators***
```
AT                  :  "@"            WS                 ;
AND                 :  "&" notEqAnd   WS                 ;
AND_AND             :  "&&"           WS                 ;
AND_EQU             :  "&="           WS                 ;
BANG                :  "!" nEq        WS                 ;
BSR                 :  ">>>" nEq      WS                 ;
BSR_EQU             :  ">>>="         WS                 ;
COLON               :  ":"            WS                 ;
COMMA               :  ","            WS                 ;
DEC                 :  "--"           WS                 ;
DIV                 :  "/" nEq        WS                 ;
DIV_EQU             :  "/="           WS                 ;
DOT                 :  "."            WS                 ;
EQU                 :  "=" nEq        WS                 ;
EQUAL               :  "=="           WS                 ;
GE                  :  ">="           WS                 ;
GT                  :  ">" notEqCar   WS                 ;
HAT                 :  "^" nEq        WS                 ;
HAT_EQU             :  "^="           WS                 ;

notEqAnd            :  not "=" not "&"                   ;

INC                 :  "++"           WS                 ;
LBRK                :  "["            WS                 ;
LE                  :  "<="           WS                 ;
LPAR                :  "("            WS                 ;
LPOINT              :  "<"            WS                 ;
LT                  :  "<" notEqCar2  WS                 ;
LWING               :  "{"            WS                 ;
MINUS               : "-" notEqSlDash WS                 ;
MINUS_EQU           :  "-="           WS                 ;
MOD                 :  "%" nEq        WS                 ;
MOD_EQU             :  "%="           WS                 ;
NOT_EQUAL           :  "!="           WS                 ;   
OR                  :  "|" notEqPipe  WS                 ;
OR_EQU              :  "|="           WS                 ;
OR_OR               :  "||"           WS                 ;
PLUS                :  "+" notEqPlus  WS                 ;
PLUS_EQU            :  "+="           WS                 ;

notEqPipe           :  not "=" not "|"                   ; 
notEqPlus           :  not "=" not "+"                   ;
notEqCar2           :  not "=" not "<"                   ;
notEqSlDash         :  not "=" not "\\" not "-"          ;

QUERY               :  "?"            WS                 ;
RBRK                :  "]"            WS                 ;
RPAR                :  ")"            WS                 ;
RPOINT              :  ">"            WS                 ;
RWING               :  "}"            WS                 ;
SEMI                :  ";"            WS                 ;
SL                  :  "<<" nEq       WS                 ;
SL_EQU              :  "<<="          WS                 ;
SR                  :  ">>" notEqCar  WS                 ;
SR_EQU              :  ">>="          WS                 ;
STAR                :  "*" nEq        WS                 ;
STAR_EQU            :  "*="           WS                 ;
TILDA               :  "~"            WS                 ;
    

notEqCar            :  not "=" not ">"                   ; 
nEq                 : not "="                            ;


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
      newline       : any "\r\n"                       ;

```
### **COPYRIGHT AND LICENSING INFORMATION**
**Copyright 2021 Brynn Harrington and Emily Hoppe**

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.

You may obtain a copy of the License [here](http://www.apache.org/licenses/LICENSE-2.0) or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.