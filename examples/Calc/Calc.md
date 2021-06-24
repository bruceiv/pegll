# **`calc` GRAMMAR**

### **AUTHORSHIP INFORMATION**
#### *Authors :* Emily Hoppe Copyright (C) 2021
#### *Adapted from :* Aaron Moss's [`Calc` Egg Grammar](https://github.com/bruceiv/egg/blob/deriv/grammars/Calc.egg)
#### *Creation Date :* June 17, 2021 
#### *Last Modified :* June 24, 2021
#### *Copyright and Licensing Information :* See end of file.

###  **GENERAL DESCRIPTION**
A modification of `calc` grammar from [Egg](https://github.com/bruceiv/egg/blob/deriv/grammars/Calc.egg) ported into GoGLL to test a simple calculator.

### **STATUS ON GRAMMAR**
#### *Markdown File Creation:* Working 
#### *Parser Generated :* Complete
#### *Test File Creation:* Incomplete
#### *Testing Results:* Unknown

### **`calc` GRAMMAR GUIDE**
The following grammar tests simple calculations, with order of operations under consideration, based on a given input.
```
package "calc"
```
`EXPR` represents the starting rule for the grammar being a semantic rule composed of a space followed by a `SUM`.
```
EXPR             : space SUM                     ;
```
The following section is composed of `SUM`, `RepPLUSorMINUS0x`, and `PLUSorMINUS`, where:
- `SUM` is a semantic rule matched with `PRODUCT` followed by `RepPLUSorMINUS0x`;
- `RepPLUSorMINUS0x` is a semantic rule matched by zero or more repetitions of `PLUSorMINUS`;
- `PLUSorMINUS` is a semantic rule matched by `PLUS`, addition, or `MINUS`, subtraction, of a `PRODUCT`.
```
SUM              : PRODUCT RepPLUSorMINUS0x      ;
RepPLUSorMINUS0x : PLUSorMINUS RepPLUSorMINUS0x 
                 / empty                         ;
PLUSorMINUS      : PLUS PRODUCT 
                 | MINUS PRODUCT                 ; 
```
The following section is composed of `PRODUCT`, `RepTIMESorDIV0x`, and `TIMESorDIV`, where:
- `PRODUCT` is a semantic rule matched with `ELEMENT` followed by `RepTIMESorDIV0x`;
- `RepTIMESorDIV0x` is a semantic rule matched by zero or more repetitions of `TIMESorDIV`;
- `TIMESorDIV` is a semantic rule matched by `TIMES`, multiplication, or `DIVIDE`, division, of an `ELEMENT`.
```
PRODUCT          : ELEMENT RepTIMESorDIV0x       ;
RepTIMESorDIV0x  : TIMESorDIVIDE RepTIMESorDIV0x 
                 / empty                         ;
TIMESorDIVIDE    : TIMES ELEMENT  
                 | DIVIDE ELEMENT                ;
```
The following section is composed of `ELEMENT`, `Number`, and `repNumber1x`, where:
- `ELEMENT` is a semantic rule matched with `SUM` enclosed by `OPEN` and `CLOSE` or a number;
- `Number` is a semantic rule matched by `repNumber1x` followed by a space;
- `repNumber1x` is a lexical rule matched by a `number` repeated one or more times.
For more information about the `number` reserved word, see the [grammar for details.](../../gogll.md)
```       
ELEMENT          : OPEN SUM CLOSE 
                 | Number                        ;
Number           : repNumber1x space             ;
repNumber1x      : < number >                    ;
```
The following section is composed of `PLUS`, `MINUS`, `TIMES`, `DIVIDE`, `OPEN`, and `CLOSE`, where:
- `PLUS` is a semantic rule matched with a '+' character followed by a space representing the addition operation;
- `MINUS` is a semantic rule matched with a '-' character followed by a space representing the subtraction operation;
- `TIMES` is a semantic rule matched with a '*' character followed by a space representing the multiplication operation;
- `DIVIDE` is a semantic rule matched with a '/' character followed by a space representing the division operation;
- `OPEN` is a semantic rule matched with a '(' character followed by a space;
- `CLOSE` is a semantic rule matched with a ')' character followed by a space.
```
PLUS             : "+" space                     ;
MINUS            : "-" space                     ;
TIMES            : "*" space                     ;
DIVIDE           : "/" space                     ;
OPEN             : "(" space                     ;                
CLOSE            : ")" space                     ;
```
`space` is a lexical rule composed of the whitespace characters ' ' and '\t'. It may be repeated zero or more times as defined by the `{}` in the GoGLL grammar. See the [grammar for details.](../../gogll.md)
```
space            : { ' ' | '\t' }                ;
```
#
### **COPYRIGHT AND LICENSING INFORMATION**
**Copyright 2021 Brynn Harrington and Emily Hoppe**

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.

You may obtain a copy of the License [here](http://www.apache.org/licenses/LICENSE-2.0) or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.