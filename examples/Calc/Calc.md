# **`calc` Grammar**

### **AUTHORSHIP INFORMATION**
#### *Authors :* Emily Hoppe Copyright (C) 2021
#### *Adapted from :* Aaron Moss's [`Calc` Egg Grammar](https://github.com/bruceiv/egg/blob/deriv/grammars/Calc.egg)
#### *Creation Date :* June 17, 2021 
#### *Last Modified :* June 22, 2021
#### *Copyright and Licensing Information :* See end of file.

###  **GENERAL DESCRIPTION**
An originally Egg Parsing grammar created by Aaron Moss ported into the GoGLL grammar to test a simple calculator. Modification of `calc` grammar from [Egg](https://github.com/bruceiv/egg/blob/deriv/grammars/Calc.egg) to calculate based on given inputs.

### **`calc` Grammar Guide**
ERRORS:
- cannot figure out how to use a whitespace token ID 
- cannot get past element - asking for EOF token in rule
- cannot figure out repetitions for PROD and SUM
       - may need to have an ordered choice to use the '{}' grouping 

See the [grammar for details.](../../gogll.md)

### **STATUS ON GRAMMAR**
#### *Markdown File Creation:* Working 
#### *Parser Generated :* Complete
#### *Test File Creation:* Incomplete
#### *Testing Results:* Unknown
```
package "calc"

EXPR   : space SUM ;

SUM    : PROD PoMRep ;

PoMRep      : PLUSorMINUS PoMRep / empty ;
PLUSorMINUS : PLUS PROD 
              | MINUS  ;

PROD   : ELEM ToDRep ;

ToDRep        : TIMESorDIVIDE ToDRep / empty ;
TIMESorDIVIDE : TIMES ELEM 
              | DIVIDE ELEM  ;
              
ELEM   : OPEN SUM CLOSE 
       | num ;

num    : < number > { ' ' | '\t' } ;

PLUS   : "+" space ;
MINUS  : "-" space ;
TIMES  : "*" space ;
DIVIDE : "/" space ;
OPEN   : "(" space ;        
CLOSE  : ")" space ;

space  : { ' ' | '\t' } ;
```
### **IN PROGRESS GRAMMARS**
**Original / Not working**
expr :  _ sum ;

sum  : prod { PLUS prod | MINUS } ;

prod : elem { TIMES elem | DIVIDE elem } ;

elem : OPEN sum CLOSE | num ;

num  :  < number > _ ;

PLUS   : '+' _ ;
MINUS  : '-' _ ;
TIMES  : '*' _ ; 
DIVIDE : '/' - ;
OPEN   : '(' _ ; 
CLOSE  : ')' _ ;

_      : { ' ' | '\t' } ;
**Partially Working**
EXPR   : SUM ;

SUM    : PROD plus PROD | minus  ;
PROD   : ELEM times ELEM | divide ELEM  ;
ELEM   : open SUM close | num ;

num    : < number > { ' ' | '\t' } ;

plus   : '+' { ' ' | '\t' } ;
minus  : '-' { ' ' | '\t' } ;
times  : '*' { ' ' | '\t' } ;
divide : '/' { ' ' | '\t' } ;
open   : '(' { ' ' | '\t' } ;        
close  : ')' { ' ' | '\t' } ;
#
### **COPYRIGHT AND LICENSING INFORMATION**
**Copyright 2021 Brynn Harrington and Emily Hoppe**

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.

You may obtain a copy of the License [here](http://www.apache.org/licenses/LICENSE-2.0) or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.