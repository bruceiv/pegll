# **`exp` Grammar**

### **AUTHORSHIP INFORMATION**
#### *Author :* Brynn Harrington and Emily Hoppe Copyright (C) 2021
#### *Adapted from :* Aaron Moss's [`eggr` Egg Grammar](https://github.com/bruceiv/egg/blob/deriv/grammars/exp.egg)
#### *Creation Date :* June 11, 2021 
#### *Last Modified :* June 18, 2021
#### *Copyright and Licensing Information :* See end of file.

###  **GENERAL DESCRIPTION**
An originally Egg Parsing grammar created by Aaron Moss ported into the GoGLL grammar to determine if a grammar, under recursive descent, will have exponential runtime.
Modification of `exp` grammar from [Egg](https://github.com/bruceiv/egg/blob/deriv/grammars/exp.egg) to test an example grammar with possible exponential runtime under recursive descent.

### **`exp` Grammar Guide**
NEED TO FINISH ONE GRAMMAR IS WORKING 

 See the [grammar for details.](../../gogll.md)
```
package "exp"

S : A '\n' ;

A : 'a' A 'b' 
  | 'a' A 'c' 
  | empty ;

```
#
### **COPYRIGHT AND LICENSING INFORMATION**
**Copyright 2021 Brynn Harrington and Emily Hoppe**

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.

You may obtain a copy of the License [here](http://www.apache.org/licenses/LICENSE-2.0) or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.