# `A~BC` Grammar
### **AUTHORSHIP INFORMATION**
#### *Author :* Aaron Moss Copyright (C) 2021
#### *Adapted from :* Aaron Moss's [`abc` Egg Grammar](https://github.com/bruceiv/egg/blob/deriv/grammars/abc.egg)
#### *Creation Date :* June 11, 2021 
#### *Last Modified :* June 18, 2021
#### *Copyright and Licensing Information :* See end of file.

###  **GENERAL DESCRIPTION**
An originally Egg Parsing grammar created by Aaron Moss ported into the GoGLL grammar to test ordered choice through the elimination of the `"ab"` branch. Modification of `abc` grammar from [Egg](https://github.com/bruceiv/egg/blob/deriv/grammars/abc.egg) to meaningfully exercise ordered choice by eliminating `"ab"` branch.

### **STATUS ON GRAMMAR**
#### *Markdown File Creation:* Complete
#### *Parser Generated :* Complete
#### *Test File Creation:* Complete
#### *Testing Results:* Passed

### **`A~BC` Grammar Guide**
The `/` syntax rule after`"ab"` and in front of `as` instructs the parser to recognize the ordered choice of these tokens. See the [grammar for details.](../../gogll.md)

```
package "axbc"
```
`AxBC` is a semantic starting rule that is composed of the semantic rule `AorB` followed by the character 'c'.
```
AxBC    : AorB "c"  ;
```
`AorB` is a semantic rule that represents the ordered choice of a string beginning with `'a'` zero or more times, with the alternate choice being the string `"ab"`. Here the lexical rule `/` is utilized for ordered choice/ See the [grammar for details.](../../gogll.md)
```
AorB    : as / "ab" ;
```
`repa0x` consists of the repeated token of `a` zero or more times. The use of the bracketed expression`{}` within this grammar is to indicate a set of tokens or identifiers repeated zero or more times. See the [grammar for details.](../../gogll.md)
```
repa0x  : { 'a' }   ;
```
#
### **COPYRIGHT AND LICENSING INFORMATION**
**Copyright 2021 Aaron Moss**

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.

You may obtain a copy of the License [here](http://www.apache.org/licenses/LICENSE-2.0) or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

