# **`miniegg` Grammar**
### **AUTHORSHIP INFORMATION**
#### *Author :* Brynn Harrington and Emily Hoppe Copyright (C) 2021
#### *Adapted from :* Aaron Moss's [`eggr` Egg Grammar](https://github.com/bruceiv/egg/blob/deriv/grammars/miniegg.egg)
#### *Creation Date :* June 11, 2021 
#### *Last Modified :* June 18, 2021
#### *Copyright and Licensing Information :* See end of file.

###  **GENERAL DESCRIPTION**
An originally Egg Parsing grammar created by Aaron Moss ported into the GoGLL grammar to determine test a given structure. Modification of `miniegg` grammar from [Egg](https://github.com/bruceiv/egg/blob/deriv/grammars/miniegg.egg) to test an example structure.

### **`miniegg` Grammar Guide**
Errors:
- not sure if id is functioning the way it should - going based on how basic IDs work

See the [grammar for details.](../../gogll.md)

### **STATUS ON GRAMMAR**
#### *Markdown File Creation:* Not working 
#### *Parser Generated :* Incomplete - parser generated from partially complete md
#### *Test File Creation:* Incomplete
#### *Testing Results:* Unknown
```
package "miniegg"

Expr    : Id neq ; 
Id      : upC Space ; 

upC     : upcase ;
Space   :  " "  
        | empty 
        ; 
    eq      : '=' ; 
    neq     : not "=" ;

```


### **IN PROGRESS GRAMMARS**
**Original / Not working**
    gram    : _ ruler ; 
    ruler   : < rule > ;
    rule    : id '=' _ exprr ; L
    id      : upcase _ ; L
        exprr   : < expr > ; L
    expr    : id not '=' ; S
    _       : [ ' ' ] ; L
**Partially Working**
Expr    : Id neq ; 
Id      : upC Space ; 

upC     : upcase ;
Space   :  " "  
        | empty 
        ; 
eq      : '=' ; 
neq     : not "=" ;
#
### **COPYRIGHT AND LICENSING INFORMATION**
**Copyright 2021 Brynn Harrington and Emily Hoppe**

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.

You may obtain a copy of the License [here](http://www.apache.org/licenses/LICENSE-2.0) or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.