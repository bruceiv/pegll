# **`nested` Grammar**

### **AUTHORSHIP INFORMATION**
#### *Author :* Brynn Harrington and Emily Hoppe Copyright (C) 2021
#### *Adapted from :* Aaron Moss's [`eggr` Egg Grammar](https://github.com/bruceiv/egg/blob/deriv/grammars/nested.egg)
#### *Creation Date :* June 11, 2021 
#### *Last Modified :* June 18, 2021
#### *Copyright and Licensing Information :* See end of file.

###  **GENERAL DESCRIPTION**
An originally Egg Parsing grammar created by Aaron Moss ported into the GoGLL grammar to determine test a given structure. Modification of `nested` grammar from [Egg](https://github.com/bruceiv/egg/blob/deriv/grammars/nested.egg) to test an example structure similar to XML.
### **STATUS ON GRAMMAR**
#### *Markdown File Creation:* Working
#### *Parser Generated :* Complete
#### *Test File Creation:* Incomplete
#### *Testing Results:* Unknown
### **`nested` Grammar Guide**
The following grammar utilizes recursion through semantic rules to test matching nested parentheses. 
```
package "nested"
```
The semantic rule `String` represents the starting rule for testing the nested parenthesis through simply calling on `Content`.
```
String      : Content ; 
```
`Content` is a semantic, recursive rule calling on `ParensOrChar` and `Content` allowing for matching the nested parenthesis in the grammar passed unless the string is `empty` as defined by the [grammar rules.](../../gogll.md) This is an ordered choice signified by the `/` operator. 
```
Content     : ParensOrChar Content 
            / empty ;
```
`Parens` is a semantic rule that calls on `Content` to match the content within the nested parenthesis. 
```
Parens      : open Content close ;
```
`ParensOrChar` is a semantic rule that represents an ordered choice between `Parens` and `char`.
```
ParensOrChar: Parens 
            / char ;
```
`open`, `close`, and `letter` represent lexical rules for the open parentheses '(', the close parentheses')', and one or more characters from the unicode letter class through the `<>` operator, respectively. 
```
open        : '(' ;
close       : ')' ;
char        : < letter > ;

```
#
### **COPYRIGHT AND LICENSING INFORMATION**
**Copyright 2021 Brynn Harrington and Emily Hoppe**

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.

You may obtain a copy of the License [here](http://www.apache.org/licenses/LICENSE-2.0) or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.