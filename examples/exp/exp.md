# **`exp` Grammar**

### **AUTHORSHIP INFORMATION**
#### *Author :* Brynn Harrington and Emily Hoppe Copyright (C) 2021
#### *Adapted from :* Aaron Moss's [`eggr` Egg Grammar](https://github.com/bruceiv/egg/blob/deriv/grammars/exp.egg)
#### *Creation Date :* June 11, 2021 
#### *Last Modified :* June 18, 2021
#### *Copyright and Licensing Information :* See end of file.

###  **GENERAL DESCRIPTION**
An originally Egg Parsing grammar created by Aaron Moss ported into the GoGLL grammar to determine if a grammar, under recursive descent, will have exponential runtime. Modification of `exp` grammar from [Egg](https://github.com/bruceiv/egg/blob/deriv/grammars/exp.egg) to test an example grammar with possible exponential runtime under recursive descent.
### **STATUS ON GRAMMAR**
#### *Markdown File Creation:* Working
#### *Parser Generated :* Complete
#### *Test File Creation:* Incomplete
#### *Testing Results:* Unknown
### **`exp` GRAMMAR GUIDE**
The following grammar will have exponential runtime under recursive descent. It utilizes three lexical rules to act ask token identifiers with one starting rule exercising the recursive descent. 
```
package "exp"
```
`EXP` represents the starting semantic rule exercising recursive descent through calling itself within the characters 'a' then 'b', 'a' then 'c', or empty, choice based on the input grammar. Here, the `|` operator represents unordered choice. See the [grammar for details.](../../gogll.md). The recursion descent 
```
EXP     : "a" EXP "b" 
        | "a" EXP "c" 
        | empty         ;
```
#
### **COPYRIGHT AND LICENSING INFORMATION**
**Copyright 2021 Brynn Harrington and Emily Hoppe**

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.

You may obtain a copy of the License [here](http://www.apache.org/licenses/LICENSE-2.0) or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.