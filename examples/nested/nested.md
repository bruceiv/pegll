# **`nested` Grammar**

### **AUTHORSHIP INFORMATION**
#### *Author :* Brynn Harrington and Emily Hoppe Copyright (C) 2021
#### *Adapted from :* Aaron Moss's [`eggr` Egg Grammar](https://github.com/bruceiv/egg/blob/deriv/grammars/nested.egg)
#### *Creation Date :* June 11, 2021 
#### *Last Modified :* June 18, 2021
#### *Copyright and Licensing Information :* See end of file.

###  **GENERAL DESCRIPTION**
An originally Egg Parsing grammar created by Aaron Moss ported into the GoGLL grammar to determine test a given structure. Modification of `nested` grammar from [Egg](https://github.com/bruceiv/egg/blob/deriv/grammars/nested.egg) to test an example structure similar to XML.

### **`miniegg` Grammar Guide**
CURRENT ERRORS: 
- only works for single parentheses 
- need to implement a zero+ number of possible nesting 

See the [grammar for details.](../../gogll.md)

### **STATUS ON GRAMMAR**
#### *Markdown File Creation:* Not working 
#### *Parser Generated :* Incomplete - files generated from partially correct grammar 
#### *Test File Creation:* Incomplete
#### *Testing Results:* Unknown
```
package "nested"

String      : Content ; 
Content     : parens / char ;
parens      : '(' ')' ;
char        : < letter > ;
```
### **IN PROGRESS GRAMMARS**
**Original / Not working**
    String      : Content ;
    Content     : { Parens / Char } ;
    Parens      : '(' Content ')' ;
    Char        : < letter > ;
**Partially Working**
There is a definite difference between how the parser recognizes uppercase and lowercase.
    String      : Content ;
    Content     : parens / char ;
    parens      : '(' ')' ;
    char        : < letter > ;
#
### **COPYRIGHT AND LICENSING INFORMATION**
**Copyright 2021 Brynn Harrington and Emily Hoppe**

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.

You may obtain a copy of the License [here](http://www.apache.org/licenses/LICENSE-2.0) or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.