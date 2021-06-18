# **`Calc` Grammar**

### **AUTHORSHIP INFORMATION**
#### *Authors :* Emily Hoppe Copyright (C) 2021
#### *Creation Date :* June 17, 2021 
#### *Last Modified :* June 18, 2021
#### *Copyright and Licensing Information :* See end of file.

###  **GENERAL DESCRIPTION**
An originally Egg Parsing grammar created by Dr. Aaron Moss ported into the GoGLL grammar to test a simple calculator. Modification of `Calc` grammar from [Egg](https://github.com/bruceiv/egg/blob/deriv/grammars/Calc.egg) to calculate based on given inputs.

### **`Calc` Grammar Guide**
NEED TO FINISH ONE GRAMMAR IS WORKING 

 See the [grammar for details.](../../gogll.md)
```
package "Calc"

expr :  _ sum ;

sum  : prod { PLUS prod | MINUS } ;

prod : elem { TIMES elem | DIVIDE elem } ;

elem : OPEN sum CLOSE | num ;

num  :  < number > _ ;

PLUS   : '+' _ ;
MINUS  : '-' _ ;
TIMES  : '*' _ ; 
DIVIDE : '/' _ ;
OPEN   : '(' _ ; 
CLOSE  : ')' _ ;

_      : { ' ' 
       |   '\t' } ;

```
#
### **COPYRIGHT AND LICENSING INFORMATION**
**Copyright 2021 Brynn Harrington and Emily Hoppe**

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.

You may obtain a copy of the License [here](http://www.apache.org/licenses/LICENSE-2.0) or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.